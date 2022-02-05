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

var domainRegistrationCredentialListMappingResourceName = "twilio_sip_domain_registration_credential_list_mapping"

func TestAccTwilioSIPDomainRegistrationCredentialListMapping_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain_registration_credential_list_mapping", domainRegistrationCredentialListMappingResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainRegistrationCredentialListMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomainRegistrationCredentialListMapping_basic(testData, friendlyName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainRegistrationCredentialListMappingExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "domain_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "credential_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSIPDomainRegistrationCredentialListMappingImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSIPDomainRegistrationCredentialListMapping_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomainRegistrationCredentialListMapping_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccTwilioSIPDomainRegistrationCredentialListMapping_invalidDomainSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomainRegistrationCredentialListMapping_invalidDomainSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of domain_sid to match regular expression "\^SD\[0-9a-fA-F\]\{32\}\$", got domain_sid`),
			},
		},
	})
}

func TestAccTwilioSIPDomainRegistrationCredentialListMapping_invalidCredentialListSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomainRegistrationCredentialListMapping_invalidCredentialListSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of credential_list_sid to match regular expression "\^CL\[0-9a-fA-F\]\{32\}\$", got credential_list_sid`),
			},
		},
	})
}

func testAccCheckTwilioSIPDomainRegistrationCredentialListMappingDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != domainRegistrationCredentialListMappingResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.Domain(rs.Primary.Attributes["domain_sid"]).Auth.Registrations.CredentialListMapping(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving SIP domain credential list mapping information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPDomainRegistrationCredentialListMappingExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.Domain(rs.Primary.Attributes["domain_sid"]).Auth.Registrations.CredentialListMapping(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving SIP domain registration credential list mapping information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPDomainRegistrationCredentialListMappingImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/SIP/Domains/%s/Auth/Registrations/CredentialListMappings/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["domain_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPDomainRegistrationCredentialListMapping_basic(testData *acceptance.TestData, friendlyName string, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "%[1]s"
  friendly_name = "%[2]s"
}

resource "twilio_sip_domain" "domain" {
  account_sid = "%[1]s"
  domain_name = "%[3]s"
}

resource "twilio_sip_domain_registration_credential_list_mapping" "domain_registration_credential_list_mapping" {
  account_sid         = twilio_sip_domain.domain.account_sid
  domain_sid          = twilio_sip_domain.domain.sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
}
`, testData.AccountSid, friendlyName, domainName)
}

func testAccTwilioSIPDomainRegistrationCredentialListMapping_invalidAccountSid() string {
	return `
resource "twilio_sip_domain_registration_credential_list_mapping" "domain_registration_credential_list_mapping" {
  account_sid         = "account_sid"
  domain_sid          = "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  credential_list_sid = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioSIPDomainRegistrationCredentialListMapping_invalidDomainSid() string {
	return `
resource "twilio_sip_domain_registration_credential_list_mapping" "domain_registration_credential_list_mapping" {
  account_sid         = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_sid          = "domain_sid"
  credential_list_sid = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioSIPDomainRegistrationCredentialListMapping_invalidCredentialListSid() string {
	return `
resource "twilio_sip_domain_registration_credential_list_mapping" "domain_registration_credential_list_mapping" {
  account_sid         = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_sid          = "SDaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  credential_list_sid = "credential_list_sid"
}
`
}
