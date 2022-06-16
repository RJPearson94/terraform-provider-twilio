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

var taskSampleResourceName = "twilio_autopilot_task_sample"

func TestAccTwilioAutopilotTaskSample_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_sample", taskSampleResourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	taggedText := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskSampleDestroy,
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
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", "voice"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotTaskSampleImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
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
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskSampleDestroy,
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
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", "voice"),
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
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", "voice"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskSample_sourceChannel(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_sample", taskSampleResourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	taggedText := "test"
	sourceChannel := "chat"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskSampleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTaskSample_sourceChannel(uniqueName, language, taggedText, sourceChannel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskSampleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", sourceChannel),
				),
			},
			{
				Config: testAccTwilioAutopilotTaskSample_basic(uniqueName, language, taggedText),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskSampleExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "source_channel", "voice"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskSample_blankLanguage(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskSample_blankLanguage(),
				ExpectError: regexp.MustCompile(`(?s)expected \"language\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskSample_blankTaggedText(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskSample_blankTaggedText(),
				ExpectError: regexp.MustCompile(`(?s)expected \"tagged_text\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskSample_invalidSourceChannel(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskSample_invalidSourceChannel(),
				ExpectError: regexp.MustCompile(`(?s)expected source_channel to be one of \[voice sms chat alexa google-assistant slack\], got test`),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskSample_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskSample_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func TestAccTwilioAutopilotTaskSample_invalidTaskSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTaskSample_invalidTaskSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of task_sid to match regular expression "\^UD\[0-9a-fA-F\]\{32\}\$", got task_sid`),
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

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.Attributes["task_sid"]).Sample(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving task sample information %s", err.Error())
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

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.Attributes["task_sid"]).Sample(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving task sample information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAutopilotTaskSampleImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s/Tasks/%s/Samples/%s", rs.Primary.Attributes["assistant_sid"], rs.Primary.Attributes["task_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotTaskSample_basic(uniqueName string, language string, taggedText string) string {
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
`, uniqueName, language, taggedText)
}

func testAccTwilioAutopilotTaskSample_sourceChannel(uniqueName string, language string, taggedText string, sourceChannel string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}

resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid  = twilio_autopilot_assistant.assistant.sid
  task_sid       = twilio_autopilot_task.task.sid
  language       = "%[2]s"
  tagged_text    = "%[3]s"
  source_channel = "%[4]s"
}
`, uniqueName, language, taggedText, sourceChannel)
}

func testAccTwilioAutopilotTaskSample_invalidAssistantSid() string {
	return `
resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = "assistant_sid"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  language      = "en-US"
  tagged_text   = "hi"
}
`
}

func testAccTwilioAutopilotTaskSample_invalidTaskSid() string {
	return `
resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "task_sid"
  language      = "en-US"
  tagged_text   = "hi"
}
`
}

func testAccTwilioAutopilotTaskSample_blankLanguage() string {
	return `
resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  language      = ""
  tagged_text   = "hi"
}
`
}

func testAccTwilioAutopilotTaskSample_blankTaggedText() string {
	return `
resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  language      = "en-US"
  tagged_text   = ""
}
`
}

func testAccTwilioAutopilotTaskSample_invalidSourceChannel() string {
	return `
resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  task_sid      = "UDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  language      = "en-US"
  tagged_text   = "hi"
  source_channel = "test"
}
`
}
