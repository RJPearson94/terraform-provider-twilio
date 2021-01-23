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

var userResourceName = "twilio_conversations_user"

func TestAccTwilioConversationsUser_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.user", userResourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsUser_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsUserExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "attributes"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_notifiable"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_online"),
					resource.TestCheckResourceAttrSet(stateResourceName, "role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioConversationsUserImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsUser_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.user", userResourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsUser_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsUserExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "attributes"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_notifiable"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_online"),
					resource.TestCheckResourceAttrSet(stateResourceName, "role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsUser_withAttributes(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsUserExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttr(stateResourceName, "attributes", `{"test":true}`),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_notifiable"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_online"),
					resource.TestCheckResourceAttrSet(stateResourceName, "role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioConversationsUserDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != userResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).User(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving user information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioConversationsUserExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

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

func testAccTwilioConversationsUserImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Users/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsUser_basic(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  identity    = "%s"
}
`, friendlyName, identity)
}

func testAccTwilioConversationsUser_withAttributes(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  identity    = "%s"
  attributes  = "{\"test\": true}"
}
`, friendlyName, identity)
}
