package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var assetsDataSourceName = "twilio_serverless_assets"

func TestAccDataSourceTwilioServerlessAssets_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.assets", assetsDataSourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessAssets_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "assets.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "assets.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "assets.0.path", "/test-asset"),
					resource.TestCheckResourceAttr(stateDataSourceName, "assets.0.visibility", visibility),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assets.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assets.0.latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assets.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assets.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assets.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessAssets_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessAssets_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessAssets_basic(uniqueName string, friendlyName string, visibility string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "%s"
}

data "twilio_serverless_assets" "assets" {
  service_sid = twilio_serverless_asset.asset.service_sid
}
`, uniqueName, friendlyName, visibility)
}

func testAccDataSourceTwilioServerlessAssets_invalidServiceSid() string {
	return `
data "twilio_serverless_assets" "assets" {
  service_sid = "service_sid"
}
`
}
