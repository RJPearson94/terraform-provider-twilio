package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var domainCredentialListMappingsDataSourceName = "twilio_sip_domain_credential_list_mappings"

func TestAccDataSourceTwilioSIPDomainCredentialListMappings_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.credential_list_mappings", domainCredentialListMappingsDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPDomainCredentialListMappings_basic(testData, friendlyName, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "domain_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "credential_list_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_list_mappings.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_list_mappings.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_list_mappings.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_list_mappings.0.date_updated"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioSIPDomainCredentialListMappings_basic(testData *acceptance.TestData, friendlyName string, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "%s"
  friendly_name = "%s"
}

resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
}

resource "twilio_sip_domain_credential_list_mapping" "credential_list_mapping" {
  account_sid         = twilio_sip_domain.domain.account_sid
  domain_sid          = twilio_sip_domain.domain.sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
}

data "twilio_sip_domain_credential_list_mappings" "credential_list_mappings" {
  account_sid = twilio_sip_domain_credential_list_mapping.credential_list_mapping.account_sid
  domain_sid  = twilio_sip_domain_credential_list_mapping.credential_list_mapping.domain_sid
}
`, testData.AccountSid, friendlyName, testData.AccountSid, domainName)
}
