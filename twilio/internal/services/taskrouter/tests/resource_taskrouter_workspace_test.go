package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var workspaceResourceName = "twilio_taskrouter_workspace"

func TestAccTwilioTaskRouterWorkspace_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace", workspaceResourceName)
	friendlyName := acctest.RandString(10)
	queueOrder := "FIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspace_basic(friendlyName, queueOrder),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckNoResourceAttr(stateResourceName, "event_filters"),
					resource.TestCheckResourceAttr(stateResourceName, "multi_task_enabled", "true"),
					resource.TestCheckNoResourceAttr(stateResourceName, "template"),
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
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioTaskRouterWorkspaceImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspace_invalidOrderQueue(t *testing.T) {
	friendlyName := acctest.RandString(10)
	queueOrder := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkspace_basic(friendlyName, queueOrder),
				ExpectError: regexp.MustCompile(`(?s)expected prioritize_queue_order to be one of \[LIFO FIFO\], got test`)},
		},
	})
}

func TestAccTwilioTaskRouterWorkspace_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace", workspaceResourceName)

	friendlyName := acctest.RandString(10)
	queueOrder := "FIFO"
	newQueueOrder := "LIFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTaskRouterWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspace_basic(friendlyName, queueOrder),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", ""),
					resource.TestCheckNoResourceAttr(stateResourceName, "event_filters"),
					resource.TestCheckResourceAttr(stateResourceName, "multi_task_enabled", "true"),
					resource.TestCheckNoResourceAttr(stateResourceName, "template"),
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
					resource.TestCheckNoResourceAttr(stateResourceName, "event_filters"),
					resource.TestCheckResourceAttr(stateResourceName, "multi_task_enabled", "true"),
					resource.TestCheckNoResourceAttr(stateResourceName, "template"),
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

func TestAccTwilioTaskRouterWorkspace_eventCallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.workspace", workspaceResourceName)

	friendlyName := acctest.RandString(10)
	eventFilters := []string{"task.created", "task.canceled"}
	callbackURL := "https://test.com/callback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspace_eventCallback(friendlyName, eventFilters, callbackURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "event_callback_url", callbackURL),
					resource.TestCheckResourceAttr(stateResourceName, "event_filters.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "event_filters.0", "task.created"),
					resource.TestCheckResourceAttr(stateResourceName, "event_filters.1", "task.canceled"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspace_invalidEventCallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	eventFilters := []string{"task.created", "task.canceled"}
	callbackURL := "callback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTaskRouterWorkspace_eventCallback(friendlyName, eventFilters, callbackURL),
				ExpectError: regexp.MustCompile(`(?s)expected "event_callback_url" to have a host, got callback`),
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

		if _, err := client.Workspace(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving workspace information %s", err.Error())
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

		if _, err := client.Workspace(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving workspace information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioTaskRouterWorkspaceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Workspaces/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioTaskRouterWorkspace_basic(friendlyName string, queueOrder string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "%s"
}
`, friendlyName, queueOrder)
}

func testAccTwilioTaskRouterWorkspace_eventCallback(friendlyName string, eventFilters []string, callbackURL string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name      = "%s"
  event_filters      = %s
  event_callback_url = "%s"
}
`, friendlyName, `["`+strings.Join(eventFilters[:], `", "`)+`"]`, callbackURL)
}
