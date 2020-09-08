package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var proxyPhoneNumberResourceName = "twilio_proxy_phone_number"

// Tests have to run sequentially as a phone number cannot be associated with more than 1 proxy service at a given time

func TestAccTwilioProxyPhoneNumber_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.phone_number", proxyPhoneNumberResourceName)

	testData := acceptance.TestAccData
	uniqueName := acctest.RandString(10)
	isReserved := true

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyPhoneNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyPhoneNumber_basic(testData, uniqueName, isReserved),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_reserved", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "iso_country"),
					resource.TestCheckResourceAttrSet(stateResourceName, "in_use"),
					resource.TestCheckResourceAttr(stateResourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.fax_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.fax_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.mms_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.mms_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_fax_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_mms_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_sms_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_voice_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sip_trunking"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sms_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sms_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.voice_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.voice_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioProxyPhoneNumberImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioProxyPhoneNumber_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.phone_number", proxyPhoneNumberResourceName)

	testData := acceptance.TestAccData
	uniqueName := acctest.RandString(10)
	isReserved := true
	newIsReserved := false

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyPhoneNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyPhoneNumber_basic(testData, uniqueName, isReserved),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_reserved", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "iso_country"),
					resource.TestCheckResourceAttrSet(stateResourceName, "in_use"),
					resource.TestCheckResourceAttr(stateResourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.fax_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.fax_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.mms_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.mms_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_fax_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_mms_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_sms_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_voice_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sip_trunking"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sms_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sms_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.voice_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.voice_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioProxyPhoneNumber_basic(acceptance.TestAccData, uniqueName, newIsReserved),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_reserved", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "iso_country"),
					resource.TestCheckResourceAttrSet(stateResourceName, "in_use"),
					resource.TestCheckResourceAttr(stateResourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.fax_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.fax_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.mms_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.mms_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_fax_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_mms_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_sms_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.restriction_voice_domestic"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sip_trunking"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sms_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.sms_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.voice_inbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "capabilities.0.voice_outbound"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioProxyPhoneNumberDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Proxy

	for _, rs := range s.RootModule().Resources {
		if rs.Type != proxyPhoneNumberResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).PhoneNumber(rs.Primary.ID).Fetch(); err != nil {
			if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
				// currently proxy returns a 400 if the proxy phone number instance does not exist
				if twilioError.Status == 400 && twilioError.Message == "Invalid Phone Number Sid" {
					return nil
				}
			}
			return fmt.Errorf("Error occurred when retrieving proxy phone number information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioProxyPhoneNumberExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Proxy

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).PhoneNumber(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving proxy phone number information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioProxyPhoneNumberImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/PhoneNumbers/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioProxyPhoneNumber_basic(testData *acceptance.TestData, uniqueName string, isReserved bool) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name = "%s"
}

resource "twilio_proxy_phone_number" "phone_number" {
  service_sid = twilio_proxy_service.service.sid
  sid         = "%s"
  is_reserved = %t
}
`, uniqueName, testData.PhoneNumberSid, isReserved)
}
