package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var serviceNotificationResourceName = "twilio_conversations_service_notification"

func TestAccTwilioConversationsServiceNotification_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_notification", serviceNotificationResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceNotification_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceNotificationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "added_to_conversation.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "removed_from_conversation.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioConversationsServiceNotificationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.ID).Configuration().Notification().Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving service notification information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsServiceNotification_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}
