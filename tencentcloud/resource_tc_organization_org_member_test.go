package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudOrganizationOrgMemberResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationOrgMember,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_organization_org_member.org_member", "id"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "name", "terraform_test"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "node_id", "59849"),
					resource.TestCheckResourceAttrSet("tencentcloud_organization_org_member.org_member", "permission_ids.#"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "policy_type", "Financial"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "remark", "for terraform test"),
				),
			},
			{
				Config: testAccOrganizationOrgMemberUpdate,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_organization_org_member.org_member", "id"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "name", "terraform_test_1"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "node_id", "59849"),
					resource.TestCheckResourceAttrSet("tencentcloud_organization_org_member.org_member", "permission_ids.#"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "policy_type", "Financial"),
					resource.TestCheckResourceAttr("tencentcloud_organization_org_member.org_member", "remark", "for terraform test"),
				),
			},
			{
				ResourceName:      "tencentcloud_organization_org_member.org_member",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccOrganizationOrgMember = `

resource "tencentcloud_organization_org_member" "org_member" {
  name            = "terraform_test"
  node_id         = 59849
  permission_ids  = [
    1,
    2,
    3,
    4,
  ]
  policy_type     = "Financial"
  remark          = "for terraform test"
}

`
const testAccOrganizationOrgMemberUpdate = `

resource "tencentcloud_organization_org_member" "org_member" {
  name            = "terraform_test_1"
  node_id         = 59849
  permission_ids  = [
    1,
    2,
    3,
    4,
  ]
  policy_type     = "Financial"
  remark          = "for terraform test"
}

`
