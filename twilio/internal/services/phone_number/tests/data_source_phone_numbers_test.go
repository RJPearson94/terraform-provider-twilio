package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumbersDataSourceName = "twilio_phone_numbers"

func TestAccDataSourceTwilioPhoneNumbers_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_numbers", phoneNumbersDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioPhoneNumbers_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.phone_number"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.address_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.address_requirements"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.beta"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.capabilities.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.emergency_address_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.emergency_status"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.messaging.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.trunk_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.voice_and_fax.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.identity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.bundle_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.status"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.status_callback_method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.origin"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_updated"),
				),
			},
		},
	})
}

func testAccTwilioPhoneNumbers_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
data "twilio_phone_numbers" "phone_numbers" {
  account_sid = "%s"
}
`, testData.AccountSid)
}
