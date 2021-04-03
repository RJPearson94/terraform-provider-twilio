package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var accountBalanceDataSourceName = "twilio_account_balance"

func TestAccDataSourceTwilioAccountBalance_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.balance", accountBalanceDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountBalance_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "account_sid", regexp.MustCompile(`^AC(.+)$`)),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "balance"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "currency"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountBalance_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountBalance_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAccountBalance_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
data "twilio_account_balance" "balance" {
  account_sid = "%s"
}
`, testData.AccountSid)
}

func testAccDataSourceTwilioAccountBalance_invalidAccountSid() string {
	return `
data "twilio_account_balance" "balance" {
  account_sid = "account_sid"
}
`
}
