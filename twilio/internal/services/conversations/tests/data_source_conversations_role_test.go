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

var roleDataSourceName = "twilio_conversations_role"

func TestAccDataSourceTwilioConversationsRole_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.role", roleDataSourceName)
	friendlyName := acctest.RandString(10)
	typeName := "conversation"
	permissions := []string{"sendMessage"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsRole_basic(friendlyName, typeName, permissions),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSource, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSource, "type", typeName),
					resource.TestCheckResourceAttr(stateDataSource, "permissions.#", "1"),
					resource.TestCheckResourceAttr(stateDataSource, "permissions.0", permissions[0]),
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsRole_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsRole_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsRole_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsRole_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^RL\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsRole_basic(friendlyName string, typeName string, permissions []string) string {
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

data "twilio_conversations_role" "role" {
  service_sid = twilio_conversations_service.service.sid
  sid         = twilio_conversations_role.role.sid
}
`, friendlyName, friendlyName, typeName, `["`+strings.Join(permissions, `","`)+`"]`)
}

func testAccDataSourceTwilioConversationsRole_invalidServiceSid() string {
	return `
data "twilio_conversations_role" "role" {
  service_sid = "service_sid"
  sid         = "RLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioConversationsRole_invalidSid() string {
	return `
data "twilio_conversations_role" "role" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
