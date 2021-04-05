package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var serviceNotificationDataSourceName = "twilio_conversations_service_notification"

func TestAccDataSourceTwilioConversationsServiceNotification_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service_notification", serviceNotificationDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "new_message.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "added_to_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "removed_from_conversation.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsServiceNotification_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsServiceNotification_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsServiceNotification_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

data "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}

func testAccDataSourceTwilioConversationsServiceNotification_invalidServiceSid() string {
	return `
data "twilio_conversations_service_notification" "service_notification" {
  service_sid = "service_sid"
}
`
}
