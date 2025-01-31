package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTsfGroupsDataSource_basic -v
func TestAccTencentCloudTsfGroupsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfGroupsDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_tsf_groups.groups"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.total_count"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.alias"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.application_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.application_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.application_type"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.cluster_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.cluster_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.create_time"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.group_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.group_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.group_resource_type"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.microservice_type"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.namespace_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.namespace_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.startup_parameters"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.update_time"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_tsf_groups.groups", "result.0.content.0.updated_time"),
				),
			},
		},
	})
}

const testAccTsfGroupsDataSourceVar = `
variable "application_id" {
	default = "` + defaultTsfApplicationId + `"
}
variable "cluster_id" {
	default = "` + defaultTsfClustId + `"
}
variable "namespace_id" {
	default = "` + defaultNamespaceId + `"
}
variable "group_id" {
	default = "` + defaultTsfGroupId + `"
}

`

const testAccTsfGroupsDataSource = testAccTsfGroupsDataSourceVar + `

data "tencentcloud_tsf_groups" "groups" {
	search_word = "keep"
	application_id = var.application_id
	order_by = "createTime"
	order_type = 0
	namespace_id = var.namespace_id
	cluster_id = var.cluster_id
	group_resource_type_list = ["DEF"]
	status = "Running"
	group_id_list = [var.group_id]
}

`
