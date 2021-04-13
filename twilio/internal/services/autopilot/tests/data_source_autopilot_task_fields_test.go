package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var taskFieldsDataSourceName = "twilio_autopilot_task_fields"

func TestAccDataSourceTwilioAutopilotTaskFields_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_fields", taskFieldsDataSourceName)
	uniqueName := acctest.RandString(10)
	fieldType := "Twilio.YES_NO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotTaskFields_basic(uniqueName, fieldType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "fields.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "fields.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "fields.0.unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "fields.0.field_type", fieldType),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "fields.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "fields.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "fields.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskFields_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTaskFields_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskFields_invalidTaskSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTaskFields_invalidTaskSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of task_sid to match regular expression "\^UD\[0-9a-fA-F\]\{32\}\$", got task_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotTaskFields_basic(uniqueName string, fieldType string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}

resource "twilio_autopilot_task_field" "task_field" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  task_sid      = twilio_autopilot_task.task.sid
  unique_name   = "%[1]s"
  field_type    = "%[2]s"
}

data "twilio_autopilot_task_fields" "task_fields" {
  assistant_sid = twilio_autopilot_task_field.task_field.assistant_sid
  task_sid      = twilio_autopilot_task_field.task_field.task_sid
}
`, uniqueName, fieldType)
}

func testAccDataSourceTwilioAutopilotTaskFields_invalidAssistantSid() string {
	return `
data "twilio_autopilot_task_fields" "task_fields" {
  assistant_sid = "assistant_sid"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioAutopilotTaskFields_invalidTaskSid() string {
	return `
data "twilio_autopilot_task_fields" "task_fields" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "task_sid"
}
`
}
