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

const taskQueueResourceName = "twilio_taskrouter_task_queue"

func TestAccTwilioTaskRouterTaskQueue_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)
	workspaceStateResourceName := "twilio_taskrouter_workspace.workspace"

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "workspace_sid", workspaceStateResourceName, "sid"),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", "1==1"),
					resource.TestCheckResourceAttr(stateResourceName, "task_order", "FIFO"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioTaskRouterTaskQueueImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)
	workspaceStateResourceName := "twilio_taskrouter_workspace.workspace"

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "workspace_sid", workspaceStateResourceName, "sid"),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", "1==1"),
					resource.TestCheckResourceAttr(stateResourceName, "task_order", "FIFO"),
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
					resource.TestCheckResourceAttrPair(stateResourceName, "workspace_sid", workspaceStateResourceName, "sid"),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", "1==1"),
					resource.TestCheckResourceAttr(stateResourceName, "task_order", "FIFO"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterTaskQueue_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_maxReservedWorkers(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)

	friendlyName := acctest.RandString(10)
	newTaskOrder := "LIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "task_order", "FIFO"),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_taskOrder(friendlyName, newTaskOrder),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "task_order", newTaskOrder),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "task_order", "FIFO"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_invalidMaxReservedWorkersOf0(t *testing.T) {
	friendlyName := acctest.RandString(10)
	maxReservedWorkers := 0

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterTaskQueue_maxReservedWorkers(friendlyName, maxReservedWorkers),
				ExpectError: regexp.MustCompile(`(?s)expected max_reserved_workers to be in the range \(1 - 50\), got 0`),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_invalidMaxReservedWorkersOf51(t *testing.T) {
	friendlyName := acctest.RandString(10)
	maxReservedWorkers := 51

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterTaskQueue_maxReservedWorkers(friendlyName, maxReservedWorkers),
				ExpectError: regexp.MustCompile(`(?s)expected max_reserved_workers to be in the range \(1 - 50\), got 51`),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_targetWorkers(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)

	friendlyName := acctest.RandString(10)
	newTargetWorkers := `(languages HAS 'english')`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", "1==1"),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_targetWorkers(friendlyName, newTargetWorkers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", newTargetWorkers),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "target_workers", "1==1"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_taskOrder(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)

	friendlyName := acctest.RandString(10)
	newMaxReservedWorkers := 10

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_maxReservedWorkers(friendlyName, newMaxReservedWorkers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "10"),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max_reserved_workers", "1"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_invalidTaskOrder(t *testing.T) {
	friendlyName := acctest.RandString(10)
	taskOrder := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterTaskQueue_taskOrder(friendlyName, taskOrder),
				ExpectError: regexp.MustCompile(`(?s)expected task_order to be one of \["LIFO" "FIFO"\], got test`),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_assignmentActivitySid(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)
	activityStateResourceName := "twilio_taskrouter_activity.activity"

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_assignmentActivitySid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttrPair(stateResourceName, "assignment_activity_sid", activityStateResourceName, "sid"),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_detachAssignmentActivitySid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_activity_sid", ""),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_invalidAssignmentActivitySid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterTaskQueue_invalidAssignmentActivitySid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of assignment_activity_sid to match regular expression "\^WA\[0-9a-fA-F\]\{32\}\$", got assignment_activity_sid`),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_reservationActivitySid(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task_queue", taskQueueResourceName)
	activityStateResourceName := "twilio_taskrouter_activity.activity"

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterTaskQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_reservationActivitySid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttrPair(stateResourceName, "reservation_activity_sid", activityStateResourceName, "sid"),
				),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_detachReservationActivitySid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", "")),
			},
			{
				Config: testAccTwilioTaskRouterTaskQueue_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterTaskQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reservation_activity_sid", ""),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_invalidReservationActivitySid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterTaskQueue_invalidReservationActivitySid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of reservation_activity_sid to match regular expression "\^WA\[0-9a-fA-F\]\{32\}\$", got reservation_activity_sid`),
			},
		},
	})
}

func TestAccTwilioTaskRouterTaskQueue_blankFriendlyName(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterTaskQueue_blankFriendlyName(),
				ExpectError: regexp.MustCompile(`(?s)expected \"friendly_name\" to not be an empty string, got `),
			},
		},
	})
}

func testAccCheckTwilioTaskRouterTaskQueueDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

	for _, rs := range s.RootModule().Resources {
		if rs.Type != taskQueueResourceName {
			continue
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).TaskQueue(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving task queue information %s", err.Error())
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

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).TaskQueue(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving task queue information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioTaskRouterTaskQueueImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Workspaces/%s/TaskQueues/%s", rs.Primary.Attributes["workspace_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioTaskRouterTaskQueue_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}
`, friendlyName)
}

func testAccTwilioTaskRouterTaskQueue_maxReservedWorkers(friendlyName string, maxReservedWorkers int) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid        = twilio_taskrouter_workspace.workspace.sid
  friendly_name        = "%[1]s"
  max_reserved_workers = "%[2]d"
}
`, friendlyName, maxReservedWorkers)
}

func testAccTwilioTaskRouterTaskQueue_taskOrder(friendlyName string, taskOrder string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  task_order    = "%[2]s"
}
`, friendlyName, taskOrder)
}

func testAccTwilioTaskRouterTaskQueue_targetWorkers(friendlyName string, targetWorkers string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid  = twilio_taskrouter_workspace.workspace.sid
  friendly_name  = "%[1]s"
  target_workers = "%[2]s"
}
`, friendlyName, targetWorkers)
}

func testAccTwilioTaskRouterTaskQueue_assignmentActivitySid(friendlyName string) string {
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

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid           = twilio_taskrouter_workspace.workspace.sid
  friendly_name           = "%[1]s"
  assignment_activity_sid = twilio_taskrouter_activity.activity.sid
}
`, friendlyName)
}

// This is required because terraform will destroy the activity resource before updating the task queue referencing it which causes an error to occur
func testAccTwilioTaskRouterTaskQueue_detachAssignmentActivitySid(friendlyName string) string {
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

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}
`, friendlyName)
}

func testAccTwilioTaskRouterTaskQueue_invalidAssignmentActivitySid() string {
	return `
resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid           = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name           = "invalid_assignment_activity_sid"
  assignment_activity_sid = "assignment_activity_sid"
}
`
}

func testAccTwilioTaskRouterTaskQueue_reservationActivitySid(friendlyName string) string {
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

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid            = twilio_taskrouter_workspace.workspace.sid
  friendly_name            = "%[1]s"
  reservation_activity_sid = twilio_taskrouter_activity.activity.sid
}
`, friendlyName)
}

// This is required because terraform will destroy the activity resource before updating the task queue referencing it which causes an error to occur
func testAccTwilioTaskRouterTaskQueue_detachReservationActivitySid(friendlyName string) string {
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

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}
`, friendlyName)
}

func testAccTwilioTaskRouterTaskQueue_invalidReservationActivitySid() string {
	return `
resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid            = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name            = "invalid_reservation_activity_sid"
  reservation_activity_sid = "reservation_activity_sid"
}
`
}

func testAccTwilioTaskRouterTaskQueue_invalidWorkspaceSid() string {
	return `
resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = "workspace_sid"
  friendly_name = "invalid_workspace_sid"
}
`
}

func testAccTwilioTaskRouterTaskQueue_blankFriendlyName() string {
	return `
resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name = ""
}
`
}
