package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var functionVersionResourceName = "twilio_serverless_function_version"

func TestAccTwilioServerlessFunctionVersion_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.function_version", functionVersionResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessFunctionVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessFunctionVersion_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionVersionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/javascript"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "helloWorld.js"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-function"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", visibility),
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

func TestAccTwilioServerlessFunctionVersion_invalidVisibility(t *testing.T) {
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessFunctionVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunctionVersion_basic(uniqueName, friendlyName, visibility),
				ExpectError: regexp.MustCompile("config is invalid: expected visibility to be one of \\[public protected private\\], got test"),
			},
		},
	})
}

func testAccCheckTwilioServerlessFunctionVersionDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != functionVersionResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Function(rs.Primary.Attributes["function_sid"]).Version(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving function version information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioServerlessFunctionVersionExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		// Function versions cannot be destroyed however the supporting resources (service, function) will be destroyed so this will verify the version is no longer present
		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Function(rs.Primary.Attributes["function_sid"]).Version(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving function version information %s", err)
		}

		return nil
	}
}

func testAccTwilioServerlessFunctionVersion_basic(uniqueName string, friendlyName string, visibility string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
	unique_name   = "service-%s"
	friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
	service_sid   = twilio_serverless_service.service.sid
	friendly_name = "%s"
}

resource "twilio_serverless_function_version" "function_version" {
	service_sid  = twilio_serverless_service.service.sid
	function_sid = twilio_serverless_function.function.sid
	content 	 = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
	content_type = "application/javascript"
	content_file_name = "helloWorld.js"
	path         = "/test-function"
	visibility   = "%s"
  }`, uniqueName, friendlyName, visibility)
}
