package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var flowResourceName = "twilio_flex_flow"

func TestAccTwilioFlexFlow_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.flow", flowResourceName)

	friendlyName := acctest.RandString(10)
	channelType := "web"
	integrationType := "external"
	integrationURL := "https://test.com/external"
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioFlexFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexFlow_basic(testData, friendlyName, channelType, integrationType, integrationURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "channel_type", channelType),
					resource.TestCheckResourceAttr(stateResourceName, "chat_service_sid", testData.FlexChannelServiceSid),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", integrationType),
					resource.TestCheckResourceAttr(stateResourceName, "integration.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.url", integrationURL),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "contact_identity", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "janitor_enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "long_lived"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.channel", ""),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.creation_on_message", ""),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.flow_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.priority"),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.timeout"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.workspace_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioFlexFlow_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.flow", flowResourceName)

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)
	channelType := "web"
	integrationType := "external"
	integrationURL := "https://test.com/external"
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioFlexFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexFlow_basic(testData, friendlyName, channelType, integrationType, integrationURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "channel_type", channelType),
					resource.TestCheckResourceAttr(stateResourceName, "chat_service_sid", testData.FlexChannelServiceSid),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", integrationType),
					resource.TestCheckResourceAttr(stateResourceName, "integration.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.url", integrationURL),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "contact_identity", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "janitor_enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "long_lived"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.channel", ""),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.creation_on_message", ""),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.flow_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.priority"),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.timeout"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.workspace_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioFlexFlow_basic(testData, newFriendlyName, channelType, integrationType, integrationURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "channel_type", channelType),
					resource.TestCheckResourceAttr(stateResourceName, "chat_service_sid", testData.FlexChannelServiceSid),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", integrationType),
					resource.TestCheckResourceAttr(stateResourceName, "integration.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.url", integrationURL),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "contact_identity", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "janitor_enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "long_lived"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.channel", ""),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.creation_on_message", ""),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.flow_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.priority"),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "integration.0.timeout"),
					resource.TestCheckResourceAttr(stateResourceName, "integration.0.workspace_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioFlexFlowDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

	for _, rs := range s.RootModule().Resources {
		if rs.Type != flowResourceName {
			continue
		}

		if _, err := client.FlexFlow(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving flow information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioFlexFlowExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.FlexFlow(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving flow information %s", err)
		}

		return nil
	}
}

func testAccTwilioFlexFlow_basic(testData *acceptance.TestData, friendlyName string, channelType string, integrationType string, integrationURL string) string {
	return fmt.Sprintf(`
resource "twilio_flex_flow" "flow" {
	friendly_name = "%s"
	chat_service_sid = "%s"
	channel_type = "%s"
	integration_type = "%s"
	integration {
		url = "%s"
	}
}`, friendlyName, testData.FlexChannelServiceSid, channelType, integrationType, integrationURL)
}
