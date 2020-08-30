package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var taskQueueDataSourceName = "twilio_taskrouter_task_queue"

func TestAccDataSourceTwilioTaskRouterTaskQueue_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_queue", taskQueueDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "event_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "target_workers", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_order"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterTaskQueue_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%s"
}

data "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_task_queue.task_queue.workspace_sid
  sid           = twilio_taskrouter_task_queue.task_queue.sid
}
`, friendlyName, friendlyName)
}
