package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCssTimeShiftRecordDetailDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCssTimeShiftRecordDetailDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_css_time_shift_record_detail.time_shift_record_detail"),
				),
			},
		},
	})
}

const testAccCssTimeShiftRecordDetailDataSource = `

data "tencentcloud_css_time_shift_record_detail" "time_shift_record_detail" {
  domain        = "177154.push.tlivecloud.com"
  app_name      = "qqq"
  stream_name   = "live"
  start_time    = 1698768000
  end_time      = 1698820641
  domain_group  = "tf-test"
  trans_code_id = 0
}

`
