package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var functionResourceName = "twilio_serverless_function"

func TestAccTwilioServerlessFunction_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.function", functionResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessFunction_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/javascript"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "helloWorld.js"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-function"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", visibility),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:            stateResourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccTwilioServerlessFunctionImportStateIdFunc(stateResourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"content_file_name", "content_type", "source_hash"},
			},
		},
	})
}

func TestAccTwilioServerlessFunction_multipleFunctions(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.function", functionResourceName)
	stateResourceName2 := fmt.Sprintf("%s.function2", functionResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessFunction_multipleFunctions(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionExists(stateResourceName),
					testAccCheckTwilioServerlessFunctionExists(stateResourceName2),
				),
			},
		},
	})
}

func TestAccTwilioServerlessAssetFunction_invalidVisibility(t *testing.T) {
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_basic(uniqueName, friendlyName, visibility),
				ExpectError: regexp.MustCompile(`(?s)expected visibility to be one of \[public protected private\], got test`),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.function", functionResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)
	visibility := "private"
	newVisibility := "protected"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessFunction_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/javascript"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "helloWorld.js"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-function"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", visibility),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioServerlessFunction_basic(uniqueName, newFriendlyName, newVisibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/javascript"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "helloWorld.js"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-function"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", newVisibility),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioServerlessFunctionDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != functionResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Function(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving function information %s", err.Error())
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Function(rs.Primary.ID).Version(rs.Primary.Attributes["latest_version_sid"]).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving function version information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioServerlessFunctionExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Function(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving function information %s", err.Error())
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Function(rs.Primary.ID).Version(rs.Primary.Attributes["latest_version_sid"]).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving function version information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioServerlessFunctionImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Functions/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioServerlessFunction_basic(uniqueName string, friendlyName string, visibility string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "%s"
}
`, uniqueName, friendlyName, visibility)
}

func testAccTwilioServerlessFunction_multipleFunctions(uniqueName string, friendlyName string, visibility string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "%s"
}

resource "twilio_serverless_function" "function2" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s-2"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function-2"
  visibility        = "%s"
}
`, uniqueName, friendlyName, visibility, friendlyName, visibility)
}
