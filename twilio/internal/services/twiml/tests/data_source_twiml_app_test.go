package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const appDataSourceName = "twilio_twiml_app"

func TestAccDataSourceTwilioTwimlApp_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.app", appDataSourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTwimlApp_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.0.url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.0.method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.0.caller_id_lookup", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTwimlApp_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTwimlApp_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioTwimlApp_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTwimlApp_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^AP\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTwimlApp_basic(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
}

data "twilio_twiml_app" "app" {
  account_sid = twilio_twiml_app.app.account_sid
  sid         = twilio_twiml_app.app.sid
}
`, testData.AccountSid)
}

func testAccDataSourceTwilioTwimlApp_invalidAccountSid() string {
	return `
data "twilio_twiml_app" "app" {
  account_sid = "account_sid"
  sid         = "APaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioTwimlApp_invalidSid() string {
	return `
data "twilio_twiml_app" "app" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
