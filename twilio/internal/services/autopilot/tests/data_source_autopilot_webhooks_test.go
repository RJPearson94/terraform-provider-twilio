package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var webhooksDataSourceName = "twilio_autopilot_webhooks"

func TestAccDataSourceTwilioAutopilotWebhooks_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhooks", webhooksDataSourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/webhook"
	events := []string{"onDialogueStart", "onDialogueEnd"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
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

func testAccDataSourceTwilioAutopilotWebhooks_basic(uniqueName string, url string, events []string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
  webhook_url   = "%s"
  events        = %s
}

data "twilio_autopilot_webhooks" "webhooks" {
  assistant_sid = twilio_autopilot_webhook.webhook.assistant_sid
}
`, uniqueName, uniqueName, url, "[\""+strings.Join(events, "\",\"")+"\"]")
}
