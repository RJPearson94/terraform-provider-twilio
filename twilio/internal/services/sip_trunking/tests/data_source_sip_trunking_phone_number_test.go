package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumberDataSourceName = "twilio_sip_trunking_phone_number"

func TestAccDataSourceTwilioSIPTrunkingPhoneNumber_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_number", phoneNumberDataSourceName)
	testData := acceptance.TestAccData

	// Run tests in sequence to prevent the same value 2 SIP trunks mutating the same phone number simultaneously
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingPhoneNumber_complete(testData),
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
					resource.TestCheckNoResourceAttr(stateDataSourceName, "fax"),
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

func testAccTwilioSIPTrunkingPhoneNumber_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = twilio_sip_trunking_trunk.trunk.sid
  sid       = "%s"
}

data "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = twilio_sip_trunking_phone_number.phone_number.trunk_sid
  sid       = twilio_sip_trunking_phone_number.phone_number.sid
}
`, testData.PhoneNumberSid)
}
