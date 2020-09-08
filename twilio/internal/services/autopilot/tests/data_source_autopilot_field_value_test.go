package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var fieldValueDataSourceName = "twilio_autopilot_field_value"

func TestAccDataSourceTwilioAutopilotFieldValue_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.field_value", fieldValueDataSourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	value := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotFieldValue_basic(uniqueName, language, value),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "language", language),
					resource.TestCheckResourceAttr(stateDataSourceName, "value", value),
					resource.TestCheckResourceAttr(stateDataSourceName, "synonym_of", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_type_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotFieldValue_basic(uniqueName string, language string, value string) string {
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

data "twilio_autopilot_field_value" "field_value" {
  assistant_sid  = twilio_autopilot_field_value.field_value.assistant_sid
  field_type_sid = twilio_autopilot_field_value.field_value.field_type_sid
  sid            = twilio_autopilot_field_value.field_value.sid
}
`, uniqueName, uniqueName, language, value)
}
