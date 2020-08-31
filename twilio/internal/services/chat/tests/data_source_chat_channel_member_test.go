package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var channelMemberDataSourceName = "twilio_chat_channel_member"

func TestAccDataSourceTwilioChatChannelMember_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.member", channelMemberDataSourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatChannelMember_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "attributes"),
					resource.TestCheckResourceAttr(stateDataSource, "identity", identity),
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "channel_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "last_consumed_message_index"),
					resource.TestCheckResourceAttr(stateDataSource, "last_consumption_timestamp", ""),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannelMember_basic(friendlyName string, identity string) string {
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

data "twilio_chat_channel_member" "member" {
  service_sid = twilio_chat_channel_member.member.service_sid
  channel_sid = twilio_chat_channel_member.member.channel_sid
  sid         = twilio_chat_channel_member.member.sid
}
`, friendlyName, friendlyName, identity)
}
