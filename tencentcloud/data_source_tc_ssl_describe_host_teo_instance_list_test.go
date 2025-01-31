package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudSslDescribeHostTeoInstanceListDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckCommon(t, ACCOUNT_TYPE_SSL)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSslDescribeHostTeoInstanceListDataSource,
				Check: resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.tencentcloud_ssl_describe_host_teo_instance_list.describe_host_teo_instance_list"),
					resource.TestCheckResourceAttr("data.tencentcloud_ssl_describe_host_teo_instance_list.describe_host_teo_instance_list", "certificate_id", "9D3qRt7r"),
					resource.TestCheckResourceAttr("data.tencentcloud_ssl_describe_host_teo_instance_list.describe_host_teo_instance_list", "resource_type", "teo"),
				),
			},
		},
	})
}

const testAccSslDescribeHostTeoInstanceListDataSource = `

data "tencentcloud_ssl_describe_host_teo_instance_list" "describe_host_teo_instance_list" {
  certificate_id = "9D3qRt7r"
  resource_type = "teo"
}
`
