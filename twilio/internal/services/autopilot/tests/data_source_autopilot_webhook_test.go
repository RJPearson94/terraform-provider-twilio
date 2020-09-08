package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var webhookDataSourceName = "twilio_autopilot_webhook"

func TestAccDataSourceTwilioAutopilotWebhook_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhook", webhookDataSourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/webhook"
	events := []string{"onDialogueStart", "onDialogueEnd"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotWebhook_basic(uniqueName, url, events),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "events.#", "2"),
					resource.TestCheckResourceAttr(stateDataSourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttr(stateDataSourceName, "events.1", "onDialogueEnd"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhook_url", url),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhook_method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotWebhook_basic(uniqueName string, url string, events []string) string {
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

data "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_webhook.webhook.assistant_sid
  sid           = twilio_autopilot_webhook.webhook.sid
}
`, uniqueName, uniqueName, url, "[\""+strings.Join(events, "\",\"")+"\"]")
}
