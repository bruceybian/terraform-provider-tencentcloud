/*
Provides a resource to create a mps event

Example Usage

```hcl
resource "tencentcloud_mps_event" "event" {
  event_name = "you-event-name"
  description = "event description"
}
```

Import

mps event can be imported using the id, e.g.

```
terraform import tencentcloud_mps_event.event event_id
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	mps "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mps/v20190612"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudMpsEvent() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudMpsEventCreate,
		Read:   resourceTencentCloudMpsEventRead,
		Update: resourceTencentCloudMpsEventUpdate,
		Delete: resourceTencentCloudMpsEventDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"event_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Event name.",
			},

			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Event description.",
			},
		},
	}
}

func resourceTencentCloudMpsEventCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mps_event.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = mps.NewCreateStreamLinkEventRequest()
		response = mps.NewCreateStreamLinkEventResponse()
		eventId  string
	)
	if v, ok := d.GetOk("event_name"); ok {
		request.EventName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseMpsClient().CreateStreamLinkEvent(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create mps event failed, reason:%+v", logId, err)
		return err
	}

	if response.Response.Info != nil {
		eventId = *response.Response.Info.EventId
	}

	d.SetId(eventId)

	return resourceTencentCloudMpsEventRead(d, meta)
}

func resourceTencentCloudMpsEventRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mps_event.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := MpsService{client: meta.(*TencentCloudClient).apiV3Conn}

	eventId := d.Id()

	event, err := service.DescribeMpsEventById(ctx, eventId)
	if err != nil {
		return err
	}

	if event == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `MpsEvent` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if event.EventName != nil {
		_ = d.Set("event_name", event.EventName)
	}

	if event.Description != nil {
		_ = d.Set("description", event.Description)
	}

	return nil
}

func resourceTencentCloudMpsEventUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mps_event.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := mps.NewModifyStreamLinkEventRequest()

	eventId := d.Id()

	request.EventId = &eventId

	if d.HasChange("event_name") {
		if v, ok := d.GetOk("event_name"); ok {
			request.EventName = helper.String(v.(string))
		}
	}

	if d.HasChange("description") {
		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseMpsClient().ModifyStreamLinkEvent(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update mps event failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudMpsEventRead(d, meta)
}

func resourceTencentCloudMpsEventDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mps_event.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := MpsService{client: meta.(*TencentCloudClient).apiV3Conn}
	eventId := d.Id()

	if err := service.DeleteMpsEventById(ctx, eventId); err != nil {
		return err
	}

	return nil
}
