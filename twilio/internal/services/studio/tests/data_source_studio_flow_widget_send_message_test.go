package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetSendMessage_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_message.send_message"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendMessage_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendMessage","properties":{"body":"Hello World","from":"{{flow.channel.address}}","to":"{{contact.channel.address}}"},"transitions":[{"event":"failed"},{"event":"sent"}],"type":"send-message"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSendMessage_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_message.send_message"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendMessage_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendMessage","properties":{"attributes":"{\"channelSid\":\"{{trigger.message.ChannelSid}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"name\":\"{{trigger.message.ChannelAttributes.from}}\"}","body":"Hello World","channel":"{{trigger.message.ChannelSid}}","from":"{{flow.channel.address}}","media_url":"https://test.com","offset":{"x":10,"y":20},"service":"{{trigger.message.InstanceSid}}","to":"{{contact.channel.address}}"},"transitions":[{"event":"failed","next":"SendMessage"},{"event":"sent","next":"SendMessage"}],"type":"send-message"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetSendMessage_basic() string {
	return `
data "twilio_studio_flow_widget_send_message" "send_message" {
	name = "SendMessage"
	body = "Hello World"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSendMessage_complete() string {
	return `
data "twilio_studio_flow_widget_send_message" "send_message" {
	name = "SendMessage"

	transitions {
		failed = "SendMessage"
		sent = "SendMessage"
	}

	attributes = jsonencode({
		"name":"{{trigger.message.ChannelAttributes.from}}",
		"channelType":"{{trigger.message.ChannelAttributes.channel_type}}",
		"channelSid":"{{trigger.message.ChannelSid}}"
	})
	body = "Hello World"
    channel_sid = "{{trigger.message.ChannelSid}}"
	from = "{{flow.channel.address}}"
    media_url = "https://test.com"
	service_sid = "{{trigger.message.InstanceSid}}"
	to = "{{contact.channel.address}}"

	offset {
		x = 10
		y = 20
	}
}
`
}
