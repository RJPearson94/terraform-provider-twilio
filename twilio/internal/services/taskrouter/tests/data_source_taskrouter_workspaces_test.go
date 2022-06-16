package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const workspacesDataSourceName = "twilio_taskrouter_workspaces"

func TestAccDataSourceTwilioTaskRouterWorkspaces_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workspaces", workspacesDataSourceName)
	friendlyName := acctest.RandString(10)
	queueOrder := "FIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkspaces_basic(friendlyName, queueOrder),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "workspaces.#", regexp.MustCompile("^[1-9]([0-9]+)?$")),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkspaces_withFriendlyName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workspaces", workspacesDataSourceName)
	friendlyName := acctest.RandString(10)
	queueOrder := "FIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkspaces_withFriendlyName(friendlyName, queueOrder),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workspaces.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workspaces.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workspaces.0.event_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "workspaces.0.event_filters.#", "0"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workspaces.0.multi_task_enabled", "true"),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "workspaces.0.template"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workspaces.0.prioritize_queue_order", queueOrder),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.default_activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.timeout_activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspaces.0.url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkspaces_basic(friendlyName string, queueOrder string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "%s"
}

data "twilio_taskrouter_workspaces" "workspaces" {
  depends_on = [
    twilio_taskrouter_workspace.workspace
  ]
}
`, friendlyName, queueOrder)
}

func testAccDataSourceTwilioTaskRouterWorkspaces_withFriendlyName(friendlyName string, queueOrder string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "%[2]s"
}

data "twilio_taskrouter_workspaces" "workspaces" {
  friendly_name = "%[1]s"

  depends_on = [
    twilio_taskrouter_workspace.workspace
  ]
}
`, friendlyName, queueOrder)
}
