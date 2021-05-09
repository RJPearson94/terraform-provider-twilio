package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetSendToFlex_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_to_flex.send_to_flex"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendToFlex_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendToFlex","properties":{"channel":"TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","workflow":"WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"callComplete"},{"event":"callFailure"},{"event":"failedToEnqueue"}],"type":"send-to-flex"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSendToFlex_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_send_to_flex.send_to_flex"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSendToFlex_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SendToFlex","properties":{"attributes":"{\"channelSid\":\"{{trigger.message.ChannelSid}}\",\"channelType\":\"{{trigger.message.ChannelAttributes.channel_type}}\",\"name\":\"{{trigger.message.ChannelAttributes.from}}\"}","channel":"TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","offset":{"x":10,"y":20},"priority":"10","timeout":"3600","waitUrl":"https://test.com/hold","waitUrlMethod":"POST","workflow":"WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"callComplete","next":"SendToFlex"},{"event":"callFailure","next":"SendToFlex"},{"event":"failedToEnqueue","next":"SendToFlex"}],"type":"send-to-flex"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetSendToFlex_basic() string {
	return `
data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
	name = "SendToFlex"
	channel_sid = "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSendToFlex_complete() string {
	return `
data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
	name = "SendToFlex"
	
	transitions {
		call_complete = "SendToFlex"
		call_failure = "SendToFlex"
		failed_to_enqueue = "SendToFlex"
	}

	attributes = jsonencode({
		"name":"{{trigger.message.ChannelAttributes.from}}",
		"channelType":"{{trigger.message.ChannelAttributes.channel_type}}",
		"channelSid":"{{trigger.message.ChannelSid}}"
	})
	channel_sid = "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	priority = "10"
    timeout = "3600"
    wait_url = "https://test.com/hold"
	wait_url_method = "POST"
	workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	offset {
		x = 10
		y = 20
	}
}
`
}
