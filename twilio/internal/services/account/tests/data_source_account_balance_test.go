package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var accountBalanceDataSourceName = "twilio_account_balance"

func TestAccDataSourceTwilioAccountBalance_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.balance", accountBalanceDataSourceName)
	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountBalance_complete(testData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "account_sid", regexp.MustCompile(`^AC(.+)$`)),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "balance"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "currency"),
				),
			},
		},
	})
}

func testAccTwilioAccountBalance_complete(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
data "twilio_account_balance" "balance" {
  account_sid = "%s"
}
`, testData.AccountSid)
}
