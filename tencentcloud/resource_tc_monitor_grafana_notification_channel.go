/*
Provides a resource to create a monitor grafanaNotificationChannel

Example Usage

```hcl
variable "availability_zone" {
  default = "ap-guangzhou-6"
}

resource "tencentcloud_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  name       = "tf_monitor_vpc"
}

resource "tencentcloud_subnet" "subnet" {
  vpc_id            = tencentcloud_vpc.vpc.id
  availability_zone = var.availability_zone
  name              = "tf_monitor_subnet"
  cidr_block        = "10.0.1.0/24"
}

resource "tencentcloud_monitor_grafana_instance" "foo" {
  instance_name         = "test-grafana"
  vpc_id                = tencentcloud_vpc.vpc.id
  subnet_ids            = [tencentcloud_subnet.subnet.id]
  grafana_init_password = "1234567890"
  enable_internet = false

  tags = {
    "createdBy" = "test"
  }
}

resource "tencentcloud_monitor_alarm_notice" "foo" {
  name                  = "tf_alarm_notice"
  notice_type           = "ALL"
  notice_language       = "zh-CN"

  user_notices    {
      receiver_type              = "USER"
      start_time                 = 0
      end_time                   = 1
      notice_way                 = ["SMS","EMAIL"]
      user_ids                   = [10001]
      group_ids                  = []
      phone_order                = [10001]
      phone_circle_times         = 2
      phone_circle_interval      = 50
      phone_inner_interval       = 60
      need_phone_arrive_notice   = 1
      phone_call_type            = "CIRCLE"
      weekday                    =[1,2,3,4,5,6,7]
  }

  url_notices {
      url    = "https://www.mytest.com/validate"
      end_time =  0
      start_time = 1
      weekday = [1,2,3,4,5,6,7]
  }
}

resource "tencentcloud_monitor_grafana_notification_channel" "grafanaNotificationChannel" {
  instance_id   = tencentcloud_monitor_grafana_instance.foo.id
  channel_name  = "tf-channel"
  org_id        = 1
  receivers     = [tencentcloud_monitor_alarm_notice.foo.amp_consumer_id]
  extra_org_ids = ["1"]
}

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
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudMonitorGrafanaNotificationChannel() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTencentCloudMonitorGrafanaNotificationChannelRead,
		Create: resourceTencentCloudMonitorGrafanaNotificationChannelCreate,
		Update: resourceTencentCloudMonitorGrafanaNotificationChannelUpdate,
		Delete: resourceTencentCloudMonitorGrafanaNotificationChannelDelete,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "grafana instance id.",
			},

			"channel_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "plugin id.",
			},

			"channel_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "channel name.",
			},

			"org_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Grafana organization which channel will be installed, default to 1 representing Main Org.",
			},

			"receivers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Computed:    true,
				Description: "cloud monitor notification template notice-id list.",
			},

			"extra_org_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "extra grafana organization id list, default to 1 representing Main Org.",
			},
		},
	}
}

func resourceTencentCloudMonitorGrafanaNotificationChannelCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_monitor_grafana_notification_channel.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = monitor.NewCreateGrafanaNotificationChannelRequest()
		response   *monitor.CreateGrafanaNotificationChannelResponse
		channelId  string
		instanceId string
	)

	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("channel_name"); ok {
		request.ChannelName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("org_id"); ok {
		request.OrgId = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("receivers"); ok {
		receiversSet := v.(*schema.Set).List()
		for i := range receiversSet {
			receivers := receiversSet[i].(string)
			request.Receivers = append(request.Receivers, &receivers)
		}
	}

	if v, ok := d.GetOk("extra_org_ids"); ok {
		extraOrgIdsSet := v.(*schema.Set).List()
		for i := range extraOrgIdsSet {
			extraOrgIds := extraOrgIdsSet[i].(string)
			request.ExtraOrgIds = append(request.ExtraOrgIds, &extraOrgIds)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseMonitorClient().CreateGrafanaNotificationChannel(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create monitor grafanaNotificationChannel failed, reason:%+v", logId, err)
		return err
	}

	channelId = *response.Response.ChannelId

	d.SetId(channelId + FILED_SP + instanceId)
	return resourceTencentCloudMonitorGrafanaNotificationChannelRead(d, meta)
}

func resourceTencentCloudMonitorGrafanaNotificationChannelRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_monitor_grafana_notification_channel.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := MonitorService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	channelId := idSplit[0]
	instanceId := idSplit[1]

	grafanaNotificationChannel, err := service.DescribeMonitorGrafanaNotificationChannel(ctx, channelId, instanceId)

	if err != nil {
		return err
	}

	if grafanaNotificationChannel == nil {
		d.SetId("")
		return fmt.Errorf("resource `grafanaNotificationChannel` %s does not exist", channelId)
	}

	_ = d.Set("instance_id", instanceId)

	if grafanaNotificationChannel.ChannelId != nil {
		_ = d.Set("channel_id", grafanaNotificationChannel.ChannelId)
	}

	if grafanaNotificationChannel.ChannelName != nil {
		_ = d.Set("channel_name", grafanaNotificationChannel.ChannelName)
	}

	_ = d.Set("receivers", grafanaNotificationChannel.Receivers)

	return nil
}

func resourceTencentCloudMonitorGrafanaNotificationChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_monitor_grafana_notification_channel.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := monitor.NewUpdateGrafanaNotificationChannelRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	channelId := idSplit[0]
	instanceId := idSplit[1]

	request.ChannelId = &channelId
	request.InstanceId = &instanceId

	if v, ok := d.GetOk("channel_name"); ok {
		request.ChannelName = helper.String(v.(string))
	}

	if d.HasChange("instance_id") {
		return fmt.Errorf("`instance_id` do not support change now.")
	}

	if d.HasChange("channel_name") {
		return fmt.Errorf("`channel_name` do not support change now.")
	}

	if d.HasChange("org_id") {
		return fmt.Errorf("`org_id` do not support change now.")
	}

	if d.HasChange("receivers") {
		return fmt.Errorf("`receivers` do not support change now.")
	}

	if d.HasChange("extra_org_ids") {
		return fmt.Errorf("`extra_org_ids` do not support change now.")
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseMonitorClient().UpdateGrafanaNotificationChannel(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudMonitorGrafanaNotificationChannelRead(d, meta)
}

func resourceTencentCloudMonitorGrafanaNotificationChannelDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_monitor_grafana_notification_channel.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := MonitorService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	channelId := idSplit[0]
	instanceId := idSplit[1]

	if err := service.DeleteMonitorGrafanaNotificationChannelById(ctx, channelId, instanceId); err != nil {
		return err
	}

	return nil
}
