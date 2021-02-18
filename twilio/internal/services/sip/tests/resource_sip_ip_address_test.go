package tests

import (
	"fmt"
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
					resource.TestCheckResourceAttrSet(stateResourceName, "cidr_length_prefix"),
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
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "ip_address", ipAddress),
					resource.TestCheckResourceAttrSet(stateResourceName, "cidr_length_prefix"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ip_access_control_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioSIPIPAddress_cidrLengthPrefix(testData, friendlyName, ipAddress, newCidrLengthPrefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPIPAddressExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "ip_address", ipAddress),
					resource.TestCheckResourceAttr(stateResourceName, "cidr_length_prefix", "8"),
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
  account_sid   = "%s"
  friendly_name = "%s"
}

resource "twilio_sip_ip_address" "ip_address" {
  account_sid                = twilio_sip_ip_access_control_list.ip_access_control_list.account_sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
  friendly_name              = "%s"
  ip_address                 = "%s"
}
`, testData.AccountSid, friendlyName, friendlyName, ipAddress)
}

func testAccTwilioSIPIPAddress_cidrLengthPrefix(testData *acceptance.TestData, friendlyName string, ipAddress string, cidrLengthPrefix int) string {
	return fmt.Sprintf(`
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "%s"
  friendly_name = "%s"
}

resource "twilio_sip_ip_address" "ip_address" {
  account_sid                = twilio_sip_ip_access_control_list.ip_access_control_list.account_sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
  friendly_name              = "%s"
  ip_address                 = "%s"
  cidr_length_prefix         = %d
}
`, testData.AccountSid, friendlyName, friendlyName, ipAddress, cidrLengthPrefix)
}
