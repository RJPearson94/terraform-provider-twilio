package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetSendToAutopilot_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_to_autopilot.send_to_autopilot"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendToAutopilot_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendToAutopilot","properties":{"autopilot_assistant_sid":"UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","body":"{{trigger.Message.Body}}","from":"{{flow.channel.address}}","timeout":14400},"transitions":[{"event":"failure"},{"event":"sessionEnded"},{"event":"timeout"}],"type":"send-to-auto-pilot"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSendToAutopilot_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_to_autopilot.send_to_autopilot"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendToAutopilot_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendToAutopilot","properties":{"autopilot_assistant_sid":"UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","body":"Hello World","chat_attributes":"{\"channelSid\":\"{{trigger.message.ChannelSid}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"name\":\"{{trigger.message.ChannelAttributes.from}}\"}","chat_channel":"{{trigger.message.ChannelSid}}","chat_service":"{{trigger.message.InstanceSid}}","from":"test","memory_parameters":[{"key":"key","value":"value"},{"key":"key2","value":"value2"}],"offset":{"x":10,"y":20},"target_task":"Task","timeout":100},"transitions":[{"event":"failure","next":"SendToAutopilot"},{"event":"sessionEnded","next":"SendToAutopilot"},{"event":"timeout","next":"SendToAutopilot"}],"type":"send-to-auto-pilot"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetSendToAutopilot_basic() string {
	return `
data "twilio_studio_flow_widget_send_to_autopilot" "send_to_autopilot" {
  name                    = "SendToAutopilot"
  autopilot_assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSendToAutopilot_complete() string {
	return `
data "twilio_studio_flow_widget_send_to_autopilot" "send_to_autopilot" {
  name = "SendToAutopilot"

  transitions {
    failure       = "SendToAutopilot"
    session_ended = "SendToAutopilot"
    timeout       = "SendToAutopilot"
  }

  attributes = jsonencode({
    "name" : "{{trigger.message.ChannelAttributes.from}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "channelSid" : "{{trigger.message.ChannelSid}}"
  })
  autopilot_assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  body                    = "Hello World"
  channel_sid             = "{{trigger.message.ChannelSid}}"
  from                    = "test"
  memory_parameters {
    key   = "key"
    value = "value"
  }
  memory_parameters {
    key   = "key2"
    value = "value2"
  }
  service_sid = "{{trigger.message.InstanceSid}}"
  target_task = "Task"
  timeout     = 100

  offset {
    x = 10
    y = 20
  }
}
`
}
