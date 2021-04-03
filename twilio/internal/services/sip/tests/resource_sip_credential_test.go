package tests

import (
	"fmt"
	"regexp"
	"strings"
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
	password := "A1" + acctest.RandString(10)

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

func TestAccTwilioSIPCredential_username(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.credential", credentialResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(1)
	newUsername := acctest.RandString(32)
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
					resource.TestCheckResourceAttr(stateResourceName, "username", username),
				),
			},
			{
				Config: testAccTwilioSIPCredential_basic(testData, friendlyName, newUsername, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPCredentialExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "username", newUsername),
				),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidUsernameWithLengthOf0(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := ""
	password := "A1" + acctest.RandString(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				ExpectError: regexp.MustCompile(`(?s)expected length of username to be in the range \(1 - 32\), got `),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidUsernameWithLengthOf33(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	password := "A1" + acctest.RandString(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				ExpectError: regexp.MustCompile(`(?s)expected length of username to be in the range \(1 - 32\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioSIPCredential_password(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.credential", credentialResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(1)
	password := "A1" + acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPCredentialDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPCredentialExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "password", password),
				),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidPasswordWith11Characters(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(1)
	password := "A1" + acctest.RandString(9)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				ExpectError: regexp.MustCompile(`(?s)invalid value for password \(Must contain at least 12 characters\)`),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidPasswordWithNoUppercaseCharacter(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(1)
	password := strings.ToLower("1" + acctest.RandString(11))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				ExpectError: regexp.MustCompile(`(?s)invalid value for password \(Must contain a uppercase letter\)`),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidPasswordWithNoLowercaseCharacter(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(1)
	password := strings.ToUpper("1" + acctest.RandString(11))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				ExpectError: regexp.MustCompile(`(?s)invalid value for password \(Must contain a lowercase letter\)`),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidPasswordWithNoNumber(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(1)
	password := "A" + acctest.RandStringFromCharSet(11, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_basic(testData, friendlyName, username, password),
				ExpectError: regexp.MustCompile(`(?s)invalid value for password \(Must contain a number\)`),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccTwilioSIPCredential_invalidCredentialListSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPCredential_invalidCredentialListSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of credential_list_sid to match regular expression "\^CL\[0-9a-fA-F\]\{32\}\$", got credential_list_sid`),
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

func testAccTwilioSIPCredential_invalidAccountSid() string {
	password := "A1" + acctest.RandString(12)

	return fmt.Sprintf(`
resource "twilio_sip_credential" "credential" {
  account_sid         = "account_sid"
  credential_list_sid = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  username            = "test"
  password            = "%s"
}
`, password)
}

func testAccTwilioSIPCredential_invalidCredentialListSid() string {
	password := "A1" + acctest.RandString(12)

	return fmt.Sprintf(`
resource "twilio_sip_credential" "credential" {
  account_sid         = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  credential_list_sid = "credential_list_sid"
  username            = "test"
  password            = "%s"
}
`, password)
}
