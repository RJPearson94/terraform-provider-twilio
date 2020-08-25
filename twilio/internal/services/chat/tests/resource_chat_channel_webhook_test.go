package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var channelWebhookResourceName = "twilio_chat_channel_webhook"

func TestAccTwilioChatChannelWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", channelWebhookResourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "http://localhost:3000/current"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioChatChannelWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelWebhook_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", "onMessageSent"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioChatChannelWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", channelWebhookResourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "http://localhost:3000/current"
	newWebhookURL := "http://localhost:3000/new"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioChatChannelWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelWebhook_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", "onMessageSent"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioChatChannelWebhook_basic(friendlyName, newWebhookURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", newWebhookURL),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", "onMessageSent"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioChatChannelWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

	for _, rs := range s.RootModule().Resources {
		if rs.Type != channelWebhookResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Channel(rs.Primary.Attributes["channel_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
				// currently programmable chat returns a 403 if the service instance does not exist
				if twilioError.Status == 403 && twilioError.Message == "Service instance not found" {
					return nil
				}
			}
			return fmt.Errorf("Error occurred when retrieving channel webhook information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioChatChannelWebhookExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Channel(rs.Primary.Attributes["channel_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving channel webhook information %s", err)
		}

		return nil
	}
}

func testAccTwilioChatChannelWebhook_basic(friendlyName string, webhookUrl string) string {
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
`, friendlyName, friendlyName, webhookUrl)
}
