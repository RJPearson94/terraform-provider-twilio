package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetTrigger_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_trigger.trigger"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetTrigger_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"Trigger","properties":{},"transitions":[{"event":"incomingCall"},{"event":"incomingMessage"},{"event":"incomingParent"},{"event":"incomingRequest"}],"type":"trigger"}`),
					helper.ValidateFlowWidgetTrigger(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetTrigger_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_trigger.trigger"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetTrigger_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"Trigger","properties":{"offset":{"x":10,"y":20}},"transitions":[{"event":"incomingCall","next":"Next"},{"event":"incomingMessage","next":"Next"},{"event":"incomingParent","next":"Next"},{"event":"incomingRequest","next":"Next"}],"type":"trigger"}`),
					helper.ValidateFlowWidgetTrigger(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetTrigger_basic() string {
	return `
data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetTrigger_complete() string {
	return `
data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"

  transitions {
    incoming_call    = "Next"
    incoming_message = "Next"
    incoming_request = "Next"
    incoming_parent  = "Next"
  }

  offset {
    x = 10
    y = 20
  }
}
`
}
