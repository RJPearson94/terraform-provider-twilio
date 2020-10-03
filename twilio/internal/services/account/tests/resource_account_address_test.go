package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var addressResourceName = "twilio_account_address"

func TestAccTwilioAccountAddress_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address", addressResourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		Providers:         acceptance.TestAccProviders,
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAccountAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountAddress_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "customer_name", testData.CustomerName),
					resource.TestCheckResourceAttr(stateResourceName, "street", testData.Address.Street),
					resource.TestCheckResourceAttr(stateResourceName, "city", testData.Address.City),
					resource.TestCheckResourceAttr(stateResourceName, "region", testData.Address.Region),
					resource.TestCheckResourceAttr(stateResourceName, "postal_code", testData.Address.PostalCode),
					resource.TestCheckResourceAttr(stateResourceName, "iso_country", testData.Address.IsoCountry),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "validated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "verified"),
					resource.TestCheckResourceAttrSet(stateResourceName, "emergency_enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "street_secondary", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAccountAddressImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAccountAddress_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address", addressResourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAccountAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountAddress_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "customer_name", testData.CustomerName),
					resource.TestCheckResourceAttr(stateResourceName, "street", testData.Address.Street),
					resource.TestCheckResourceAttr(stateResourceName, "city", testData.Address.City),
					resource.TestCheckResourceAttr(stateResourceName, "region", testData.Address.Region),
					resource.TestCheckResourceAttr(stateResourceName, "postal_code", testData.Address.PostalCode),
					resource.TestCheckResourceAttr(stateResourceName, "iso_country", testData.Address.IsoCountry),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "validated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "verified"),
					resource.TestCheckResourceAttrSet(stateResourceName, "emergency_enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "street_secondary", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioAccountAddress_withStreetSecondary(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "customer_name", testData.CustomerName),
					resource.TestCheckResourceAttr(stateResourceName, "street", testData.Address.Street),
					resource.TestCheckResourceAttr(stateResourceName, "street_secondary", testData.Address.StreetSecondary),
					resource.TestCheckResourceAttr(stateResourceName, "city", testData.Address.City),
					resource.TestCheckResourceAttr(stateResourceName, "region", testData.Address.Region),
					resource.TestCheckResourceAttr(stateResourceName, "postal_code", testData.Address.PostalCode),
					resource.TestCheckResourceAttr(stateResourceName, "iso_country", testData.Address.IsoCountry),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "validated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "verified"),
					resource.TestCheckResourceAttrSet(stateResourceName, "emergency_enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccCheckTwilioAccountAddressDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != addressResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Address(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving address information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioAccountAddressExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Address(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving address information %s", err)
		}

		return nil
	}
}

func testAccTwilioAccountAddressImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/Addresses/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAccountAddress_basic(testData *acceptance.TestData) string {
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
`, testData.AccountSid, testData.CustomerName, testData.Address.Street, testData.Address.City, testData.Address.Region, testData.Address.PostalCode, testData.Address.IsoCountry)
}

func testAccTwilioAccountAddress_withStreetSecondary(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_account_address" "address" {
  account_sid      = "%s"
  customer_name    = "%s"
  street           = "%s"
  street_secondary = "%s"
  city             = "%s"
  region           = "%s"
  postal_code      = "%s"
  iso_country      = "%s"
}
`, testData.AccountSid, testData.CustomerName, testData.Address.Street, testData.Address.StreetSecondary, testData.Address.City, testData.Address.Region, testData.Address.PostalCode, testData.Address.IsoCountry)
}
