/*
Provides a resource to create a PrivateDns zone_vpc_attachment

~> **NOTE:**  If you need to bind account A to account B's VPC resources, you need to first grant role authorization to account A.

Example Usage

Append VPC associated with private dns zone

```hcl
resource "tencentcloud_private_dns_zone" "example" {
  domain = "domain.com"
  remark = "remark."

  dns_forward_status   = "DISABLED"
  cname_speedup_status = "ENABLED"

  tags = {
    createdBy : "terraform"
  }
}

resource "tencentcloud_vpc" "vpc" {
  name       = "vpc-example"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloud_private_dns_zone_vpc_attachment" "example" {
  zone_id = tencentcloud_private_dns_zone.example.id

  vpc_set {
    uniq_vpc_id = tencentcloud_vpc.vpc.id
    region      = "ap-guangzhou"
  }
}
```

Add VPC information for associated accounts in the private dns zone

```hcl
resource "tencentcloud_private_dns_zone_vpc_attachment" "example" {
  zone_id = tencentcloud_private_dns_zone.example.id

  account_vpc_set {
    uniq_vpc_id = "vpc-82znjzn3"
    region      = "ap-guangzhou"
    uin         = "100017155920"
  }
}
```

Import

PrivateDns zone_vpc_attachment can be imported using the id, e.g.

```
terraform import tencentcloud_private_dns_zone_vpc_attachment.example zone-6t11lof0#vpc-jdx11z0t
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
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudPrivateDnsZoneVpcAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudPrivateDnsZoneVpcAttachmentCreate,
		Read:   resourceTencentCloudPrivateDnsZoneVpcAttachmentRead,
		Delete: resourceTencentCloudPrivateDnsZoneVpcAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "PrivateZone ID.",
			},
			"vpc_set": {
				Optional:     true,
				ForceNew:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"account_vpc_set"},
				Type:         schema.TypeList,
				Description:  "New add vpc info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uniq_vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Uniq Vpc Id.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Vpc region.",
						},
					},
				},
			},
			"account_vpc_set": {
				Optional:     true,
				ForceNew:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"vpc_set"},
				Type:         schema.TypeList,
				Description:  "New add account vpc info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uniq_vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Uniq Vpc Id.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Vpc region.",
						},
						"uin": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Vpc owner uin. To grant role authorization to this account.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudPrivateDnsZoneVpcAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_private_dns_zone_vpc_attachment.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId     = getLogId(contextNil)
		request   = privatedns.NewAddSpecifyPrivateZoneVpcRequest()
		zoneId    string
		uniqVpcId string
	)

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
		zoneId = v.(string)
	}

	if v, ok := d.GetOk("vpc_set"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			vpcInfo := new(privatedns.VpcInfo)
			if v, ok := dMap["uniq_vpc_id"]; ok {
				vpcInfo.UniqVpcId = helper.String(v.(string))
				uniqVpcId = v.(string)
			}

			if v, ok := dMap["region"]; ok {
				vpcInfo.Region = helper.String(v.(string))
			}

			request.VpcSet = append(request.VpcSet, vpcInfo)
		}
	}

	if v, ok := d.GetOk("account_vpc_set"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			accountVpcInfo := new(privatedns.AccountVpcInfo)
			if v, ok := dMap["uniq_vpc_id"]; ok {
				accountVpcInfo.UniqVpcId = helper.String(v.(string))
				uniqVpcId = v.(string)
			}

			if v, ok := dMap["region"]; ok {
				accountVpcInfo.Region = helper.String(v.(string))
			}

			if v, ok := dMap["uin"]; ok {
				accountVpcInfo.Uin = helper.String(v.(string))
			}

			request.AccountVpcSet = append(request.AccountVpcSet, accountVpcInfo)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UsePrivateDnsClient().AddSpecifyPrivateZoneVpc(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create PrivateDns ZoneVpcAttachment failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{zoneId, uniqVpcId}, FILED_SP))

	return resourceTencentCloudPrivateDnsZoneVpcAttachmentRead(d, meta)
}

func resourceTencentCloudPrivateDnsZoneVpcAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_private_dns_zone_vpc_attachment.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = PrivateDnsService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	zoneId := idSplit[0]
	uniqVpcId := idSplit[1]

	ZoneVpcAttachment, err := service.DescribePrivateDnsZoneVpcAttachmentById(ctx, zoneId)
	if err != nil {
		return err
	}

	if ZoneVpcAttachment == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `PrivateDnsZoneVpcAttachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if ZoneVpcAttachment.ZoneId != nil {
		_ = d.Set("zone_id", ZoneVpcAttachment.ZoneId)
	}

	if ZoneVpcAttachment.VpcSet != nil {
		vpcSetList := []interface{}{}
		for _, vpcSet := range ZoneVpcAttachment.VpcSet {
			vpcSetMap := map[string]interface{}{}

			if *vpcSet.UniqVpcId == uniqVpcId {
				vpcSetMap["uniq_vpc_id"] = *vpcSet.UniqVpcId
				vpcSetMap["region"] = *vpcSet.Region
				vpcSetList = append(vpcSetList, vpcSetMap)
				break
			}
		}

		_ = d.Set("vpc_set", vpcSetList)
	}

	if ZoneVpcAttachment.AccountVpcSet != nil {
		accountVpcSetList := []interface{}{}
		for _, accountVpcSet := range ZoneVpcAttachment.AccountVpcSet {
			accountVpcSetMap := map[string]interface{}{}

			if *accountVpcSet.UniqVpcId == uniqVpcId {
				accountVpcSetMap["uniq_vpc_id"] = *accountVpcSet.UniqVpcId
				accountVpcSetMap["region"] = *accountVpcSet.Region
				accountVpcSetMap["uin"] = *accountVpcSet.Uin
				accountVpcSetList = append(accountVpcSetList, accountVpcSetMap)
				break
			}
		}

		_ = d.Set("account_vpc_set", accountVpcSetList)
	}

	return nil
}

func resourceTencentCloudPrivateDnsZoneVpcAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_private_dns_zone_vpc_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = PrivateDnsService{client: meta.(*TencentCloudClient).apiV3Conn}
		region  string
		uin     string
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	zoneId := idSplit[0]
	uniqVpcId := idSplit[1]

	// get vpc detail
	ZoneVpcAttachment, err := service.DescribePrivateDnsZoneVpcAttachmentById(ctx, zoneId)
	if err != nil {
		return err
	}

	if ZoneVpcAttachment == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `PrivateDnsZoneVpcAttachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if ZoneVpcAttachment.VpcSet != nil {
		for _, vpcSet := range ZoneVpcAttachment.VpcSet {
			if *vpcSet.UniqVpcId == uniqVpcId {
				region = *vpcSet.Region
				break
			}
		}
	}

	if ZoneVpcAttachment.AccountVpcSet != nil {
		for _, accountVpcSet := range ZoneVpcAttachment.AccountVpcSet {
			if *accountVpcSet.UniqVpcId == uniqVpcId {
				region = *accountVpcSet.Region
				uin = *accountVpcSet.Uin
				break
			}
		}
	}

	if err = service.DeletePrivateDnsZoneVpcAttachmentById(ctx, zoneId, uniqVpcId, region, uin); err != nil {
		return err
	}

	return nil
}
