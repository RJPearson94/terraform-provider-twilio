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

var webhooksDataSourceName = "twilio_autopilot_webhooks"

func TestAccDataSourceTwilioAutopilotWebhooks_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhooks", webhooksDataSourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/webhook"
	events := []string{"onDialogueStart", "onDialogueEnd"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotWebhooks_basic(uniqueName, url, events),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.events.#", "2"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.events.0", "onDialogueStart"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.events.1", "onDialogueEnd"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.webhook_url", url),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.webhook_method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotWebhooks_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotWebhooks_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotWebhooks_basic(uniqueName string, url string, events []string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  webhook_url   = "%[2]s"
  events        = %[3]s
}

data "twilio_autopilot_webhooks" "webhooks" {
  assistant_sid = twilio_autopilot_webhook.webhook.assistant_sid
}
`, uniqueName, url, `["`+strings.Join(events, `","`)+`"]`)
}

func testAccDataSourceTwilioAutopilotWebhooks_invalidAssistantSid() string {
	return `
data "twilio_autopilot_webhooks" "webhooks" {
  assistant_sid = "assistant_sid"
}
`
}
