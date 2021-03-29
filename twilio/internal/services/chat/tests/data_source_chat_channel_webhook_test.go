package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var channelWebhookDataSourceName = "twilio_chat_channel_webhook"

func TestAccDataSourceTwilioChatChannelWebhook_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhook", channelWebhookDataSourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatChannelWebhook_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "type", "webhook"),
					resource.TestCheckResourceAttr(stateDataSourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "configuration.0.webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateDataSourceName, "configuration.0.filters.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "configuration.0.filters.0", "onMessageSent"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "configuration.0.method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "configuration.0.retry_count"),
					resource.TestCheckResourceAttr(stateDataSourceName, "configuration.0.flow_sid", ""),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "configuration.0.triggers"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannelWebhook_basic(friendlyName string, webhookUrl string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%s"
  type          = "private"
}

resource "twilio_chat_channel_webhook" "webhook" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  webhook_url = "%s"
  filters     = ["onMessageSent"]
}

data "twilio_chat_channel_webhook" "webhook" {
  service_sid = twilio_chat_channel_webhook.webhook.service_sid
  channel_sid = twilio_chat_channel_webhook.webhook.channel_sid
  sid         = twilio_chat_channel_webhook.webhook.sid
}
`, friendlyName, friendlyName, webhookUrl)
}
