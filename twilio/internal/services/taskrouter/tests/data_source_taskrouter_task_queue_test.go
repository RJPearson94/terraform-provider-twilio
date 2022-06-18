package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const taskQueueDataSourceName = "twilio_taskrouter_task_queue"

func TestAccDataSourceTwilioTaskRouterTaskQueue_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_queue", taskQueueDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "target_workers", "1==1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "task_order", "FIFO"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterTaskQueue_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterTaskQueue_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterTaskQueue_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterTaskQueue_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^WQ\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterTaskQueue_basic(friendlyName string) string {
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

data "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_task_queue.task_queue.workspace_sid
  sid           = twilio_taskrouter_task_queue.task_queue.sid
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterTaskQueue_invalidWorkspaceSid() string {
	return `
data "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = "workspace_sid"
  sid           = "WQaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioTaskRouterTaskQueue_invalidSid() string {
	return `
data "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid           = "sid"
}
`
}
