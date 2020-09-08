package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var environmentsDataSourceName = "twilio_serverless_environments"

func TestAccDataSourceTwilioServerlessEnvironments_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.environments", environmentsDataSourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessEnvironments_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "environments.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "environments.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "environments.0.unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "environments.0.domain_suffix", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "environments.0.build_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "environments.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "environments.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "environments.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessEnvironments_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%s"
}

data "twilio_serverless_environments" "environments" {
  service_sid = twilio_serverless_environment.environment.service_sid
}
`, uniqueName, uniqueName)
}
