package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var buildsDataSourceName = "twilio_serverless_builds"

func TestAccDataSourceTwilioServerlessBuilds_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.builds", buildsDataSourceName)
	uniqueName := acctest.RandString(10)
	version := "3.6.2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessBuilds_basic(uniqueName, version),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "builds.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "builds.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "builds.0.asset_versions.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "builds.0.function_versions.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "builds.0.dependencies.%", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "builds.0.dependencies.twilio", version),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "builds.0.status"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "builds.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "builds.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "builds.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessBuilds_basic(uniqueName string, twilioVersion string) string {
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
    "twilio" : "%s"
  }
}

data "twilio_serverless_builds" "builds" {
  service_sid = twilio_serverless_build.build.service_sid
}
`, uniqueName, twilioVersion)
}
