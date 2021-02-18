package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var domainIPAccessControlListMappingDataSourceName = "twilio_sip_domain_ip_access_control_list_mapping"

func TestAccDataSourceTwilioSIPDomainIPAccessControlListMapping_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.ip_access_control_list_mapping", domainIPAccessControlListMappingDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_basic(testData, friendlyName, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "domain_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_basic(testData *acceptance.TestData, friendlyName string, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "%s"
  friendly_name = "%s"
}

resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
}

resource "twilio_sip_domain_ip_access_control_list_mapping" "ip_access_control_list_mapping" {
  account_sid                = twilio_sip_domain.domain.account_sid
  domain_sid                 = twilio_sip_domain.domain.sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
}

data "twilio_sip_domain_ip_access_control_list_mapping" "ip_access_control_list_mapping" {
  account_sid = twilio_sip_domain_ip_access_control_list_mapping.ip_access_control_list_mapping.account_sid
  domain_sid  = twilio_sip_domain_ip_access_control_list_mapping.ip_access_control_list_mapping.domain_sid
  sid         = twilio_sip_domain_ip_access_control_list_mapping.ip_access_control_list_mapping.sid
}
`, testData.AccountSid, friendlyName, testData.AccountSid, domainName)
}
