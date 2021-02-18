package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var ipAddressDataSourceName = "twilio_sip_ip_address"

func TestAccDataSourceTwilioSIPIPAddress_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.ip_address", ipAddressDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	ipAddress := "127.0.0.1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPIPAddress_basic(testData, friendlyName, ipAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "ip_address", ipAddress),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "cidr_length_prefix"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_list_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPIPAddress_invalidIPAddress(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	ipAddress := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPIPAddress_basic(testData, friendlyName, ipAddress),
				ExpectError: regexp.MustCompile(`(?s)expected ip_address to contain a valid IP, got: test`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPIPAddress_basic(testData *acceptance.TestData, friendlyName string, ipAddress string) string {
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

data "twilio_sip_ip_address" "ip_address" {
  account_sid                = twilio_sip_ip_address.ip_address.account_sid
  ip_access_control_list_sid = twilio_sip_ip_address.ip_address.ip_access_control_list_sid
  sid                        = twilio_sip_ip_address.ip_address.sid
}
`, testData.AccountSid, friendlyName, friendlyName, ipAddress)
}
