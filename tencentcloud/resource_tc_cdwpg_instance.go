/*
Provides a resource to create a cdwpg instance

Example Usage

```hcl
resource "tencentcloud_cdwpg_instance" "instance" {
	instance_name  = "test_cdwpg"
	zone           = "ap-guangzhou-6"
	user_vpc_id    = "vpc-xxxxxx"
	user_subnet_id = "subnet-xxxxxx"
	charge_properties {
	  renew_flag  = 0
	  time_span   = 1
	  time_unit   = "h"
	  charge_type = "POSTPAID_BY_HOUR"

	}
	admin_password = "xxxxxx"
	resources {
	  spec_name = "S_4_16_H_CN"
	  count     = 2
	  disk_spec {
		disk_type  = "CLOUD_HSSD"
		disk_size  = 200
		disk_count = 1
	  }
	  type = "cn"

	}
	resources {
	  spec_name = "S_4_16_H_CN"
	  count     = 2
	  disk_spec {
		disk_type  = "CLOUD_HSSD"
		disk_size  = 20
		disk_count = 10
	  }
	  type = "dn"

	}
	tags = {
	  "tagKey" = "tagValue"
	}
}
```

Import

cdwpg instance can be imported using the id, e.g.

```
terraform import tencentcloud_cdwpg_instance.instance instance_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cdwpg "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdwpg/v20201230"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudCdwpgInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCdwpgInstanceCreate,
		Read:   resourceTencentCloudCdwpgInstanceRead,
		Update: resourceTencentCloudCdwpgInstanceUpdate,
		Delete: resourceTencentCloudCdwpgInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "instance name.",
			},

			"zone": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Availability Zone.",
			},

			"user_vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "private network.",
			},

			"user_subnet_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "subnet.",
			},

			"charge_properties": {
				Required:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "instance billing mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"renew_flag": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "0-no automatic renewal,1-automatic renewalNote: This field may return null, indicating that a valid value cannot be obtained.",
						},
						"time_span": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Time RangeNote: This field may return null, indicating that a valid value cannot be obtained.",
						},
						"time_unit": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Time Unit,Generally h and mNote: This field may return null, indicating that a valid value cannot be obtained.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Charge type, vaild values: PREPAID, POSTPAID_BY_HOUR.",
						},
					},
				},
			},

			"admin_password": {
				Required:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
				Description: "cluster password.",
			},

			"resources": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "resource information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spec_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "resource name.",
						},
						"count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "resource count.",
						},
						"disk_spec": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "disk Information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disk_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "disk type.",
									},
									"disk_size": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "disk size.",
									},
									"disk_count": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "disk count.",
									},
								},
							},
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "resource type.",
						},
					},
				},
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudCdwpgInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cdwpg_instance.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = cdwpg.NewCreateInstanceByApiRequest()
		response   = cdwpg.NewCreateInstanceByApiResponse()
		instanceId string
	)
	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("zone"); ok {
		request.Zone = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_vpc_id"); ok {
		request.UserVPCId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_subnet_id"); ok {
		request.UserSubnetId = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "charge_properties"); ok {
		chargeProperties := cdwpg.ChargeProperties{}
		if v, ok := dMap["renew_flag"]; ok {
			chargeProperties.RenewFlag = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["time_span"]; ok {
			chargeProperties.TimeSpan = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["time_unit"]; ok {
			chargeProperties.TimeUnit = helper.String(v.(string))
		}
		if v, ok := dMap["charge_type"]; ok {
			chargeProperties.ChargeType = helper.String(v.(string))
		}
		request.ChargeProperties = &chargeProperties
	}

	if v, ok := d.GetOk("admin_password"); ok {
		request.AdminPassword = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resources"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			resourceSpecNew := cdwpg.ResourceSpecNew{}
			if v, ok := dMap["spec_name"]; ok {
				resourceSpecNew.SpecName = helper.String(v.(string))
			}
			if v, ok := dMap["count"]; ok {
				resourceSpecNew.Count = helper.IntInt64(v.(int))
			}
			if diskSpecMap, ok := helper.InterfaceToMap(dMap, "disk_spec"); ok {
				cBSSpec := cdwpg.CBSSpec{}
				if v, ok := diskSpecMap["disk_type"]; ok {
					cBSSpec.DiskType = helper.String(v.(string))
				}
				if v, ok := diskSpecMap["disk_size"]; ok {
					cBSSpec.DiskSize = helper.IntInt64(v.(int))
				}
				if v, ok := diskSpecMap["disk_count"]; ok {
					cBSSpec.DiskCount = helper.IntInt64(v.(int))
				}
				resourceSpecNew.DiskSpec = &cBSSpec
			}
			if v, ok := dMap["type"]; ok {
				resourceSpecNew.Type = helper.String(v.(string))
			}
			request.Resources = append(request.Resources, &resourceSpecNew)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCdwpgClient().CreateInstanceByApi(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cdwpg instance failed, reason:%+v", logId, err)
		return err
	}

	instanceId = *response.Response.InstanceId
	service := CdwpgService{client: meta.(*TencentCloudClient).apiV3Conn}
	conf := BuildStateChangeConf([]string{}, []string{"Serving"}, 10*readRetryTimeout, time.Second, service.InstanceStateRefreshFunc(instanceId, []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	d.SetId(instanceId)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::cdwpg:%s:uin/:cdwpgInstance/%s", region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudCdwpgInstanceRead(d, meta)
}

func resourceTencentCloudCdwpgInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cdwpg_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CdwpgService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	instance, err := service.DescribeCdwpgInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if instance == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CdwpgInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if instance.InstanceName != nil {
		_ = d.Set("instance_name", instance.InstanceName)
	}

	if instance.Zone != nil {
		_ = d.Set("zone", instance.Zone)
	}

	if instance.UserVPCID != nil {
		_ = d.Set("user_vpc_id", instance.UserVPCID)
	}

	if instance.UserSubnetID != nil {
		_ = d.Set("user_subnet_id", instance.UserSubnetID)
	}

	if instance.ChargeProperties != nil {
		chargePropertiesMap := map[string]interface{}{}

		if instance.ChargeProperties.RenewFlag != nil {
			chargePropertiesMap["renew_flag"] = instance.ChargeProperties.RenewFlag
		}

		if instance.ChargeProperties.TimeSpan != nil {
			chargePropertiesMap["time_span"] = instance.ChargeProperties.TimeSpan
		}

		if instance.ChargeProperties.TimeUnit != nil {
			chargePropertiesMap["time_unit"] = instance.ChargeProperties.TimeUnit
		}

		if instance.ChargeProperties.ChargeType != nil {
			chargePropertiesMap["charge_type"] = instance.ChargeProperties.ChargeType
		}

		_ = d.Set("charge_properties", []interface{}{chargePropertiesMap})
	}

	if instance.Resources != nil {
		resourcesList := []interface{}{}
		for _, resource := range instance.Resources {
			resourcesMap := map[string]interface{}{}

			if resource.SpecName != nil {
				resourcesMap["spec_name"] = resource.SpecName
			}

			if resource.Count != nil {
				resourcesMap["count"] = resource.Count
			}

			if resource.DiskSpec != nil {
				diskSpecMap := map[string]interface{}{}

				if resource.DiskSpec.DiskType != nil {
					diskSpecMap["disk_type"] = resource.DiskSpec.DiskType
				}

				if resource.DiskSpec.DiskSize != nil {
					diskSpecMap["disk_size"] = resource.DiskSpec.DiskSize
				}

				if resource.DiskSpec.DiskCount != nil {
					diskSpecMap["disk_count"] = resource.DiskSpec.DiskCount
				}

				resourcesMap["disk_spec"] = []interface{}{diskSpecMap}
			}

			if resource.Type != nil {
				resourcesMap["type"] = resource.Type
			}

			resourcesList = append(resourcesList, resourcesMap)
		}

		_ = d.Set("resources", resourcesList)

	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "cdwpg", "cdwpgInstance", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudCdwpgInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cdwpg_instance.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cdwpg.NewModifyInstanceRequest()

	instanceId := d.Id()

	request.InstanceId = &instanceId

	immutableArgs := []string{"zone", "user_vpc_id", "user_subnet_id", "charge_properties", "admin_password", "resources"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if d.HasChange("instance_name") {
		if v, ok := d.GetOk("instance_name"); ok {
			request.InstanceName = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCdwpgClient().ModifyInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update cdwpg instance failed, reason:%+v", logId, err)
		return err
	}

	service := CdwpgService{client: meta.(*TencentCloudClient).apiV3Conn}
	conf := BuildStateChangeConf([]string{}, []string{"Serving"}, 10*readRetryTimeout, time.Second, service.InstanceStateRefreshFunc(instanceId, []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	if d.HasChange("tags") {
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("cdwpg", "cdwpgInstance", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudCdwpgInstanceRead(d, meta)
}

func resourceTencentCloudCdwpgInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cdwpg_instance.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CdwpgService{client: meta.(*TencentCloudClient).apiV3Conn}
	instanceId := d.Id()

	if err := service.DeleteCdwpgInstanceById(ctx, instanceId); err != nil {
		return err
	}

	conf := BuildStateChangeConf([]string{}, []string{"Deleted"}, 10*readRetryTimeout, time.Second, service.InstanceStateRefreshFunc(instanceId, []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}
	return nil
}
