package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const workspaceConfigurationResourceName = "twilio_taskrouter_workspace_configuration"

func TestAccTwilioTaskRouterWorkspaceConfiguration_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace_configuration", workspaceConfigurationResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspaceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspaceConfiguration_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkspaceConfiguration_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspaceConfiguration_invalidDefaultActivitySid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkspaceConfiguration_invalidDefaultActivitySid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of default_activity_sid to match regular expression "\^WA\[0-9a-fA-F\]\{32\}\$", got default_activity_sid`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspaceConfiguration_invalidTimeoutActivitySid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkspaceConfiguration_invalidTimeoutActivitySid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of timeout_activity_sid to match regular expression "\^WA\[0-9a-fA-F\]\{32\}\$", got timeout_activity_sid`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspaceConfiguration_withDefaultActivitySid(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace_configuration", workspaceConfigurationResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspaceConfiguration_withDefaultActivitySid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrPair(stateResourceName, "default_activity_name", "twilio_taskrouter_activity.activity", "friendly_name"),
					resource.TestCheckResourceAttrPair(stateResourceName, "default_activity_sid", "twilio_taskrouter_activity.activity", "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspaceConfiguration_withDefaultActivitySidSetToOffline(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspaceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspaceConfiguration_withTimeoutActivitySid(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace_configuration", workspaceConfigurationResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspaceConfiguration_withTimeoutActivitySid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "timeout_activity_name", "twilio_taskrouter_activity.activity", "friendly_name"),
					resource.TestCheckResourceAttrPair(stateResourceName, "timeout_activity_sid", "twilio_taskrouter_activity.activity", "sid"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspaceConfiguration_withTimeoutActivitySidSetToOffline(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspaceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
				),
			},
		},
	})
}

func testAccCheckTwilioTaskRouterWorkspaceConfigurationDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

	for _, rs := range s.RootModule().Resources {
		if rs.Type != workspaceConfigurationResourceName {
			continue
		}

		if _, err := client.Workspace(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving workspace information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioTaskRouterWorkspaceConfigurationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Workspace(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving workspace information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioTaskRouterWorkspaceConfiguration_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
}
`, friendlyName)
}

func testAccTwilioTaskRouterWorkspaceConfiguration_withDefaultActivitySid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  available     = true
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = twilio_taskrouter_workspace.workspace.sid
  default_activity_sid = twilio_taskrouter_activity.activity.sid
}
`, friendlyName)
}

func testAccTwilioTaskRouterWorkspaceConfiguration_withDefaultActivitySidSetToOffline(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

data "twilio_taskrouter_activities" "activities" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Offline"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  available     = true
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = twilio_taskrouter_workspace.workspace.sid
  default_activity_sid = data.twilio_taskrouter_activities.activities.activities[0].sid
}
`, friendlyName)
}

func testAccTwilioTaskRouterWorkspaceConfiguration_withTimeoutActivitySid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  available     = false
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = twilio_taskrouter_workspace.workspace.sid
  timeout_activity_sid = twilio_taskrouter_activity.activity.sid
}
`, friendlyName)
}

func testAccTwilioTaskRouterWorkspaceConfiguration_withTimeoutActivitySidSetToOffline(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  available     = false
}

data "twilio_taskrouter_activities" "activities" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Offline"
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = twilio_taskrouter_workspace.workspace.sid
  timeout_activity_sid = data.twilio_taskrouter_activities.activities.activities[0].sid
}
`, friendlyName)
}

func testAccTwilioTaskRouterWorkspaceConfiguration_invalidWorkspaceSid() string {
	return `
resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid = "workspace_sid"
}
`
}

func testAccTwilioTaskRouterWorkspaceConfiguration_invalidDefaultActivitySid() string {
	return `
resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  default_activity_sid = "default_activity_sid"
}
`
}

func testAccTwilioTaskRouterWorkspaceConfiguration_invalidTimeoutActivitySid() string {
	return `
resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  timeout_activity_sid = "timeout_activity_sid"
}
`
}
