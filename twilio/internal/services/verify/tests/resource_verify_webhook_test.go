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

var webhookResourceName = "twilio_verify_webhook"

func TestAccTwilioVerifyWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	friendlyName := acctest.RandString(10)
	eventType := "*"
	webhookUrl := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyWebhook_basic(friendlyName, eventType, webhookUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_types.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "event_types.0", eventType),
					resource.TestCheckResourceAttr(stateResourceName, "status", "enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "version", "v2"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrl),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioVerifyWebhookImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)
	eventType := "*"
	webhookUrl := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyWebhook_basic(friendlyName, eventType, webhookUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_types.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "event_types.0", eventType),
					resource.TestCheckResourceAttr(stateResourceName, "status", "enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "version", "v2"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrl),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioVerifyWebhook_basic(newFriendlyName, eventType, webhookUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_types.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "event_types.0", eventType),
					resource.TestCheckResourceAttr(stateResourceName, "status", "enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "version", "v2"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrl),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_status(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	friendlyName := acctest.RandString(10)
	status := "enabled"
	newStatus := "disabled"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyWebhook_status(friendlyName, status),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status", status),
				),
			},
			{
				Config: testAccTwilioVerifyWebhook_status(friendlyName, newStatus),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status", newStatus),
				),
			},
			{
				Config: testAccTwilioVerifyWebhook_basic(friendlyName, "*", "https://localhost.com/webhook"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status", "enabled"),
				),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_version(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)
	friendlyName := acctest.RandString(10)
	version := "v2"
	newVersion := "v1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyWebhook_version(friendlyName, version),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "version", version),
				),
			},
			{
				Config: testAccTwilioVerifyWebhook_version(friendlyName, newVersion),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "version", newVersion),
				),
			},
			{
				Config: testAccTwilioVerifyWebhook_basic(friendlyName, "*", "https://localhost.com/webhook"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "version", "v2"),
				),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_blankFriendlyName(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyWebhook_blankFriendlyName(),
				ExpectError: regexp.MustCompile(`(?s)expected \"friendly_name\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_invalidWebhookUrl(t *testing.T) {
	friendlyName := acctest.RandString(10)
	eventType := "*"
	webhookUrl := "webhookUrl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyWebhook_basic(friendlyName, eventType, webhookUrl),
				ExpectError: regexp.MustCompile(`(?s)expected "webhook_url" to have a host, got webhookUrl`),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_invalidEventType(t *testing.T) {
	friendlyName := acctest.RandString(10)
	eventType := "event_type"
	webhookUrl := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyWebhook_basic(friendlyName, eventType, webhookUrl),
				ExpectError: regexp.MustCompile(`(?s)expected event_types\.0 to be one of \[\* "factor\.created" "factor\.verified" "factor\.deleted" "challenge\.approved" "challenge\.denied"\], got event_type`),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyWebhook_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccCheckTwilioVerifyWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

	for _, rs := range s.RootModule().Resources {
		if rs.Type != webhookResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioVerifyWebhookExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Webhook(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioVerifyWebhookImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Webhooks/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioVerifyWebhook_basic(friendlyName string, event string, webhookUrl string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_webhook" "webhook" {
  service_sid   = twilio_verify_service.service.sid
  friendly_name = "%[1]s"
  event_types   = ["%[2]s"]
  webhook_url   = "%[3]s"
}
`, friendlyName, event, webhookUrl)
}

func testAccTwilioVerifyWebhook_invalidServiceSid() string {
	return `
resource "twilio_verify_webhook" "webhook" {
  service_sid   = "service_sid"
  friendly_name = "invalid service sid"
  event_types   = ["*"]
  webhook_url   = "https://localhost.com/webhook"
}
`
}

func testAccTwilioVerifyWebhook_blankFriendlyName() string {
	return `
resource "twilio_verify_webhook" "webhook" {
  service_sid   = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name = ""
  event_types   = ["*"]
  webhook_url   = "https://localhost.com/webhook"
}
`
}

func testAccTwilioVerifyWebhook_status(friendlyName string, status string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_webhook" "webhook" {
  service_sid   = twilio_verify_service.service.sid
  friendly_name = "%[1]s"
  event_types   = ["*"]
  webhook_url   = "https://localhost.com/webhook"
  status        = "%[2]s"
}
`, friendlyName, status)
}

func testAccTwilioVerifyWebhook_version(friendlyName string, version string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_webhook" "webhook" {
  service_sid   = twilio_verify_service.service.sid
  friendly_name = "%[1]s"
  event_types   = ["*"]
  webhook_url   = "https://localhost.com/webhook"
  version       = "%[2]s"
}
`, friendlyName, version)
}
