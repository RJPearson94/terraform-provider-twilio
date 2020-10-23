package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumberAvailableLocalDataSourceName = "twilio_phone_number_available_local_numbers"

func TestAccDataSourceTwilioPhoneNumberAvailableLocal_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.available_local_numbers", phoneNumberAvailableLocalDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioPhoneNumberAvailableLocal_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "iso_country", "GB"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "available_phone_numbers.#"),
				),
			},
		},
	})
}

func testAccTwilioPhoneNumberAvailableLocal_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
data "twilio_phone_number_available_local_numbers" "available_local_numbers" {
  account_sid = "%s"
  iso_country = "GB"
}
`, testData.AccountSid)
}
