package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const workflowDataSourceName = "twilio_taskrouter_workflow"

func TestAccDataSourceTwilioTaskRouterWorkflow_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workflow", workflowDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
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

func TestAccDataSourceTwilioTaskRouterWorkflow_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterWorkflow_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkflow_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterWorkflow_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^WW\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkflow_basic(friendlyName string) string {
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

data "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workflow.workflow.workspace_sid
  sid           = twilio_taskrouter_workflow.workflow.sid
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkflow_invalidWorkspaceSid() string {
	return `
data "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = "workspace_sid"
  sid           = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioTaskRouterWorkflow_invalidSid() string {
	return `
data "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid           = "sid"
}
`
}
