package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var assetVersionResourceName = "twilio_serverless_asset_version"

func TestAccTwilioServerlessAssetVersion_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.asset_version", assetVersionResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessAssetVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessAssetVersion_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetVersionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "e30="),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/json"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "test.json"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-asset"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", "private"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioServerlessAssetVersionDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != assetVersionResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Asset(rs.Primary.Attributes["asset_sid"]).Version(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving asset version information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioServerlessAssetVersionExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		// Asset versions cannot be destroyed however the supporting resources (service, asset) will be destroyed so this will verify the version is no longer present
		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Asset(rs.Primary.Attributes["asset_sid"]).Version(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving asset version information %s", err)
		}

		return nil
	}
}

func testAccTwilioServerlessAssetVersion_basic(uniqueName string, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
	unique_name   = "service-%s"
	friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
	service_sid   = twilio_serverless_service.service.sid
	friendly_name = "%s"
}

resource "twilio_serverless_asset_version" "asset_version" {
	service_sid       = twilio_serverless_service.service.sid
	asset_sid         = twilio_serverless_asset.asset.sid
	content           = "{}"
	content_type      = "application/json"
	content_file_name = "test.json"
	path              = "/test-asset"
	visibility        = "private"
}`, uniqueName, friendlyName)
}
