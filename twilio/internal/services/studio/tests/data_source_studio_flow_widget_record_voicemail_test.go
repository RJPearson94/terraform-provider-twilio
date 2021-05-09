package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetRecordVoicemail_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_record_voicemail.record_voicemail"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetRecordVoicemail_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"RecordVoicemail","properties":{},"transitions":[{"event":"hangup"},{"event":"noAudio"},{"event":"recordingComplete"}],"type":"record-voicemail"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetRecordVoicemail_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_record_voicemail.record_voicemail"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetRecordVoicemail_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"RecordVoicemail","properties":{"finish_on_key":"1","max_length":1000,"offset":{"x":10,"y":20},"play_beep":"true","recording_status_callback_url":"http://localhost.com/recording","timeout":10,"transcribe":true,"transcription_callback_url":"http://localhost.com/transcript","trim":"trim-silence"},"transitions":[{"event":"hangup","next":"RecordVoicemail"},{"event":"noAudio","next":"RecordVoicemail"},{"event":"recordingComplete","next":"RecordVoicemail"}],"type":"record-voicemail"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetRecordVoicemail_basic() string {
	return `
data "twilio_studio_flow_widget_record_voicemail" "record_voicemail" {
	name = "RecordVoicemail"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetRecordVoicemail_complete() string {
	return `
data "twilio_studio_flow_widget_record_voicemail" "record_voicemail" {
	name = "RecordVoicemail"

	transitions {
		hangup = "RecordVoicemail"
		no_audio = "RecordVoicemail"
		recording_complete = "RecordVoicemail"
	}

	max_length = 1000
	play_beep = "true"
    recording_status_callback_url = "http://localhost.com/recording"
    timeout = 10
	finish_on_key = "1"
	transcribe = true
    transcription_callback_url = "http://localhost.com/transcript"
	trim = "trim-silence"

    offset {
		x = 10
		y = 20
	}
}
`
}
