package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var conversationWebhookDataSourceName = "twilio_conversations_conversation_webhook"

func TestAccDataSourceTwilioConversationsConversationWebhook_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.webhook", conversationWebhookDataSourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsConversationWebhook_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "sid"),
					resource.TestCheckResourceAttr(stateDataSource, "configuration.#", "1"),
					resource.TestCheckResourceAttr(stateDataSource, "configuration.0.webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateDataSource, "configuration.0.filters.#", "1"),
					resource.TestCheckResourceAttr(stateDataSource, "configuration.0.filters.0", "onMessageAdded"),
					resource.TestCheckResourceAttrSet(stateDataSource, "configuration.0.replay_after"),
					resource.TestCheckResourceAttrSet(stateDataSource, "configuration.0.method"),
					resource.TestCheckResourceAttrSet(stateDataSource, "configuration.0.flow_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "configuration.0.triggers.#", "0"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsConversationWebhook_basic(friendlyName string, webhookUrl string) string {
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

data "twilio_conversations_conversation_webhook" "webhook" {
  service_sid      = twilio_conversations_service.service.sid
  conversation_sid = twilio_conversations_conversation.conversation.sid
  sid              = twilio_conversations_conversation_webhook.webhook.sid
}
`, friendlyName, webhookUrl)
}
