package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumberDataSourceName = "twilio_phone_number"

func TestAccDataSourceTwilioPhoneNumber_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_number", phoneNumberDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioPhoneNumber_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_number"),
					resource.TestCheckResourceAttr(stateDataSourceName, "address_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "address_requirements"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "beta"),
					resource.TestCheckResourceAttr(stateDataSourceName, "capabilities.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "emergency_address_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "emergency_status"),
					resource.TestCheckResourceAttr(stateDataSourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "trunk_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice_and_fax.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "identity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "bundle_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status"),
					resource.TestCheckResourceAttr(stateDataSourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status_callback_method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "origin"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccTwilioPhoneNumber_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
data "twilio_phone_number" "phone_number" {
  account_sid = "%s"
  sid         = "%s"
}
`, testData.AccountSid, testData.PhoneNumberSid)
}
