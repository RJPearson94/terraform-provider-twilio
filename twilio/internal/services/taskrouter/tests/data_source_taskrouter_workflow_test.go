package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var workflowDataSourceName = "twilio_taskrouter_workflow"

func TestAccDataSourceTwilioTaskRouterWorkflow_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workflow", workflowDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkflow_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "fallback_assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_reservation_timeout", "120"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "configuration"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttr(stateDataSourceName, "document_content_type", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkflow_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%ss"
}

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%s"
  configuration = <<EOF
{
  "task_routing": {
    "filters": [],
    "default_filter": {
      "queue": "${twilio_taskrouter_task_queue.task_queue.sid}"
    }
  }
}
EOF
}

data "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workflow.workflow.workspace_sid
  sid           = twilio_taskrouter_workflow.workflow.sid
}
`, friendlyName, friendlyName, friendlyName)
}
