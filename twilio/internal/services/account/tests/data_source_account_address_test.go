package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var accountAddressDataSourceName = "twilio_account_address"

func TestAccDataSourceTwilioAccountAddress_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.address", accountAddressDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountAddress_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "customer_name", testData.CustomerName),
					resource.TestCheckResourceAttr(stateDataSourceName, "street", testData.Address.Street),
					resource.TestCheckResourceAttr(stateDataSourceName, "city", testData.Address.City),
					resource.TestCheckResourceAttr(stateDataSourceName, "region", testData.Address.Region),
					resource.TestCheckResourceAttr(stateDataSourceName, "postal_code", testData.Address.PostalCode),
					resource.TestCheckResourceAttr(stateDataSourceName, "iso_country", testData.Address.IsoCountry),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "validated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "verified"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "emergency_enabled"),
					resource.TestCheckResourceAttr(stateDataSourceName, "street_secondary", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountAddress_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountAddress_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountAddress_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountAddress_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^AD\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAccountAddress_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_account_address" "address" {
  account_sid   = "%s"
  customer_name = "%s"
  street        = "%s"
  city          = "%s"
  region        = "%s"
  postal_code   = "%s"
  iso_country   = "%s"
}

data "twilio_account_address" "address" {
  account_sid = twilio_account_address.address.account_sid
  sid         = twilio_account_address.address.sid
}
`, testData.AccountSid, testData.CustomerName, testData.Address.Street, testData.Address.City, testData.Address.Region, testData.Address.PostalCode, testData.Address.IsoCountry)
}

func testAccDataSourceTwilioAccountAddress_invalidAccountSid() string {
	return `
data "twilio_account_address" "address" {
  account_sid = "account_sid"
  sid         = "ADaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioAccountAddress_invalidSid() string {
	return `
data "twilio_account_address" "address" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
