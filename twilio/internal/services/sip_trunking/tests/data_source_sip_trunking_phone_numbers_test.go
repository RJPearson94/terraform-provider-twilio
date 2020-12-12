package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumbersDataSourceName = "twilio_sip_trunking_phone_numbers"

func TestAccDataSourceTwilioSIPTrunkingPhoneNumbers_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_numbers", phoneNumbersDataSourceName)
	testData := acceptance.TestAccData

	// Run tests in sequence to prevent the same value 2 SIP trunks mutating the same phone number simultaneously
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingPhoneNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingPhoneNumbers_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.phone_number"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.address_requirements"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.beta"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.capabilities.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.messaging.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.voice.#", "1"),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "phone_numbers.0.fax"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.status_callback_method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.url"),
				),
			},
		},
	})
}

func testAccTwilioSIPTrunkingPhoneNumbers_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = twilio_sip_trunking_trunk.trunk.sid
  sid       = "%s"
}

data "twilio_sip_trunking_phone_numbers" "phone_numbers" {
  trunk_sid = twilio_sip_trunking_phone_number.phone_number.trunk_sid
}
`, testData.PhoneNumberSid)
}
