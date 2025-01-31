/*
Use this data source to query detailed information of waf attack_log_histogram

Example Usage

Obtain the specified domain name log information

```hcl
data "tencentcloud_waf_attack_log_histogram" "example" {
  domain       = "domain.com"
  start_time   = "2023-09-01 00:00:00"
  end_time     = "2023-09-29 00:00:00"
  query_string = "method:GET"
}
```

Obtain all domain name log information

```hcl
data "tencentcloud_waf_attack_log_histogram" "example" {
  domain       = "all"
  start_time   = "2023-09-01 00:00:00"
  end_time     = "2023-09-29 00:00:00"
  query_string = "method:GET"
}
```
*/
package tencentcloud

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	waf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf/v20180125"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudWafAttackLogHistogram() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudWafAttackLogHistogramRead,
		Schema: map[string]*schema.Schema{
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain for query, all domain use all.",
			},
			"start_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Begin time.",
			},
			"end_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "End time.",
			},
			"query_string": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Lucene grammar.",
			},
			"data": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "The statistics detail.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of logs.",
						},
						"time_stamp": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Timestamp.",
						},
					},
				},
			},
			"period": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Period.",
			},
			"total_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "total count.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudWafAttackLogHistogramRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_waf_attack_log_histogram.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = WafService{client: meta.(*TencentCloudClient).apiV3Conn}
		data    *waf.GetAttackHistogramResponseParams
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("domain"); ok {
		paramMap["Domain"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("start_time"); ok {
		paramMap["StartTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_time"); ok {
		paramMap["EndTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("query_string"); ok {
		paramMap["QueryString"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafAttackLogHistogramByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		data = result
		return nil
	})

	if err != nil {
		return err
	}

	if data.Period != nil {
		_ = d.Set("period", data.Period)
	}

	if data.TotalCount != nil {
		_ = d.Set("total_count", data.TotalCount)
	}

	tmpList := make([]map[string]interface{}, 0, len(data.Data))

	if data.Data != nil {
		for _, logHistogramInfo := range data.Data {
			logHistogramInfoMap := map[string]interface{}{}

			if logHistogramInfo.Count != nil {
				logHistogramInfoMap["count"] = logHistogramInfo.Count
			}

			if logHistogramInfo.TimeStamp != nil {
				logHistogramInfoMap["time_stamp"] = logHistogramInfo.TimeStamp
			}

			tmpList = append(tmpList, logHistogramInfoMap)
		}

		_ = d.Set("data", tmpList)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), d); e != nil {
			return e
		}
	}

	return nil
}
