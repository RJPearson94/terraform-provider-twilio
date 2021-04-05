package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var conversationWebhooksDataSourceName = "twilio_conversations_conversation_webhooks"

func TestAccDataSourceTwilioConversationsConversationWebhooks_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.webhooks", conversationWebhooksDataSourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsConversationWebhooks_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversation_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "webhooks.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSource, "webhooks.0.sid"),
					resource.TestCheckResourceAttr(stateDataSource, "webhooks.0.target", "webhook"),
					resource.TestCheckResourceAttr(stateDataSource, "webhooks.0.configuration.#", "1"),
					resource.TestCheckResourceAttr(stateDataSource, "webhooks.0.configuration.0.webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateDataSource, "webhooks.0.configuration.0.filters.#", "1"),
					resource.TestCheckResourceAttr(stateDataSource, "webhooks.0.configuration.0.filters.0", "onMessageAdded"),
					resource.TestCheckResourceAttrSet(stateDataSource, "webhooks.0.configuration.0.replay_after"),
					resource.TestCheckResourceAttrSet(stateDataSource, "webhooks.0.configuration.0.method"),
					resource.TestCheckResourceAttrSet(stateDataSource, "webhooks.0.configuration.0.flow_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "webhooks.0.configuration.0.triggers.#", "0"),
					resource.TestCheckResourceAttrSet(stateDataSource, "webhooks.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "webhooks.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "webhooks.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsConversationWebhooks_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsConversationWebhooks_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsConversationWebhooks_invalidConversationSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsConversationWebhooks_invalidConversationSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of conversation_sid to match regular expression "\^CH\[0-9a-fA-F\]\{32\}\$", got conversation_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsConversationWebhooks_basic(friendlyName string, webhooksUrl string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}

resource "twilio_conversations_conversation_webhook" "webhook" {
  service_sid      = twilio_conversations_service.service.sid
  conversation_sid = twilio_conversations_conversation.conversation.sid
  webhook_url      = "%s"
  filters          = ["onMessageAdded"]
}

data "twilio_conversations_conversation_webhooks" "webhooks" {
  service_sid      = twilio_conversations_conversation_webhook.webhook.service_sid
  conversation_sid = twilio_conversations_conversation_webhook.webhook.conversation_sid
}
`, friendlyName, webhooksUrl)
}

func testAccDataSourceTwilioConversationsConversationWebhooks_invalidServiceSid() string {
	return `
data "twilio_conversations_conversation_webhooks" "webhooks" {
  service_sid      = "service_sid"
  conversation_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioConversationsConversationWebhooks_invalidConversationSid() string {
	return `
data "twilio_conversations_conversation_webhooks" "webhooks" {
  service_sid      = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  conversation_sid = "conversation_sid"
}
`
}
