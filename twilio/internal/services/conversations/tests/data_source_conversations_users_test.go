package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var usersDataSourceName = "twilio_conversations_users"

func TestAccDataSourceTwilioConversationsUsers_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.users", usersDataSourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsUsers_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "users.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.sid"),
					resource.TestCheckResourceAttr(stateDataSource, "users.0.identity", identity),
					resource.TestCheckResourceAttr(stateDataSource, "users.0.friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.attributes"),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.is_notifiable"),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.is_online"),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "users.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsUsers_basic(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  identity    = "%s"
}

data "twilio_conversations_users" "users" {
  service_sid = twilio_conversations_user.user.service_sid
}
`, friendlyName, identity)
}
