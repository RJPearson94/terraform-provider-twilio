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

var resourceName = "twilio_iam_api_key"

func TestAccTwilioIAMAPIKey_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
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

func TestAccTwilioIAMAPIKey_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
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

func TestAccTwilioIAMAPIKey_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(64)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAPIKey_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
			{
				Config: testAccTwilioAPIKey_friendlyName(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioAPIKey_friendlyName(testData, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
			{
				Config: testAccTwilioAPIKey_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioIAMAPIKey_invalidFriendlyNameWithLengthOf0(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAPIKey_friendlyName(testData, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 64\), got `),
			},
		},
	})
}

func TestAccTwilioIAMAPIKey_invalidFriendlyNameWithLengthOf65(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := "7y80krlx0npe98jtdhahyvx8jvfz09x21x226uxj8gowkun6dgl2p1xj819qjzgtt"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAPIKey_friendlyName(testData, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 64\), got 7y80krlx0npe98jtdhahyvx8jvfz09x21x226uxj8gowkun6dgl2p1xj819qjzgtt`),
			},
		},
	})
}

func TestAccTwilioIAMAPIKey_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAPIKey_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
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
			return fmt.Errorf("Error occurred when retrieving account key information %s", err.Error())
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
			return fmt.Errorf("Error occurred when retrieving account key information %s", err.Error())
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

func testAccTwilioAPIKey_invalidAccountSid() string {
	return `
resource "twilio_iam_api_key" "api_key" {
  account_sid   = "account_sid"
  friendly_name = "invalid_account_sid"
}
`
}
