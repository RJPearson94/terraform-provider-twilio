package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const appsDataSourceName = "twilio_twiml_apps"

func TestAccDataSourceTwilioTwimlApps_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.apps", appsDataSourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTwimlApps_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "apps.#", regexp.MustCompile("^[1-9]([0-9]+)?$")),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTwimlApps_friendlyName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.apps", appsDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTwimlApps_friendlyName(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "apps.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.messaging.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.messaging.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.messaging.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.messaging.0.url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.messaging.0.method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.0.caller_id_lookup", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.0.url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "apps.0.voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "apps.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "apps.0.date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTwimlApps_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTwimlApps_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTwimlApps_basic(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
}

data "twilio_twiml_apps" "apps" {
  account_sid = twilio_twiml_app.app.account_sid
}
`, testData.AccountSid)
}

func testAccDataSourceTwilioTwimlApps_friendlyName(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid   = "%[1]s"
  friendly_name = "%[2]s"
}

data "twilio_twiml_apps" "apps" {
  account_sid   = twilio_twiml_app.app.account_sid
  friendly_name = "%[2]s"
}
`, testData.AccountSid, friendlyName)
}

func testAccDataSourceTwilioTwimlApps_invalidAccountSid() string {
	return `
data "twilio_twiml_apps" "apps" {
  account_sid = "account_sid"
}
`
}
