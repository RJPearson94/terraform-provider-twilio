package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var assistantDataSourceName = "twilio_autopilot_assistant"

func TestAccDataSourceTwilioAutopilotAssistant_sid(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.assistant", assistantDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotAssistant_sid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "latest_model_build_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "unique_name"),
					resource.TestCheckResourceAttr(stateDataSourceName, "callback_events.#", "0"),
					resource.TestCheckResourceAttr(stateDataSourceName, "callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "log_queries"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "development_stage"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "needs_model_build"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "defaults"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "stylesheet"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotAssistant_uniqueName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.assistant", assistantDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAutopilotAssistant_uniqueName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "latest_model_build_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "unique_name"),
					resource.TestCheckResourceAttr(stateDataSourceName, "callback_events.#", "0"),
					resource.TestCheckResourceAttr(stateDataSourceName, "callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "log_queries"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "development_stage"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "needs_model_build"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "defaults"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "stylesheet"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAutopilotAssistant_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAutopilotAssistant_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAutopilotAssistant_sid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "%s"
}

data "twilio_autopilot_assistant" "assistant" {
  sid = twilio_autopilot_assistant.assistant.sid
}
`, friendlyName)
}

func testAccDataSourceTwilioAutopilotAssistant_uniqueName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "%[1]s"
  unique_name   = "%[1]s"
}

data "twilio_autopilot_assistant" "assistant" {
  unique_name = twilio_autopilot_assistant.assistant.unique_name
}
`, friendlyName)
}

func testAccDataSourceTwilioAutopilotAssistant_invalidSid() string {
	return `
data "twilio_autopilot_assistant" "assistant" {
  sid = "sid"
}
`
}
