package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDnspodDomainAliasResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspodDomainAlias,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloud_dnspod_domain_alias.domain_alias", "domain", "iac-tf.cloud"),
					resource.TestCheckResourceAttr("tencentcloud_dnspod_domain_alias.domain_alias", "domain_alias", "iac-tf-alias.cloud"),
				),
			},
			// {
			// 	Config: testAccDnspodDomainAliasUp,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr("tencentcloud_dnspod_domain_alias.domain_alias", "domain", "iac-tf.cloud"),
			// 		resource.TestCheckResourceAttr("tencentcloud_dnspod_domain_alias.domain_alias", "domain_alias", "iac-tf-alias2.cloud"),
			// 	),
			// },
			{
				ResourceName:      "tencentcloud_dnspod_domain_alias.domain_alias",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccDnspodDomainAlias = `

resource "tencentcloud_dnspod_domain_alias" "domain_alias" {
  domain = "iac-tf.cloud"
  domain_alias = "iac-tf-alias.cloud"
}

`

// const testAccDnspodDomainAliasUp = `

// resource "tencentcloud_dnspod_domain_alias" "domain_alias" {
//   domain_alias = "iac-tf-alias2.cloud"
//   domain = "iac-tf.cloud"
// }

// `
