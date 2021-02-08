package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var deploymentsDataSourceName = "twilio_serverless_deployments"

func TestAccDataSourceTwilioServerlessDeployments_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.deployments", deploymentsDataSourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessDeployments_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "environment_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "deployments.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "deployments.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "deployments.0.build_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "deployments.0.date_created"),
					resource.TestCheckResourceAttr(stateDataSourceName, "deployments.0.date_updated", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "deployments.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessDeployments_basic(uniqueName string) string {
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

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }
  dependencies = {
    "twilio"                  = "3.6.3"
    "fs"                      = "0.0.1-security"
    "lodash"                  = "4.17.11"
    "util"                    = "0.11.0"
    "xmldom"                  = "0.1.27"
    "@twilio/runtime-handler" = "1.0.1"
  }

  polling {
    enabled = true
  }
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%s"
}

resource "twilio_serverless_deployment" "deployment" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  build_sid       = twilio_serverless_build.build.sid
}

data "twilio_serverless_deployments" "deployments" {
  service_sid     = twilio_serverless_deployment.deployment.service_sid
  environment_sid = twilio_serverless_deployment.deployment.environment_sid
}
`, uniqueName, uniqueName)
}
