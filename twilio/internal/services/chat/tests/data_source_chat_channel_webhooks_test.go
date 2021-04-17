package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var channelWebhooksDataSourceName = "twilio_chat_channel_webhooks"

func TestAccDataSourceTwilioChatChannelWebhooks_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhooks", channelWebhooksDataSourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
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

func TestAccDataSourceTwilioChatChannelWebhooks_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannelWebhooks_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioChatChannelWebhooks_invalidChannelSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannelWebhooks_invalidChannelSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of channel_sid to match regular expression "\^CH\[0-9a-fA-F\]\{32\}\$", got channel_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannelWebhooks_basic(friendlyName string, webhookUrl string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%[1]s"
  type          = "private"
}

resource "twilio_chat_channel_webhook" "webhook" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  webhook_url = "%[2]s"
  filters     = ["onMessageSent"]
}

data "twilio_chat_channel_webhooks" "webhooks" {
  service_sid = twilio_chat_channel_webhook.webhook.service_sid
  channel_sid = twilio_chat_channel_webhook.webhook.channel_sid
}
`, friendlyName, webhookUrl)
}

func testAccDataSourceTwilioChatChannelWebhooks_invalidServiceSid() string {
	return `
data "twilio_chat_channel_webhooks" "webhooks" {
  service_sid = "service_sid"
  channel_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioChatChannelWebhooks_invalidChannelSid() string {
	return `
data "twilio_chat_channel_webhooks" "webhooks" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid = "channel_sid"
}
`
}
