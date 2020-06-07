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

var workspaceActivityResourceName = "twilio_taskrouter_workspace_activity"

func TestAccTwilioTaskRouterWorkspaceActivity_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.activity", workspaceActivityResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioTaskRouterWorkspaceActivityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspaceActivity_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceActivityExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "available", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioTaskRouterWorkspaceActivity_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.activity", workspaceActivityResourceName)

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioTaskRouterWorkspaceActivityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspaceActivity_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceActivityExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "available", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspaceActivity_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceActivityExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "available", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioTaskRouterWorkspaceActivityDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

	for _, rs := range s.RootModule().Resources {
		if rs.Type != workspaceActivityResourceName {
			continue
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).Activity(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving workspace activity information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioTaskRouterWorkspaceActivityExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).Activity(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving workspace activity information %s", err)
		}

		return nil
	}
}

func testAccTwilioTaskRouterWorkspaceActivity_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
	friendly_name          = "%s"
	multi_task_enabled     = true
	prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_workspace_activity" "activity" {
	workspace_sid = twilio_taskrouter_workspace.workspace.sid
	friendly_name = "%s"
	available 	  = true
}`, friendlyName, friendlyName)
}
