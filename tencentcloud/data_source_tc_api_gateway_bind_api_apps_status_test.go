package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudApiGatewayBindApiAppsStatusDataSource_basic -v
func TestAccTencentCloudApiGatewayBindApiAppsStatusDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayBindApiAppsStatusDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_api_gateway_bind_api_apps_status.example"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_api_gateway_bind_api_apps_status.example", "service_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_api_gateway_bind_api_apps_status.example", "api_ids.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_api_gateway_bind_api_apps_status.example", "filters.#"),
				),
			},
		},
	})
}

const testAccApiGatewayBindApiAppsStatusDataSource = `
data "tencentcloud_api_gateway_bind_api_apps_status" "example" {
  service_id = "service-nxz6yync"
  api_ids    = ["api-0cvmf4x4", "api-jvqlzolk"]
  filters {
    name   = "ApiAppId"
    values = ["app-krljp4wn"]
  }
}
`
