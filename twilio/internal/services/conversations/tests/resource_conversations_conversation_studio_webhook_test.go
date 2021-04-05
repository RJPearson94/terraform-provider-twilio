package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var conversationStudioWebhookResourceName = "twilio_conversations_conversation_studio_webhook"

func TestAccTwilioConversationsConversationStudioWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.studio_webhook", conversationStudioWebhookResourceName)
	friendlyName := acctest.RandString(10)
	flowSid := "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationStudioWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversationStudioWebhook_basic(friendlyName, flowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", flowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "conversation_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "replay_after"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioConversationsConversationStudioWebhookImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsConversationStudioWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.studio_webhook", conversationStudioWebhookResourceName)
	friendlyName := acctest.RandString(10)
	flowSid := "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	newFlowSid := "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationStudioWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversationStudioWebhook_basic(friendlyName, flowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", flowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "conversation_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "replay_after"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsConversationStudioWebhook_basic(friendlyName, newFlowSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationStudioWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "flow_sid", newFlowSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "conversation_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "replay_after"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsConversationStudioWebhook_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversationStudioWebhook_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsConversationStudioWebhook_invalidConversationSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversationStudioWebhook_invalidConversationSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of conversation_sid to match regular expression "\^CH\[0-9a-fA-F\]\{32\}\$", got conversation_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsConversationStudioWebhook_invalidFlowSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversationStudioWebhook_invalidFlowSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of flow_sid to match regular expression "\^FW\[0-9a-fA-F\]\{32\}\$", got flow_sid`),
			},
		},
	})
}

func testAccCheckTwilioConversationsConversationStudioWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != conversationStudioWebhookResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Conversation(rs.Primary.Attributes["conversation_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioConversationsConversationStudioWebhookExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Conversation(rs.Primary.Attributes["conversation_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsConversationStudioWebhookImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Conversations/%s/Webhooks/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["conversation_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsConversationStudioWebhook_basic(friendlyName string, flowSid string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}

resource "twilio_conversations_conversation_studio_webhook" "studio_webhook" {
  service_sid      = twilio_conversations_service.service.sid
  conversation_sid = twilio_conversations_conversation.conversation.sid
  flow_sid         = "%s"
}
`, friendlyName, flowSid)
}

func testAccTwilioConversationsConversationStudioWebhook_invalidServiceSid() string {
	return `
resource "twilio_conversations_conversation_studio_webhook" "studio_webhook" {
  service_sid      = "service_sid"
  conversation_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  flow_sid         = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioConversationsConversationStudioWebhook_invalidConversationSid() string {
	return `
resource "twilio_conversations_conversation_studio_webhook" "studio_webhook" {
  service_sid      = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  conversation_sid = "conversation_sid"
  flow_sid         = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioConversationsConversationStudioWebhook_invalidFlowSid() string {
	return `
resource "twilio_conversations_conversation_studio_webhook" "studio_webhook" {
  service_sid      = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  conversation_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  flow_sid         = "flow_sid"
}
`
}
