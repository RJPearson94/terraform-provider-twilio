package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var fieldValueResourceName = "twilio_autopilot_field_value"

func TestAccTwilioAutopilotFieldValue_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.field_value", fieldValueResourceName)
	uniqueName := acctest.RandString(10)
	language := "en-US"
	value := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotFieldValueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotFieldValue_basic(uniqueName, language, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotFieldValueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "language", language),
					resource.TestCheckResourceAttr(stateResourceName, "value", value),
					resource.TestCheckResourceAttr(stateResourceName, "synonym_of", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "field_type_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotFieldValueImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTwilioAutopilotFieldValueDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != fieldValueResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).FieldType(rs.Primary.Attributes["field_type_sid"]).FieldValue(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving field value information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotFieldValueExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).FieldType(rs.Primary.Attributes["field_type_sid"]).FieldValue(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving field value information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAutopilotFieldValueImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s/FieldTypes/%s/FieldValues/%s", rs.Primary.Attributes["assistant_sid"], rs.Primary.Attributes["field_type_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotFieldValue_basic(uniqueName string, language string, value string) string {
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
`, uniqueName, uniqueName, language, value)
}
