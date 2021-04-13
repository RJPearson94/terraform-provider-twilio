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

var fieldTypeResourceName = "twilio_autopilot_field_type"

func TestAccTwilioAutopilotFieldType_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.field_type", fieldTypeResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotFieldTypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotFieldType_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotFieldTypeExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotFieldTypeImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAutopilotFieldType_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.field_type", fieldTypeResourceName)
	uniqueName := acctest.RandString(1)
	newUniqueName := acctest.RandString(64)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotFieldTypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotFieldType_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotFieldTypeExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotFieldType_basic(newUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotFieldTypeExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotFieldType_invalidUniqueNameWith0Characters(t *testing.T) {
	uniqueName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotFieldType_uniqueNameWithStubbedAssistantSid(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got `),
			},
		},
	})
}

func TestAccTwilioAutopilotFieldType_invalidUniqueNameWith65Characters(t *testing.T) {
	uniqueName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotFieldType_uniqueNameWithStubbedAssistantSid(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioAutopilotFieldType_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.field_type", fieldTypeResourceName)
	uniqueName := acctest.RandString(10)
	fieldTypeFriendlyName := ""
	newFieldTypeFriendlyName := acctest.RandString(255)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotFieldTypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotFieldType_friendlyName(uniqueName, fieldTypeFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotFieldTypeExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", fieldTypeFriendlyName),
				),
			},
			{
				Config: testAccTwilioAutopilotFieldType_friendlyName(uniqueName, newFieldTypeFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotFieldTypeExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFieldTypeFriendlyName),
				),
			},
			{
				Config: testAccTwilioAutopilotFieldType_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotFieldTypeExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotFieldType_invalidFriendlyNameWith256Characters(t *testing.T) {
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotFieldType_friendlyNameWithStubbedAssistantSid(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(0 - 255\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioAutopilotFieldType_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotFieldType_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func testAccCheckTwilioAutopilotFieldTypeDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != fieldTypeResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).FieldType(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving field type information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotFieldTypeExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).FieldType(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving field type information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAutopilotFieldTypeImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s/FieldTypes/%s", rs.Primary.Attributes["assistant_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotFieldType_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}
`, uniqueName)
}

func testAccTwilioAutopilotFieldType_uniqueNameWithStubbedAssistantSid(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "%s"
}
`, uniqueName)
}

func testAccTwilioAutopilotFieldType_friendlyName(uniqueName string, fieldTypeFriendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  friendly_name = "%[2]s"

}
`, uniqueName, fieldTypeFriendlyName)
}

func testAccTwilioAutopilotFieldType_friendlyNameWithStubbedAssistantSid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "friendly_name_with_stubbed_assistant_sid"
  friendly_name = "%s"
}
`, friendlyName)
}

func testAccTwilioAutopilotFieldType_invalidAssistantSid() string {
	return `
resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = "assistant_sid"
  unique_name   = "invalid_assistant_sid"
}
`
}
