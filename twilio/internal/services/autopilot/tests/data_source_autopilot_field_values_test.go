package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var fieldValuesDataSourceName = "twilio_autopilot_field_values"

func TestAccDataSourceTwilioAutopilotFieldValues_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.field_values", fieldValuesDataSourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	value := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotFieldValues_basic(uniqueName, language, value),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_type_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_values.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_values.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_values.0.language", language),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_values.0.value", value),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_values.0.synonym_of", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_values.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_values.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_values.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotFieldValues_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotFieldValues_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotFieldValues_invalidFieldTypeSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotFieldValues_invalidFieldTypeSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of field_type_sid to match regular expression "\^UB\[0-9a-fA-F\]\{32\}\$", got field_type_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotFieldValues_basic(uniqueName string, language string, value string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}

resource "twilio_autopilot_field_value" "field_value" {
  assistant_sid  = twilio_autopilot_assistant.assistant.sid
  field_type_sid = twilio_autopilot_field_type.field_type.sid
  language       = "%[2]s"
  value          = "%[3]s"
}

data "twilio_autopilot_field_values" "field_values" {
  assistant_sid  = twilio_autopilot_field_value.field_value.assistant_sid
  field_type_sid = twilio_autopilot_field_value.field_value.field_type_sid
}
`, uniqueName, language, value)
}

func testAccDataSourceTwilioAutopilotFieldValues_invalidAssistantSid() string {
	return `
data "twilio_autopilot_field_values" "field_values" {
  assistant_sid  = "assistant_sid"
  field_type_sid = "UBaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioAutopilotFieldValues_invalidFieldTypeSid() string {
	return `
data "twilio_autopilot_field_values" "field_values" {
  assistant_sid  = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  field_type_sid = "field_type_sid"
}
`
}
