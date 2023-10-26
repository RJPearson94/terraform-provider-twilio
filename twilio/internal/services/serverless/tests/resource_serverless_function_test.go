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

func TestAccTwilioServerlessFunctionFunction_invalidVisibility(t *testing.T) {
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_basic(uniqueName, friendlyName, visibility),
				ExpectError: regexp.MustCompile(`(?s)expected visibility to be one of \["public" "protected" "private"\], got test`),
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

func TestAccTwilioServerlessFunction_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.function", functionResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(255)
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
				),
			},
			{
				Config: testAccTwilioServerlessFunction_basic(uniqueName, newFriendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_invalidFriendlyNameWith0Characters(t *testing.T) {
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_friendlyNameWithStubbedServiceSid(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 255\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_invalidFriendlyNameWith256Characters(t *testing.T) {
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_friendlyNameWithStubbedServiceSid(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 255\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_path(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.function", functionResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(1)
	visibility := "private"
	path := "/a"
	newPath := "/" + acctest.RandString(254)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessFunction_path(uniqueName, friendlyName, visibility, path),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "path", path),
				),
			},
			{
				Config: testAccTwilioServerlessFunction_path(uniqueName, friendlyName, visibility, newPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessFunctionExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "path", newPath),
				),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_invalidPathWith0Characters(t *testing.T) {
	path := ""
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_pathWithWithStubbedServiceSid(path),
				ExpectError: regexp.MustCompile(`(?s)expected length of path to be in the range \(1 - 255\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_invalidPathWith256Characters(t *testing.T) {
	path := "/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_pathWithWithStubbedServiceSid(path),
				ExpectError: regexp.MustCompile(`(?s)expected length of path to be in the range \(1 - 255\), got /aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_invalidContentType(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessFunction_blankContentType(),
				ExpectError: regexp.MustCompile(`(?s)expected \"content_type\" to not be an empty string, got `),
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
  unique_name   = "%s"
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
  unique_name   = "%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%[2]s"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "%[3]s"
}

resource "twilio_serverless_function" "function2" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%[2]s-2"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function-2"
  visibility        = "%[3]s"
}
`, uniqueName, friendlyName, visibility)
}

func testAccTwilioServerlessFunction_path(uniqueName string, friendlyName string, visibility string, path string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "%s"
  visibility        = "%s"
}
`, uniqueName, friendlyName, path, visibility)
}

func testAccTwilioServerlessFunction_invalidServiceSid() string {
	return `
resource "twilio_serverless_function" "function2" {
  service_sid       = "service_sid"
  friendly_name     = "invalid_service_sid"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/invalid-service-sid"
  visibility        = "private"
}
`
}

func testAccTwilioServerlessFunction_pathWithWithStubbedServiceSid(path string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_function" "function" {
  service_sid       = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name     = "invalid_path"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "%s"
  visibility        = "private"
}
`, path)
}

func testAccTwilioServerlessFunction_friendlyNameWithStubbedServiceSid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_function" "function" {
  service_sid       = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name     = "%s"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/friendly-name"
  visibility        = "private"
}
`, friendlyName)
}

func testAccTwilioServerlessFunction_blankContentType() string {
	return `
resource "twilio_serverless_function" "function" {
  service_sid       = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name     = "invalid_content_type"
  content           = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type      = ""
  content_file_name = "helloWorld.js"
  path              = "/invalid-content-type"
  visibility        = "private"
}
`
}
