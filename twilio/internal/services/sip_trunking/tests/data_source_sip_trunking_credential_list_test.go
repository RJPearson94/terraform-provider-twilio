package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const credentialListDataSourceName = "twilio_sip_trunking_credential_list"

func TestAccDataSourceTwilioSIPTrunkingCredentialList_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.credential_list", credentialListDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPTrunkingCredentialList_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trunk_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPTrunkingCredentialList_invalidTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingCredentialList_invalidTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of trunk_sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got trunk_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPTrunkingCredentialList_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingCredentialList_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^CL\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPTrunkingCredentialList_basic(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "%s"
  friendly_name = "%s"
}

resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid           = twilio_sip_trunking_trunk.trunk.sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
}

data "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid = twilio_sip_trunking_credential_list.credential_list.trunk_sid
  sid       = twilio_sip_trunking_credential_list.credential_list.sid
}
`, testData.AccountSid, friendlyName)
}

func testAccDataSourceTwilioSIPTrunkingCredentialList_invalidTrunkSid() string {
	return `
data "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid = "trunk_sid"
  sid       = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPTrunkingCredentialList_invalidSid() string {
	return `
data "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid = "TKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid       = "sid"
}
`
}
