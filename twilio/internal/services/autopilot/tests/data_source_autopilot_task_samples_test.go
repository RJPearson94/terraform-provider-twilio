package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var taskSamplesDataSourceName = "twilio_autopilot_task_samples"

func TestAccDataSourceTwilioAutopilotTaskSamples_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_samples", taskSamplesDataSourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	taggedText := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
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
					resource.TestCheckResourceAttr(stateDataSourceName, "samples.0.source_channel", "voice"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "samples.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "samples.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "samples.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskSamples_language(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_samples", taskSamplesDataSourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	taggedText := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotTaskSamples_language(uniqueName, language, taggedText),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "samples.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskSamples_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTaskSamples_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskSamples_invalidTaskSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTaskSamples_invalidTaskSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of task_sid to match regular expression "\^UD\[0-9a-fA-F\]\{32\}\$", got task_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotTaskSamples_basic(uniqueName string, language string, taggedText string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}

resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  task_sid      = twilio_autopilot_task.task.sid
  language      = "%[2]s"
  tagged_text   = "%[3]s"
}

data "twilio_autopilot_task_samples" "task_samples" {
  assistant_sid = twilio_autopilot_task_sample.task_sample.assistant_sid
  task_sid      = twilio_autopilot_task_sample.task_sample.task_sid
}
`, uniqueName, language, taggedText)
}

func testAccDataSourceTwilioAutopilotTaskSamples_language(uniqueName string, language string, taggedText string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}

resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  task_sid      = twilio_autopilot_task.task.sid
  language      = "%[2]s"
  tagged_text   = "%[3]s"
}

resource "twilio_autopilot_task_sample" "task_sample_2" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  task_sid      = twilio_autopilot_task.task.sid
  language      = "%[2]s"
  tagged_text   = "%[3]s 2"
}

data "twilio_autopilot_task_samples" "task_samples" {
  assistant_sid = twilio_autopilot_task.task.assistant_sid
  task_sid      = twilio_autopilot_task.task.sid
	language      = "%[2]s"

	depends_on = [
		twilio_autopilot_task_sample.task_sample,
		twilio_autopilot_task_sample.task_sample_2,
	]
}
`, uniqueName, language, taggedText)
}

func testAccDataSourceTwilioAutopilotTaskSamples_invalidAssistantSid() string {
	return `
data "twilio_autopilot_task_samples" "task_samples" {
  assistant_sid = "assistant_sid"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioAutopilotTaskSamples_invalidTaskSid() string {
	return `
data "twilio_autopilot_task_samples" "task_samples" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "task_sid"
}
`
}
