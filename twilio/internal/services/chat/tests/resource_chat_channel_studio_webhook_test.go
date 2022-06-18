package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var channelStudioWebhookResourceName = "twilio_chat_channel_studio_webhook"

func TestAccTwilioChatChannelStudioWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.studio_webhook", channelStudioWebhookResourceName)
	friendlyName := acctest.RandString(10)
	flowSid := "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatChannelStudioWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelStudioWebhook_basic(friendlyName, flowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", flowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "retry_count", "3"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioChatChannelStudioWebhookImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
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
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatChannelStudioWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelStudioWebhook_basic(friendlyName, flowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", flowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "retry_count", "3"),
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
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "retry_count", "3"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioChatChannelStudioWebhook_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatChannelStudioWebhook_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioChatChannelStudioWebhook_invalidChannelSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatChannelStudioWebhook_invalidChannelSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of channel_sid to match regular expression "\^CH\[0-9a-fA-F\]\{32\}\$", got channel_sid`),
			},
		},
	})
}

func TestAccTwilioChatChannelStudioWebhook_invalidFlowSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatChannelStudioWebhook_invalidFlowSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of flow_sid to match regular expression "\^FW\[0-9a-fA-F\]\{32\}\$", got flow_sid`),
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
			return fmt.Errorf("Error occurred when retrieving channel webhook information %s", err.Error())
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
			return fmt.Errorf("Error occurred when retrieving channel webhook information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioChatChannelStudioWebhookImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Channels/%s/Webhooks/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["channel_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioChatChannelStudioWebhook_basic(friendlyName string, flowSid string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%[1]s"
  type          = "private"
}

resource "twilio_chat_channel_studio_webhook" "studio_webhook" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  flow_sid    = "%[2]s"
}
`, friendlyName, flowSid)
}

func testAccTwilioChatChannelStudioWebhook_invalidServiceSid() string {
	return `
resource "twilio_chat_channel_studio_webhook" "studio_webhook" {
  service_sid = "service_sid"
  channel_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  flow_sid    = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioChatChannelStudioWebhook_invalidChannelSid() string {
	return `
resource "twilio_chat_channel_studio_webhook" "studio_webhook" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid = "channel_sid"
  flow_sid    = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioChatChannelStudioWebhook_invalidFlowSid() string {
	return `
resource "twilio_chat_channel_studio_webhook" "studio_webhook" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  flow_sid    = "flow_sid"
}
`
}
