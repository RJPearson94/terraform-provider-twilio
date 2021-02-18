package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var credentialListsDataSourceName = "twilio_sip_trunking_credential_lists"

func TestAccDataSourceTwilioSIPTrunkingCredentialLists_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.credential_lists", credentialListsDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPTrunkingCredentialLists_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trunk_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "credential_lists.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_lists.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_lists.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_lists.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_lists.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_lists.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioSIPTrunkingCredentialLists_basic(testData *acceptance.TestData, friendlyName string) string {
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

data "twilio_sip_trunking_credential_lists" "credential_lists" {
  trunk_sid = twilio_sip_trunking_credential_list.credential_list.trunk_sid
}
`, testData.AccountSid, friendlyName)
}
