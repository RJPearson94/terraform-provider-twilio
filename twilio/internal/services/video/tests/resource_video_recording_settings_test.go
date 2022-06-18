package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var recordingSettingsResourceName = "twilio_video_recording_settings"

func TestAccTwilioVideoRecordingSettings_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.recording_settings", recordingSettingsResourceName)

	friendlyName := "Basic Recording Settings"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVideoRecordingSettings_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoRecordingSettingsExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "aws_credentials_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "aws_s3_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "aws_storage_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "encryption_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "encryption_key_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioVideoRecordingSettings_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.recording_settings", recordingSettingsResourceName)
	friendlyName := "Composition Settings"
	newFriendlyName := "Basic Recording Settings"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVideoRecordingSettings_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoRecordingSettingsExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "aws_credentials_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "aws_s3_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "aws_storage_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "encryption_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "encryption_key_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioVideoRecordingSettings_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoRecordingSettingsExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "aws_credentials_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "aws_s3_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "aws_storage_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "encryption_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "encryption_key_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioVideoRecordingSettings_invalidEncryptionKeySid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVideoRecordingSettings_invalidEncryptionKeySid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of encryption_key_sid to match regular expression "\^CR\[0-9a-fA-F\]\{32\}\$", got encryption_key_sid`),
			},
		},
	})
}

func TestAccTwilioVideoRecordingSettings_invalidAWSCredentialSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVideoRecordingSettings_invalidAWSCredentialSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of aws_credentials_sid to match regular expression "\^CR\[0-9a-fA-F\]\{32\}\$", got aws_credentials_sid`),
			},
		},
	})
}

func TestAccTwilioVideoRecordingSettings_invalidAWSS3URL(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVideoRecordingSettings_invalidAWSS3URL(),
				ExpectError: regexp.MustCompile(`(?s)expected "aws_s3_url" to have a host, got aws_s3_url`),
			},
		},
	})
}

func testAccCheckTwilioVideoRecordingSettingsExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Video

		// Ensure we have enough information in state to look up in API
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.RecordingSettings().Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving recording settings information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioVideoRecordingSettings_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_video_recording_settings" "recording_settings" {
  friendly_name = "%s"
}
`, friendlyName)
}

func testAccTwilioVideoRecordingSettings_invalidEncryptionKeySid() string {
	return `
resource "twilio_video_recording_settings" "recording_settings" {
  friendly_name      = "Invalid Encryption Key SID"
  encryption_key_sid = "encryption_key_sid"
}
`
}

func testAccTwilioVideoRecordingSettings_invalidAWSCredentialSid() string {
	return `
resource "twilio_video_recording_settings" "recording_settings" {
  friendly_name       = "Invalid AWS Credential SID"
  aws_credentials_sid = "aws_credentials_sid"
}
`
}

func testAccTwilioVideoRecordingSettings_invalidAWSS3URL() string {
	return `
resource "twilio_video_recording_settings" "recording_settings" {
  friendly_name = "Invalid AWS S3 URL"
  aws_s3_url    = "aws_s3_url"
}
`
}
