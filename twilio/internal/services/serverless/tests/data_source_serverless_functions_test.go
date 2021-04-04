package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var functionsDataSourceName = "twilio_serverless_functions"

func TestAccDataSourceTwilioServerlessFunctions_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.functions", functionsDataSourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessFunctions_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "functions.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "functions.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "functions.0.content", "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="),
					resource.TestCheckResourceAttr(stateDataSourceName, "functions.0.path", "/test-function"),
					resource.TestCheckResourceAttr(stateDataSourceName, "functions.0.visibility", visibility),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "functions.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "functions.0.latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "functions.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "functions.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "functions.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessFunctions_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessFunctions_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessFunctions_basic(uniqueName string, friendlyName string, visibility string) string {
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

data "twilio_serverless_functions" "functions" {
  service_sid = twilio_serverless_function.function.service_sid
}
`, uniqueName, friendlyName, visibility)
}

func testAccDataSourceTwilioServerlessFunctions_invalidServiceSid() string {
	return `
data "twilio_serverless_functions" "functions" {
  service_sid = "service_sid"
}
`
}
