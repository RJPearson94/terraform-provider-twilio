package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var conversationDataSourceName = "twilio_conversations_conversation"

func TestAccDataSourceTwilioConversationsConversation_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.conversation", conversationDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsConversation_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "unique_name", ""),
					resource.TestCheckResourceAttr(stateDataSource, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSource, "attributes"),
					resource.TestCheckResourceAttrSet(stateDataSource, "messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "state"),
					resource.TestCheckResourceAttr(stateDataSource, "timers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsConversation_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}

data "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_conversation.conversation.service_sid
  sid         = twilio_conversations_conversation.conversation.sid
}
`, friendlyName)
}
