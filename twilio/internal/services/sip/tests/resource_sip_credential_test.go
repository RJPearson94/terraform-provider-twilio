package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var credentialResourceName = "twilio_sip_credential"

func TestAccTwilioSIPCredential_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.credential", credentialResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(10)
	password := "A1" + acctest.RandString(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPCredentialDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPCredentialExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "username", username),
					resource.TestCheckResourceAttr(stateResourceName, "password", password),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "credential_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:            stateResourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccTwilioSIPCredentialImportStateIdFunc(stateResourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

func TestAccTwilioSIPCredential_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.credential", credentialResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(10)
	password := "A1" + acctest.RandString(12)
	newPassword := "B2" + acctest.RandString(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPCredentialDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPCredentialExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "username", username),
					resource.TestCheckResourceAttr(stateResourceName, "password", password),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "credential_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioSIPCredential_basic(testData, friendlyName, username, newPassword),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPCredentialExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "username", username),
					resource.TestCheckResourceAttr(stateResourceName, "password", newPassword),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "credential_list_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccCheckTwilioSIPCredentialDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != credentialResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.CredentialList(rs.Primary.Attributes["credential_list_sid"]).Credential(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving credential information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPCredentialExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.CredentialList(rs.Primary.Attributes["credential_list_sid"]).Credential(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving credential information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPCredentialImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/SIP/CredentialLists/%s/Credentials/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["credential_list_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPCredential_basic(testData *acceptance.TestData, friendlyName string, username string, password string) string {
	return fmt.Sprintf(`
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "%s"
  friendly_name = "%s"
}

resource "twilio_sip_credential" "credential" {
  account_sid         = twilio_sip_credential_list.credential_list.account_sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
  username            = "%s"
  password            = "%s"
}
`, testData.AccountSid, friendlyName, username, password)
}
