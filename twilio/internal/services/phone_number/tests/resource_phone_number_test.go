// +build high_value

package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var phoneNumberResourceName = "twilio_phone_number"

func TestAccTwilioPhoneNumber_complete(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.phone_number", phoneNumberResourceName)
	testData := acceptance.TestAccData
	url := "https://demo.twilio.com/welcome/voice/"
	newUrl := "https://demo.twilio.com/welcome/sms/reply"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioPhoneNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioPhoneNumber_complete(testData, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttr(stateResourceName, "address_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "address_requirements"),
					resource.TestCheckResourceAttrSet(stateResourceName, "beta"),
					resource.TestCheckResourceAttr(stateResourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency_address_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "emergency_status"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.application_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "trunk_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.application_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.caller_id_lookup", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", url),
					resource.TestCheckNoResourceAttr(stateResourceName, "fax"),
					resource.TestCheckResourceAttr(stateResourceName, "identity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "bundle_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "status"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "origin"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioPhoneNumberImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
			{
				Config: testAccTwilioPhoneNumber_complete(testData, newUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttr(stateResourceName, "address_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "address_requirements"),
					resource.TestCheckResourceAttrSet(stateResourceName, "beta"),
					resource.TestCheckResourceAttr(stateResourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency_address_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "emergency_status"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.application_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "trunk_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.application_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.caller_id_lookup", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", newUrl),
					resource.TestCheckNoResourceAttr(stateResourceName, "fax"),
					resource.TestCheckResourceAttr(stateResourceName, "identity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "bundle_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "status"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "origin"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccCheckTwilioPhoneNumberDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != phoneNumberResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).IncomingPhoneNumber(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving phone number %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioPhoneNumberExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).IncomingPhoneNumber(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving phone number %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioPhoneNumberImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/PhoneNumbers/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioPhoneNumber_complete(testData *acceptance.TestData, url string) string {
	return fmt.Sprintf(`
resource "twilio_phone_number" "phone_number" {
  account_sid  = "%s"
  phone_number = "%s"

  voice {
    url = "%s"
  }
}
`, testData.AccountSid, testData.PurchasablePhoneNumber, url)
}
