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

const workflowResourceName = "twilio_taskrouter_workflow"

func TestAccTwilioTaskRouterWorkflow_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workflow", workflowResourceName)
	workspaceStateResourceName := "twilio_taskrouter_workspace.workspace"

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkflowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkflow_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "task_reservation_timeout", "120"),
					resource.TestCheckResourceAttrSet(stateResourceName, "configuration"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "workspace_sid", workspaceStateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttr(stateResourceName, "document_content_type", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioTaskRouterWorkflowImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workflow", workflowResourceName)
	workspaceStateResourceName := "twilio_taskrouter_workspace.workspace"

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkflowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkflow_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "task_reservation_timeout", "120"),
					resource.TestCheckResourceAttrSet(stateResourceName, "configuration"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "workspace_sid", workspaceStateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttr(stateResourceName, "document_content_type", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkflow_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "task_reservation_timeout", "120"),
					resource.TestCheckResourceAttrSet(stateResourceName, "configuration"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "workspace_sid", workspaceStateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttr(stateResourceName, "document_content_type", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_assignmentCallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workflow", workflowResourceName)

	friendlyName := acctest.RandString(10)
	callbackURL := "https://test.com/callback"
	fallbackCallbackURL := "https://test.com/fallback-callback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkflow_assignmentCallback(friendlyName, callbackURL, fallbackCallbackURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_callback_url", callbackURL),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_assignment_callback_url", fallbackCallbackURL),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkflow_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_assignment_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "assignment_callback_url", ""),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_invalidAssignmentCallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	callbackURL := "callback"
	fallbackCallbackURL := "https://test.com/fallback-callback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_assignmentCallback(friendlyName, callbackURL, fallbackCallbackURL),
				ExpectError: regexp.MustCompile(`(?s)expected "assignment_callback_url" to have a host, got callback`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_invalidAssignmentFallbackCallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	callbackURL := "https://test.com/callback"
	fallbackCallbackURL := "fallback-callback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_assignmentCallback(friendlyName, callbackURL, fallbackCallbackURL),
				ExpectError: regexp.MustCompile(`(?s)expected "fallback_assignment_callback_url" to have a host, got fallback-callback`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_taskReservationTimeout(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workflow", workflowResourceName)

	friendlyName := acctest.RandString(10)
	taskReservationTimeout := 1
	newTaskReservationTimeout := 86400

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkflow_taskReservationTimeout(friendlyName, taskReservationTimeout),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "task_reservation_timeout", "1"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkflow_taskReservationTimeout(friendlyName, newTaskReservationTimeout),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "task_reservation_timeout", "86400"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkflow_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkflowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "task_reservation_timeout", "120"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_invalidTaskReservationTimeoutOf0(t *testing.T) {
	friendlyName := acctest.RandString(10)
	taskReservationTimeout := 0

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_taskReservationTimeout(friendlyName, taskReservationTimeout),
				ExpectError: regexp.MustCompile(`(?s)expected task_reservation_timeout to be in the range \(1 - 86400\), got 0`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_invalidTaskReservationTimeoutOf86401(t *testing.T) {
	friendlyName := acctest.RandString(10)
	taskReservationTimeout := 86401

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_taskReservationTimeout(friendlyName, taskReservationTimeout),
				ExpectError: regexp.MustCompile(`(?s)expected task_reservation_timeout to be in the range \(1 - 86400\), got 86401`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_invalidWorkspaceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_blankFriendlyName(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_blankFriendlyName(),
				ExpectError: regexp.MustCompile(`(?s)expected \"friendly_name\" to not be an empty string, got `),
			},
		},
	})
}

func testAccCheckTwilioTaskRouterWorkflowDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

	for _, rs := range s.RootModule().Resources {
		if rs.Type != workflowResourceName {
			continue
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).Workflow(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving workflow information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioTaskRouterWorkflowExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).Workflow(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving workflow information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioTaskRouterWorkflowImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Workspaces/%s/Workflows/%s", rs.Primary.Attributes["workspace_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioTaskRouterWorkflow_basic(friendlyName string) string {
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

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : "${twilio_taskrouter_task_queue.task_queue.sid}"
      }
    }
  })
}
`, friendlyName)
}

func testAccTwilioTaskRouterWorkflow_assignmentCallback(friendlyName string, url string, fallbackURL string) string {
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

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
  assignment_callback_url          = "%[2]s"
  fallback_assignment_callback_url = "%[3]s"
}
`, friendlyName, url, fallbackURL)
}

func testAccTwilioTaskRouterWorkflow_taskReservationTimeout(friendlyName string, taskReservationTimeout int) string {
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

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
  task_reservation_timeout = "%[2]d"
}
`, friendlyName, taskReservationTimeout)
}

func testAccTwilioTaskRouterWorkflow_invalidWorkspaceSid() string {
	return `
resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = "workspace_sid"
  friendly_name = "invalid_workspace_sid"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : "test_queue"
      }
    }
  })
}
`
}

func testAccTwilioTaskRouterWorkflow_blankFriendlyName() string {
	return `
resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name = ""
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : "test_queue"
      }
    }
  })
}
`
}
