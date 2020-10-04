package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var variablesDataSourceName = "twilio_serverless_variables"

func TestAccDataSourceTwilioServerlessVariables_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.variables", variablesDataSourceName)
	uniqueName := acctest.RandString(10)
	key := "test-key"
	value := "test-value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessVariables_basic(uniqueName, key, value),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "environment_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "variables.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "variables.0.key", key),
					resource.TestCheckResourceAttr(stateDataSourceName, "variables.0.value", value),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "variables.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "variables.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "variables.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "variables.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessVariables_basic(uniqueName string, key string, value string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%s"
}

resource "twilio_serverless_variable" "variable" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  key             = "%s"
  value           = "%s"
}

data "twilio_serverless_variables" "variables" {
  service_sid     = twilio_serverless_variable.variable.service_sid
  environment_sid = twilio_serverless_variable.variable.environment_sid
}
`, uniqueName, uniqueName, key, value)
}
