package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const taskQueuesDataSourceName = "twilio_taskrouter_task_queues"

func TestAccDataSourceTwilioTaskRouterTaskQueues_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_queues", taskQueuesDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterTaskQueues_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_queues.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.event_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.target_workers", "1==1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_queues.0.task_order", "FIFO"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_queues.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_queues.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_queues.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterTaskQueues_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterTaskQueues_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterTaskQueues_basic(friendlyName string) string {
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

data "twilio_taskrouter_task_queues" "task_queues" {
  workspace_sid = twilio_taskrouter_task_queue.task_queue.workspace_sid
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterTaskQueues_invalidWorkspaceSid() string {
	return `
data "twilio_taskrouter_task_queues" "task_queues" {
  workspace_sid = "workspace_sid"
}
`
}
