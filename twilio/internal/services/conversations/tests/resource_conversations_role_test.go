package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var roleResourceName = "twilio_conversations_role"

func TestAccTwilioConversationsRole_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.role", roleResourceName)
	friendlyName := acctest.RandString(10)
	typeName := "conversation"
	permissions := []string{"sendMessage"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsRole_basic(friendlyName, typeName, permissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "type", typeName),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.0", permissions[0]),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioConversationsRoleImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsRole_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.role", roleResourceName)
	friendlyName := acctest.RandString(10)
	typeName := "conversation"
	permissions := []string{"sendMessage"}
	newPermissions := []string{"sendMediaMessage", "sendMessage"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsRole_basic(friendlyName, typeName, permissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "type", typeName),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.0", permissions[0]),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsRole_basic(friendlyName, typeName, newPermissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "type", typeName),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.0", newPermissions[0]),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.1", newPermissions[1]),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsRole_invalidPermission(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsRole_invalidPermission(),
				ExpectError: regexp.MustCompile(`(?s)expected permissions.0 to be one of \["editOwnMessage" "deleteAnyMessage" "addParticipant" "editConversationAttributes" "editAnyParticipantAttributes" "editAnyMessage" "editConversationName" "editAnyMessageAttributes" "deleteOwnMessage" "editOwnMessageAttributes" "removeParticipant" "addNonChatParticipant" "editOwnParticipantAttributes" "deleteConversation" "editNotificationLevel" "sendMessage" "leaveConversation" "sendMediaMessage" "editAnyUserInfo" "removeParticipant" "createConversation" "editOwnUserInfo" "joinConversation"\], got test`),
			},
		},
	})
}

func TestAccTwilioConversationsRole_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.role", roleResourceName)

	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(256)
	typeName := "conversation"
	permissions := []string{"sendMessage"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsRole_basic(friendlyName, typeName, permissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioConversationsRole_basic(newFriendlyName, typeName, permissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
		},
	})
}

func TestAccTwilioConversationsRole_invalidFriendlyNameWith0Characters(t *testing.T) {
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsRole_friendlyNameWithStubbedService(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 256\), got `),
			},
		},
	})
}

func TestAccTwilioConversationsRole_invalidFriendlyNameWith257Characters(t *testing.T) {
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsRole_friendlyNameWithStubbedService(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 256\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioConversationsRole_invalidType(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsRole_invalidType(),
				ExpectError: regexp.MustCompile(`(?s)expected type to be one of \["conversation" "service"\], got test`),
			},
		},
	})
}

func TestAccTwilioConversationsRole_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsRole_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccCheckTwilioConversationsRoleDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != roleResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Role(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving role information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioConversationsRoleExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Role(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving role information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsRoleImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Roles/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsRole_basic(friendlyName string, typeName string, permissions []string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_role" "role" {
  service_sid   = twilio_conversations_service.service.sid
  friendly_name = "%s"
  type          = "%s"
  permissions   = %s
}
`, friendlyName, friendlyName, typeName, `["`+strings.Join(permissions, `","`)+`"]`)
}

func testAccTwilioConversationsRole_invalidType() string {
	return `
resource "twilio_conversations_role" "role" {
  service_sid   = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name = "invalid_type"
  type          = "test"
  permissions   = ["sendMessage"]
}
`
}

func testAccTwilioConversationsRole_invalidPermission() string {
	return `
resource "twilio_conversations_role" "role" {
  service_sid   = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name = "invalid_permission"
  type          = "conversation"
  permissions   = ["test"]
}
`
}

func testAccTwilioConversationsRole_friendlyNameWithStubbedService(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_role" "role" {
  service_sid   = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name = "%s"
  type          = "conversation"
  permissions   = ["sendMessage"]
}
`, friendlyName)
}

func testAccTwilioConversationsRole_invalidServiceSid() string {
	return `
resource "twilio_conversations_role" "role" {
  service_sid   = "service_sid"
  friendly_name = "invalid_service_sid"
  type          = "conversation"
  permissions   = ["sendMessage"]
}
`
}
