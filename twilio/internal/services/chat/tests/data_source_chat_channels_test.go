package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var channelsDataSourceName = "twilio_chat_channels"

func TestAccDataSourceTwilioChatChannels_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.channels", channelsDataSourceName)
	friendlyName := acctest.RandString(10)
	channelType := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatChannels_basic(friendlyName, channelType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "channels.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "channels.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "channels.0.unique_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.attributes"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.type"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.created_by"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.members_count"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.messages_count"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "channels.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioChatChannels_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatChannels_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannels_basic(friendlyName string, channelType string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%[1]s"
  type          = "%[2]s"
}

data "twilio_chat_channels" "channels" {
  service_sid = twilio_chat_channel.channel.service_sid
}
`, friendlyName, channelType)
}

func testAccDataSourceTwilioChatChannels_invalidServiceSid() string {
	return `
data "twilio_chat_channels" "channels" {
  service_sid = "service_sid"
}
`
}
