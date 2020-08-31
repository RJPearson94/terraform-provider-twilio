package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var channelMemberResourceName = "twilio_chat_channel_member"

func TestAccTwilioChatChannelMember_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.member", channelMemberResourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		Providers:         acceptance.TestAccProviders,
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioChatChannelMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelMember_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelMemberExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "attributes"),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "last_consumed_message_index"),
					resource.TestCheckResourceAttr(stateResourceName, "last_consumption_timestamp", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioChatChannelMemberImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioChatChannelMember_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.member", channelMemberResourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)
	attributes := `{"test":"test"}`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioChatChannelMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatChannelMember_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelMemberExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "attributes"),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "last_consumed_message_index"),
					resource.TestCheckResourceAttr(stateResourceName, "last_consumption_timestamp", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioChatChannelMember_attributes(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatChannelMemberExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "attributes", attributes),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "last_consumed_message_index"),
					resource.TestCheckResourceAttr(stateResourceName, "last_consumption_timestamp", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioChatChannelMemberDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

	for _, rs := range s.RootModule().Resources {
		if rs.Type != channelMemberResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Channel(rs.Primary.Attributes["channel_sid"]).Member(rs.Primary.ID).Fetch(); err != nil {
			if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
				// currently programmable chat returns a 403 if the service instance does not exist
				if twilioError.Status == 403 && twilioError.Message == "Service instance not found" {
					return nil
				}
			}
			return fmt.Errorf("Error occurred when retrieving channel member information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioChatChannelMemberExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Channel(rs.Primary.Attributes["channel_sid"]).Member(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving channel member information %s", err)
		}

		return nil
	}
}

func testAccTwilioChatChannelMemberImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Channels/%s/Members/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["channel_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioChatChannelMember_basic(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%s"
  type          = "private"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "%s"
}

resource "twilio_chat_channel_member" "member" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  identity    = twilio_chat_user.user.identity
}
`, friendlyName, friendlyName, identity)
}

func testAccTwilioChatChannelMember_attributes(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%s"
  type          = "private"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "%s"
}

resource "twilio_chat_channel_member" "member" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  identity    = twilio_chat_user.user.identity
  attributes  = <<EOF
{
	"test": "test"
}
EOF
}
`, friendlyName, friendlyName, identity)
}
