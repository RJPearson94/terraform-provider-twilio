package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var channelWebhooksDataSourceName = "twilio_chat_channel_webhooks"

func TestAccDataSourceTwilioChatChannelWebhooks_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhooks", channelWebhooksDataSourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "http://localhost:3000/current"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatChannelWebhooks_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channel_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.type", "webhook"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.configuration.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.configuration.0.webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.configuration.0.filters.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.configuration.0.filters.0", "onMessageSent"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.configuration.0.method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.configuration.0.retry_count"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.configuration.0.flow_sid", ""),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "webhooks.0.configuration.0.triggers"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannelWebhooks_basic(friendlyName string, webhookUrl string) string {
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

data "twilio_chat_channel_webhooks" "webhooks" {
	service_sid = twilio_chat_channel_webhook.webhook.service_sid
	channel_sid = twilio_chat_channel_webhook.webhook.channel_sid
  }
`, friendlyName, friendlyName, webhookUrl)
}
