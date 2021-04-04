package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var functionDataSourceName = "twilio_serverless_function"

func TestAccDataSourceTwilioServerlessFunction_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.function", functionDataSourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessFunction_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "content", "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="),
					resource.TestCheckResourceAttr(stateDataSourceName, "path", "/test-function"),
					resource.TestCheckResourceAttr(stateDataSourceName, "visibility", visibility),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessFunction_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessFunction_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessFunction_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessFunction_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^ZH\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessFunction_basic(uniqueName string, friendlyName string, visibility string) string {
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

data "twilio_serverless_function" "function" {
  service_sid = twilio_serverless_function.function.service_sid
  sid         = twilio_serverless_function.function.sid
}
`, uniqueName, friendlyName, visibility)
}

func testAccDataSourceTwilioServerlessFunction_invalidServiceSid() string {
	return `
data "twilio_serverless_function" "function" {
  service_sid = "service_sid"
  sid         = "ZHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioServerlessFunction_invalidSid() string {
	return `
data "twilio_serverless_function" "function" {
  service_sid = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
