package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const taskChannelsDataSourceName = "twilio_taskrouter_task_channels"

func TestAccDataSourceTwilioTaskRouterTaskChannels_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_channels", taskChannelsDataSourceName)
	friendlyName := acctest.RandString(10)
	uniqueName := acctest.RandString(10)

	// Twilio creates multiple task channels when a workspace is created, so can't guarantee the order
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterTaskChannels_basic(friendlyName, uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.0.unique_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.0.channel_optimized_routing"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "task_channels.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterTaskChannels_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterTaskChannels_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterTaskChannels_basic(friendlyName string, uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_channel" "task_channel" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  unique_name   = "%[2]s"
}

data "twilio_taskrouter_task_channels" "task_channels" {
  workspace_sid = twilio_taskrouter_task_channel.task_channel.workspace_sid
}
`, friendlyName, uniqueName)
}

func testAccDataSourceTwilioTaskRouterTaskChannels_invalidWorkspaceSid() string {
	return `
data "twilio_taskrouter_task_channels" "task_channels" {
  workspace_sid = "workspace_sid"
}
`
}
