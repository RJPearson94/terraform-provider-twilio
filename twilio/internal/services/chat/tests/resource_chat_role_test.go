package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var roleResourceName = "twilio_chat_role"

func TestAccTwilioChatRole_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.role", roleResourceName)
	friendlyName := acctest.RandString(10)
	permissions := []string{
		"sendMessage",
		"leaveChannel",
	}
	roleType := "channel"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatRole_basic(friendlyName, roleType, permissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "type", roleType),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.0", "sendMessage"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.1", "leaveChannel"),
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
				ImportStateIdFunc: testAccTwilioChatRoleImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioChatRole_invalidType(t *testing.T) {
	friendlyName := acctest.RandString(10)
	permissions := []string{
		"sendMessage",
		"leaveChannel",
	}
	roleType := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatRole_basic(friendlyName, roleType, permissions),
				ExpectError: regexp.MustCompile(`(?s)expected type to be one of \[channel deployment\], got test`),
			},
		},
	})
}

func TestAccTwilioChatRole_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.role", roleResourceName)
	friendlyName := acctest.RandString(10)
	permissions := []string{
		"sendMessage",
		"leaveChannel",
	}
	newPermissions := []string{
		"sendMessage",
	}
	roleType := "channel"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatRole_basic(friendlyName, roleType, permissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "type", roleType),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.0", "sendMessage"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.1", "leaveChannel"),
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
				Config: testAccTwilioChatRole_basic(friendlyName, roleType, newPermissions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatRoleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "type", roleType),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "permissions.0", "sendMessage"),
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

func TestAccTwilioChatRole_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatRole_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccCheckTwilioChatRoleDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

	for _, rs := range s.RootModule().Resources {
		if rs.Type != roleResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Role(rs.Primary.ID).Fetch(); err != nil {
			if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
				// currently programmable chat returns a 403 if the service instance does not exist
				if twilioError.Status == 403 && twilioError.Message == "Service instance not found" {
					return nil
				}
			}
			return fmt.Errorf("Error occurred when retrieving role information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioChatRoleExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

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

func testAccTwilioChatRoleImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Roles/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioChatRole_basic(friendlyName string, roleType string, permissions []string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_chat_role" "role" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%[1]s"
  type          = "%[2]s"
  permissions   = %[3]s
}
`, friendlyName, roleType, `["`+strings.Join(permissions, `","`)+`"]`)
}

func testAccTwilioChatRole_invalidServiceSid() string {
	return `
resource "twilio_chat_role" "role" {
  service_sid   = "service_sid"
  friendly_name = "invalid_service_sid"
  type          = "channel"
  permissions = [
    "sendMessage"
  ]
}
`
}
