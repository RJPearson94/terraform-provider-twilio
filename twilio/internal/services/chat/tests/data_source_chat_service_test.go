package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var serviceDataSourceName = "twilio_chat_service"

func TestAccDataSourceTwilioChatService_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service", serviceDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_channel_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_channel_role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_service_role_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "limits.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "media.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "notifications.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "notifications.0.added_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "notifications.0.removed_from_channel.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "notifications.0.invited_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "notifications.0.new_message.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatService_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

data "twilio_chat_service" "service" {
	sid = twilio_chat_service.service.sid
}
`, friendlyName)
}
