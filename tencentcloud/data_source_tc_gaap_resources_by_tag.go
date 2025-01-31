/*
Use this data source to query detailed information of gaap resources by tag

Example Usage

```hcl
data "tencentcloud_gaap_resources_by_tag" "resources_by_tag" {
  tag_key = "tagKey"
  tag_value = "tagValue"
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

func dataSourceTencentCloudGaapResourcesByTag() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudGaapResourcesByTagRead,
		Schema: map[string]*schema.Schema{
			"tag_key": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Tag key.",
			},

			"tag_value": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Tag value.",
			},

			"resource_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Resource type, where:Proxy represents the proxy;ProxyGroup represents a proxy group;RealServer represents the Real Server.If this field is not specified, all resources under the label will be queried.",
			},

			"resource_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of resources corresponding to labels.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource type, where:Proxy represents the proxy,ProxyGroup represents a proxy group,RealServer represents the real server.",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource Id.",
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

func dataSourceTencentCloudGaapResourcesByTagRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_gaap_resources_by_tag.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("tag_key"); ok {
		paramMap["TagKey"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("tag_value"); ok {
		paramMap["TagValue"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resource_type"); ok {
		paramMap["ResourceType"] = helper.String(v.(string))
	}

	service := GaapService{client: meta.(*TencentCloudClient).apiV3Conn}

	var resourceSet []*gaap.TagResourceInfo

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeGaapResourcesByTagByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		resourceSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(resourceSet))
	tmpList := make([]map[string]interface{}, 0, len(resourceSet))

	if resourceSet != nil {
		for _, tagResourceInfo := range resourceSet {
			tagResourceInfoMap := map[string]interface{}{}

			if tagResourceInfo.ResourceType != nil {
				tagResourceInfoMap["resource_type"] = tagResourceInfo.ResourceType
			}

			if tagResourceInfo.ResourceId != nil {
				tagResourceInfoMap["resource_id"] = tagResourceInfo.ResourceId
			}

			ids = append(ids, *tagResourceInfo.ResourceId)
			tmpList = append(tmpList, tagResourceInfoMap)
		}

		_ = d.Set("resource_set", tmpList)
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
