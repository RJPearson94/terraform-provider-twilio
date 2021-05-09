package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetEnqueueCall_task(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_enqueue_call.enqueue_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetEnqueueCall_task(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"EnqueueCall","properties":{"workflow_sid":"WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"callComplete"},{"event":"callFailure"},{"event":"failedToEnqueue"}],"type":"enqueue-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetEnqueueCall_queue(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_enqueue_call.enqueue_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetEnqueueCall_queue(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"EnqueueCall","properties":{"queue_name":"Test"},"transitions":[{"event":"callComplete"},{"event":"callFailure"},{"event":"failedToEnqueue"}],"type":"enqueue-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetEnqueueCall_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_enqueue_call.enqueue_call"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetEnqueueCall_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"EnqueueCall","properties":{"offset":{"x":10,"y":20},"priority":1,"task_attributes":"{\"test\":\"test\"}","timeout":10,"wait_url":"http://localhost.com","wait_url_method":"POST","workflow_sid":"WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"callComplete","next":"EnqueueCall"},{"event":"callFailure","next":"EnqueueCall"},{"event":"failedToEnqueue","next":"EnqueueCall"}],"type":"enqueue-call"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetEnqueueCall_task() string {
	return `
data "twilio_studio_flow_widget_enqueue_call" "enqueue_call" {
	name = "EnqueueCall"
	workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetEnqueueCall_queue() string {
	return `
data "twilio_studio_flow_widget_enqueue_call" "enqueue_call" {
	name = "EnqueueCall"
	queue_name = "Test"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetEnqueueCall_complete() string {
	return `
data "twilio_studio_flow_widget_enqueue_call" "enqueue_call" {
	name = "EnqueueCall"

	transitions {
		call_complete = "EnqueueCall"
		call_failure = "EnqueueCall"
		failed_to_enqueue = "EnqueueCall"
	}
	
    priority = 1
    task_attributes = jsonencode({
		"test": "test"
	})
    timeout = 10
    wait_url = "http://localhost.com"
    wait_url_method = "POST"
	workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

    offset {
		x = 10
		y = 20
	}
}
`
}
