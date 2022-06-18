package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var webhookResourceName = "twilio_autopilot_webhook"

func TestAccTwilioAutopilotWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotWebhook_basic(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotWebhookImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	uniqueName := acctest.RandString(1)
	newUniqueName := acctest.RandString(64)
	url := "http://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotWebhook_basic(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotWebhook_basic(newUniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_invalidUniqueNameWith0Characters(t *testing.T) {
	uniqueName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_withStubbedAssistantSid(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got `),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_invalidUniqueNameWith65Characters(t *testing.T) {
	uniqueName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_withStubbedAssistantSid(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_events(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	uniqueName := acctest.RandString(1)
	url := "http://localhost.com/webhook"
	events := []string{"onDialogueStart", "onDialogueEnd"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotWebhook_events(uniqueName, url, events),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttr(stateResourceName, "events.1", "onDialogueEnd"),
				),
			},
			{
				Config: testAccTwilioAutopilotWebhook_basic(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_invalidEvent(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_invalidEvent(),
				ExpectError: regexp.MustCompile(`(?s)expected events.0 to be one of \[onDialogueStart onDialogueEnd onDialogueTaskStart onDialogueTaskEnd onDialogueTurn onCollectAttempt onActionsFetch\], got test`),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_webhookMethod(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	uniqueName := acctest.RandString(1)
	method := "GET"
	url := "http://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotWebhook_method(uniqueName, method),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", method),
				),
			},
			{
				Config: testAccTwilioAutopilotWebhook_basic(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_invalidWebhookMethod(t *testing.T) {
	uniqueName := acctest.RandString(10)
	method := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_method(uniqueName, method),
				ExpectError: regexp.MustCompile(`(?s)expected webhook_method to be one of \[GET POST\], got test`),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_invalidWebhookURL(t *testing.T) {
	uniqueName := acctest.RandString(10)
	url := "webhookURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_basic(uniqueName, url),
				ExpectError: regexp.MustCompile(`(?s)expected "webhook_url" to have a host, got webhookURL`),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_invalidAssistantSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_invalidAssistantSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
			},
		},
	})
}

func testAccCheckTwilioAutopilotWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != webhookResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotWebhookExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAutopilotWebhookImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s/Webhooks/%s", rs.Primary.Attributes["assistant_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotWebhook_basic(uniqueName string, url string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  webhook_url   = "%[2]s"
  events = [
    "onDialogueStart"
  ]
}
`, uniqueName, url)
}

func testAccTwilioAutopilotWebhook_events(uniqueName string, url string, events []string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  webhook_url   = "%[2]s"
  events        = %[3]s
}
`, uniqueName, url, `["`+strings.Join(events, `","`)+`"]`)
}

func testAccTwilioAutopilotWebhook_method(uniqueName string, method string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid  = twilio_autopilot_assistant.assistant.sid
  unique_name    = "%[1]s"
  webhook_url    = "http://localhost.com/webhook"
  webhook_method = "%[2]s"
  events = [
    "onDialogueStart",
    "onDialogueEnd"
  ]
}
`, uniqueName, method)
}

func testAccTwilioAutopilotWebhook_invalidAssistantSid() string {
	return `
resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = "assistant_sid"
  unique_name   = "invalid_account_sid"
  webhook_url   = "http://localhost.com/webhook"
  events = [
    "onDialogueStart"
  ]
}
`
}

func testAccTwilioAutopilotWebhook_withStubbedAssistantSid(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "%s"
  webhook_url   = "http://localhost.com/webhook"
  events = [
    "onDialogueStart"
  ]
}
`, uniqueName)
}

func testAccTwilioAutopilotWebhook_invalidEvent() string {
	return `
resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "invalid_event"
  webhook_url   = "http://localhost.com/webhook"
  events = [
    "test"
  ]
}
`
}
