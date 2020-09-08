package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumbersDataSourceName = "twilio_messaging_phone_numbers"

func TestAccDataSourceTwilioMessagingPhoneNumbers_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_numbers", phoneNumbersDataSourceName)
	friendlyName := acctest.RandString(10)
	testData := acceptance.TestAccData

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioMessagingPhoneNumbers_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "phone_numbers.0.sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.capabilities.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.country_code"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.phone_number"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_numbers.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioMessagingPhoneNumbers_basic(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "service-%s"
}

resource "twilio_messaging_phone_number" "phone_number" {
  service_sid = twilio_messaging_service.service.sid
  sid         = "%s"
}

data "twilio_messaging_phone_numbers" "phone_numbers" {
  service_sid = twilio_messaging_phone_number.phone_number.service_sid
}
`, friendlyName, testData.PhoneNumberSid)
}
