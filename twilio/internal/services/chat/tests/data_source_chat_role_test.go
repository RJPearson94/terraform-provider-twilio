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
		ProviderFactories: acceptance.TestAccProviderFactories,
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

func TestAccDataSourceTwilioChatRole_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatRole_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioChatRole_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioChatRole_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^RL\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioChatRole_basic(friendlyName string, roleType string, permissions []string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_chat_role" "role" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "%[1]s"
  type          = "%[2]s"
  permissions   = %[3]s
}

data "twilio_chat_role" "role" {
  service_sid = twilio_chat_role.role.service_sid
  sid         = twilio_chat_role.role.sid
}
`, friendlyName, roleType, `["`+strings.Join(permissions, `","`)+`"]`)
}

func testAccDataSourceTwilioChatRole_invalidServiceSid() string {
	return `
data "twilio_chat_role" "role" {
  service_sid = "service_sid"
  sid         = "RLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioChatRole_invalidSid() string {
	return `
data "twilio_chat_role" "role" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
