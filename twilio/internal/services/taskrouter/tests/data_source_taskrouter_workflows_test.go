package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const workflowsDataSourceName = "twilio_taskrouter_workflows"

func TestAccDataSourceTwilioTaskRouterWorkflows_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workflows", workflowsDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkflows_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workflows.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkflows_withFriendlyName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workflows", workflowsDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkflows_withFriendlyName(friendlyName),
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

func TestAccDataSourceTwilioTaskRouterWorkflows_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterWorkflows_invalidWorkflowSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkflows_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
}

resource "twilio_taskrouter_workflow" "workflow_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
}

data "twilio_taskrouter_workflows" "workflows" {
  workspace_sid = twilio_taskrouter_workflow.workflow.workspace_sid
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkflows_withFriendlyName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
}

resource "twilio_taskrouter_workflow" "workflow_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
}

data "twilio_taskrouter_workflows" "workflows" {
  workspace_sid = twilio_taskrouter_workflow.workflow.workspace_sid
  friendly_name = "%[1]s"
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkflows_invalidWorkflowSid() string {
	return `
data "twilio_taskrouter_workflows" "workflows" {
  workspace_sid = "workspace_sid"
}
`
}
