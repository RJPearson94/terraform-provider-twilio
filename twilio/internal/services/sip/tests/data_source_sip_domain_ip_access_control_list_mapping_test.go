package tests

import (
	"fmt"
	"regexp"
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

func TestAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidDomainSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidDomainSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of domain_sid to match regular expression "\^SD\[0-9a-fA-F\]\{32\}\$", got domain_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^AL\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_basic(testData *acceptance.TestData, friendlyName string, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "%[1]s"
  friendly_name = "%[2]s"
}

resource "twilio_sip_domain" "domain" {
  account_sid = "%[1]s"
  domain_name = "%[3]s"
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
`, testData.AccountSid, friendlyName, domainName)
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidAccountSid() string {
	return `
data "twilio_sip_domain_ip_access_control_list_mapping" "ip_access_control_list_mapping" {
  account_sid = "account_sid"
  domain_sid  = "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "ALaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidDomainSid() string {
	return `
data "twilio_sip_domain_ip_access_control_list_mapping" "ip_access_control_list_mapping" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_sid  = "domain_sid"
  sid         = "ALaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMapping_invalidSid() string {
	return `
data "twilio_sip_domain_ip_access_control_list_mapping" "ip_access_control_list_mapping" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_sid  = "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
