package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var userDataSourceName = "twilio_conversations_user"

func TestAccDataSourceTwilioConversationsUser_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.user", userDataSourceName)
	friendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsUser_basic(friendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSource, "identity", identity),
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSource, "attributes"),
					resource.TestCheckResourceAttrSet(stateDataSource, "is_notifiable"),
					resource.TestCheckResourceAttrSet(stateDataSource, "is_online"),
					resource.TestCheckResourceAttrSet(stateDataSource, "role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsUser_basic(friendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  identity    = "%s"
}

data "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  sid         = twilio_conversations_user.user.sid
}
`, friendlyName, identity)
}
