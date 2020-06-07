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

var taskQueueResourceName = "twilio_taskrouter_task_queue"

func TestAccTwilioTaskRouterTaskQueue_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "task_order"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "task_order"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "task_order"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioTaskRouterTaskQueueDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

	for _, rs := range s.RootModule().Resources {
		if rs.Type != ActivityResourceName {
			continue
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).TaskQueue(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving task queue information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioTaskRouterTaskQueueExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).TaskQueue(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving task queue information %s", err)
		}

		return nil
	}
}

func testAccTwilioTaskRouterTaskQueue_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
	friendly_name          = "%s"
	multi_task_enabled     = true
	prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
	workspace_sid = twilio_taskrouter_workspace.workspace.sid
	friendly_name = "%s"
}`, friendlyName, friendlyName)
}
