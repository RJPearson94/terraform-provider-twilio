package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var ipAccessControlListResourceName = "twilio_sip_trunking_ip_access_control_list"

func TestAccTwilioSIPTrunkingIPAccessControlList_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.ip_access_control_list", ipAccessControlListResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingIPAccessControlListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingIPAccessControlList_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingIPAccessControlListExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trunk_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ip_access_control_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSIPTrunkingIPAccessControlListImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSIPTrunkingIPAccessControlList_invalidTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingIPAccessControlList_invalidTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of trunk_sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got trunk_sid`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingIPAccessControlList_invalidIPAccessControlListSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingIPAccessControlList_invalidIPAccessControlListSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of ip_access_control_list_sid to match regular expression "\^AL\[0-9a-fA-F\]\{32\}\$", got ip_access_control_list_sid`),
			},
		},
	})
}

func testAccCheckTwilioSIPTrunkingIPAccessControlListDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

	for _, rs := range s.RootModule().Resources {
		if rs.Type != ipAccessControlListResourceName {
			continue
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).IpAccessControlList(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving IP access control list information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPTrunkingIPAccessControlListExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).IpAccessControlList(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving IP access control list information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPTrunkingIPAccessControlListImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Trunks/%s/IpAccessControlLists/%s", rs.Primary.Attributes["trunk_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPTrunkingIPAccessControlList_basic(testData *acceptance.TestData, friendlyName string) string {
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
`, testData.AccountSid, friendlyName)
}

func testAccTwilioSIPTrunkingIPAccessControlList_invalidTrunkSid() string {
	return `
resource "twilio_sip_trunking_ip_access_control_list" "ip_access_control_list" {
  trunk_sid                  = "trunk_sid"
  ip_access_control_list_sid = "ALaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioSIPTrunkingIPAccessControlList_invalidIPAccessControlListSid() string {
	return `
resource "twilio_sip_trunking_ip_access_control_list" "ip_access_control_list" {
  trunk_sid                  = "TKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  ip_access_control_list_sid = "ip_access_control_list_sid"
}
`
}
