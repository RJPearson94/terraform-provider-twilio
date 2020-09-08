package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var workersDataSourceName = "twilio_taskrouter_workers"

func TestAccDataSourceTwilioTaskRouterWorkers_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkers_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%s"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_worker.worker.workspace_sid
}
`, friendlyName, friendlyName)
}
