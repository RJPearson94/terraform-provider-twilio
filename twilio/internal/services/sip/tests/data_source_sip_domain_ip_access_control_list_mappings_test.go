package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var domainIPAccessControlListMappingsDataSourceName = "twilio_sip_domain_ip_access_control_list_mappings"

func TestAccDataSourceTwilioSIPDomainIPAccessControlListMappings_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.ip_access_control_list_mappings", domainIPAccessControlListMappingsDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPDomainIPAccessControlListMappings_basic(testData, friendlyName, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "domain_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "ip_access_control_list_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_list_mappings.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_list_mappings.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_list_mappings.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_list_mappings.0.date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPDomainIPAccessControlListMappings_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainIPAccessControlListMappings_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPDomainIPAccessControlListMappings_invalidDomainSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainIPAccessControlListMappings_invalidDomainSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of domain_sid to match regular expression "\^SD\[0-9a-fA-F\]\{32\}\$", got domain_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMappings_basic(testData *acceptance.TestData, friendlyName string, domainName string) string {
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

data "twilio_sip_domain_ip_access_control_list_mappings" "ip_access_control_list_mappings" {
  account_sid = twilio_sip_domain_ip_access_control_list_mapping.ip_access_control_list_mapping.account_sid
  domain_sid  = twilio_sip_domain_ip_access_control_list_mapping.ip_access_control_list_mapping.domain_sid
}
`, testData.AccountSid, friendlyName, testData.AccountSid, domainName)
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMappings_invalidAccountSid() string {
	return `
data "twilio_sip_domain_ip_access_control_list_mappings" "ip_access_control_list_mappings" {
  account_sid = "account_sid"
  domain_sid  = "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPDomainIPAccessControlListMappings_invalidDomainSid() string {
	return `
data "twilio_sip_domain_ip_access_control_list_mappings" "ip_access_control_list_mappings" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_sid  = "domain_sid"
}
`
}
