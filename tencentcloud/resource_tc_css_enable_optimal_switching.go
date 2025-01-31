/*
Provides a resource to create a css enable_optimal_switching

~> **NOTE:** This resource is only valid when the push stream. When the push stream ends, it will be deleted.

Example Usage

```hcl
resource "tencentcloud_css_enable_optimal_switching" "enable_optimal_switching" {
  stream_name     = "1308919341_test"
  enable_switch   = 1
  host_group_name = "test-group"
}
```

Import

css domain can be imported using the id, e.g.

```
terraform import tencentcloud_css_enable_optimal_switching.enable_optimal_switching streamName
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	css "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudCssEnableOptimalSwitching() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCssEnableOptimalSwitchingCreate,
		Read:   resourceTencentCloudCssEnableOptimalSwitchingRead,
		Update: resourceTencentCloudCssEnableOptimalSwitchingUpdate,
		Delete: resourceTencentCloudCssEnableOptimalSwitchingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"stream_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Stream id.",
			},

			"enable_switch": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "`0`:disabled, `1`:enable.",
			},

			"host_group_name": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Group name.",
			},
		},
	}
}

func resourceTencentCloudCssEnableOptimalSwitchingCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_enable_optimal_switching.create")()
	defer inconsistentCheck(d, meta)()

	var streamName string
	if v, ok := d.GetOk("stream_name"); ok {
		streamName = v.(string)
	}

	d.SetId(streamName)

	return resourceTencentCloudCssEnableOptimalSwitchingUpdate(d, meta)
}

func resourceTencentCloudCssEnableOptimalSwitchingRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_enable_optimal_switching.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CssService{client: meta.(*TencentCloudClient).apiV3Conn}

	streamName := d.Id()

	backupStream, err := service.DescribeCssBackupStreamByStreamName(ctx, streamName)
	if err != nil {
		return err
	}

	if backupStream == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CssBackupStream` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("stream_name", streamName)

	if backupStream.OptimalEnable != nil {
		_ = d.Set("enable_switch", backupStream.OptimalEnable)
	}

	if backupStream.HostGroupName != nil {
		_ = d.Set("host_group_name", backupStream.HostGroupName)
	}

	return nil
}

func resourceTencentCloudCssEnableOptimalSwitchingUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_enable_optimal_switching.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := css.NewEnableOptimalSwitchingRequest()

	streamName := d.Id()

	request.StreamName = &streamName

	if v, ok := d.GetOkExists("enable_switch"); ok {
		request.EnableSwitch = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("host_group_name"); ok {
		request.HostGroupName = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCssClient().EnableOptimalSwitching(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate css enableOptimalSwitching failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCssEnableOptimalSwitchingRead(d, meta)
}

func resourceTencentCloudCssEnableOptimalSwitchingDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_enable_optimal_switching.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
