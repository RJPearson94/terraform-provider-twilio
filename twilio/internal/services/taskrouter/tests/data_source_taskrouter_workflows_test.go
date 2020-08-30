package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var workflowsDataSourceName = "twilio_taskrouter_workflows"

func TestAccDataSourceTwilioTaskRouterWorkflows_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workflows", workflowsDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkflows_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workflows.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workflows.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workflows.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workflows.0.fallback_assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "workflows.0.assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "workflows.0.task_reservation_timeout", "120"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workflows.0.configuration"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workflows.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workflows.0.date_updated"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workflows.0.document_content_type", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workflows.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkflows_basic(friendlyName string) string {
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

data "twilio_taskrouter_workflows" "workflows" {
  workspace_sid = twilio_taskrouter_workflow.workflow.workspace_sid
}
`, friendlyName, friendlyName, friendlyName)
}
