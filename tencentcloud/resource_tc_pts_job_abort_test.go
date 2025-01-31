package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixPtsJobAbortResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPtsJobAbort,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloud_pts_job_abort.job_abort", "id"),
				),
			},
		},
	})
}

const testAccPtsJobAbort = `

resource "tencentcloud_pts_job_abort" "job_abort" {
  job_id       = "job-my644ozi"
  project_id   = "project-45vw7v82"
  scenario_id  = "scenario-22q19f3k"
}

`
