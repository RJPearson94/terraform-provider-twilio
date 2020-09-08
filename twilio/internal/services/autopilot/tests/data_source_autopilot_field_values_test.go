package tests

import (
	"fmt"
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

func testAccDataSourceTwilioAutopilotFieldValues_basic(uniqueName string, language string, value string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
}

resource "twilio_autopilot_field_value" "field_value" {
  assistant_sid  = twilio_autopilot_assistant.assistant.sid
  field_type_sid = twilio_autopilot_field_type.field_type.sid
  language       = "%s"
  value          = "%s"
}

data "twilio_autopilot_field_values" "field_values" {
  assistant_sid  = twilio_autopilot_field_value.field_value.assistant_sid
  field_type_sid = twilio_autopilot_field_value.field_value.field_type_sid
}
`, uniqueName, uniqueName, language, value)
}
