/*
Provides a resource to create a gaap global domain dns

Example Usage

```hcl
resource "tencentcloud_gaap_global_domain_dns" "global_domain_dns" {
	domain_id = "dm-xxxxxx"
	proxy_id_list = ["link-xxxxxx"]
	nation_country_inner_codes = ["101001"]
}
```

Import

gaap global_domain_dns can be imported using the id, e.g.

```
terraform import tencentcloud_gaap_global_domain_dns.global_domain_dns ${domainId}#${dnsRecordId}
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gaap "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudGaapGlobalDomainDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudGaapGlobalDomainDnsCreate,
		Read:   resourceTencentCloudGaapGlobalDomainDnsRead,
		Update: resourceTencentCloudGaapGlobalDomainDnsUpdate,
		Delete: resourceTencentCloudGaapGlobalDomainDnsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain Id.",
			},

			"proxy_id_list": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Proxy Id List.",
			},

			"nation_country_inner_codes": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Nation Country Inner Codes.",
			},
		},
	}
}

func resourceTencentCloudGaapGlobalDomainDnsCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_global_domain_dns.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		request                    = gaap.NewCreateGlobalDomainDnsRequest()
		domainId                   string
		proxyIdList                []string
		nationCountryInnerCodeList []string
	)
	if v, ok := d.GetOk("domain_id"); ok {
		domainId = v.(string)
		request.DomainId = helper.String(domainId)
	}

	if v, ok := d.GetOk("proxy_id_list"); ok {
		proxyIdSet := v.(*schema.Set).List()
		for i := range proxyIdSet {
			proxyId := proxyIdSet[i].(string)
			proxyIdList = append(proxyIdList, proxyId)
		}
		request.ProxyIdList = helper.Strings(proxyIdList)

	}

	if v, ok := d.GetOk("nation_country_inner_codes"); ok {
		nationCountryInnerCodesSet := v.(*schema.Set).List()
		for i := range nationCountryInnerCodesSet {
			nationCountryInnerCodeList = append(nationCountryInnerCodeList, nationCountryInnerCodesSet[i].(string))
		}
		request.NationCountryInnerCodes = helper.Strings(nationCountryInnerCodeList)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseGaapClient().CreateGlobalDomainDns(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create gaap globalDomainDns failed, reason:%+v", logId, err)
		return err
	}

	service := GaapService{client: meta.(*TencentCloudClient).apiV3Conn}
	parmas := map[string]interface{}{
		"AccessList":      proxyIdList,
		"CountryAreaList": nationCountryInnerCodeList,
	}
	globalDomainDns, err := service.DescribeGaapGlobalDomainDnsById(ctx, domainId, parmas)
	if err != nil {
		return err
	}
	dnsRecordId := strconv.Itoa(int(*globalDomainDns.DnsRecordId))

	conf := BuildStateChangeConf([]string{}, []string{"1"}, 1*readRetryTimeout, time.Second, service.GlobalDomainDnsStateRefreshFunc(domainId, int(*globalDomainDns.DnsRecordId), []string{}))
	if _, e := conf.WaitForState(); e != nil {
		return e
	}
	d.SetId(domainId + FILED_SP + dnsRecordId)

	return resourceTencentCloudGaapGlobalDomainDnsRead(d, meta)
}

func resourceTencentCloudGaapGlobalDomainDnsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_global_domain_dns.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := GaapService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domainId := idSplit[0]
	dnsRecordIdInt, err := strconv.Atoi(idSplit[1])
	if err != nil {
		return err
	}
	params := make(map[string]interface{})
	params["DnsRecordId"] = dnsRecordIdInt
	globalDomainDns, err := service.DescribeGaapGlobalDomainDnsById(ctx, idSplit[0], params)
	if err != nil {
		return err
	}

	if globalDomainDns == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `GaapGlobalDomainDns` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if globalDomainDns.AccessList != nil {
		proxyIdList := make([]string, 0)
		for _, item := range globalDomainDns.AccessList {
			proxyIdList = append(proxyIdList, *item.ProxyId)
		}
		_ = d.Set("proxy_id_list", proxyIdList)
	}

	if globalDomainDns.CountryAreaList != nil {
		nationCountryInnerCodes := make([]string, 0)
		for _, item := range globalDomainDns.CountryAreaList {
			nationCountryInnerCodes = append(nationCountryInnerCodes, *item.NationCountryInnerCode)
		}
		_ = d.Set("nation_country_inner_codes", nationCountryInnerCodes)
	}
	_ = d.Set("domain_id", domainId)

	return nil
}

func resourceTencentCloudGaapGlobalDomainDnsUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_global_domain_dns.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := gaap.NewModifyGlobalDomainDnsRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domainId := idSplit[0]
	dnsRecordIdInt, err := strconv.Atoi(idSplit[1])
	if err != nil {
		return err
	}
	request.DomainId = helper.String(domainId)
	request.DnsRecordId = helper.IntUint64(dnsRecordIdInt)

	immutableArgs := []string{"domain_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if d.HasChange("proxy_id_list") {
		if v, ok := d.GetOk("proxy_id_list"); ok {
			proxyIdListSet := v.(*schema.Set).List()
			for i := range proxyIdListSet {
				proxyIdList := proxyIdListSet[i].(string)
				request.ProxyIdList = append(request.ProxyIdList, &proxyIdList)
			}
		}
	}

	if d.HasChange("nation_country_inner_codes") {
		if v, ok := d.GetOk("nation_country_inner_codes"); ok {
			nationCountryInnerCodesSet := v.(*schema.Set).List()
			for i := range nationCountryInnerCodesSet {
				nationCountryInnerCodes := nationCountryInnerCodesSet[i].(string)
				request.NationCountryInnerCodes = append(request.NationCountryInnerCodes, &nationCountryInnerCodes)
			}
		}
	}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseGaapClient().ModifyGlobalDomainDns(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update gaap globalDomainDns failed, reason:%+v", logId, err)
		return err
	}

	service := GaapService{client: meta.(*TencentCloudClient).apiV3Conn}
	conf := BuildStateChangeConf([]string{}, []string{"1"}, 1*readRetryTimeout, time.Second, service.GlobalDomainDnsStateRefreshFunc(domainId, dnsRecordIdInt, []string{}))
	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	return resourceTencentCloudGaapGlobalDomainDnsRead(d, meta)
}

func resourceTencentCloudGaapGlobalDomainDnsDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_gaap_global_domain_dns.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := GaapService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	dnsRecordIdInt, err := strconv.Atoi(idSplit[1])
	if err != nil {
		return err
	}
	if err := service.DeleteGaapGlobalDomainDnsById(ctx, dnsRecordIdInt); err != nil {
		return err
	}

	return nil
}
