package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var modelBuildDataSourceName = "twilio_autopilot_model_build"

func TestAccDataSourceTwilioAutopilotModelBuild_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.model_build", modelBuildDataSourceName)
	uniqueName := acctest.RandString(10)
	modelBuildUniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotModelBuild_basic(uniqueName, modelBuildUniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", modelBuildUniqueName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "build_duration"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "error_code"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotModelBuild_basic(uniqueName string, modelBuildUniqueName string) string {
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
  language      = "en-US"
  tagged_text   = "test"
}

resource "twilio_autopilot_model_build" "model_build" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"

  depends_on = [twilio_autopilot_task_sample.task_sample]
}

data "twilio_autopilot_model_build" "model_build" {
  assistant_sid = twilio_autopilot_model_build.model_build.assistant_sid
  sid           = twilio_autopilot_model_build.model_build.sid
}
`, uniqueName, uniqueName, modelBuildUniqueName)
}
