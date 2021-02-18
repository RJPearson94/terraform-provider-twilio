package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var ipAddressesDataSourceName = "twilio_sip_ip_addresses"

func TestAccDataSourceTwilioSIPIPAddresses_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.ip_addresses", ipAddressesDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	ipAddress := "127.0.0.1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPIPAddresses_basic(testData, friendlyName, ipAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_list_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "ip_addresses.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "ip_addresses.0.ip_address", ipAddress),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_addresses.0.cidr_length_prefix"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_addresses.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_addresses.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_addresses.0.date_updated"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioSIPIPAddresses_basic(testData *acceptance.TestData, friendlyName string, ipAddress string) string {
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

data "twilio_sip_ip_addresses" "ip_addresses" {
  account_sid                = twilio_sip_ip_address.ip_address.account_sid
  ip_access_control_list_sid = twilio_sip_ip_address.ip_address.ip_access_control_list_sid
}
`, testData.AccountSid, friendlyName, friendlyName, ipAddress)
}
