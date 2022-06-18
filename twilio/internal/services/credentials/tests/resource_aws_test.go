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

var awsResourceName = "twilio_credentials_aws"

func TestAccTwilioCredentialsAWS_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.aws", awsResourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAWSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAWS_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAWSExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "aws_access_key_id", testData.AWSAccessKeyID),
					resource.TestCheckResourceAttr(stateResourceName, "aws_secret_access_key", testData.AWSSecretAccessKey),
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

func TestAccTwilioCredentialsAWS_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.aws", awsResourceName)

	testData := acceptance.TestAccData
	friendlyName := ""
	newFriendlyName := acctest.RandString(64)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAWSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAWS_friendlyName(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAWSExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioAWS_friendlyName(testData, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAWSExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
			{
				Config: testAccTwilioAWS_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAWSExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioCredentialsAWS_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAWS_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccCheckTwilioAWSDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Accounts

	for _, rs := range s.RootModule().Resources {
		if rs.Type != awsResourceName {
			continue
		}

		if _, err := client.Credentials.AWSCredential(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving aws credential information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAWSExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Accounts

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Credentials.AWSCredential(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving aws credential information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAWS_basic(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_credentials_aws" "aws" {
  aws_access_key_id     = "%s"
  aws_secret_access_key = "%s"
}
`, testData.AWSAccessKeyID, testData.AWSSecretAccessKey)
}

func testAccTwilioAWS_friendlyName(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_credentials_aws" "aws" {
  aws_access_key_id     = "%s"
  aws_secret_access_key = "%s"
  friendly_name         = "%s"
}
`, testData.AWSAccessKeyID, testData.AWSSecretAccessKey, friendlyName)
}

func testAccTwilioAWS_invalidAccountSid() string {
	return `
resource "twilio_credentials_aws" "aws" {
  aws_access_key_id     = "aws_access_key_id"
  aws_secret_access_key = "aws_secret_access_key"
  account_sid           = "account_sid"
}
`
}
