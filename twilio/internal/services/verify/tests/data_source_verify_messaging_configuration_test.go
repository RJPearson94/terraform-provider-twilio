package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const messagingConfigurationDataSourceName = "twilio_verify_messaging_configuration"

func TestAccDataSourceTwilioVerifyMessagingConfiguration_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.messaging_configuration", messagingConfigurationDataSourceName)
	friendlyName := acctest.RandString(10)
	countryCode := "GB"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyMessagingConfiguration_basic(friendlyName, countryCode),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "country_code", countryCode),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyMessagingConfiguration_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyMessagingConfiguration_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyMessagingConfiguration_basic(friendlyName string, countryCode string) string {
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

data "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid  = twilio_verify_service.service.sid
  country_code = twilio_verify_messaging_configuration.messaging_configuration.country_code
}
`, friendlyName, countryCode)
}

func testAccDataSourceTwilioVerifyMessagingConfiguration_invalidServiceSid() string {
	return `
data "twilio_verify_messaging_configuration" "messaging_configuration" {
  country_code = "GB"
  service_sid  = "service_sid"
}
`
}
