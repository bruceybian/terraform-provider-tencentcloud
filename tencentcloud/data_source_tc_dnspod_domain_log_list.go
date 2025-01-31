/*
Use this data source to query detailed information of dnspod domain_log_list

Example Usage

```hcl
data "tencentcloud_dnspod_domain_log_list" "domain_log_list" {
  domain = "iac-tf.cloud"
  domain_id = 123
}
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	// dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudDnspodDomainLogList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudDnspodDomainLogListRead,
		Schema: map[string]*schema.Schema{
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},

			"domain_id": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Domain ID. The parameter DomainId has a higher priority than the parameter Domain. If the parameter DomainId is passed, the parameter Domain will be ignored. You can find all Domains and DomainIds through the DescribeDomainList interface.",
			},

			"log_list": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Domain Operation Log List. Note: This field may return null, indicating that no valid value can be obtained.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudDnspodDomainLogListRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_dnspod_domain_log_list.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var domain string

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
		paramMap["Domain"] = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("domain_id"); ok {
		paramMap["DomainId"] = helper.IntUint64(v.(int))
	}

	service := DnspodService{client: meta.(*TencentCloudClient).apiV3Conn}

	var logList []*string

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeDnspodDomainLogListByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		logList = result
		return nil
	})
	if err != nil {
		return err
	}

	// ids := make([]string, 0, len(logList))
	if logList != nil {
		_ = d.Set("log_list", logList)
	}

	d.SetId(helper.DataResourceIdHash(domain))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), logList); e != nil {
			return e
		}
	}
	return nil
}
