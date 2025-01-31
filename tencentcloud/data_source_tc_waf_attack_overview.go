/*
Use this data source to query detailed information of waf attack_overview

Example Usage

Basic Query

```hcl
data "tencentcloud_waf_attack_overview" "example" {
  from_time = "2023-09-01 00:00:00"
  to_time   = "2023-09-07 00:00:00"
}
```

Query by filter

```hcl
data "tencentcloud_waf_attack_overview" "example" {
  from_time   = "2023-09-01 00:00:00"
  to_time     = "2023-09-07 00:00:00"
  appid       = 1304251372
  domain      = "test.com"
  edition     = "clb-waf"
  instance_id = "waf_2kxtlbky00b2v1fn"
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

func dataSourceTencentCloudWafAttackOverview() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudWafAttackOverviewRead,
		Schema: map[string]*schema.Schema{
			"from_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Begin time.",
			},
			"to_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "End time.",
			},
			"appid": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "App id.",
			},
			"domain": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"edition": {
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(EDITION_TYPE),
				Description:  "support `sparta-waf`, `clb-waf`, otherwise not filter.",
			},
			"instance_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Waf instanceId, otherwise not filter.",
			},
			"access_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Access count.",
			},
			"attack_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Attack count.",
			},
			"acl_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Access control count.",
			},
			"cc_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "CC attack count.",
			},
			"bot_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Bot attack count.",
			},
			"api_assets_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Api asset count.",
			},
			"api_risk_event_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Number of API risk events.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudWafAttackOverviewRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_waf_attack_overview.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId          = getLogId(contextNil)
		ctx            = context.WithValue(context.TODO(), logIdKey, logId)
		service        = WafService{client: meta.(*TencentCloudClient).apiV3Conn}
		attackOverview *waf.DescribeAttackOverviewResponseParams
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("from_time"); ok {
		paramMap["FromTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("to_time"); ok {
		paramMap["ToTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("appid"); ok {
		paramMap["Appid"] = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("domain"); ok {
		paramMap["Domain"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("edition"); ok {
		paramMap["Edition"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["InstanceID"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafAttackOverviewByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		attackOverview = result
		return nil
	})

	if err != nil {
		return err
	}

	if attackOverview.AccessCount != nil {
		_ = d.Set("access_count", attackOverview.AccessCount)
	}

	if attackOverview.AttackCount != nil {
		_ = d.Set("attack_count", attackOverview.AttackCount)
	}

	if attackOverview.ACLCount != nil {
		_ = d.Set("acl_count", attackOverview.ACLCount)
	}

	if attackOverview.CCCount != nil {
		_ = d.Set("cc_count", attackOverview.CCCount)
	}

	if attackOverview.BotCount != nil {
		_ = d.Set("bot_count", attackOverview.BotCount)
	}

	if attackOverview.ApiAssetsCount != nil {
		_ = d.Set("api_assets_count", attackOverview.ApiAssetsCount)
	}

	if attackOverview.ApiRiskEventCount != nil {
		_ = d.Set("api_risk_event_count", attackOverview.ApiRiskEventCount)
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
