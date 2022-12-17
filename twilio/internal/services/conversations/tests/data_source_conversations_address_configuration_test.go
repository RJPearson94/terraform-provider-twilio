package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var addressConfigurationDataSourceName = "twilio_conversations_address_configuration"

func TestAccDataSourceTwilioConversationsAddressConfiguration_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.address_configuration", addressConfigurationDataSourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsAddressConfiguration_basic(address, addressType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "address", address),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.#", "1"),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.0.service_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSource, "auto_creation.0.enabled"),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.0.flow_sid", ""),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.0.retry_count", ""),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.0.integration_type", "default"),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.0.webhook_filters.#", "0"),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.0.webhook_method", ""),
					resource.TestCheckResourceAttr(stateDataSource, "auto_creation.0.webhook_url", ""),
					resource.TestCheckResourceAttr(stateDataSource, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateDataSource, "type", addressType),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsAddressConfiguration_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsAddressConfiguration_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^IG\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsAddressConfiguration_basic(address string, addressType string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_address_configuration_default" "address_configuration_default" {
  address = "%[1]s"
  type    = "%[2]s"
}

data "twilio_conversations_address_configuration" "address_configuration" {
  sid = twilio_conversations_address_configuration_default.address_configuration_default.sid
}
`, address, addressType)
}

func testAccDataSourceTwilioConversationsAddressConfiguration_invalidSid() string {
	return `
data "twilio_conversations_address_configuration" "address_configuration" {
  sid = "sid"
}
`
}
