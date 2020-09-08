package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var flowDataSourceName = "twilio_flex_flow"

func TestAccDataSourceTwilioFlexFlow_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.flow", flowDataSourceName)

	friendlyName := acctest.RandString(10)
	channelType := "web"
	integrationType := "external"
	integrationURL := "https://test.com/external"
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioFlexFlow_basic(testData, friendlyName, channelType, integrationType, integrationURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "channel_type", channelType),
					resource.TestCheckResourceAttr(stateDataSourceName, "chat_service_sid", testData.FlexChannelServiceSid),
					resource.TestCheckResourceAttr(stateDataSourceName, "integration_type", integrationType),
					resource.TestCheckResourceAttr(stateDataSourceName, "integration.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "integration.0.url", integrationURL),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "contact_identity", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "janitor_enabled"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "long_lived"),
					resource.TestCheckResourceAttr(stateDataSourceName, "integration.0.channel", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "integration.0.creation_on_message"),
					resource.TestCheckResourceAttr(stateDataSourceName, "integration.0.flow_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "integration.0.priority"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "integration.0.retry_count"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "integration.0.timeout"),
					resource.TestCheckResourceAttr(stateDataSourceName, "integration.0.workspace_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioFlexFlow_basic(testData *acceptance.TestData, friendlyName string, channelType string, integrationType string, integrationURL string) string {
	return fmt.Sprintf(`
resource "twilio_flex_flow" "flow" {
  friendly_name    = "%s"
  chat_service_sid = "%s"
  channel_type     = "%s"
  integration_type = "%s"
  integration {
    url = "%s"
  }
}

data "twilio_flex_flow" "flow" {
  sid = twilio_flex_flow.flow.sid
}
`, friendlyName, testData.FlexChannelServiceSid, channelType, integrationType, integrationURL)
}
