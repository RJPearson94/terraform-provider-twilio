package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var workspaceResourceName = "twilio_taskrouter_workspace"

func TestAccTwilioTaskRouterWorkspace_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace", workspaceResourceName)
	friendlyName := acctest.RandString(10)
	queueOrder := "FIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioTaskRouterWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspace_basic(friendlyName, queueOrder),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "events_filter", ""),
					resource.TestCheckResourceAttr(stateResourceName, "multi_task_enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "prioritize_queue_order", queueOrder),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspace_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace", workspaceResourceName)

	friendlyName := acctest.RandString(10)
	queueOrder := "FIFO"
	newQueueOrder := "LIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioTaskRouterWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspace_basic(friendlyName, queueOrder),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "events_filter", ""),
					resource.TestCheckResourceAttr(stateResourceName, "multi_task_enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "prioritize_queue_order", queueOrder),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspace_basic(friendlyName, newQueueOrder),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "events_filter", ""),
					resource.TestCheckResourceAttr(stateResourceName, "multi_task_enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "template", ""),
					resource.TestCheckResourceAttr(stateResourceName, "prioritize_queue_order", newQueueOrder),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "timeout_activity_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioTaskRouterWorkspaceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

	for _, rs := range s.RootModule().Resources {
		if rs.Type != workspaceResourceName {
			continue
		}

		if _, err := client.Workspace(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving workspace information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioTaskRouterWorkspaceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Workspace(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving workspace information %s", err)
		}

		return nil
	}
}

func testAccTwilioTaskRouterWorkspace_basic(friendlyName string, queueOrder string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
	friendly_name          = "%s"
	multi_task_enabled     = true
	prioritize_queue_order = "%s"
}`, friendlyName, queueOrder)
}
