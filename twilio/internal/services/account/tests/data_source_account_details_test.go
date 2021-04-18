package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var accountDetailsDataSourceName = "twilio_account_details"

func TestAccDataSourceTwilioAccountDetails_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.details", accountDetailsDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountDetails_providerAccountSid(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "sid", testData.AccountSid),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "owner_account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "auth_token"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "type"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountDetails_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.details", accountDetailsDataSourceName)
	subAccountStateResourceName := "twilio_account_sub_account.sub_account"

	friendlyName := acctest.RandString(10)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountDetails_complete(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(stateDataSourceName, "sid", subAccountStateResourceName, "sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "owner_account_sid", testData.AccountSid),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "auth_token"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "type"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountDetails_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountDetails_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

// Create Sub Account to prevent leaking auth token of the parent account
func testAccDataSourceTwilioAccountDetails_complete(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_account_sub_account" "sub_account" {
  friendly_name = "%s"
}

data "twilio_account_details" "details" {
  sid = twilio_account_sub_account.sub_account.sid
}
`, friendlyName)
}

func testAccDataSourceTwilioAccountDetails_providerAccountSid() string {
	return `
data "twilio_account_details" "details" {}
`
}

func testAccDataSourceTwilioAccountDetails_invalidSid() string {
	return `
data "twilio_account_details" "details" {
  sid = "sid"
}
`
}
