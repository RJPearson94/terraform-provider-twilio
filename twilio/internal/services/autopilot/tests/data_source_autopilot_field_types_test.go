package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var fieldTypesDataSourceName = "twilio_autopilot_field_types"

func TestAccDataSourceTwilioAutopilotFieldTypes_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.field_types", fieldTypesDataSourceName)
	uniqueName := acctest.RandString(10)
	fieldTypeFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotFieldTypes_basic(uniqueName, fieldTypeFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_types.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_types.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_types.0.unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "field_types.0.friendly_name", fieldTypeFriendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_types.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_types.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "field_types.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotFieldTypes_basic(uniqueName string, fieldTypeFriendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
  friendly_name = "%s"
}

data "twilio_autopilot_field_types" "field_types" {
  assistant_sid = twilio_autopilot_field_type.field_type.assistant_sid
}
`, uniqueName, uniqueName, fieldTypeFriendlyName)
}
