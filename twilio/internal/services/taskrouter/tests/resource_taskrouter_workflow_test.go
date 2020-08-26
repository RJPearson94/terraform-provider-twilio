package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var workflowResourceName = "twilio_taskrouter_workflow"

func TestAccTwilioTaskRouterWorkflow_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workflow", workflowResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
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
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttr(stateResourceName, "document_content_type", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workflow", workflowResourceName)

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
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
		ProviderFactories: acceptance.TestAccProviderFactories(),
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
		},
	})
}

func TestAccTwilioTaskRouterWorkflow_invalidAssignmentCallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	callbackURL := "callback"
	fallbackCallbackURL := "https://test.com/fallback-callback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_assignmentCallback(friendlyName, callbackURL, fallbackCallbackURL),
				ExpectError: regexp.MustCompile("config is invalid: expected \"assignment_callback_url\" to have a host, got callback"),
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
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkflow_assignmentCallback(friendlyName, callbackURL, fallbackCallbackURL),
				ExpectError: regexp.MustCompile("config is invalid: expected \"fallback_assignment_callback_url\" to have a host, got fallback-callback"),
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
			return fmt.Errorf("Error occurred when retrieving workflow information %s", err)
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
			return fmt.Errorf("Error occurred when retrieving workflow information %s", err)
		}

		return nil
	}
}

func testAccTwilioTaskRouterWorkflow_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%ss"
}

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%s"
  configuration = <<EOF
{
  "task_routing": {
    "filters": [],
    "default_filter": {
      "queue": "${twilio_taskrouter_task_queue.task_queue.sid}"
    }
  }
}
EOF
}
`, friendlyName, friendlyName, friendlyName)
}

func testAccTwilioTaskRouterWorkflow_assignmentCallback(friendlyName string, url string, fallbackURL string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%s"
}

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%s"
  configuration = <<EOF
{
  "task_routing": {
    "filters": [],
    "default_filter": {
      "queue": "${twilio_taskrouter_task_queue.task_queue.sid}"
    }
  }
}
EOF
  assignment_callback_url = "%s"
  fallback_assignment_callback_url = "%s"
}
`, friendlyName, friendlyName, friendlyName, url, fallbackURL)
}
