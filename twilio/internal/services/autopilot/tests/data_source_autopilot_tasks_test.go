package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var tasksDataSourceName = "twilio_autopilot_tasks"

func TestAccDataSourceTwilioAutopilotTasks_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.tasks", tasksDataSourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotTasks_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "tasks.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "tasks.0.unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "tasks.0.friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "tasks.0.actions_url"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "tasks.0.actions"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "tasks.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "tasks.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "tasks.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTasks_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTasks_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotTasks_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}

data "twilio_autopilot_tasks" "tasks" {
  assistant_sid = twilio_autopilot_task.task.assistant_sid
}
`, uniqueName)
}

func testAccDataSourceTwilioAutopilotTasks_invalidAssistantSid() string {
	return `
data "twilio_autopilot_tasks" "tasks" {
  assistant_sid = "assistant_sid"
}
`
}
