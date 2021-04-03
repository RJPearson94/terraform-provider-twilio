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

const originationURLDataSourceName = "twilio_sip_trunking_origination_url"

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
				Config: testAccDataSourceTwilioSIPTrunkingOriginationURL_complete(friendlyName, enabled, priority, sipURL, weight),
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

func TestAccDataSourceTwilioSIPTrunkingOriginationURL_invalidTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingOriginationURL_invalidTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of trunk_sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got trunk_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPTrunkingOriginationURL_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPTrunkingOriginationURL_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^OU\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPTrunkingOriginationURL_complete(friendlyName string, enabled bool, priority int, sipURL string, weight int) string {
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

func testAccDataSourceTwilioSIPTrunkingOriginationURL_invalidTrunkSid() string {
	return `
data "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid = "trunk_sid"
  sid       = "OUaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPTrunkingOriginationURL_invalidSid() string {
	return `
data "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid = "TKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid       = "sid"
}
`
}
