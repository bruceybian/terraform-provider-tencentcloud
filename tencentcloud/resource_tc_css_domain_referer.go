/*
Provides a resource to create a css domain_referer

Example Usage

```hcl
resource "tencentcloud_css_domain_referer" "domain_referer" {
  allow_empty = 1
  domain_name = "test122.jingxhu.top"
  enable      = 0
  rules       = "example.com"
  type        = 1
}
```

Import

css domain_referer can be imported using the id, e.g.

```
terraform import tencentcloud_css_domain_referer.domain_referer domainName
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

func resourceTencentCloudCssDomainReferer() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCssDomainRefererCreate,
		Read:   resourceTencentCloudCssDomainRefererRead,
		Update: resourceTencentCloudCssDomainRefererUpdate,
		Delete: resourceTencentCloudCssDomainRefererDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain Name.",
			},

			"enable": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.",
			},

			"type": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "List type: 0: blacklist, 1: whitelist.",
			},

			"allow_empty": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Allow blank referers, 0: not allowed, 1: allowed.",
			},

			"rules": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The list of referers to; separate.",
			},
		},
	}
}

func resourceTencentCloudCssDomainRefererCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_domain_referer.create")()
	defer inconsistentCheck(d, meta)()

	var domainName string
	if v, ok := d.GetOk("domain_name"); ok {
		domainName = v.(string)
	}

	d.SetId(domainName)

	return resourceTencentCloudCssDomainRefererUpdate(d, meta)
}

func resourceTencentCloudCssDomainRefererRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_domain_referer.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CssService{client: meta.(*TencentCloudClient).apiV3Conn}

	domainName := d.Id()

	domainReferer, err := service.DescribeCssDomainRefererById(ctx, domainName)
	if err != nil {
		return err
	}

	if domainReferer == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CssDomainReferer` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if domainReferer.DomainName != nil {
		_ = d.Set("domain_name", domainReferer.DomainName)
	}

	if domainReferer.Enable != nil {
		_ = d.Set("enable", domainReferer.Enable)
	}

	if domainReferer.Type != nil {
		_ = d.Set("type", domainReferer.Type)
	}

	if domainReferer.AllowEmpty != nil {
		_ = d.Set("allow_empty", domainReferer.AllowEmpty)
	}

	if domainReferer.Rules != nil {
		_ = d.Set("rules", domainReferer.Rules)
	}

	return nil
}

func resourceTencentCloudCssDomainRefererUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_domain_referer.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := css.NewModifyLiveDomainRefererRequest()

	domainName := d.Id()

	request.DomainName = &domainName

	if v, ok := d.GetOkExists("enable"); ok {
		request.Enable = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("type"); ok {
		request.Type = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("allow_empty"); ok {
		request.AllowEmpty = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("rules"); ok {
		request.Rules = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCssClient().ModifyLiveDomainReferer(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update css domainReferer failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCssDomainRefererRead(d, meta)
}

func resourceTencentCloudCssDomainRefererDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_domain_referer.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
