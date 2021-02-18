package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var ipAccessControlListsDataSourceName = "twilio_sip_trunking_ip_access_control_lists"

func TestAccDataSourceTwilioSIPTrunkingIPAccessControlLists_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.ip_access_control_lists", ipAccessControlListsDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPTrunkingIPAccessControlLists_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trunk_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "ip_access_control_lists.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_lists.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_lists.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_lists.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_lists.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ip_access_control_lists.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioSIPTrunkingIPAccessControlLists_basic(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "%s"
  friendly_name = "%s"
}

resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_ip_access_control_list" "ip_access_control_list" {
  trunk_sid                  = twilio_sip_trunking_trunk.trunk.sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
}

data "twilio_sip_trunking_ip_access_control_lists" "ip_access_control_lists" {
  trunk_sid = twilio_sip_trunking_ip_access_control_list.ip_access_control_list.trunk_sid
}
`, testData.AccountSid, friendlyName)
}
