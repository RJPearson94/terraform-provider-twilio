package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var credentialDataSourceName = "twilio_sip_credential"

func TestAccDataSourceTwilioSIPCredential_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.credential", credentialDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(10)
	password := "A1" + acctest.RandString(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPCredential_basic(testData, friendlyName, username, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttr(stateDataSourceName, "username", username),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_list_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPCredential_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPCredential_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPCredential_invalidCredentialListSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPCredential_invalidCredentialListSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of credential_list_sid to match regular expression "\^CL\[0-9a-fA-F\]\{32\}\$", got credential_list_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioSIPCredential_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSIPCredential_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^CR\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSIPCredential_basic(testData *acceptance.TestData, friendlyName string, username string, password string) string {
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

data "twilio_sip_credential" "credential" {
  account_sid         = twilio_sip_credential.credential.account_sid
  credential_list_sid = twilio_sip_credential.credential.credential_list_sid
  sid                 = twilio_sip_credential.credential.sid
}
`, testData.AccountSid, friendlyName, username, password)
}

func testAccDataSourceTwilioSIPCredential_invalidAccountSid() string {
	return `
data "twilio_sip_credential" "credential" {
  account_sid         = "account_sid"
  credential_list_sid = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid                 = "CRaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPCredential_invalidCredentialListSid() string {
	return `
data "twilio_sip_credential" "credential" {
  account_sid         = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  credential_list_sid = "credential_list_sid"
  sid                 = "CRaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioSIPCredential_invalidSid() string {
	return `
data "twilio_sip_credential" "credential" {
  account_sid         = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  credential_list_sid = "CLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid                 = "sid"
}
`
}
