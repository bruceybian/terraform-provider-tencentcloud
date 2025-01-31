package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamOidcConfigDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamOidcConfigDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.tencentcloud_cam_oidc_config.oidc_config")),
			},
		},
	})
}

const testAccCamOidcConfigDataSource = `

data "tencentcloud_cam_oidc_config" "oidc_config" {
  name = "cls-kzilgv5m"
}

output "identity_key" {
  value = data.tencentcloud_cam_oidc_config.oidc_config.identity_key
}

output "identity_url" {
  value = data.tencentcloud_cam_oidc_config.oidc_config.identity_url
}
`
