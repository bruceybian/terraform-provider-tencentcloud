package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDlcRenewDataEngineResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDlcRenewDataEngine,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_dlc_renew_data_engine_operation.renew_data_engine", "id")),
			},
		},
	})
}

const testAccDlcRenewDataEngine = `

resource "tencentcloud_dlc_renew_data_engine_operation" "renew_data_engine" {
  data_engine_name = "test-iac"
  time_span = 1
  pay_mode = 1
  time_unit = "m"
  renew_flag = 1
}

`
