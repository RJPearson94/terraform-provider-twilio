package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var taskFieldDataSourceName = "twilio_autopilot_task_field"

func TestAccDataSourceTwilioAutopilotTaskField_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_field", taskFieldDataSourceName)
	uniqueName := acctest.RandString(10)
	fieldType := "Twilio.YES_NO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotTaskField_basic(uniqueName, fieldType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_type", fieldType),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskField_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTaskField_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskField_invalidTaskSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTaskField_invalidTaskSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of task_sid to match regular expression "\^UD\[0-9a-fA-F\]\{32\}\$", got task_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotTaskField_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotTaskField_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^UE\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotTaskField_basic(uniqueName string, fieldType string) string {
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

data "twilio_autopilot_task_field" "task_field" {
  assistant_sid = twilio_autopilot_task_field.task_field.assistant_sid
  task_sid      = twilio_autopilot_task_field.task_field.task_sid
  sid           = twilio_autopilot_task_field.task_field.sid
}
`, uniqueName, fieldType)
}

func testAccDataSourceTwilioAutopilotTaskField_invalidAssistantSid() string {
	return `
data "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "assistant_sid"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid           = "UEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioAutopilotTaskField_invalidTaskSid() string {
	return `
data "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "task_sid"
  sid           = "UEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioAutopilotTaskField_invalidSid() string {
	return `
data "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid           = "sid"
}
`
}
