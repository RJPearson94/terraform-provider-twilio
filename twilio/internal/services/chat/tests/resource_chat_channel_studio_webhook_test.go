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

var channelStudioWebhookResourceName = "twilio_chat_channel_studio_webhook"

func TestAccTwilioChatChannelStudioWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.studio_webhook", channelStudioWebhookResourceName)
	friendlyName := acctest.RandString(10)
	flowSid := "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioChatChannelStudioWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelStudioWebhook_basic(friendlyName, flowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", flowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioChatChannelStudioWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.studio_webhook", channelStudioWebhookResourceName)
	friendlyName := acctest.RandString(10)
	flowSid := "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	newFlowSid := "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioChatChannelStudioWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelStudioWebhook_basic(friendlyName, flowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", flowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioChatChannelStudioWebhook_basic(friendlyName, newFlowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", newFlowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "retry_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioChatChannelStudioWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

	for _, rs := range s.RootModule().Resources {
		if rs.Type != channelStudioWebhookResourceName {
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

func testAccCheckTwilioChatChannelStudioWebhookExists(name string) resource.TestCheckFunc {
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

func testAccTwilioChatChannelStudioWebhook_basic(friendlyName string, flowSid string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
	friendly_name = "%s"
}

resource "twilio_chat_channel" "channel" {
	service_sid   = twilio_chat_service.service.sid
	friendly_name = "%s"
	type		  = "private"
}

resource "twilio_chat_channel_studio_webhook" "studio_webhook" {
	service_sid   = twilio_chat_service.service.sid
	channel_sid   = twilio_chat_channel.channel.sid
	flow_sid      = "%s"
}
`, friendlyName, friendlyName, flowSid)
}
