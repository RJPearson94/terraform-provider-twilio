package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var proxyPhoneNumberDataSourceName = "twilio_proxy_phone_number"

// Tests have to run sequentially as a phone number cannot be associated with more than 1 proxy service at a given time

func TestAccDataSourceTwilioProxyPhoneNumber_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_number", proxyPhoneNumberDataSourceName)

	testData := acceptance.TestAccData
	uniqueName := acctest.RandString(10)
	isReserved := true

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioProxyPhoneNumber_basic(testData, uniqueName, isReserved),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "is_reserved", "true"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "iso_country"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "in_use"),
					resource.TestCheckResourceAttr(stateDataSourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.fax_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.fax_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.mms_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.mms_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.restriction_fax_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.restriction_mms_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.restriction_sms_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.restriction_voice_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.sip_trunking"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.sms_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.sms_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.voice_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.0.voice_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioProxyPhoneNumber_basic(testData *acceptance.TestData, uniqueName string, isReserved bool) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name = "%s"
}

resource "twilio_proxy_phone_number" "phone_number" {
  service_sid = twilio_proxy_service.service.sid
  sid         = "%s"
  is_reserved = %t
}

data "twilio_proxy_phone_number" "phone_number" {
  service_sid = twilio_proxy_phone_number.phone_number.service_sid
  sid         = twilio_proxy_phone_number.phone_number.sid
}
`, uniqueName, testData.PhoneNumberSid, isReserved)
}
