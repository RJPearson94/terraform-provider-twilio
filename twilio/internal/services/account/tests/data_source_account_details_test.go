package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var accountDetailsDataSourceName = "twilio_account_details"

func TestAccDataSourceTwilioAccountDetails_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.details", accountDetailsDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAccountDetailsSubAccountDestroy,
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

func testAccCheckTwilioAccountDetailsSubAccountDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != subAccountResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving account information %s", err)
		}
	}

	return nil
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
