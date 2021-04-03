package tests

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const originationURLsDataSourceName = "twilio_sip_trunking_origination_urls"

func TestAccDataSourceTwilioSIPTrunkingOriginationURLs_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.origination_urls", originationURLsDataSourceName)

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
				Config: testAccDataSourceTwilioSIPTrunkingOriginationURLs_complete(friendlyName, enabled, priority, sipURL, weight),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "origination_urls.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "origination_urls.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "origination_urls.0.enabled", strconv.FormatBool(enabled)),
					resource.TestCheckResourceAttr(stateDataSourceName, "origination_urls.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "origination_urls.0.priority", strconv.Itoa(priority)),
					resource.TestCheckResourceAttr(stateDataSourceName, "origination_urls.0.sip_url", sipURL),
					resource.TestCheckResourceAttr(stateDataSourceName, "origination_urls.0.weight", strconv.Itoa(weight)),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "origination_urls.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "origination_urls.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "origination_urls.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPTrunkingOriginationURLs_invalidTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingOriginationURLs_invalidTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of trunk_sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got trunk_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPTrunkingOriginationURLs_complete(friendlyName string, enabled bool, priority int, sipURL string, weight int) string {
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

data "twilio_sip_trunking_origination_urls" "origination_urls" {
  trunk_sid = twilio_sip_trunking_origination_url.origination_url.trunk_sid
}
`, friendlyName, enabled, priority, sipURL, weight)
}

func testAccDataSourceTwilioSIPTrunkingOriginationURLs_invalidTrunkSid() string {
	return `
data "twilio_sip_trunking_origination_urls" "origination_urls" {
  trunk_sid = "trunk_sid"
}
`
}
