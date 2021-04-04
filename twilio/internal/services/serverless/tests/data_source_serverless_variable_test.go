package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var variableDataSourceName = "twilio_serverless_variable"

func TestAccDataSourceTwilioServerlessVariable_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.variable", variableDataSourceName)
	uniqueName := acctest.RandString(10)
	key := "test-key"
	value := "test-value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessVariable_basic(uniqueName, key, value),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "key", key),
					resource.TestCheckResourceAttr(stateDataSourceName, "value", value),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "environment_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessVariable_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessVariable_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessVariable_invalidEnvironmentSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessVariable_invalidEnvironmentSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of environment_sid to match regular expression "\^ZE\[0-9a-fA-F\]\{32\}\$", got environment_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioServerlessVariable_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioServerlessVariable_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^ZV\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessVariable_basic(uniqueName string, key string, value string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%[1]s"
}

resource "twilio_serverless_variable" "variable" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  key             = "%[2]s"
  value           = "%[3]s"
}

data "twilio_serverless_variable" "variable" {
  service_sid     = twilio_serverless_variable.variable.service_sid
  environment_sid = twilio_serverless_variable.variable.environment_sid
  sid             = twilio_serverless_variable.variable.sid
}
`, uniqueName, key, value)
}

func testAccDataSourceTwilioServerlessVariable_invalidServiceSid() string {
	return `
data "twilio_serverless_variable" "variable" {
  service_sid     = "service_sid"
  environment_sid = "ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid             = "ZVaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioServerlessVariable_invalidEnvironmentSid() string {
	return `
data "twilio_serverless_variable" "variable" {
  service_sid     = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  environment_sid = "environment_sid"
  sid             = "ZVaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioServerlessVariable_invalidSid() string {
	return `
data "twilio_serverless_variable" "variable" {
  service_sid     = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  environment_sid = "ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid             = "sid"
}
`
}
