package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var roleDataSourceName = "twilio_chat_role"

func TestAccDataSourceTwilioChatRole_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.role", roleDataSourceName)
	friendlyName := acctest.RandString(10)
	permissions := []string{
		"sendMessage",
		"leaveChannel",
	}
	roleType := "channel"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatRole_basic(friendlyName, roleType, permissions),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "type", roleType),
					resource.TestCheckResourceAttr(stateDataSourceName, "permissions.#", "2"),
					resource.TestCheckResourceAttr(stateDataSourceName, "permissions.0", "sendMessage"),
					resource.TestCheckResourceAttr(stateDataSourceName, "permissions.1", "leaveChannel"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatRole_basic(friendlyName string, roleType string, permissions []string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}

resource "twilio_chat_role" "role" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%s"
  type          = "%s"
  permissions   = %s
}

data "twilio_chat_role" "role" {
  service_sid = twilio_chat_role.role.service_sid
  sid         = twilio_chat_role.role.sid
}
`, friendlyName, friendlyName, roleType, `["`+strings.Join(permissions, `","`)+`"]`)
}
