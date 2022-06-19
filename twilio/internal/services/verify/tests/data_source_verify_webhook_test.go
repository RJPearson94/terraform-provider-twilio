package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const webhookDataSourceName = "twilio_verify_webhook"

func TestAccDataSourceTwilioVerifyWebhook_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhook", webhookDataSourceName)
	friendlyName := acctest.RandString(10)
	eventType := "*"
	webhookUrl := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyWebhook_basic(friendlyName, eventType, webhookUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "event_types.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "event_types.0", eventType),
					resource.TestCheckResourceAttr(stateDataSourceName, "status", "enabled"),
					resource.TestCheckResourceAttr(stateDataSourceName, "version", "v2"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhook_url", webhookUrl),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyWebhook_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyWebhook_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^YW\[0-9a-fA-F\]\{32\}\$", got webhook_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyWebhook_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyWebhook_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyWebhook_basic(friendlyName string, event string, webhookUrl string) string {
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

data "twilio_verify_webhook" "webhook" {
  sid         = twilio_verify_webhook.webhook.sid
  service_sid = twilio_verify_webhook.webhook.service_sid
}


`, friendlyName, event, webhookUrl)
}

func testAccDataSourceTwilioVerifyWebhook_invalidSid() string {
	return `
data "twilio_verify_webhook" "webhook" {
  sid         = "webhook_sid"
  service_sid = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioVerifyWebhook_invalidServiceSid() string {
	return `
data "twilio_verify_webhook" "webhook" {
  sid         = "YWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  service_sid = "service_sid"
}
`
}
