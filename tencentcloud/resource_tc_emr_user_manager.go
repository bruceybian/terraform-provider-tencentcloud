/*
Provides a resource to create a emr user

Example Usage

```hcl
data "tencentcloud_emr" "my_emr" {
  display_strategy = "clusterList"
}

resource "tencentcloud_emr_user_manager" "user_manager" {
  instance_id = data.tencentcloud_emr.my_emr.clusters.0.cluster_id
  user_name   = "tf-test"
  user_group  = "group1"
  password    = "tf@123456"
}

```

Import

emr user_manager can be imported using the id, e.g.

```
terraform import tencentcloud_emr_user_manager.user_manager instanceId#userName
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	emr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/emr/v20190103"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudEmrUserManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudEmrUserManagerCreate,
		Read:   resourceTencentCloudEmrUserManagerRead,
		Update: resourceTencentCloudEmrUserManagerUpdate,
		Delete: resourceTencentCloudEmrUserManagerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Cluster string ID.",
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Username.",
			},
			"user_group": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "User group membership.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PassWord.",
			},
			"user_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User type.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time.",
			},
			"support_download_keytab": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If support download keytab.",
			},
			"download_keytab_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Download keytab url.",
			},
		},
	}
}

func resourceTencentCloudEmrUserManagerCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_emr_user_manager.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = emr.NewAddUsersForUserManagerRequest()
		response   = emr.NewAddUsersForUserManagerResponse()
		instanceId string
		userName   string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
		instanceId = v.(string)
	}

	userInfoForUserManager := emr.UserInfoForUserManager{}
	if v, ok := d.GetOk("user_name"); ok {
		userInfoForUserManager.UserName = helper.String(v.(string))
		userName = v.(string)
	}
	if v, ok := d.GetOk("user_group"); ok {
		userInfoForUserManager.UserGroup = helper.String(v.(string))
	}
	if v, ok := d.GetOk("password"); ok {
		userInfoForUserManager.PassWord = helper.String(v.(string))
	}
	request.UserManagerUserList = append(request.UserManagerUserList, &userInfoForUserManager)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseEmrClient().AddUsersForUserManager(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create emr userManager failed, reason:%+v", logId, err)
		return err
	}

	if len(response.Response.FailedUserList) > 0 {
		return fmt.Errorf("add user failed, please try again.")
	}

	d.SetId(instanceId + FILED_SP + userName)

	return resourceTencentCloudEmrUserManagerRead(d, meta)
}

func resourceTencentCloudEmrUserManagerRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_emr_user_manager.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := EMRService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	userName := idSplit[1]

	userManager, err := service.DescribeEmrUserManagerById(ctx, instanceId, userName)
	if err != nil {
		return err
	}

	if userManager == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `EmrUserManager` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("instance_id", instanceId)

	if userManager.UserManagerUserList != nil {
		for _, userManager := range userManager.UserManagerUserList {
			if userManager.UserName != nil {
				_ = d.Set("user_name", userManager.UserName)
			}

			if userManager.UserGroup != nil {
				_ = d.Set("user_group", userManager.UserGroup)
			}

			if userManager.UserType != nil {
				_ = d.Set("user_type", userManager.UserType)
			}

			if userManager.CreateTime != nil {
				_ = d.Set("create_time", userManager.CreateTime)
			}

			if userManager.SupportDownLoadKeyTab != nil {
				_ = d.Set("support_download_keytab", userManager.SupportDownLoadKeyTab)
			}

			if userManager.DownLoadKeyTabUrl != nil {
				_ = d.Set("download_keytab_url", userManager.DownLoadKeyTabUrl)
			}
		}
	}

	return nil
}

func resourceTencentCloudEmrUserManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_emr_user_manager.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := emr.NewModifyUserManagerPwdRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	userName := idSplit[1]

	request.InstanceId = &instanceId
	request.UserName = &userName

	if d.HasChange("password") {
		if v, ok := d.GetOk("password"); ok {
			request.PassWord = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseEmrClient().ModifyUserManagerPwd(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update emr userManager failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudEmrUserManagerRead(d, meta)
}

func resourceTencentCloudEmrUserManagerDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_emr_user_manager.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := EMRService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	userName := idSplit[1]

	if err := service.DeleteEmrUserManagerById(ctx, instanceId, userName); err != nil {
		return err
	}

	return nil
}
