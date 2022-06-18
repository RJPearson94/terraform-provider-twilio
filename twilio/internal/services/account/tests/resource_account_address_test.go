package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var addressResourceName = "twilio_account_address"

func TestAccTwilioAccountAddress_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address", addressResourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
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
					resource.TestCheckResourceAttr(stateResourceName, "emergency_enabled", "false"),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
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
					resource.TestCheckResourceAttr(stateResourceName, "emergency_enabled", "false"),
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
					resource.TestCheckResourceAttr(stateResourceName, "emergency_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioAccountAddress_streetSecondary(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address", addressResourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountAddress_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "street_secondary", ""),
				),
			},
			{
				Config: testAccTwilioAccountAddress_withStreetSecondary(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "street_secondary", testData.Address.StreetSecondary),
				),
			},
			{
				Config: testAccTwilioAccountAddress_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "street_secondary", ""),
				),
			},
		},
	})
}

func TestAccTwilioAccountAddress_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address", addressResourceName)
	testData := acceptance.TestAccData
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountAddress_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
			{
				Config: testAccTwilioAccountAddress_friendlyName(testData, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
			{
				Config: testAccTwilioAccountAddress_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioAccountAddress_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountAddress_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccTwilioAccountAddress_blankCustomerName(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountAddress_blankCustomerName(),
				ExpectError: regexp.MustCompile(`(?s)expected \"customer_name\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAccountAddress_blankStreet(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountAddress_blankStreet(),
				ExpectError: regexp.MustCompile(`(?s)expected \"street\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAccountAddress_blankCity(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountAddress_blankCity(),
				ExpectError: regexp.MustCompile(`(?s)expected \"city\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAccountAddress_blankRegion(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountAddress_blankRegion(),
				ExpectError: regexp.MustCompile(`(?s)expected \"region\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAccountAddress_blankPostalCode(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountAddress_blankPostalCode(),
				ExpectError: regexp.MustCompile(`(?s)expected \"postal_code\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioAccountAddress_blankIsoCountry(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountAddress_blankIsoCountry(),
				ExpectError: regexp.MustCompile(`(?s)expected \"iso_country\" to not be an empty string, got `),
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

func testAccTwilioAccountAddress_friendlyName(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_account_address" "address" {
  account_sid   = "%s"
  customer_name = "%s"
  street        = "%s"
  city          = "%s"
  region        = "%s"
  postal_code   = "%s"
  iso_country   = "%s"
  friendly_name = "%s"
}
`, testData.AccountSid, testData.CustomerName, testData.Address.Street, testData.Address.City, testData.Address.Region, testData.Address.PostalCode, testData.Address.IsoCountry, friendlyName)
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

func testAccTwilioAccountAddress_invalidAccountSid() string {
	return `
resource "twilio_account_address" "address" {
  account_sid      = "account_sid"
  customer_name    = "customer_name"
  street           = "street"
  street_secondary = "street_secondary"
  city             = "city"
  region           = "region"
  postal_code      = "postal_code"
  iso_country      = "iso_country"
}
`
}

func testAccTwilioAccountAddress_blankCustomerName() string {
	return `
resource "twilio_account_address" "address" {
  account_sid   = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  customer_name = ""
  street        = "street"
  city          = "city"
  region        = "region"
  postal_code   = "postal_code"
  iso_country   = "iso_country"
}
`
}

func testAccTwilioAccountAddress_blankStreet() string {
	return `
resource "twilio_account_address" "address" {
  account_sid   = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  customer_name = "customer_name"
  street        = ""
  city          = "city"
  region        = "region"
  postal_code   = "postal_code"
  iso_country   = "iso_country"
}
`
}

func testAccTwilioAccountAddress_blankCity() string {
	return `
resource "twilio_account_address" "address" {
  account_sid   = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  customer_name = "customer_name"
  street        = "street"
  city          = ""
  region        = "region"
  postal_code   = "postal_code"
  iso_country   = "iso_country"
}
`
}

func testAccTwilioAccountAddress_blankRegion() string {
	return `
resource "twilio_account_address" "address" {
  account_sid   = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  customer_name = "customer_name"
  street        = "street"
  city          = "city"
  region        = ""
  postal_code   = "postal_code"
  iso_country   = "iso_country"
}
`
}

func testAccTwilioAccountAddress_blankPostalCode() string {
	return `
resource "twilio_account_address" "address" {
  account_sid   = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  customer_name = "customer_name"
  street        = "street"
  city          = "city"
  region        = "region"
  postal_code   = ""
  iso_country   = "iso_country"
}
`
}

func testAccTwilioAccountAddress_blankIsoCountry() string {
	return `
resource "twilio_account_address" "address" {
  account_sid   = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  customer_name = "customer_name"
  street        = "street"
  city          = "city"
  region        = "region"
  postal_code   = "postal_code"
  iso_country   = ""
}
`
}
