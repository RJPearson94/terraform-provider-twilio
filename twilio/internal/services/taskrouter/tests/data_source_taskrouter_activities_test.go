package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const activitiesDataSourceName = "twilio_taskrouter_activities"

func TestAccDataSourceTwilioTaskRouterActivities_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.activities", activitiesDataSourceName)
	friendlyName := acctest.RandString(10)

	// Twilio creates multiple activities when a workspace is created, so can't guarantee the order
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterActivities_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.#"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterActivities_withAvailable(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.activities", activitiesDataSourceName)
	friendlyName := acctest.RandString(10)

	// Twilio creates multiple activities when a workspace is created, so can't guarantee the order
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterActivities_withAvailableFalse(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.#"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterActivities_withFriendlyName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.activities", activitiesDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterActivities_withFriendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "activities.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "activities.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "activities.0.available", "true"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "activities.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterActivities_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterActivities_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterActivities_basic(friendlyName string) string {
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

resource "twilio_taskrouter_activity" "activity_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
  available     = true
}

data "twilio_taskrouter_activities" "activities" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid

	depends_on = [
		twilio_taskrouter_activity.activity,
		twilio_taskrouter_activity.activity_2
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterActivities_withFriendlyName(friendlyName string) string {
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

resource "twilio_taskrouter_activity" "activity_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
  available     = true
}

data "twilio_taskrouter_activities" "activities" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	friendly_name = "%[1]s"

	depends_on = [
		twilio_taskrouter_activity.activity,
		twilio_taskrouter_activity.activity_2
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterActivities_withAvailableFalse(friendlyName string) string {
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

resource "twilio_taskrouter_activity" "activity_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
  available     = false
}

data "twilio_taskrouter_activities" "activities" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	available     = false

	depends_on = [
		twilio_taskrouter_activity.activity,
		twilio_taskrouter_activity.activity_2
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterActivities_invalidWorkspaceSid() string {
	return `
data "twilio_taskrouter_activities" "activities" {
  workspace_sid = "workspace_sid"
}
`
}
