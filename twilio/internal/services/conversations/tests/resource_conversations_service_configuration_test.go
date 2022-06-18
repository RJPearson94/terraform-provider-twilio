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

var serviceConfigurationResourceName = "twilio_conversations_service_configuration"

func TestAccTwilioConversationsServiceConfiguration_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_configuration", serviceConfigurationResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_chat_service_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_conversation_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_conversation_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "reachability_enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsServiceConfiguration_reachabilityIndicator(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_configuration", serviceConfigurationResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_enabled", "false"),
				),
			},
			{
				Config: testAccTwilioConversationsServiceConfiguration_reachabilityIndicatorTrue(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_enabled", "true"),
				),
			},
			{
				Config: testAccTwilioConversationsServiceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_enabled", "false"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsServiceConfiguration_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsServiceConfiguration_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsServiceConfiguration_invalidDefaultChatServiceRoleSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsServiceConfiguration_invalidDefaultChatServiceRoleSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of default_chat_service_role_sid to match regular expression "\^RL\[0-9a-fA-F\]\{32\}\$", got default_chat_service_role_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsServiceConfiguration_invalidDefaultConversationCreatorRole(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsServiceConfiguration_invalidDefaultConversationCreatorRole(),
				ExpectError: regexp.MustCompile(`(?s)expected value of default_conversation_creator_role_sid to match regular expression "\^RL\[0-9a-fA-F\]\{32\}\$", got default_conversation_creator_role_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsServiceConfiguration_invalidDefaultConversationRoleSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsServiceConfiguration_invalidDefaultConversationRoleSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of default_conversation_role_sid to match regular expression "\^RL\[0-9a-fA-F\]\{32\}\$", got default_conversation_role_sid`),
			},
		},
	})
}

func testAccCheckTwilioConversationsServiceConfigurationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.ID).Configuration().Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving service configuration information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsServiceConfiguration_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_configuration" "service_configuration" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}

func testAccTwilioConversationsServiceConfiguration_reachabilityIndicatorTrue(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_configuration" "service_configuration" {
  service_sid          = twilio_conversations_service.service.sid
  reachability_enabled = true
}
`, friendlyName)
}

func testAccTwilioConversationsServiceConfiguration_invalidServiceSid() string {
	return `
resource "twilio_conversations_service_configuration" "service_configuration" {
  service_sid = "service_sid"
}
`
}

func testAccTwilioConversationsServiceConfiguration_invalidDefaultChatServiceRoleSid() string {
	return `
resource "twilio_conversations_service_configuration" "service_configuration" {
  service_sid                   = "service_sid"
  default_chat_service_role_sid = "default_chat_service_role_sid"
}
`
}

func testAccTwilioConversationsServiceConfiguration_invalidDefaultConversationCreatorRole() string {
	return `
resource "twilio_conversations_service_configuration" "service_configuration" {
  service_sid                           = "service_sid"
  default_conversation_creator_role_sid = "default_conversation_creator_role_sid"
}
`
}

func testAccTwilioConversationsServiceConfiguration_invalidDefaultConversationRoleSid() string {
	return `
resource "twilio_conversations_service_configuration" "service_configuration" {
  service_sid                   = "service_sid"
  default_conversation_role_sid = "default_conversation_role_sid"
}
`
}
