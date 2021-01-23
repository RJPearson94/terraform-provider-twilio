package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var conversationResourceName = "twilio_conversations_conversation"

func TestAccTwilioConversationsConversation_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.conversation", conversationResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversation_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "attributes"),
					resource.TestCheckResourceAttrSet(stateResourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "state"),
					resource.TestCheckResourceAttr(stateResourceName, "timers.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioConversationsConversationImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsConversation_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.conversation", conversationResourceName)
	friendlyName := acctest.RandString(10)

	closedTimer := "PT20M"
	inactiveTimer := "PT15M"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversation_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "attributes"),
					resource.TestCheckResourceAttrSet(stateResourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "state"),
					resource.TestCheckResourceAttr(stateResourceName, "timers.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_withTimers(friendlyName, closedTimer, inactiveTimer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "attributes"),
					resource.TestCheckResourceAttrSet(stateResourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "state"),
					resource.TestCheckResourceAttr(stateResourceName, "timers.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "timers.0.closed", closedTimer),
					resource.TestCheckResourceAttrSet(stateResourceName, "timers.0.date_closed"),
					resource.TestCheckResourceAttr(stateResourceName, "timers.0.inactive", inactiveTimer),
					resource.TestCheckResourceAttrSet(stateResourceName, "timers.0.date_inactive"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioConversationsConversationDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != conversationResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Conversation(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving conversation information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioConversationsConversationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Conversation(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving conversation information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsConversationImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Conversations/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsConversation_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}

func testAccTwilioConversationsConversation_withTimers(friendlyName string, closedTimer string, inactiveTimer string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
  timers {
    closed   = "%s"
    inactive = "%s"
  }
}
`, friendlyName, closedTimer, inactiveTimer)
}
