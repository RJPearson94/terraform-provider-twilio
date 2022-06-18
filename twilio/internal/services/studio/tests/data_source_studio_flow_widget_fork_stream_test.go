package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetForkStream_start(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_fork_stream.fork_stream"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetForkStream_start(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ForkStream","properties":{"stream_action":"start","stream_name":"test","stream_track":"inbound_track","stream_transport_type":"websocket","stream_url":"wss://test.com"},"transitions":[{"event":"next"}],"type":"fork-stream"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetForkStream_stop(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_fork_stream.fork_stream"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetForkStream_stop(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ForkStream","properties":{"stream_action":"stop","stream_transport_type":"websocket"},"transitions":[{"event":"next"}],"type":"fork-stream"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetForkStream_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_fork_stream.fork_stream"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetForkStream_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ForkStream","properties":{"offset":{"x":10,"y":20},"stream_action":"start","stream_connector":"connector","stream_name":"test","stream_parameters":[{"key":"key","value":"value"},{"key":"key2","value":"value2"}],"stream_track":"inbound_track","stream_transport_type":"siprec"},"transitions":[{"event":"next","next":"ForkStream"}],"type":"fork-stream"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetForkStream_start() string {
	return `
data "twilio_studio_flow_widget_fork_stream" "fork_stream" {
  name = "ForkStream"

  stream_action         = "start"
  stream_name           = "test"
  stream_track          = "inbound_track"
  stream_transport_type = "websocket"
  stream_url            = "wss://test.com"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetForkStream_stop() string {
	return `
data "twilio_studio_flow_widget_fork_stream" "fork_stream" {
  name = "ForkStream"

  stream_transport_type = "websocket"
  stream_action         = "stop"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetForkStream_complete() string {
	return `
data "twilio_studio_flow_widget_fork_stream" "fork_stream" {
  name = "ForkStream"

  transitions {
    next = "ForkStream"
  }

  stream_action    = "start"
  stream_connector = "connector"
  stream_name      = "test"
  stream_parameters {
    key   = "key"
    value = "value"
  }
  stream_parameters {
    key   = "key2"
    value = "value2"
  }
  stream_track          = "inbound_track"
  stream_transport_type = "siprec"

  offset {
    x = 10
    y = 20
  }
}
`
}
