package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var conversationsDataSourceName = "twilio_conversations_conversations"

func TestAccDataSourceTwilioConversationsConversations_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.conversations", conversationsDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsConversations_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSource, "conversations.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversations.0.sid"),
					resource.TestCheckResourceAttr(stateDataSource, "conversations.0.unique_name", ""),
					resource.TestCheckResourceAttr(stateDataSource, "conversations.0.friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversations.0.attributes"),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversations.0.messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversations.0.state"),
					resource.TestCheckResourceAttr(stateDataSource, "conversations.0.timers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversations.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversations.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSource, "conversations.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioConversationsConversations_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioConversationsConversations_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsConversations_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}

data "twilio_conversations_conversations" "conversations" {
  service_sid = twilio_conversations_conversation.conversation.service_sid
}
`, friendlyName)
}

func testAccDataSourceTwilioConversationsConversations_invalidServiceSid() string {
	return `
data "twilio_conversations_conversations" "conversations" {
  service_sid = "service_sid"
}
`
}
