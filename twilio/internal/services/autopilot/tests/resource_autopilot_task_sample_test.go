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

var taskSampleResourceName = "twilio_autopilot_task_sample"

func TestAccTwilioAutopilotTaskSample_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_sample", taskSampleResourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	taggedText := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioAutopilotTaskSampleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTaskSample_basic(uniqueName, language, taggedText),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskSampleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "language", language),
					resource.TestCheckResourceAttr(stateResourceName, "tagged_text", taggedText),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "task_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskSample_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_sample", taskSampleResourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	taggedText := "test"
	newTaggedText := "new test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioAutopilotTaskSampleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTaskSample_basic(uniqueName, language, taggedText),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskSampleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "language", language),
					resource.TestCheckResourceAttr(stateResourceName, "tagged_text", taggedText),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "task_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotTaskSample_basic(uniqueName, language, newTaggedText),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskSampleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "language", language),
					resource.TestCheckResourceAttr(stateResourceName, "tagged_text", newTaggedText),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "task_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioAutopilotTaskSampleDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != taskSampleResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.Attributes["task_sid"]).Sample(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving task sample information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotTaskSampleExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.Attributes["task_sid"]).Sample(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving task sample information %s", err)
		}

		return nil
	}
}

func testAccTwilioAutopilotTaskSample_basic(uniqueName string, language string, taggedText string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
	unique_name = "%s"
}

resource "twilio_autopilot_task" "task" {
	assistant_sid = twilio_autopilot_assistant.assistant.sid
	unique_name = "%s"
}

resource "twilio_autopilot_task_sample" "task_sample" {
	assistant_sid = twilio_autopilot_assistant.assistant.sid
	task_sid = twilio_autopilot_task.task.sid
	language = "%s"
	tagged_text = "%s"
}`, uniqueName, uniqueName, language, taggedText)
}
