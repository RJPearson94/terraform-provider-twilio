package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_play(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_gather_input_on_call.gather_input_on_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_play(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"GatherInputOnCall","properties":{"play":"http://localhost.com"},"transitions":[{"event":"keypress"},{"event":"speech"},{"event":"timeout"}],"type":"gather-input-on-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_say(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_gather_input_on_call.gather_input_on_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_say(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"GatherInputOnCall","properties":{"say":"Hello World"},"transitions":[{"event":"keypress"},{"event":"speech"},{"event":"timeout"}],"type":"gather-input-on-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_gather_input_on_call.gather_input_on_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"GatherInputOnCall","properties":{"finish_on_key":"1","gather_language":"en-US","hints":"test,test2","language":"en-US","loop":1,"number_of_digits":3,"offset":{"x":10,"y":20},"profanity_filter":"true","say":"Hello World","speech_model":"phone_call","speech_timeout":"auto","stop_gather":true,"timeout":5,"voice":"alice"},"transitions":[{"event":"keypress","next":"GatherInputOnCall"},{"event":"speech","next":"GatherInputOnCall"},{"event":"timeout","next":"GatherInputOnCall"}],"type":"gather-input-on-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_play() string {
	return `
data "twilio_studio_flow_widget_gather_input_on_call" "gather_input_on_call" {
	name = "GatherInputOnCall"

	play = "http://localhost.com"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_say() string {
	return `
data "twilio_studio_flow_widget_gather_input_on_call" "gather_input_on_call" {
	name = "GatherInputOnCall"

	say = "Hello World"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetGatherInputOnCall_complete() string {
	return `
data "twilio_studio_flow_widget_gather_input_on_call" "gather_input_on_call" {
	name = "GatherInputOnCall"

	transitions {
		keypress = "GatherInputOnCall"
		speech = "GatherInputOnCall"
		timeout = "GatherInputOnCall"
	}

    finish_on_key = "1"
    gather_language = "en-US"
    hints = [
		"test",
		"test2"
	]
	language = "en-US"
    loop = 1
    number_of_digits = 3
    profanity_filter = "true"
	say = "Hello World"
    speech_model = "phone_call"
    speech_timeout = "auto"
    stop_gather = true
    timeout = 5
	voice = "alice"

    offset {
		x = 10
		y = 20
	}
}
`
}
