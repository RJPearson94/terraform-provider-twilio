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

var ipAddressResourceName = "twilio_sip_ip_address"

func TestAccTwilioSIPIPAddress_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.ip_address", ipAddressResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	ipAddress := "127.0.0.1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPIPAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPIPAddress_basic(testData, friendlyName, ipAddress),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPIPAddressExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "ip_address", ipAddress),
					resource.TestCheckResourceAttr(stateResourceName, "cidr_length_prefix", "32"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ip_access_control_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSIPIPAddressImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSIPIPAddress_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.ip_address", ipAddressResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	ipAddress := "127.0.0.1"
	newIPAddress := "0.0.0.0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPIPAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPIPAddress_basic(testData, friendlyName, ipAddress),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPIPAddressExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "ip_address", ipAddress),
					resource.TestCheckResourceAttr(stateResourceName, "cidr_length_prefix", "32"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ip_access_control_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioSIPIPAddress_basic(testData, friendlyName, newIPAddress),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPIPAddressExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "ip_address", newIPAddress),
					resource.TestCheckResourceAttr(stateResourceName, "cidr_length_prefix", "32"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ip_access_control_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioSIPIPAddress_cidrLengthPrefix(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.ip_address", ipAddressResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	ipAddress := "127.0.0.1"
	newCidrLengthPrefix := 8

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPIPAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPIPAddress_basic(testData, friendlyName, ipAddress),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPIPAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "cidr_length_prefix", "32"),
				),
			},
			{
				Config: testAccTwilioSIPIPAddress_cidrLengthPrefix(testData, friendlyName, ipAddress, newCidrLengthPrefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPIPAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "cidr_length_prefix", "8"),
				),
			},
			{
				Config: testAccTwilioSIPIPAddress_basic(testData, friendlyName, ipAddress),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPIPAddressExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "cidr_length_prefix", "32"),
				),
			},
		},
	})
}

func TestAccTwilioSIPIPAddress_blankFriendlyName(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := ""
	ipAddress := "127.0.0.1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPIPAddress_basic(testData, friendlyName, ipAddress),
				ExpectError: regexp.MustCompile(`(?s)expected \"friendly_name\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioSIPIPAddress_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPIPAddress_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccTwilioSIPIPAddress_invalidIPAccessControlListSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPIPAddress_invalidIPAccessControlListSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of ip_access_control_list_sid to match regular expression "\^AL\[0-9a-fA-F\]\{32\}\$", got ip_access_control_list_sid`),
			},
		},
	})
}

func testAccCheckTwilioSIPIPAddressDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != ipAddressResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.IpAccessControlList(rs.Primary.Attributes["ip_access_control_list_sid"]).IpAddress(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving the SIP IP address information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPIPAddressExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.IpAccessControlList(rs.Primary.Attributes["ip_access_control_list_sid"]).IpAddress(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving the SIP IP address information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPIPAddressImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/SIP/IpAccessControlLists/%s/IpAddresses/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["ip_access_control_list_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPIPAddress_basic(testData *acceptance.TestData, friendlyName string, ipAddress string) string {
	return fmt.Sprintf(`
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "%[1]s"
  friendly_name = "%[2]s"
}

resource "twilio_sip_ip_address" "ip_address" {
  account_sid                = twilio_sip_ip_access_control_list.ip_access_control_list.account_sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
  friendly_name              = "%[2]s"
  ip_address                 = "%[3]s"
}
`, testData.AccountSid, friendlyName, ipAddress)
}

func testAccTwilioSIPIPAddress_cidrLengthPrefix(testData *acceptance.TestData, friendlyName string, ipAddress string, cidrLengthPrefix int) string {
	return fmt.Sprintf(`
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "%[1]s"
  friendly_name = "%[2]s"
}

resource "twilio_sip_ip_address" "ip_address" {
  account_sid                = twilio_sip_ip_access_control_list.ip_access_control_list.account_sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
  friendly_name              = "%[2]s"
  ip_address                 = "%[3]s"
  cidr_length_prefix         = %[4]d
}
`, testData.AccountSid, friendlyName, ipAddress, cidrLengthPrefix)
}

func testAccTwilioSIPIPAddress_invalidAccountSid() string {
	return `
resource "twilio_sip_ip_address" "ip_address" {
  account_sid                = "account_sid"
  ip_access_control_list_sid = "ALaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name              = "invalid_account_sid"
  ip_address                 = "127.0.0.1"
}
`
}

func testAccTwilioSIPIPAddress_invalidIPAccessControlListSid() string {
	return `
resource "twilio_sip_ip_address" "ip_address" {
  account_sid                = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  ip_access_control_list_sid = "ip_access_control_list_sid"
  friendly_name              = "invalid_account_sid"
  ip_address                 = "127.0.0.1"
}
`
}
