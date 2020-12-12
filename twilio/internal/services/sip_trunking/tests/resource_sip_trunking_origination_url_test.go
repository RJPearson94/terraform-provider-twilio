package tests

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var originationURLResourceName = "twilio_sip_trunking_origination_url"

func TestAccTwilioSIPTrunkingOriginationURL_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.origination_url", originationURLResourceName)

	friendlyName := acctest.RandString(10)
	weight := 0
	priority := 0
	enabled := false
	sipURL := "sip:test@test.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingOriginationURLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingOriginationURL_basic(friendlyName, enabled, priority, sipURL, weight),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingOriginationURLExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", strconv.FormatBool(enabled)),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "priority", strconv.Itoa(priority)),
					resource.TestCheckResourceAttr(stateResourceName, "sip_url", sipURL),
					resource.TestCheckResourceAttr(stateResourceName, "weight", strconv.Itoa(weight)),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSIPTrunkingOriginationURLImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSIPTrunkingOriginationURL_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.origination_url", originationURLResourceName)

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)
	weight := 0
	priority := 0
	enabled := false
	sipURL := "sip:test@test.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingOriginationURLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingOriginationURL_basic(friendlyName, enabled, priority, sipURL, weight),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingOriginationURLExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", strconv.FormatBool(enabled)),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "priority", strconv.Itoa(priority)),
					resource.TestCheckResourceAttr(stateResourceName, "sip_url", sipURL),
					resource.TestCheckResourceAttr(stateResourceName, "weight", strconv.Itoa(weight)),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingOriginationURL_basic(newFriendlyName, enabled, priority, sipURL, weight),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingOriginationURLExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trunk_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", strconv.FormatBool(enabled)),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "priority", strconv.Itoa(priority)),
					resource.TestCheckResourceAttr(stateResourceName, "sip_url", sipURL),
					resource.TestCheckResourceAttr(stateResourceName, "weight", strconv.Itoa(weight)),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioSIPTrunkingOriginationURLDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

	for _, rs := range s.RootModule().Resources {
		if rs.Type != originationURLResourceName {
			continue
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).OriginationURL(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving origination url information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPTrunkingOriginationURLExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).OriginationURL(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving origination url information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPTrunkingOriginationURLImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Trunks/%s/OriginationUrls/%s", rs.Primary.Attributes["trunk_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPTrunkingOriginationURL_basic(friendlyName string, enabled bool, priority int, sipURL string, weight int) string {
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
`, friendlyName, enabled, priority, sipURL, weight)
}
