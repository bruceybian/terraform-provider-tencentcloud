/*
Use this data source to query detailed information of gaap real servers status

Example Usage

```hcl
data "tencentcloud_gaap_real_servers_status" "real_servers_status" {
  real_server_ids = ["rs-qcygnwpd"]
}
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gaap "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudGaapRealServersStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudGaapRealServersStatusRead,
		Schema: map[string]*schema.Schema{
			"real_server_ids": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Real Server Ids.",
			},

			"real_server_status_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Real Server Status Set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"real_server_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Real Server Id.",
						},
						"bind_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Bind Status, 0 indicates unbound, 1 indicates bound by rules or listeners.",
						},
						"proxy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bind the proxy ID of this real server, which is an empty string when not bound.",
						},
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bind the group ID of this real server, which is an empty string when not bound.Note: This field may return null, indicating that a valid value cannot be obtained.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudGaapRealServersStatusRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_gaap_real_servers_status.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("real_server_ids"); ok {
		realServerIdsSet := v.(*schema.Set).List()
		paramMap["RealServerIds"] = helper.InterfacesStringsPoint(realServerIdsSet)
	}

	service := GaapService{client: meta.(*TencentCloudClient).apiV3Conn}

	var realServerStatusSet []*gaap.RealServerStatus

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeGaapRealServersStatusByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		realServerStatusSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(realServerStatusSet))
	tmpList := make([]map[string]interface{}, 0, len(realServerStatusSet))

	if realServerStatusSet != nil {
		for _, realServerStatus := range realServerStatusSet {
			realServerStatusMap := map[string]interface{}{}

			if realServerStatus.RealServerId != nil {
				realServerStatusMap["real_server_id"] = realServerStatus.RealServerId
			}

			if realServerStatus.BindStatus != nil {
				realServerStatusMap["bind_status"] = realServerStatus.BindStatus
			}

			if realServerStatus.ProxyId != nil {
				realServerStatusMap["proxy_id"] = realServerStatus.ProxyId
			}

			if realServerStatus.GroupId != nil {
				realServerStatusMap["group_id"] = realServerStatus.GroupId
			}

			ids = append(ids, *realServerStatus.RealServerId)
			tmpList = append(tmpList, realServerStatusMap)
		}

		_ = d.Set("real_server_status_set", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
