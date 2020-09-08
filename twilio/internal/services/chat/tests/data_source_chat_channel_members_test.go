package tests

import (
	"fmt"
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

func testAccDataSourceTwilioChatChannelMembers_basic(friendlyName string, identity string) string {
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

data "twilio_chat_channel_members" "members" {
  service_sid = twilio_chat_channel_member.member.service_sid
  channel_sid = twilio_chat_channel_member.member.channel_sid
}
`, friendlyName, friendlyName, identity)
}
