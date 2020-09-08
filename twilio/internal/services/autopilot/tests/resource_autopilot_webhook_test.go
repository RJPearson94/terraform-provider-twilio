package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var webhookResourceName = "twilio_autopilot_webhook"

func TestAccTwilioAutopilotWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/webhook"
	events := []string{"onDialogueStart", "onDialogueEnd"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		Providers:         acceptance.TestAccProviders,
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAutopilotWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotWebhook_basic(uniqueName, url, events),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttr(stateResourceName, "events.1", "onDialogueEnd"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", url),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "webhook_method"),
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

func TestAccTwilioAutopilotWebhook_invalidWebhookURL(t *testing.T) {
	uniqueName := acctest.RandString(10)
	url := "webhookURL"
	events := []string{"onDialogueStart"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_basic(uniqueName, url, events),
				ExpectError: regexp.MustCompile(`(?s)expected "webhook_url" to have a host, got webhookURL`),
			},
		},
	})
}

func TestAccTwilioAutopilotAssistant_invalidWebhookMethod(t *testing.T) {
	uniqueName := acctest.RandString(10)
	method := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAutopilotAssistantDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotWebhook_method(uniqueName, method),
				ExpectError: regexp.MustCompile(`(?s)expected webhook_method to be one of \[GET POST\], got test`),
			},
		},
	})
}

func TestAccTwilioAutopilotWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	uniqueName := acctest.RandString(10)
	newUniqueName := acctest.RandString(10)
	url := "http://localhost/webhook"
	events := []string{"onDialogueStart"}
	newEvents := []string{"onDialogueStart", "onDialogueEnd"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAutopilotWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotWebhook_basic(uniqueName, url, events),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", url),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "webhook_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotWebhook_basic(newUniqueName, url, newEvents),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "events.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "events.0", "onDialogueStart"),
					resource.TestCheckResourceAttr(stateResourceName, "events.1", "onDialogueEnd"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", url),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "webhook_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
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
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err)
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
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err)
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

func testAccTwilioAutopilotWebhook_basic(uniqueName string, url string, events []string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
  webhook_url   = "%s"
  events        = %s
}
`, uniqueName, uniqueName, url, "[\""+strings.Join(events, "\",\"")+"\"]")
}

func testAccTwilioAutopilotWebhook_method(uniqueName string, method string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid  = twilio_autopilot_assistant.assistant.sid
  unique_name    = "%s"
  webhook_url    = "http://localhost/webhook"
  webhook_method = "%s"
  events = [
    "onDialogueStart",
    "onDialogueEnd"
  ]
}
`, uniqueName, uniqueName, method)
}
