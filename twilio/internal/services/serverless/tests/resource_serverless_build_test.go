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

var buildResourceName = "twilio_serverless_build"

func TestAccTwilioServerlessBuild_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.build", buildResourceName)
	uniqueName := acctest.RandString(10)
	version := "3.6.2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessBuildDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessBuild_basic(uniqueName, version),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "asset_version_sids.#", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "function_version_sids.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "dependencies.%", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "dependencies.twilio", version),
					resource.TestCheckResourceAttrSet(stateResourceName, "status"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioServerlessBuildDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != buildResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Build(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving build information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioServerlessBuildExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Build(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving build information %s", err)
		}

		return nil
	}
}

func testAccTwilioServerlessBuild_basic(uniqueName string, twilioVersion string) string {
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
		"twilio" : "%s"
	}
}`, uniqueName, twilioVersion)
}
