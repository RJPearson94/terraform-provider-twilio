package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var channelMembersDataSourceName = "twilio_chat_channel_members"

func TestAccDataSourceTwilioChatChannelMembers_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.members", channelMembersDataSourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatChannelMembers_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "channel_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "members.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSource, "members.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "members.0.attributes"),
					resource.TestCheckResourceAttr(stateDataSource, "members.0.identity", identity),
					resource.TestCheckResourceAttrSet(stateDataSource, "members.0.role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "members.0.last_consumed_message_index"),
					resource.TestCheckResourceAttr(stateDataSource, "members.0.last_consumption_timestamp", ""),
					resource.TestCheckResourceAttrSet(stateDataSource, "members.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "members.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "members.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioChatChannelMembers_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannelMembers_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioChatChannelMembers_invalidChannelSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannelMembers_invalidChannelSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of channel_sid to match regular expression "\^CH\[0-9a-fA-F\]\{32\}\$", got channel_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannelMembers_basic(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%[1]s"
  type          = "private"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "%[2]s"
}

resource "twilio_chat_channel_member" "member" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  identity    = twilio_chat_user.user.identity
}

data "twilio_chat_channel_members" "members" {
  service_sid = twilio_chat_channel_member.member.service_sid
  channel_sid = twilio_chat_channel_member.member.channel_sid
}
`, friendlyName, identity)
}

func testAccDataSourceTwilioChatChannelMembers_invalidServiceSid() string {
	return `
data "twilio_chat_channel_members" "members" {
  service_sid = "service_sid"
  channel_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioChatChannelMembers_invalidChannelSid() string {
	return `
data "twilio_chat_channel_members" "members" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid = "channel_sid"
}
`
}
