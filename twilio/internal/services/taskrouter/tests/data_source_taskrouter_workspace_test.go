package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const workspaceDataSourceName = "twilio_taskrouter_workspace"

func TestAccDataSourceTwilioTaskRouterWorkspace_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workspace", workspaceDataSourceName)
	friendlyName := acctest.RandString(10)
	queueOrder := "FIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkspace_basic(friendlyName, queueOrder),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "event_callback_url", ""),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "event_filters"),
					resource.TestCheckResourceAttr(stateDataSourceName, "multi_task_enabled", "true"),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "template"),
					resource.TestCheckResourceAttr(stateDataSourceName, "prioritize_queue_order", queueOrder),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "timeout_activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkspace_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterWorkspace_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkspace_basic(friendlyName string, queueOrder string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "%s"
}

data "twilio_taskrouter_workspace" "workspace" {
  sid = twilio_taskrouter_workspace.workspace.sid
}
`, friendlyName, queueOrder)
}

func testAccDataSourceTwilioTaskRouterWorkspace_invalidSid() string {
	return `
data "twilio_taskrouter_workspace" "workspace" {
  sid = "workspace_sid"
}
`
}
