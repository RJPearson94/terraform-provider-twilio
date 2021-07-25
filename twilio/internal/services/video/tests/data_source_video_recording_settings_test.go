package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var recordingSettingsDataSourceName = "twilio_video_recording_settings"

func TestAccDataSourceTwilioVideoRecordingSettings_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.recording_settings", recordingSettingsDataSourceName)
	friendlyName := "Basic Recording Settings"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountRecordingSettings_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "aws_credentials_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "aws_s3_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "aws_storage_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "encryption_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "encryption_key_sid", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioAccountRecordingSettings_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_video_recording_settings" "recording_settings" {
  friendly_name = "%s"
}

data "twilio_video_recording_settings" "recording_settings" {
  depends_on = [
		twilio_video_recording_settings.recording_settings
	]
}
`, friendlyName)
}
