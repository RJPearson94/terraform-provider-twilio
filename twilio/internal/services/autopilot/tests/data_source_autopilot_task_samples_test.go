package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var taskSamplesDataSourceName = "twilio_autopilot_task_samples"

func TestAccDataSourceTwilioAutopilotTaskSamples_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_samples", taskSamplesDataSourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	taggedText := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotTaskSamples_basic(uniqueName, language, taggedText),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "samples.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "samples.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "samples.0.language", language),
					resource.TestCheckResourceAttr(stateDataSourceName, "samples.0.tagged_text", taggedText),
					resource.TestCheckResourceAttr(stateDataSourceName, "samples.0.source_channel", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "samples.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "samples.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "samples.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotTaskSamples_basic(uniqueName string, language string, taggedText string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
}

resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  task_sid      = twilio_autopilot_task.task.sid
  language      = "%s"
  tagged_text   = "%s"
}

data "twilio_autopilot_task_samples" "task_samples" {
  assistant_sid = twilio_autopilot_task_sample.task_sample.assistant_sid
  task_sid      = twilio_autopilot_task_sample.task_sample.task_sid
}
`, uniqueName, uniqueName, language, taggedText)
}
