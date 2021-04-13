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

var assistantResourceName = "twilio_autopilot_assistant"

func TestAccTwilioAutopilotAssistant_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_default(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "latest_model_build_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "unique_name"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_events.#", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "log_queries", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "development_stage"),
					resource.TestCheckResourceAttrSet(stateResourceName, "needs_model_build"),
					resource.TestCheckResourceAttrSet(stateResourceName, "defaults"),
					resource.TestCheckResourceAttrSet(stateResourceName, "stylesheet"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotAssistantImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)
	unqiueName := ""
	newUnqiueName := acctest.RandString(64)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_uniqueName(unqiueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "unique_name"), // Twilio will generate the unique name if you don't supply a unique name
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "latest_model_build_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "callback_events.#", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "log_queries", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "development_stage"),
					resource.TestCheckResourceAttrSet(stateResourceName, "needs_model_build"),
					resource.TestCheckResourceAttrSet(stateResourceName, "defaults"),
					resource.TestCheckResourceAttrSet(stateResourceName, "stylesheet"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotAssistant_uniqueName(newUnqiueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUnqiueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "latest_model_build_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "callback_events.#", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "log_queries", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "development_stage"),
					resource.TestCheckResourceAttrSet(stateResourceName, "needs_model_build"),
					resource.TestCheckResourceAttrSet(stateResourceName, "defaults"),
					resource.TestCheckResourceAttrSet(stateResourceName, "stylesheet"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)
	friendlyName := ""
	newFriendlyName := acctest.RandString(255)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_friendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioAutopilotAssistant_friendlyName(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
			{
				Config: testAccTwilioAutopilotAssistant_default(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_developmentStage(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)
	friendlyName := acctest.RandString(10)
	developmentStage := "in-production"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_developmentStage(friendlyName, developmentStage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "development_stage", developmentStage),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_invalidDevelopmentStage(t *testing.T) {
	friendlyName := acctest.RandString(10)
	developmentStage := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotAssistant_developmentStage(friendlyName, developmentStage),
				ExpectError: regexp.MustCompile(`(?s)expected development_stage to be one of \[in-development in-production\], got test`),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_defaults(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_defaults(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "defaults", "{\"defaults\":{\"assistant_initiation\":\"\",\"fallback\":\"http://localhost/fallback\"}}"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_invalidDefaults(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotAssistant_invalidDefaultsString(),
				ExpectError: regexp.MustCompile(`(?s)"defaults" contains an invalid JSON`),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_stylesheet(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_stylesheet(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "stylesheet", "{\"style_sheet\":{\"voice\":{\"say_voice\":\"Polly.Matthew\"}}}"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_invalidStylesheet(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotAssistant_invalidStylesheetString(),
				ExpectError: regexp.MustCompile(`(?s)"stylesheet" contains an invalid JSON`),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_callbackEvents(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)
	friendlyName := acctest.RandString(10)
	callbackURL := "http://localhost.com/callback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_callbackEvents(friendlyName, callbackURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "callback_events.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_events.0", "model_build_completed"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_events.1", "model_build_failed"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", callbackURL),
				),
			},
			{
				Config: testAccTwilioAutopilotAssistant_default(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "callback_events.#", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_invalidCallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	callbackURL := "callbackURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotAssistant_callbackEvents(friendlyName, callbackURL),
				ExpectError: regexp.MustCompile(`(?s)expected "callback_url" to have a host, got callbackURL`),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_logQueries(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.assistant", assistantResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotAssistant_logQueriesFalse(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "log_queries", "false"),
				),
			},
			{
				Config: testAccTwilioAutopilotAssistant_default(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotAssistantExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "log_queries", "true"),
				),
			},
		},
	})
}

func testAccCheckTwilioAutopilotAssistantDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != assistantResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving assistant information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotAssistantExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving assistant information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAutopilotAssistantImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotAssistant_default() string {
	return `resource "twilio_autopilot_assistant" "assistant" {}`
}

func testAccTwilioAutopilotAssistant_friendlyName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "%s"
}
`, friendlyName)
}

func testAccTwilioAutopilotAssistant_uniqueName(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}
`, uniqueName)
}

func testAccTwilioAutopilotAssistant_developmentStage(friendlyName string, developmentStage string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name     = "%s"
  development_stage = "%s"
}
`, friendlyName, developmentStage)
}

func testAccTwilioAutopilotAssistant_defaults(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "%s"
  defaults = jsonencode({
    "defaults" : {
      "assistant_initiation" : "",
      "fallback" : "http://localhost/fallback"
    }
  })
}
`, friendlyName)
}

func testAccTwilioAutopilotAssistant_invalidDefaultsString() string {
	return `
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "invalid_defaults_json_string"
  defaults      = "defaults"
}
`
}

func testAccTwilioAutopilotAssistant_stylesheet(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "%s"
  stylesheet = jsonencode({
    "style_sheet" : {
      "voice" : {
        "say_voice" : "Polly.Matthew"
      }
    }
  })
}
`, friendlyName)
}

func testAccTwilioAutopilotAssistant_invalidStylesheetString() string {
	return `
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "invalid_stylesheet_json_string"
  stylesheet    = "stylesheet"
}
`
}

func testAccTwilioAutopilotAssistant_callbackEvents(friendlyName string, callbackURL string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "%s"
  callback_url  = "%s"
  callback_events = [
    "model_build_completed",
    "model_build_failed"
  ]
}
`, friendlyName, callbackURL)
}

func testAccTwilioAutopilotAssistant_logQueriesFalse() string {
	return `
resource "twilio_autopilot_assistant" "assistant" {
	log_queries = false
}`
}
