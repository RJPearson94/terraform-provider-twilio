package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var channelDataSourceName = "twilio_chat_channel"

func TestAccDataSourceTwilioChatChannel_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.channel", channelDataSourceName)
	friendlyName := acctest.RandString(10)
	channelType := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatChannel_basic(friendlyName, channelType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "attributes"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "type"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "created_by"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "members_count"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "messages_count"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatChannel_basic(friendlyName string, channelType string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%s"
  type          = "%s"
}

data "twilio_chat_channel" "channel" {
  service_sid = twilio_chat_channel.channel.service_sid
  sid         = twilio_chat_channel.channel.sid
}
`, friendlyName, friendlyName, channelType)
}
