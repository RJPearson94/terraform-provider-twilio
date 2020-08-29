package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var fieldTypeDataSourceName = "twilio_autopilot_field_type"

func TestAccDataSourceTwilioAutopilotFieldType_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.field_type", fieldTypeDataSourceName)
	uniqueName := acctest.RandString(10)
	fieldTypeFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotFieldType_basic(uniqueName, fieldTypeFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", fieldTypeFriendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotFieldType_basic(uniqueName string, fieldTypeFriendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
  friendly_name = "%s"
}

data "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_field_type.field_type.assistant_sid
  sid           = twilio_autopilot_field_type.field_type.sid
}
`, uniqueName, uniqueName, fieldTypeFriendlyName)
}
