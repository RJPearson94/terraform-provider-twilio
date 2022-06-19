package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const messagingConfigurationsDataSourceName = "twilio_verify_messaging_configurations"

func TestAccDataSourceTwilioVerifyMessagingConfigurations_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.messaging_configurations", messagingConfigurationsDataSourceName)
	friendlyName := acctest.RandString(10)
	countryCode := "GB"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyMessagingConfigurations_basic(friendlyName, countryCode),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging_configurations.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging_configurations.0.country_code", countryCode),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "messaging_configurations.0.messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "messaging_configurations.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "messaging_configurations.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "messaging_configurations.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyMessagingConfigurations_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyMessagingConfigurations_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyMessagingConfigurations_basic(friendlyName string, countryCode string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid           = twilio_verify_service.service.sid
  messaging_service_sid = twilio_messaging_service.service.sid
  country_code          = "%[2]s"
}

data "twilio_verify_messaging_configurations" "messaging_configurations" {
  service_sid = twilio_verify_messaging_configuration.messaging_configuration.service_sid
}
`, friendlyName, countryCode)
}

func testAccDataSourceTwilioVerifyMessagingConfigurations_invalidServiceSid() string {
	return `
data "twilio_verify_messaging_configurations" "messaging_configurations" {
  service_sid = "service_sid"
}
`
}
