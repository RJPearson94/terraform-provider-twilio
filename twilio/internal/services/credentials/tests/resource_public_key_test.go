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

var publicKeyResourceName = "twilio_credentials_public_key"

func TestAccTwilioCredentialsPublicKey_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.public_key", publicKeyResourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioPublicKey_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioPublicKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "public_key", testData.PublicKey),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioCredentialsPublicKey_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.public_key", publicKeyResourceName)

	testData := acceptance.TestAccData
	friendlyName := ""
	newFriendlyName := acctest.RandString(64)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioPublicKey_friendlyName(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioPublicKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioPublicKey_friendlyName(testData, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioPublicKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
			{
				Config: testAccTwilioPublicKey_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioPublicKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioCredentialsPublicKey_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioPublicKey_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccCheckTwilioPublicKeyDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Accounts

	for _, rs := range s.RootModule().Resources {
		if rs.Type != publicKeyResourceName {
			continue
		}

		if _, err := client.Credentials.PublicKey(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving public key information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioPublicKeyExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Accounts

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Credentials.PublicKey(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving public key information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioPublicKey_basic(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_credentials_public_key" "public_key" {
  public_key = "%s"
}
`, testData.PublicKey)
}

func testAccTwilioPublicKey_friendlyName(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_credentials_public_key" "public_key" {
  public_key    = "%s"
  friendly_name = "%s"
}
`, testData.PublicKey, friendlyName)
}

func testAccTwilioPublicKey_invalidAccountSid() string {
	return `
resource "twilio_credentials_public_key" "public_key" {
  public_key = "invalid_account_sid"
	account_sid = "account_sid"
}
`
}
