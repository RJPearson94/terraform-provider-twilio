package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var modelBuildResourceName = "twilio_autopilot_model_build"

func TestAccTwilioAutopilotModelBuild_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.model_build", modelBuildResourceName)
	uniqueName := acctest.RandString(10)
	modelBuildUniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		Providers:         acceptance.TestAccProviders,
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAutopilotModelBuildDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotModelBuild_basic(uniqueName, modelBuildUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotModelBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", modelBuildUniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_duration"),
					resource.TestCheckResourceAttrSet(stateResourceName, "status"),
					resource.TestCheckResourceAttrSet(stateResourceName, "error_code"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotModelBuildImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAutopilotModelBuild_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.model_build", modelBuildResourceName)
	uniqueName := acctest.RandString(10)
	modelBuildUniqueName := acctest.RandString(10)
	newModelBuildUniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAutopilotModelBuildDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotModelBuild_basic(uniqueName, modelBuildUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotModelBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", modelBuildUniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_duration"),
					resource.TestCheckResourceAttrSet(stateResourceName, "status"),
					resource.TestCheckResourceAttrSet(stateResourceName, "error_code"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotModelBuild_basic(uniqueName, newModelBuildUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotModelBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newModelBuildUniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_duration"),
					resource.TestCheckResourceAttrSet(stateResourceName, "status"),
					resource.TestCheckResourceAttrSet(stateResourceName, "error_code"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioAutopilotModelBuildDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != modelBuildResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).ModelBuild(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving field type information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotModelBuildExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).ModelBuild(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving field type information %s", err)
		}

		return nil
	}
}

func testAccTwilioAutopilotModelBuildImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s/ModelBuilds/%s", rs.Primary.Attributes["assistant_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotModelBuild_basic(uniqueName string, modelBuildUniqueName string) string {
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
`, uniqueName, uniqueName, modelBuildUniqueName)
}
