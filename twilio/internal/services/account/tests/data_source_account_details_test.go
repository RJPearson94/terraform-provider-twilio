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

func TestAccDataSourceTwilioAccountDetails_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.details", accountDetailsDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountDetails_complete(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "sid", regexp.MustCompile(`^AC(.+)$`)),
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

// Create Sub Account to prevent leaking auth token of the parent account
func testAccTwilioAccountDetails_complete(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_account_sub_account" "sub_account" {
  friendly_name = "%s"
}

data "twilio_account_details" "details" {
  sid = twilio_account_sub_account.sub_account.sid
}
`, friendlyName)
}
