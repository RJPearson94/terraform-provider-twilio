package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetConnectVirtualAgent_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_virtual_agent.connect_virtual_agent"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectVirtualAgent_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectVirtualAgent","properties":{"connector":"test-connector"},"transitions":[{"event":"hangup"},{"event":"return"}],"type":"connect-virtual-agent"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetConnectVirtualAgent_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_virtual_agent.connect_virtual_agent"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectVirtualAgent_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectVirtualAgent","properties":{"connector":"test-connector","language":"en-US","offset":{"x":10,"y":20},"sentiment_analysis":"true","status_callback":"https://test.com"},"transitions":[{"event":"hangup","next":"ConnectVirtualAgent"},{"event":"return","next":"ConnectVirtualAgent"}],"type":"connect-virtual-agent"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetConnectVirtualAgent_basic() string {
	return `
data "twilio_studio_flow_widget_connect_virtual_agent" "connect_virtual_agent" {
	name = "ConnectVirtualAgent"
	connector = "test-connector"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetConnectVirtualAgent_complete() string {
	return `
data "twilio_studio_flow_widget_connect_virtual_agent" "connect_virtual_agent" {
	name = "ConnectVirtualAgent"

	transitions {
		hangup = "ConnectVirtualAgent"
		return = "ConnectVirtualAgent"
	}

	connector = "test-connector"
    sentiment_analysis = "true"
    language = "en-US"
    status_callback_url = "https://test.com"

    offset {
		x = 10
		y = 20
	}
}
`
}
