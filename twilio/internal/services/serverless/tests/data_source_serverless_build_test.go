package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var buildDataSourceName = "twilio_serverless_build"

func TestAccDataSourceTwilioServerlessBuild_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.build", buildDataSourceName)
	uniqueName := acctest.RandString(10)
	version := "3.6.2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessBuild_basic(uniqueName, version),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "asset_versions.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "function_versions.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "dependencies.%", "6"),
					resource.TestCheckResourceAttr(stateDataSourceName, "dependencies.twilio", version),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "runtime"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessBuild_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessBuild_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessBuild_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessBuild_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^ZB\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessBuild_basic(uniqueName string, twilioVersion string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = <<EOF
exports.handler = function (context, event, callback) {
	callback(null, "Hello World");
};
EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "private"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }
  asset_version {
    sid = twilio_serverless_asset.asset.latest_version_sid
  }
  dependencies = {
    "twilio"                  = "%s"
    "fs"                      = "0.0.1-security"
    "lodash"                  = "4.17.11"
    "util"                    = "0.11.0"
    "xmldom"                  = "0.1.27"
    "@twilio/runtime-handler" = "1.0.1"
  }
}

data "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_build.build.service_sid
  sid         = twilio_serverless_build.build.sid
}
`, uniqueName, twilioVersion)
}

func testAccDataSourceTwilioServerlessBuild_invalidServiceSid() string {
	return `
data "twilio_serverless_build" "build" {
  service_sid = "service_sid"
  sid         = "ZBaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioServerlessBuild_invalidSid() string {
	return `
data "twilio_serverless_build" "build" {
  service_sid = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
