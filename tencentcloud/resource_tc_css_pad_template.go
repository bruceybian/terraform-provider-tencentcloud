/*
Provides a resource to create a css pad_template

Example Usage

```hcl
resource "tencentcloud_css_pad_template" "pad_template" {
  description   = "pad template"
  max_duration  = 120000
  template_name = "tf-pad"
  type          = 1
  url           = "https://livewatermark-1251132611.cos.ap-guangzhou.myqcloud.com/1308919341/watermark_img_1698736540399_1441698123618_.pic.jpg"
  wait_duration = 2000
}
```

Import

css pad_template can be imported using the id, e.g.

```
terraform import tencentcloud_css_pad_template.pad_template templateId
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	css "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudCssPadTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCssPadTemplateCreate,
		Read:   resourceTencentCloudCssPadTemplateRead,
		Update: resourceTencentCloudCssPadTemplateUpdate,
		Delete: resourceTencentCloudCssPadTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"template_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Template namelimit 255 bytes.",
			},

			"url": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Pad content.",
			},

			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Description content.limit length 1024 bytes.",
			},

			"wait_duration": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Stop stream wait time.limit: 0 - 30000 ms.",
			},

			"max_duration": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Max pad duration.limit: 0 - 9999999 ms.",
			},

			"type": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Pad content type.1: picture.2: video.default: 1.",
			},
		},
	}
}

func resourceTencentCloudCssPadTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_pad_template.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = css.NewCreateLivePadTemplateRequest()
		response   = css.NewCreateLivePadTemplateResponse()
		templateId int64
	)
	if v, ok := d.GetOk("template_name"); ok {
		request.TemplateName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("url"); ok {
		request.Url = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("wait_duration"); ok {
		request.WaitDuration = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("max_duration"); ok {
		request.MaxDuration = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("type"); ok {
		request.Type = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCssClient().CreateLivePadTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create css padTemplate failed, reason:%+v", logId, err)
		return err
	}

	templateId = *response.Response.TemplateId
	d.SetId(helper.Int64ToStr(templateId))

	return resourceTencentCloudCssPadTemplateRead(d, meta)
}

func resourceTencentCloudCssPadTemplateRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_pad_template.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CssService{client: meta.(*TencentCloudClient).apiV3Conn}

	templateId := d.Id()
	templateIdInt64, err := strconv.ParseInt(templateId, 10, 64)
	if err != nil {
		return fmt.Errorf("TemplateId format type error: %s", err.Error())
	}

	padTemplate, err := service.DescribeCssPadTemplateById(ctx, templateIdInt64)
	if err != nil {
		return err
	}

	if padTemplate == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CssPadTemplate` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if padTemplate.TemplateName != nil {
		_ = d.Set("template_name", padTemplate.TemplateName)
	}

	if padTemplate.Url != nil {
		_ = d.Set("url", padTemplate.Url)
	}

	if padTemplate.Description != nil {
		_ = d.Set("description", padTemplate.Description)
	}

	if padTemplate.WaitDuration != nil {
		_ = d.Set("wait_duration", padTemplate.WaitDuration)
	}

	if padTemplate.MaxDuration != nil {
		_ = d.Set("max_duration", padTemplate.MaxDuration)
	}

	if padTemplate.Type != nil {
		_ = d.Set("type", padTemplate.Type)
	}

	return nil
}

func resourceTencentCloudCssPadTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_pad_template.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := css.NewModifyLivePadTemplateRequest()

	templateId := d.Id()
	templateIdInt64, _ := strconv.ParseInt(templateId, 10, 64)
	templateIdUint := uint64(templateIdInt64)

	request.TemplateId = &templateIdUint

	if v, ok := d.GetOk("template_name"); ok {
		request.TemplateName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("url"); ok {
		request.Url = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("wait_duration"); ok {
		request.WaitDuration = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("max_duration"); ok {
		request.MaxDuration = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("type"); ok {
		request.Type = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCssClient().ModifyLivePadTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update css padTemplate failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCssPadTemplateRead(d, meta)
}

func resourceTencentCloudCssPadTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_pad_template.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CssService{client: meta.(*TencentCloudClient).apiV3Conn}
	templateId := d.Id()
	templateIdInt64, _ := strconv.ParseInt(templateId, 10, 64)

	if err := service.DeleteCssPadTemplateById(ctx, templateIdInt64); err != nil {
		return err
	}

	return nil
}
