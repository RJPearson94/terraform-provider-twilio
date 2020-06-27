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

var deploymentResourceName = "twilio_serverless_deployment"

func TestAccTwilioServerlessDeployment_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.deployment", deploymentResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessDeployment_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessDeploymentExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "environment_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioServerlessDeploymentDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != deploymentResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.Attributes["environment_sid"]).Deployment(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving deployment information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioServerlessDeploymentExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.Attributes["environment_sid"]).Deployment(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving deployment information %s", err)
		}

		return nil
	}
}

func testAccTwilioServerlessDeployment_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
	unique_name   = "service-%s"
	friendly_name = "test"
}
  
resource "twilio_serverless_function" "function" {
	service_sid   = twilio_serverless_service.service.sid
	friendly_name = "test"
}

resource "twilio_serverless_function_version" "function_version" {
	service_sid  = twilio_serverless_service.service.sid
	function_sid = twilio_serverless_function.function.sid
	content      = <<EOF
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
	service_sid           = twilio_serverless_service.service.sid
	function_version_sids = [twilio_serverless_function_version.function_version.sid]
	dependencies = {
		"twilio" : "3.6.3"
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
}`, uniqueName, uniqueName)
}
