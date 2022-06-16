package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const phoneNumberDataSourceName = "twilio_sip_trunking_phone_number"

func TestAccDataSourceTwilioSIPTrunkingPhoneNumber_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_number", phoneNumberDataSourceName)
	testData := acceptance.TestAccData

	// Run tests in sequence to prevent the same value 2 SIP trunks mutating the same phone number simultaneously
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPTrunkingPhoneNumber_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "address_requirements"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "beta"),
					resource.TestCheckResourceAttr(stateDataSourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "fax.#", "0"),
					resource.TestCheckResourceAttr(stateDataSourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status_callback_method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPTrunkingPhoneNumber_invalidTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingPhoneNumber_invalidTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of trunk_sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got trunk_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPTrunkingPhoneNumber_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingPhoneNumber_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^PN\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPTrunkingPhoneNumber_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid        = twilio_sip_trunking_trunk.trunk.sid
  phone_number_sid = "%s"
}

data "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = twilio_sip_trunking_phone_number.phone_number.trunk_sid
  sid       = twilio_sip_trunking_phone_number.phone_number.sid
}
`, testData.PhoneNumberSid)
}

func testAccDataSourceTwilioSIPTrunkingPhoneNumber_invalidTrunkSid() string {
	return `
data "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = "trunk_sid"
  sid       = "PNaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPTrunkingPhoneNumber_invalidSid() string {
	return `
data "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = "TKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid       = "sid"
}
`
}
