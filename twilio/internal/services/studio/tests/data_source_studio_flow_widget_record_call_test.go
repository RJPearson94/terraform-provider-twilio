package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetRecordCall_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_record_call.record_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetRecordCall_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"RecordCall","properties":{"record_call":false},"transitions":[{"event":"failed"},{"event":"success"}],"type":"record-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetRecordCall_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_record_call.record_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetRecordCall_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"RecordCall","properties":{"offset":{"x":10,"y":20},"record_call":true,"recording_channels":"mono","recording_status_callback_events":"in-progress completed","recording_status_callback_method":"GET","recording_status_callback":"http://localhost.com","trim":"do-not-trim"},"transitions":[{"event":"failed","next":"RecordCall"},{"event":"success","next":"RecordCall"}],"type":"record-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetRecordCall_basic() string {
	return `
data "twilio_studio_flow_widget_record_call" "record_call" {
	name = "RecordCall"
	record_call = false
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetRecordCall_complete() string {
	return `
data "twilio_studio_flow_widget_record_call" "record_call" {
	name = "RecordCall"

	transitions {
		failed = "RecordCall"
		success = "RecordCall"
	}

	record_call = true
    trim = "do-not-trim"
    recording_status_callback_url = "http://localhost.com"
    recording_status_callback_method = "GET"
    recording_status_callback_events = [
		"in-progress", 
		"completed"
	]
    recording_channels = "mono"

    offset {
		x = 10
		y = 20
	}
}
`
}
