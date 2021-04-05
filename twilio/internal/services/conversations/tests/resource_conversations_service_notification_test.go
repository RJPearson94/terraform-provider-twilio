package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var serviceNotificationResourceName = "twilio_conversations_service_notification"

func TestAccTwilioConversationsServiceNotification_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_notification", serviceNotificationResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "log_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.sound", ""),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.badge_count_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.sound", ""),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.sound", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsServiceNotification_newMessage(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_notification", serviceNotificationResourceName)
	friendlyName := acctest.RandString(10)
	enabled := true
	template := "$${CONVERSATION}:$${PARTICIPANT}: $${MESSAGE}"
	sound := "bell"
	badgeCountEnabled := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceNotification_newMessageBlock(friendlyName, enabled, template, sound, badgeCountEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.template", "${CONVERSATION}:${PARTICIPANT}: ${MESSAGE}"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.sound", sound),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.badge_count_enabled", "true"),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_emptyNewMessageBlock(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.sound", ""),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.badge_count_enabled", "false"),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.sound", ""),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.0.badge_count_enabled", "false"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsServiceNotification_addedToConversation(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_notification", serviceNotificationResourceName)
	friendlyName := acctest.RandString(10)
	enabled := true
	template := "You have been added to the conversation $${CONVERSATION} by $${PARTICIPANT}"
	sound := "bell"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceNotification_addedToConversationBlock(friendlyName, enabled, template, sound),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.template", "You have been added to the conversation ${CONVERSATION} by ${PARTICIPANT}"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.sound", sound),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_emptyAddedToConversationBlock(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.sound", ""),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.0.sound", ""),
				),
			},
		},
	})
}

func TestAccTwilioConversationsServiceNotification_removedFromConversation(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_notification", serviceNotificationResourceName)
	friendlyName := acctest.RandString(10)
	enabled := true
	template := "$${PARTICIPANT} has removed you from the conversation $${CONVERSATION}"
	sound := "bell"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceNotification_removedFromConversationBlock(friendlyName, enabled, template, sound),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.template", "${PARTICIPANT} has removed you from the conversation ${CONVERSATION}"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.sound", sound),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_emptyRemovedFromConversationBlock(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.sound", ""),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.0.sound", ""),
				),
			},
		},
	})
}

func TestAccTwilioConversationsServiceNotification_logEnabled(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_notification", serviceNotificationResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "log_enabled", "false"),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_logsEnabledTrue(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "log_enabled", "true"),
				),
			},
			{
				Config: testAccTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "log_enabled", "false"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsServiceNotification_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsServiceNotification_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccCheckTwilioConversationsServiceNotificationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.ID).Configuration().Notification().Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving service notification information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsServiceNotification_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}

func testAccTwilioConversationsServiceNotification_newMessageBlock(friendlyName string, enabled bool, template string, sound string, badgetCountEnabled bool) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid

  new_message {
    enabled             = %t
    template            = "%s"
    sound               = "%s"
    badge_count_enabled = %t
  }
}
`, friendlyName, enabled, template, sound, badgetCountEnabled)
}

func testAccTwilioConversationsServiceNotification_emptyNewMessageBlock(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid

  new_message {}
}
`, friendlyName)
}

func testAccTwilioConversationsServiceNotification_addedToConversationBlock(friendlyName string, enabled bool, template string, sound string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid

  added_to_conversation {
    enabled  = %t
    template = "%s"
    sound    = "%s"
  }
}
`, friendlyName, enabled, template, sound)
}

func testAccTwilioConversationsServiceNotification_emptyAddedToConversationBlock(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid

  added_to_conversation {}
}
`, friendlyName)
}

func testAccTwilioConversationsServiceNotification_removedFromConversationBlock(friendlyName string, enabled bool, template string, sound string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid

  removed_from_conversation {
    enabled  = %t
    template = "%s"
    sound    = "%s"
  }
}
`, friendlyName, enabled, template, sound)
}

func testAccTwilioConversationsServiceNotification_emptyRemovedFromConversationBlock(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid

  removed_from_conversation {}
}
`, friendlyName)
}

func testAccTwilioConversationsServiceNotification_logsEnabledTrue(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid
  log_enabled = true
}
`, friendlyName)
}

func testAccTwilioConversationsServiceNotification_invalidServiceSid() string {
	return `
resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = "service_sid"
}
`
}
