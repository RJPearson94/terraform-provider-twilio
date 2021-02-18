package tests

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var originationURLDataSourceName = "twilio_sip_trunking_origination_url"

func TestAccDataSourceTwilioSIPTrunkingOriginationURL_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.origination_url", originationURLDataSourceName)

	friendlyName := acctest.RandString(10)
	weight := 0
	priority := 0
	enabled := false
	sipURL := "sip:test@test.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingOriginationURL_complete(friendlyName, enabled, priority, sipURL, weight),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "enabled", strconv.FormatBool(enabled)),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "priority", strconv.Itoa(priority)),
					resource.TestCheckResourceAttr(stateDataSourceName, "sip_url", sipURL),
					resource.TestCheckResourceAttr(stateDataSourceName, "weight", strconv.Itoa(weight)),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccTwilioSIPTrunkingOriginationURL_complete(friendlyName string, enabled bool, priority int, sipURL string, weight int) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid     = twilio_sip_trunking_trunk.trunk.sid
  friendly_name = "%s"
  enabled       = %t
  priority      = %d
  sip_url       = "%s"
  weight        = %d
}

data "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid = twilio_sip_trunking_origination_url.origination_url.trunk_sid
  sid       = twilio_sip_trunking_origination_url.origination_url.sid
}
`, friendlyName, enabled, priority, sipURL, weight)
}
