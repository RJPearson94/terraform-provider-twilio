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

var userResourceName = "twilio_chat_user"

func TestAccTwilioChatUser_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.user", userResourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatUser_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatUserExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_notifiable"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_online"),
					resource.TestCheckResourceAttrSet(stateResourceName, "joined_channels_count"),
					resource.TestCheckResourceAttrSet(stateResourceName, "role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioChatUserImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioChatUser_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.user", userResourceName)
	friendlyName := acctest.RandString(10)
	userFriendlyName := ""
	newUserFriendlyName := acctest.RandString(256)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatUser_friendlyName(friendlyName, identity, userFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatUserExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", userFriendlyName),
				),
			},
			{
				Config: testAccTwilioChatUser_friendlyName(friendlyName, identity, newUserFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatUserExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newUserFriendlyName),
				),
			},
			{
				Config: testAccTwilioChatUser_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatUserExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioChatUser_role(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.user", userResourceName)
	roleStateResourceName := "twilio_chat_role.role"
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatUser_role(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatUserExists(stateResourceName),
					resource.TestCheckResourceAttrPair(stateResourceName, "role_sid", roleStateResourceName, "sid"),
				),
			},
		},
	})
}

func TestAccTwilioChatUser_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatUser_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioChatUser_invalidRoleSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatUser_invalidRoleSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of role_sid to match regular expression "\^RL\[0-9a-fA-F\]\{32\}\$", got role_sid`),
			},
		},
	})
}

func testAccCheckTwilioChatUserDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

	for _, rs := range s.RootModule().Resources {
		if rs.Type != userResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).User(rs.Primary.ID).Fetch(); err != nil {
			if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
				// currently programmable chat returns a 403 if the service instance does not exist
				if twilioError.Status == 403 && twilioError.Message == "Service instance not found" {
					return nil
				}
			}
			return fmt.Errorf("Error occurred when retrieving user information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioChatUserExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).User(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving user information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioChatUserImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Users/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioChatUser_basic(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "%s"
}
`, friendlyName, identity)
}

func testAccTwilioChatUser_friendlyName(friendlyName string, identity string, userFriendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_user" "user" {
  service_sid   = twilio_chat_service.service.sid
  identity      = "%s"
  friendly_name = "%s"
  attributes    = "{\"test\":\"test\"}"
}
`, friendlyName, identity, userFriendlyName)
}

func testAccTwilioChatUser_role(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_chat_role" "role" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%[1]s"
  type          = "deployment"
  permissions   = ["createChannel", "editOwnUserInfo", "joinChannel"]
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "%[2]s"
  role_sid    = twilio_chat_role.role.sid
}
`, friendlyName, identity)
}

func testAccTwilioChatUser_invalidServiceSid() string {
	return `
resource "twilio_chat_user" "user" {
  service_sid = "service_sid"
  identity    = "invalid_service_sid"
}
`
}

func testAccTwilioChatUser_invalidRoleSid() string {
	return `
resource "twilio_chat_user" "user" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  identity    = "invalid_role_sid"
  role_sid    = "role_sid"
}
`
}
