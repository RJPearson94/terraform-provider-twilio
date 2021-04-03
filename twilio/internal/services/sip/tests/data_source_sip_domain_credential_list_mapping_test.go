package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var domainCredentialListMappingDataSourceName = "twilio_sip_domain_credential_list_mapping"

func TestAccDataSourceTwilioSIPDomainCredentialListMapping_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.credential_list_mapping", domainCredentialListMappingDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPDomainCredentialListMapping_basic(testData, friendlyName, domainName),
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

func TestAccDataSourceTwilioSIPDomainCredentialListMapping_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainCredentialListMapping_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPDomainCredentialListMapping_invalidDomainSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainCredentialListMapping_invalidDomainSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of domain_sid to match regular expression "\^SD\[0-9a-fA-F\]\{32\}\$", got domain_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPDomainCredentialListMapping_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPDomainCredentialListMapping_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^CL\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPDomainCredentialListMapping_basic(testData *acceptance.TestData, friendlyName string, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "%[1]s"
  friendly_name = "%[2]s"
}

resource "twilio_sip_domain" "domain" {
  account_sid = "%[1]s"
  domain_name = "%[3]s"
}

resource "twilio_sip_domain_credential_list_mapping" "credential_list_mapping" {
  account_sid         = twilio_sip_domain.domain.account_sid
  domain_sid          = twilio_sip_domain.domain.sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
}

data "twilio_sip_domain_credential_list_mapping" "credential_list_mapping" {
  account_sid = twilio_sip_domain_credential_list_mapping.credential_list_mapping.account_sid
  domain_sid  = twilio_sip_domain_credential_list_mapping.credential_list_mapping.domain_sid
  sid         = twilio_sip_domain_credential_list_mapping.credential_list_mapping.sid
}
`, testData.AccountSid, friendlyName, domainName)
}

func testAccDataSourceTwilioSIPDomainCredentialListMapping_invalidAccountSid() string {
	return `
data "twilio_sip_domain_credential_list_mapping" "credential_list_mapping" {
  account_sid = "account_sid"
  domain_sid  = "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPDomainCredentialListMapping_invalidDomainSid() string {
	return `
data "twilio_sip_domain_credential_list_mapping" "credential_list_mapping" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_sid  = "domain_sid"
  sid         = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPDomainCredentialListMapping_invalidSid() string {
	return `
data "twilio_sip_domain_credential_list_mapping" "credential_list_mapping" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_sid  = "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
