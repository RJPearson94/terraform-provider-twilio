package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var channelsDataSourceName = "twilio_chat_channels"

func TestAccDataSourceTwilioChatChannels_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.channels", channelsDataSourceName)
	friendlyName := acctest.RandString(10)
	channelType := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
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

func testAccDataSourceTwilioChatChannels_basic(friendlyName string, channelType string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%s"
  type          = "%s"
}

data "twilio_chat_channels" "channels" {
	service_sid = twilio_chat_channel.channel.service_sid
}
`, friendlyName, friendlyName, channelType)
}
