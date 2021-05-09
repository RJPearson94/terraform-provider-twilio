package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetSayPlay_say(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_say_play.say_play"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSayPlay_say(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SayPlay","properties":{"say":"Hello World"},"transitions":[{"event":"audioComplete"}],"type":"say-play"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSayPlay_play(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_say_play.say_play"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSayPlay_play(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SayPlay","properties":{"play":"http://localhost.com"},"transitions":[{"event":"audioComplete"}],"type":"say-play"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSayPlay_digits(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_say_play.say_play"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSayPlay_digits(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SayPlay","properties":{"digits":"123"},"transitions":[{"event":"audioComplete"}],"type":"say-play"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSayPlay_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_say_play.say_play"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSayPlay_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SayPlay","properties":{"language":"en-US","loop":2,"offset":{"x":10,"y":20},"say":"Test","voice":"alice"},"transitions":[{"event":"audioComplete","next":"SayPlay"}],"type":"say-play"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetSayPlay_say() string {
	return `
data "twilio_studio_flow_widget_say_play" "say_play" {
	name = "SayPlay"
	say = "Hello World"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSayPlay_play() string {
	return `
data "twilio_studio_flow_widget_say_play" "say_play" {
	name = "SayPlay"
	play = "http://localhost.com"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSayPlay_digits() string {
	return `
data "twilio_studio_flow_widget_say_play" "say_play" {
	name = "SayPlay"
	digits = "123"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSayPlay_complete() string {
	return `
data "twilio_studio_flow_widget_say_play" "say_play" {
	name = "SayPlay"

	transitions {
		audio_complete = "SayPlay"
	}
	
	language = "en-US"
	loop = 2
    say = "Test"
    voice = "alice"

    offset {
		x = 10
		y = 20
	}
}
`
}
