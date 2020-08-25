package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var resourceName = "twilio_iam_api_key"

func TestAccTwilioIAMAPIKey_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAPIKey_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "account_sid", testData.AccountSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioIAMAPIKey_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAPIKey_friendlyName(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "account_sid", testData.AccountSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioIAMAPIKey_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAPIKey_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "account_sid", testData.AccountSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioAPIKey_friendlyName(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "account_sid", testData.AccountSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccCheckTwilioAPIKeyDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Key(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving account key information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioAPIKeyExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Key(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving account key information %s", err)
		}

		return nil
	}
}

func testAccTwilioAPIKey_basic(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_iam_api_key" "api_key" {
  account_sid = "%s"
}
`, testData.AccountSid)
}

func testAccTwilioAPIKey_friendlyName(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_iam_api_key" "api_key" {
  account_sid   = "%s"
  friendly_name = "%s"
}
`, testData.AccountSid, friendlyName)
}
