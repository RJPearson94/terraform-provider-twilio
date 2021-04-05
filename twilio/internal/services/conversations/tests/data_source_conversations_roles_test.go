package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var rolesDataSourceName = "twilio_conversations_roles"

func TestAccDataSourceTwilioConversationsRoles_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.roles", rolesDataSourceName)
	friendlyName := acctest.RandString(10)
	typeName := "conversation"
	permissions := []string{"sendMessage"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsRoles_basic(friendlyName, typeName, permissions),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.#"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.0.type"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.0.permissions.#"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "roles.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsRoles_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsRoles_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsRoles_basic(friendlyName string, typeName string, permissions []string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_role" "role" {
  service_sid   = twilio_conversations_service.service.sid
  friendly_name = "%s"
  type          = "%s"
  permissions   = %s
}

data "twilio_conversations_roles" "roles" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName, friendlyName, typeName, `["`+strings.Join(permissions, `","`)+`"]`)
}

func testAccDataSourceTwilioConversationsRoles_invalidServiceSid() string {
	return `
data "twilio_conversations_roles" "rolse" {
  service_sid = "service_sid"
}
`
}
