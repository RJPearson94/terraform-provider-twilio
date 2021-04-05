package tests

import (
	"fmt"
	"regexp"
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

func TestAccDataSourceTwilioConversationsUser_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsUser_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsUser_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsUser_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^US\[0-9a-fA-F\]\{32\}\$", got sid`),
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

func testAccDataSourceTwilioConversationsUser_invalidServiceSid() string {
	return `
data "twilio_conversations_user" "user" {
  service_sid = "service_sid"
  sid         = "USaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioConversationsUser_invalidSid() string {
	return `
data "twilio_conversations_user" "user" {
  service_sid = "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
