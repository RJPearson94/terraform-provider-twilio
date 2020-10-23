package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var phoneNumberAvailableTollFreeDataSourceName = "twilio_phone_number_available_toll_free_numbers"

func TestAccDataSourceTwilioPhoneNumberAvailableTollFree_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.available_toll_free_numbers", phoneNumberAvailableTollFreeDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioPhoneNumberAvailableTollFree_complete(testData),
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

func testAccTwilioPhoneNumberAvailableTollFree_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
data "twilio_phone_number_available_toll_free_numbers" "available_toll_free_numbers" {
  account_sid = "%s"
  iso_country = "GB"
}
`, testData.AccountSid)
}
