package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const phoneNumberResourceName = "twilio_sip_trunking_phone_number"

func TestAccTwilioSIPTrunkingPhoneNumber_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.phone_number", phoneNumberResourceName)
	testData := acceptance.TestAccData

	// Run tests in sequence to prevent the same value 2 SIP trunks mutating the same phone number simultaneously
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingPhoneNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingPhoneNumber_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateResourceName, "address_requirements"),
					resource.TestCheckResourceAttrSet(stateResourceName, "beta"),
					resource.TestCheckResourceAttr(stateResourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "fax.#", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "status_callback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSIPTrunkingPhoneNumberImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTwilioSIPTrunkingPhoneNumberDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

	for _, rs := range s.RootModule().Resources {
		if rs.Type != phoneNumberResourceName {
			continue
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).PhoneNumber(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving phone number information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPTrunkingPhoneNumberExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).PhoneNumber(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving phone number information %s", err.Error())
		}

		return nil
	}
}

func TestAccDataSourceTwilioSIPTrunkingTrunk_invalidTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingPhoneNumber_invalidTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of trunk_sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got trunk_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPTrunkingTrunk_invalidPhoneNumberSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingPhoneNumber_invalidPhoneNumberSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of phone_number_sid to match regular expression "\^PN\[0-9a-fA-F\]\{32\}\$", got phone_number_sid`),
			},
		},
	})
}

func testAccTwilioSIPTrunkingPhoneNumberImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Trunks/%s/PhoneNumbers/%s", rs.Primary.Attributes["trunk_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPTrunkingPhoneNumber_basic(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid        = twilio_sip_trunking_trunk.trunk.sid
  phone_number_sid = "%s"
}
`, testData.PhoneNumberSid)
}

func testAccTwilioSIPTrunkingPhoneNumber_invalidTrunkSid() string {
	return `
resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = "trunk_sid"
  phone_number_sid       = "PNaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioSIPTrunkingPhoneNumber_invalidPhoneNumberSid() string {
	return `
resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = "TKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  phone_number_sid       = "phone_number_sid"
}
`
}
