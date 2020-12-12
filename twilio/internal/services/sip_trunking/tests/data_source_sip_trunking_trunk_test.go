package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var trunkDataSourceName = "twilio_sip_trunking_trunk"

func TestAccDataSourceTwilioSIPTrunkingTrunk_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.trunk", trunkDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "cnam_lookup_enabled"),
					resource.TestCheckResourceAttr(stateDataSourceName, "disaster_recovery_method", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "disaster_recovery_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "domain_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "recording.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "recording.0.mode"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "recording.0.trim"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "secure"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "transfer_mode"),
					resource.TestCheckResourceAttr(stateDataSourceName, "auth_type", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "auth_type_set.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccTwilioSIPTrunkingTrunk_complete(testData *acceptance.TestData) string {
	return `
resource "twilio_sip_trunking_trunk" "trunk" {}

data "twilio_sip_trunking_trunk" "trunk" {
	sid = twilio_sip_trunking_trunk.trunk.sid
}
`
}
