package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var proxyPhoneNumbersDataSourceName = "twilio_proxy_phone_numbers"

// Tests have to run sequentially as a phone number cannot be associated with more than 1 proxy service at a given time

func TestAccDataSourceTwilioProxyPhoneNumbers_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_numbers", proxyPhoneNumbersDataSourceName)

	testData := acceptance.TestAccData
	uniqueName := acctest.RandString(10)
	isReserved := true

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioProxyPhoneNumbers_basic(testData, uniqueName, isReserved),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.is_reserved", "true"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.phone_number"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.iso_country"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.in_use"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.capabilities.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.fax_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.fax_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.mms_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.mms_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.restriction_fax_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.restriction_mms_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.restriction_sms_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.restriction_voice_domestic"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.sip_trunking"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.sms_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.sms_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.voice_inbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.0.voice_outbound"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioProxyPhoneNumbers_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioProxyPhoneNumbers_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^KS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioProxyPhoneNumbers_basic(testData *acceptance.TestData, uniqueName string, isReserved bool) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name = "%s"
}

resource "twilio_proxy_phone_number" "phone_number" {
  service_sid = twilio_proxy_service.service.sid
  sid         = "%s"
  is_reserved = %t
}

data "twilio_proxy_phone_numbers" "phone_numbers" {
  service_sid = twilio_proxy_phone_number.phone_number.service_sid
}
`, uniqueName, testData.PhoneNumberSid, isReserved)
}

func testAccDataSourceTwilioProxyPhoneNumbers_invalidServiceSid() string {
	return `
data "twilio_proxy_phone_numbers" "phone_numbers" {
  service_sid = "service_sid"
}
`
}
