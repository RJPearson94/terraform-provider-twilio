package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumberDataSourceName = "twilio_messaging_phone_number"

func TestAccDataSourceTwilioMessagingPhoneNumber_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.phone_number", phoneNumberDataSourceName)
	friendlyName := acctest.RandString(10)
	testData := acceptance.TestAccData

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioMessagingPhoneNumber_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "capabilities.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "country_code"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioMessagingPhoneNumber_basic(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "service-%s"
}

resource "twilio_messaging_phone_number" "phone_number" {
  service_sid = twilio_messaging_service.service.sid
  sid         = "%s"
}

data "twilio_messaging_phone_number" "phone_number" {
  service_sid = twilio_messaging_phone_number.phone_number.service_sid
  sid         = twilio_messaging_phone_number.phone_number.sid
}
`, friendlyName, testData.PhoneNumberSid)
}
