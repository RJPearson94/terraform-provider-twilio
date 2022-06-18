package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetSendAndWaitForReply_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_and_wait_for_reply.send_and_wait_for_reply"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendAndWaitForReply_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendAndWaitForReply","properties":{"body":"Hello World","from":"{{flow.channel.address}}","timeout":"3600"},"transitions":[{"event":"deliveryFailure"},{"event":"incomingMessage","next":"SendAndWaitForReply"},{"event":"timeout"}],"type":"send-and-wait-for-reply"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSendAndWaitForReply_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_and_wait_for_reply.send_and_wait_for_reply"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendAndWaitForReply_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendAndWaitForReply","properties":{"attributes":"{\"channelSid\":\"{{trigger.message.ChannelSid}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"name\":\"{{trigger.message.ChannelAttributes.from}}\"}","body":"Hello World","channel":"{{trigger.message.ChannelSid}}","from":"{{flow.channel.address}}","media_url":"https://test.com","offset":{"x":10,"y":20},"service":"{{trigger.message.InstanceSid}}","timeout":"300"},"transitions":[{"event":"deliveryFailure","next":"SendAndWaitForReply"},{"event":"incomingMessage","next":"SendAndWaitForReply"},{"event":"timeout","next":"SendAndWaitForReply"}],"type":"send-and-wait-for-reply"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetSendAndWaitForReply_basic() string {
	return `
data "twilio_studio_flow_widget_send_and_wait_for_reply" "send_and_wait_for_reply" {
  name = "SendAndWaitForReply"

  transitions {
    incoming_message = "SendAndWaitForReply"
  }

  body = "Hello World"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSendAndWaitForReply_complete() string {
	return `
data "twilio_studio_flow_widget_send_and_wait_for_reply" "send_and_wait_for_reply" {
  name = "SendAndWaitForReply"

  transitions {
    delivery_failure = "SendAndWaitForReply"
    incoming_message = "SendAndWaitForReply"
    timeout          = "SendAndWaitForReply"
  }

  attributes = jsonencode({
    "channelSid" : "{{trigger.message.ChannelSid}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "name" : "{{trigger.message.ChannelAttributes.from}}"
  })
  body        = "Hello World"
  channel_sid = "{{trigger.message.ChannelSid}}"
  from        = "{{flow.channel.address}}"
  media_url   = "https://test.com"
  service_sid = "{{trigger.message.InstanceSid}}"
  timeout     = "300"

  offset {
    x = 10
    y = 20
  }
}
`
}
