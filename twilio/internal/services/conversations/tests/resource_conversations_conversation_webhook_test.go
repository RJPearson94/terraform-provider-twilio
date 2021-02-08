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

var conversationWebhookResourceName = "twilio_conversations_conversation_webhook"

func TestAccTwilioConversationsConversationWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", conversationWebhookResourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "http://localhost:3000/current"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversationWebhook_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", "onMessageAdded"),
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
				ImportStateIdFunc: testAccTwilioConversationsConversationWebhookImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsConversationWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", conversationWebhookResourceName)
	friendlyName := acctest.RandString(10)
	webhookURL := "http://localhost:3000/current"
	newWebhookURL := "http://localhost:3000/new"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversationWebhook_basic(friendlyName, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookURL),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", "onMessageAdded"),
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
				Config: testAccTwilioConversationsConversationWebhook_basic(friendlyName, newWebhookURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", newWebhookURL),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", "onMessageAdded"),
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

func TestAccTwilioConversationsConversationWebhook_invalidWebhookURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	webhookURL := "webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversationWebhook_basic(friendlyName, webhookURL),
				ExpectError: regexp.MustCompile(`(?s)expected "webhook_url" to have a host, got webhook`),
			},
		},
	})
}

func TestAccTwilioConversationsConversationWebhook_invalidMethod(t *testing.T) {
	friendlyName := acctest.RandString(10)
	method := "DELETE"
	webhookURL := "http://localhost:3000/current"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversationWebhook_withMethod(friendlyName, method, webhookURL),
				ExpectError: regexp.MustCompile(`(?s)expected method to be one of \[GET POST\], got DELETE`),
			},
		},
	})
}

func testAccCheckTwilioConversationsConversationWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != conversationWebhookResourceName {
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

func testAccCheckTwilioConversationsConversationWebhookExists(name string) resource.TestCheckFunc {
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

func testAccTwilioConversationsConversationWebhookImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Conversations/%s/Webhooks/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["conversation_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsConversationWebhook_basic(friendlyName string, webhookUrl string) string {
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
`, friendlyName, webhookUrl)
}

func testAccTwilioConversationsConversationWebhook_withMethod(friendlyName string, method string, webhookUrl string) string {
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
  method           = "%s"
  webhook_url      = "%s"
  filters          = ["onMessageAdded"]
}
`, friendlyName, method, webhookUrl)
}
