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

var credentialListResourceName = "twilio_sip_trunking_credential_list"

func TestAccTwilioSIPTrunkingCredentialList_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.credential_list", credentialListResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingCredentialListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingCredentialList_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingCredentialListExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trunk_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "credential_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSIPTrunkingCredentialListImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSIPTrunkingCredentialList_invalidTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingCredentialList_invalidTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of trunk_sid to match regular expression "\^TK\[0-9a-fA-F\]\{32\}\$", got trunk_sid`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingCredentialList_invalidCredentialListSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingCredentialList_invalidCredentialListSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of credential_list_sid to match regular expression "\^CL\[0-9a-fA-F\]\{32\}\$", got credential_list_sid`),
			},
		},
	})
}

func testAccCheckTwilioSIPTrunkingCredentialListDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

	for _, rs := range s.RootModule().Resources {
		if rs.Type != credentialListResourceName {
			continue
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).CredentialList(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving credential list information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPTrunkingCredentialListExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Trunk(rs.Primary.Attributes["trunk_sid"]).CredentialList(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving credential list information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPTrunkingCredentialListImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Trunks/%s/CredentialLists/%s", rs.Primary.Attributes["trunk_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPTrunkingCredentialList_basic(testData *acceptance.TestData, friendlyName string) string {
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
`, testData.AccountSid, friendlyName)
}

func testAccTwilioSIPTrunkingCredentialList_invalidTrunkSid() string {
	return `
resource "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid           = "trunk_sid"
  credential_list_sid = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioSIPTrunkingCredentialList_invalidCredentialListSid() string {
	return `
resource "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid           = "TKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  credential_list_sid = "credential_list_sid"
}
`
}
