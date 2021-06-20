package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const taskChannelDataSourceName = "twilio_taskrouter_task_channel"

func TestAccDataSourceTwilioTaskRouterTaskChannel_sid(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_channel", taskChannelDataSourceName)
	friendlyName := acctest.RandString(10)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterTaskChannel_sid(friendlyName, uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "channel_optimized_routing", "false"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterTaskChannel_uniqueName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.task_channel", taskChannelDataSourceName)
	friendlyName := acctest.RandString(10)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterTaskChannel_uniqueName(friendlyName, uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "channel_optimized_routing", "false"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterTaskChannel_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterTaskChannel_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterTaskChannel_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterTaskChannel_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^TC\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterTaskChannel_sid(friendlyName string, uniqueName string) string {
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

data "twilio_taskrouter_task_channel" "task_channel" {
  workspace_sid = twilio_taskrouter_task_channel.task_channel.workspace_sid
  sid           = twilio_taskrouter_task_channel.task_channel.sid
}
`, friendlyName, uniqueName)
}

func testAccDataSourceTwilioTaskRouterTaskChannel_uniqueName(friendlyName string, uniqueName string) string {
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

data "twilio_taskrouter_task_channel" "task_channel" {
  workspace_sid = twilio_taskrouter_task_channel.task_channel.workspace_sid
  unique_name   = twilio_taskrouter_task_channel.task_channel.unique_name
}
`, friendlyName, uniqueName)
}

func testAccDataSourceTwilioTaskRouterTaskChannel_invalidWorkspaceSid() string {
	return `
data "twilio_taskrouter_task_channel" "task_channel" {
	workspace_sid = "workspace_sid"
	sid = "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioTaskRouterTaskChannel_invalidSid() string {
	return `
data "twilio_taskrouter_task_channel" "task_channel" {
	workspace_sid = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	sid = "sid"
}
`
}
