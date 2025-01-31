package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCssXp2pDetailInfoListDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCssXp2pDetailInfoListDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_css_xp2p_detail_info_list.xp2p_detail_info_list"),
				),
			},
		},
	})
}

const testAccCssXp2pDetailInfoListDataSource = `

data "tencentcloud_css_xp2p_detail_info_list" "xp2p_detail_info_list" {
  query_time   = "2023-11-01T14:55:01+08:00"
  type         = ["live"]
}

`
