package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var taskFieldResourceName = "twilio_autopilot_task_field"

func TestAccTwilioAutopilotTaskField_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_field", taskFieldResourceName)
	uniqueName := acctest.RandString(10)
	fieldType := "Twilio.YES_NO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskFieldDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTaskField_basic(uniqueName, fieldType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskFieldExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "field_type", fieldType),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "task_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotTaskFieldImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAutopilotTaskField_uniqueName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_field", taskFieldResourceName)
	uniqueName := acctest.RandString(1)
	newUniqueName := acctest.RandString(64)
	fieldType := "Twilio.YES_NO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskFieldDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTaskField_basic(uniqueName, fieldType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskFieldExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
				),
			},
			{
				Config: testAccTwilioAutopilotTaskField_basic(newUniqueName, fieldType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskFieldExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskField_invalidUniqueNameWith0Characters(t *testing.T) {
	uniqueName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskField_withStubbedAssistantAndTaskSids(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got `),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskField_invalidUniqueNameWith65Characters(t *testing.T) {
	uniqueName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskField_withStubbedAssistantAndTaskSids(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskField_blankFieldType(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskField_blankFieldType(),
				ExpectError: regexp.MustCompile(`(?s)expected \"field_type\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskField_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskField_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskField_invalidTaskSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskField_invalidTaskSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of task_sid to match regular expression "\^UD\[0-9a-fA-F\]\{32\}\$", got task_sid`),
			},
		},
	})
}

func testAccCheckTwilioAutopilotTaskFieldDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != taskFieldResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.Attributes["task_sid"]).Field(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving task field information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotTaskFieldExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.Attributes["task_sid"]).Field(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving task field information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAutopilotTaskFieldImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s/Tasks/%s/Fields/%s", rs.Primary.Attributes["assistant_sid"], rs.Primary.Attributes["task_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotTaskField_basic(uniqueName string, fieldType string) string {
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
`, uniqueName, fieldType)
}

func testAccTwilioAutopilotTaskField_withStubbedAssistantAndTaskSids(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "%s"
  field_type    = "Twilio.YES_NO"
}
`, uniqueName)
}

func testAccTwilioAutopilotTaskField_invalidAssistantSid() string {
	return `
resource "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "assistant_sid"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "invalid_assistant_sid"
  field_type    = "Twilio.YES_NO"
}
`
}

func testAccTwilioAutopilotTaskField_invalidTaskSid() string {
	return `
resource "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "task_sid"
  unique_name   = "invalid_task_sid"
  field_type    = "Twilio.YES_NO"
}
`
}

func testAccTwilioAutopilotTaskField_blankFieldType() string {
	return `
resource "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "invalid_field_type"
  field_type    = ""
}
`
}
