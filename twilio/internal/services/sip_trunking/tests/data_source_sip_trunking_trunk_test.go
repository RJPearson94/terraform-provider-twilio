package tests

import (
	"fmt"
	"regexp"
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
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPTrunkingTrunk_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttr(stateDataSourceName, "cnam_lookup_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "disaster_recovery_method", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "disaster_recovery_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "domain_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "recording.#"),
					resource.TestCheckResourceAttr(stateDataSourceName, "recording.0.mode", "do-not-record"),
					resource.TestCheckResourceAttr(stateDataSourceName, "recording.0.trim", "do-not-trim"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "secure"),
					resource.TestCheckResourceAttr(stateDataSourceName, "transfer_mode", "disable-all"),
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

func TestAccDataSourceTwilioSIPTrunkingTrunk_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingTrunk_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPTrunkingTrunk_complete(testData *acceptance.TestData) string {
	return `
resource "twilio_sip_trunking_trunk" "trunk" {}

data "twilio_sip_trunking_trunk" "trunk" {
	sid = twilio_sip_trunking_trunk.trunk.sid
}
`
}

func testAccDataSourceTwilioSIPTrunkingTrunk_invalidSid() string {
	return `
data "twilio_sip_trunking_trunk" "trunk" {
  sid       = "sid"
}
`
}
