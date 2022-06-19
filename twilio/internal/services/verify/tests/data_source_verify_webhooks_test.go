package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const webhooksDataSourceName = "twilio_verify_webhooks"

func TestAccDataSourceTwilioVerifyWebhooks_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.webhooks", webhooksDataSourceName)
	friendlyName := acctest.RandString(10)
	eventType := "*"
	webhookUrl := "https://localhost.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyWebhooks_basic(friendlyName, eventType, webhookUrl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.event_types.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.event_types.0", eventType),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.status", "enabled"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.version", "v2"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.webhook_url", webhookUrl),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks.0.webhook_method", "POST"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhooks.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyWebhooks_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyWebhooks_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyWebhooks_basic(friendlyName string, event string, webhookUrl string) string {
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

data "twilio_verify_webhooks" "webhooks" {
  service_sid = twilio_verify_webhook.webhook.service_sid
}
`, friendlyName, event, webhookUrl)
}

func testAccDataSourceTwilioVerifyWebhooks_invalidServiceSid() string {
	return `
data "twilio_verify_webhooks" "webhooks" {
  service_sid = "service_sid"
}
`
}
