/*
Use this data source to query detailed information of cam secret_last_used_time

Example Usage

```hcl
data "tencentcloud_cam_secret_last_used_time" "secret_last_used_time" {
  secret_id_list = ["xxxx"]
  }
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudCamSecretLastUsedTime() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamSecretLastUsedTimeRead,
		Schema: map[string]*schema.Schema{
			"secret_id_list": {
				Required:  true,
				Sensitive: true,
				Type:      schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Query the key ID list. Supports up to 10.",
			},

			"secret_id_last_used_rows": {
				Computed:    true,
				Sensitive:   true,
				Type:        schema.TypeList,
				Description: "Last used time list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secret Id.",
						},
						"last_used_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last used date (with 1 day delay).",
						},
						"last_secret_used_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last used timestamp.",
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

func dataSourceTencentCloudCamSecretLastUsedTimeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_cam_secret_last_used_time.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("secret_id_list"); ok {
		secretIdListSet := v.(*schema.Set).List()
		paramMap["SecretIdList"] = helper.InterfacesStringsPoint(secretIdListSet)
	}

	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	var secretIdLastUsedRows []*cam.SecretIdLastUsed

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCamSecretLastUsedTimeByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		secretIdLastUsedRows = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(secretIdLastUsedRows))
	tmpList := make([]map[string]interface{}, 0, len(secretIdLastUsedRows))

	if secretIdLastUsedRows != nil {
		for _, secretIdLastUsed := range secretIdLastUsedRows {
			secretIdLastUsedMap := map[string]interface{}{}

			if secretIdLastUsed.SecretId != nil {
				secretIdLastUsedMap["secret_id"] = secretIdLastUsed.SecretId
			}

			if secretIdLastUsed.LastUsedDate != nil {
				secretIdLastUsedMap["last_used_date"] = secretIdLastUsed.LastUsedDate
			}

			if secretIdLastUsed.LastSecretUsedDate != nil {
				secretIdLastUsedMap["last_secret_used_date"] = secretIdLastUsed.LastSecretUsedDate
			}

			ids = append(ids, *secretIdLastUsed.SecretId)
			tmpList = append(tmpList, secretIdLastUsedMap)
		}

		_ = d.Set("secret_id_last_used_rows", tmpList)
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
