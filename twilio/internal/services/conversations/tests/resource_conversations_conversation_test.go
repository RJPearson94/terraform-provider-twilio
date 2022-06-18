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
					resource.TestCheckResourceAttr(stateResourceName, "attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "state", "active"),
					resource.TestCheckResourceAttr(stateResourceName, "timers.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:            stateResourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccTwilioConversationsConversationImportStateIdFunc(stateResourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timers.0.closed", "timers.0.inactive"},
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
					resource.TestCheckResourceAttr(stateResourceName, "attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "state", "active"),
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
					resource.TestCheckResourceAttr(stateResourceName, "attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "state", "active"),
					resource.TestCheckResourceAttr(stateResourceName, "timers.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_attributes(t *testing.T) {
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
					resource.TestCheckResourceAttr(stateResourceName, "attributes", "{}"),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_attributes(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "attributes", "{\"key\":\"value\"}"),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "attributes", "{}"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_invalidAttributesString(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversation_invalidAttributesString(),
				ExpectError: regexp.MustCompile(`(?s)"attributes" contains an invalid JSON`),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_uniqueName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.conversation", conversationResourceName)
	friendlyName := acctest.RandString(10)
	uniqueName := ""
	newUniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversation_uniqueName(friendlyName, uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_uniqueName(friendlyName, newUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.conversation", conversationResourceName)
	friendlyName := acctest.RandString(10)
	conversationFriendlyName := ""
	newConversationFriendlyName := acctest.RandString(256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversation_friendlyName(friendlyName, conversationFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", conversationFriendlyName),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_friendlyName(friendlyName, newConversationFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newConversationFriendlyName),
				),
			},
			// This is currently disabled as the API is currently not clearing down the old value when a blank string is supplied
			// {
			// 	Config: testAccTwilioConversationsConversation_basic(friendlyName),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckTwilioConversationsConversationExists(stateResourceName),
			// 		resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
			// 	),
			// },
		},
	})
}

func TestAccTwilioConversationsConversation_invalidFriendlyNameWith257Characters(t *testing.T) {
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversation_friendlyNameWithStubbedServiceSid(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(0 - 256\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_state(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.conversation", conversationResourceName)
	friendlyName := acctest.RandString(10)
	newState := "inactive"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversation_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "state", "active"),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_state(friendlyName, newState),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "state", newState),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "state", "active"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_invalidState(t *testing.T) {
	friendlyName := acctest.RandString(10)
	state := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversation_state(friendlyName, state),
				ExpectError: regexp.MustCompile(`(?s)expected state to be one of \[active inactive closed\], got test`),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_messagingService(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.conversation", conversationResourceName)
	messagingServiceStateResourceName := "twilio_messaging_service.service"
	configurationDataStateResourceName := "data.twilio_conversations_configuration.configuration"

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsConversationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConversation_customMessagingService(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttrPair(stateResourceName, "messaging_service_sid", messagingServiceStateResourceName, "sid"),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_detachMessagingService(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttrPair(stateResourceName, "messaging_service_sid", messagingServiceStateResourceName, "sid"),
				),
			},
			{
				Config: testAccTwilioConversationsConversation_updateMessagingServiceToBeConversationDefault(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConversationExists(stateResourceName),
					resource.TestCheckResourceAttrPair(stateResourceName, "messaging_service_sid", configurationDataStateResourceName, "default_messaging_service_sid"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_invalidMessagingServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversation_invalidMessagingServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of messaging_service_sid to match regular expression "\^MG\[0-9a-fA-F\]\{32\}\$", got messaging_service_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsConversation_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsConversation_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
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

func testAccTwilioConversationsConversation_friendlyName(friendlyName string, conversationFriendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid   = twilio_conversations_service.service.sid
  friendly_name = "%s"
}
`, friendlyName, conversationFriendlyName)
}

func testAccTwilioConversationsConversation_uniqueName(friendlyName string, uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
  unique_name = "%s"
}
`, friendlyName, uniqueName)
}

func testAccTwilioConversationsConversation_friendlyNameWithStubbedServiceSid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_conversation" "conversation" {
  service_sid   = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name = "%s"
}
`, friendlyName)
}

func testAccTwilioConversationsConversation_state(friendlyName string, state string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
  state       = "%s"
}
`, friendlyName, state)
}

func testAccTwilioConversationsConversation_attributes(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
  attributes = jsonencode({
    "key" : "value"
  })
}
`, friendlyName)
}

func testAccTwilioConversationsConversation_invalidAttributesString() string {
	return `
resource "twilio_conversations_conversation" "conversation" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  attributes  = "attributes"
}
`
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

func testAccTwilioConversationsConversation_invalidServiceSid() string {
	return `
resource "twilio_conversations_conversation" "conversation" {
  service_sid = "service_sid"
}
`
}

func testAccTwilioConversationsConversation_customMessagingService(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid           = twilio_conversations_service.service.sid
  messaging_service_sid = twilio_messaging_service.service.sid
}
`, friendlyName)
}

func testAccTwilioConversationsConversation_detachMessagingService(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}

func testAccTwilioConversationsConversation_updateMessagingServiceToBeConversationDefault(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "%[1]s"
}

data "twilio_conversations_configuration" "configuration" {}

resource "twilio_conversations_conversation" "conversation" {
  service_sid           = twilio_conversations_service.service.sid
  messaging_service_sid = data.twilio_conversations_configuration.configuration.default_messaging_service_sid
}
`, friendlyName)
}

func testAccTwilioConversationsConversation_invalidMessagingServiceSid() string {
	return `
resource "twilio_conversations_conversation" "conversation" {
  service_sid           = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  messaging_service_sid = "messaging_service_sid"
}
`
}
