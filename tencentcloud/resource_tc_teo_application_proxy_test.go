package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=tencentcloud_teo_zone
	resource.AddTestSweepers("tencentcloud_teo_application_proxy", &resource.Sweeper{
		Name: "tencentcloud_teo_application_proxy",
		F:    testSweepApplicationProxy,
	})
}

func testSweepApplicationProxy(region string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cli, _ := sharedClientForRegion(region)
	client := cli.(*TencentCloudClient).apiV3Conn
	service := TeoService{client}

	for {
		proxy, err := service.DescribeTeoApplicationProxy(ctx, "", "")
		if err != nil {
			return err
		}

		if proxy == nil {
			return nil
		}

		err = service.DeleteTeoApplicationProxyById(ctx, *proxy.ZoneId, *proxy.ProxyId)
		if err != nil {
			return err
		}
	}
}

// go test -i; go test -test.run TestAccTencentCloudTeoApplicationProxy_basic -v
func TestAccTencentCloudTeoApplicationProxy_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PRIVATE) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApplicationProxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTeoApplicationProxy,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationProxyExists("tencentcloud_teo_application_proxy.basic"),
					resource.TestCheckResourceAttrSet("tencentcloud_teo_application_proxy.basic", "zone_id"),
					resource.TestCheckResourceAttr("tencentcloud_teo_application_proxy.basic", "accelerate_type", "1"),
					resource.TestCheckResourceAttr("tencentcloud_teo_application_proxy.basic", "security_type", "1"),
					resource.TestCheckResourceAttr("tencentcloud_teo_application_proxy.basic", "plat_type", "domain"),
					resource.TestCheckResourceAttr("tencentcloud_teo_application_proxy.basic", "proxy_name", "test-instance"),
					resource.TestCheckResourceAttr("tencentcloud_teo_application_proxy.basic", "proxy_type", "instance"),
					resource.TestCheckResourceAttr("tencentcloud_teo_application_proxy.basic", "session_persist_time", "2400"),
				),
			},
			{
				ResourceName:      "tencentcloud_teo_application_proxy.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckApplicationProxyDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TeoService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloud_teo_application_proxy" {
			continue
		}
		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		zoneId := idSplit[0]
		proxyId := idSplit[1]

		agents, err := service.DescribeTeoApplicationProxy(ctx, zoneId, proxyId)
		if agents != nil {
			return fmt.Errorf("zone ApplicationProxy %s still exists", rs.Primary.ID)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func testAccCheckApplicationProxyExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		zoneId := idSplit[0]
		proxyId := idSplit[1]

		service := TeoService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		agents, err := service.DescribeTeoApplicationProxy(ctx, zoneId, proxyId)
		if agents == nil {
			return fmt.Errorf("zone ApplicationProxy %s is not found", rs.Primary.ID)
		}
		if err != nil {
			return err
		}

		return nil
	}
}

const testAccTeoApplicationProxy = testAccTeoZone + `

resource "tencentcloud_teo_application_proxy" "basic" {
  zone_id = tencentcloud_teo_zone.basic.id

  accelerate_type      = 1
  security_type        = 1
  plat_type            = "domain"
  proxy_name           = "test-instance"
  proxy_type           = "instance"
  session_persist_time = 2400
}

`
