package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var rolesDataSourceName = "twilio_chat_roles"

func TestAccDataSourceTwilioChatRoles_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.roles", rolesDataSourceName)
	friendlyName := acctest.RandString(10)
	permissions := []string{
		"sendMessage",
		"leaveChannel",
	}
	roleType := "channel"

	// Twilio creates some default roles when the channel is created so cant guarantee the order
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioChatRoles_basic(friendlyName, roleType, permissions),
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.0.type"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.0.permissions.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "roles.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioChatRoles_basic(friendlyName string, roleType string, permissions []string) string {
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

data "twilio_chat_roles" "roles" {
	service_sid   = twilio_chat_role.role.service_sid
  }
`, friendlyName, friendlyName, roleType, `["`+strings.Join(permissions, `","`)+`"]`)
}
