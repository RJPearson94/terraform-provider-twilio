package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const serviceDataSourceName = "twilio_verify_service"

func TestAccDataSourceTwilioVerifyService_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service", serviceDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "code_length", "6"),
					resource.TestCheckResourceAttr(stateDataSourceName, "custom_code_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "do_not_share_warning_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "dtmf_input_required", "true"),
					resource.TestCheckResourceAttr(stateDataSourceName, "lookup_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "mailer_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "psd2_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "push.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "push.0.apn_credential_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "push.0.fcm_credential_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "totp.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "totp.0.issuer"),
					resource.TestCheckResourceAttr(stateDataSourceName, "totp.0.time_step", "30"),
					resource.TestCheckResourceAttr(stateDataSourceName, "totp.0.code_length", "6"),
					resource.TestCheckResourceAttr(stateDataSourceName, "totp.0.skew", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "default_template_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "skip_sms_to_landlines", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "tts_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyService_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyService_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyService_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

data "twilio_verify_service" "service" {
  sid = twilio_verify_service.service.sid
}
`, friendlyName)
}

func testAccDataSourceTwilioVerifyService_invalidSid() string {
	return `
data "twilio_verify_service" "service" {
  sid = "service_sid"
}
`
}
