package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var channelMemberDataSourceName = "twilio_chat_channel_member"

func TestAccDataSourceTwilioChatChannelMember_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.member", channelMemberDataSourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
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
					resource.TestCheckNoResourceAttr(stateDataSource, "last_consumption_timestamp"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioChatChannelMember_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannelMember_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioChatChannelMember_invalidChannelSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannelMember_invalidChannelSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of channel_sid to match regular expression "\^CH\[0-9a-fA-F\]\{32\}\$", got channel_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioChatChannelMember_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannelMember_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^MB\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannelMember_basic(friendlyName string, identity string) string {
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

data "twilio_chat_channel_member" "member" {
  service_sid = twilio_chat_channel_member.member.service_sid
  channel_sid = twilio_chat_channel_member.member.channel_sid
  sid         = twilio_chat_channel_member.member.sid
}
`, friendlyName, identity)
}

func testAccDataSourceTwilioChatChannelMember_invalidServiceSid() string {
	return `
data "twilio_chat_channel_member" "member" {
  service_sid = "service_sid"
  channel_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "MBaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioChatChannelMember_invalidChannelSid() string {
	return `
data "twilio_chat_channel_member" "member" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid = "channel_sid"
  sid         = "MBaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioChatChannelMember_invalidSid() string {
	return `
data "twilio_chat_channel_member" "member" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid = "CHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
