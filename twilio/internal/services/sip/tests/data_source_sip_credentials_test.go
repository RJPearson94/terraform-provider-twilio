package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var credentialsDataSourceName = "twilio_sip_credentials"

func TestAccDataSourceTwilioSIPCredentials_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.credentials", credentialsDataSourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	username := acctest.RandString(10)
	password := "A1" + acctest.RandString(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPCredentials_basic(testData, friendlyName, username, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credential_list_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "credentials.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "credentials.0.username", username),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credentials.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credentials.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "credentials.0.date_updated"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioSIPCredentials_basic(testData *acceptance.TestData, friendlyName string, username string, password string) string {
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

data "twilio_sip_credentials" "credentials" {
  account_sid         = twilio_sip_credential.credential.account_sid
  credential_list_sid = twilio_sip_credential.credential.credential_list_sid
}
`, testData.AccountSid, friendlyName, username, password)
}
