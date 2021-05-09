package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetMakeOutgoingCall_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_make_outgoing_call.make_outgoing_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetMakeOutgoingCall_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"MakeOutgoingCall","properties":{"from":"{{flow.channel.address}}","to":"{{contact.channel.address}}"},"transitions":[{"event":"answered"},{"event":"busy"},{"event":"failed"},{"event":"noAnswer"}],"type":"make-outgoing-call-v2"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetMakeOutgoingCall_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_make_outgoing_call.make_outgoing_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetMakeOutgoingCall_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"MakeOutgoingCall","properties":{"detect_answering_machine":true,"from":"{{flow.channel.address}}","machine_detection":"Enable","machine_detection_silence_timeout":"2000","machine_detection_speech_end_threshold":"500","machine_detection_speech_threshold":"1000","machine_detection_timeout":"10","offset":{"x":10,"y":20},"record":true,"send_digits":"1234","sip_auth_password":"test2","sip_auth_username":"test","timeout":5,"to":"{{contact.channel.address}}","trim":"trim-silence"},"transitions":[{"event":"answered","next":"MakeOutgoingCall"},{"event":"busy","next":"MakeOutgoingCall"},{"event":"failed","next":"MakeOutgoingCall"},{"event":"noAnswer","next":"MakeOutgoingCall"}],"type":"make-outgoing-call-v2"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetMakeOutgoingCall_basic() string {
	return `
data "twilio_studio_flow_widget_make_outgoing_call" "make_outgoing_call" {
	name = "MakeOutgoingCall"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetMakeOutgoingCall_complete() string {
	return `
data "twilio_studio_flow_widget_make_outgoing_call" "make_outgoing_call" {
	name = "MakeOutgoingCall"

	transitions {
		answered = "MakeOutgoingCall"
		busy = "MakeOutgoingCall"
		failed = "MakeOutgoingCall"
		no_answer = "MakeOutgoingCall"
	}

	detect_answering_machine = true
	from = "{{flow.channel.address}}"
    to = "{{contact.channel.address}}"
    machine_detection = "Enable"
	machine_detection_speech_end_threshold = "500"
    machine_detection_speech_threshold = "1000"
	machine_detection_silence_timeout = "2000"
    machine_detection_timeout = "10"
	record = true
    send_digits = "1234"
	sip_auth_password = "test2"
    sip_auth_username = "test"
    timeout = 5
    trim = "trim-silence"

    offset {
		x = 10
		y = 20
	}
}
`
}
