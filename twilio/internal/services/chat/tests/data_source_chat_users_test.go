package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var usersDataSourceName = "twilio_chat_users"

func TestAccDataSourceTwilioChatUsers_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.users", usersDataSourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatUsers_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "users.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "users.0.identity", identity),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "users.0.friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.attributes"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.is_notifiable"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.is_online"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.joined_channels_count"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "users.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatUsers_basic(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "%s"
}

data "twilio_chat_users" "users" {
  service_sid = twilio_chat_user.user.service_sid
}
`, friendlyName, identity)
}
