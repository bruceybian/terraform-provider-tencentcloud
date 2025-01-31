package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudSsmSshKeyPairValueDataSource_basic -v
func TestAccTencentCloudSsmSshKeyPairValueDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSsmSshKeyPairValueDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_ssm_ssh_key_pair_value.example"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_ssm_ssh_key_pair_value.example", "secret_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_ssm_ssh_key_pair_value.example", "ssh_key_id"),
				),
			},
		},
	})
}

const testAccSsmSshKeyPairValueDataSource = `
data "tencentcloud_ssm_ssh_key_pair_value" "example" {
  secret_name = "keep_terraform"
  ssh_key_id  = "skey-2ae2snwd"
}
`
