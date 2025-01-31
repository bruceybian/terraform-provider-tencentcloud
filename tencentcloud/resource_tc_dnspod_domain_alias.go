/*
Provides a resource to create a dnspod domain_alias

Example Usage

```hcl
resource "tencentcloud_dnspod_domain_alias" "domain_alias" {
  domain_alias = "dnspod.com"
  domain = "dnspod.cn"
}
```

Import

dnspod domain_alias can be imported using the id, e.g.

```
terraform import tencentcloud_dnspod_domain_alias.domain_alias domain#domain_alias_id
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
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudDnspodDomainAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDnspodDomainAliasCreate,
		Read:   resourceTencentCloudDnspodDomainAliasRead,
		Delete: resourceTencentCloudDnspodDomainAliasDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain_alias": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Domain alias.",
			},

			"domain": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},

			"domain_alias_id": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Domain alias ID.",
			},
		},
	}
}

func resourceTencentCloudDnspodDomainAliasCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_dnspod_domain_alias.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request       = dnspod.NewCreateDomainAliasRequest()
		response      = dnspod.NewCreateDomainAliasResponse()
		domain        string
		domainAliasId int64
	)
	if v, ok := d.GetOk("domain_alias"); ok {
		request.DomainAlias = helper.String(v.(string))
	}

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
		request.Domain = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDnsPodClient().CreateDomainAlias(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create dnspod domain_alias failed, reason:%+v", logId, err)
		return err
	}

	domainAliasId = *response.Response.DomainAliasId
	// d.SetId(helper.String(domainAlias))
	d.SetId(strings.Join([]string{domain, helper.Int64ToStr(domainAliasId)}, FILED_SP))

	return resourceTencentCloudDnspodDomainAliasRead(d, meta)
}

func resourceTencentCloudDnspodDomainAliasRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_dnspod_domain_alias.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DnspodService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	domain := idSplit[0]
	domainAliasId := helper.StrToInt64(idSplit[1])

	domainAliasInfo, err := service.DescribeDnspodDomainAliasById(ctx, domain, domainAliasId)
	if err != nil {
		return err
	}

	if domainAliasInfo == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DnspodDomainAlias` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)

	if domainAliasInfo.DomainAlias != nil {
		_ = d.Set("domain_alias", domainAliasInfo.DomainAlias)
	}

	if domainAliasInfo.Id != nil {
		_ = d.Set("domain_alias_id", domainAliasInfo.Id)
	}

	return nil
}

func resourceTencentCloudDnspodDomainAliasDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_dnspod_domain_alias.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DnspodService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	domainAlias := idSplit[0]
	domainAliasId := helper.StrToInt64(idSplit[1])

	if err := service.DeleteDnspodDomainAliasById(ctx, domainAlias, domainAliasId); err != nil {
		return err
	}

	return nil
}
