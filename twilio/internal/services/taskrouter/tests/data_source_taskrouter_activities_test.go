package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var activitiesDataSourceName = "twilio_taskrouter_activities"

func TestAccDataSourceTwilioTaskRouterActivities_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.activities", activitiesDataSourceName)
	friendlyName := acctest.RandString(10)

	// Twilio creates multiple activities when a workspace is created, so can't guarantee the order
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterActivities_basic(friendlyName),
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

func testAccDataSourceTwilioTaskRouterActivities_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%s"
  available     = true
}

data "twilio_taskrouter_activities" "activities" {
  workspace_sid = twilio_taskrouter_activity.activity.workspace_sid
}
`, friendlyName, friendlyName)
}
