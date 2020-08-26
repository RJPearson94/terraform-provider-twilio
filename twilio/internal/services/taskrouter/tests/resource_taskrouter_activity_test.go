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

var ActivityResourceName = "twilio_taskrouter_activity"

func TestAccTwilioTaskRouterActivity_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.activity", ActivityResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		Providers:         acceptance.TestAccProviders,
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioTaskRouterActivityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterActivity_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterActivityExists(stateResourceName),
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
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioTaskRouterActivityImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioTaskRouterActivity_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.activity", ActivityResourceName)

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioTaskRouterActivityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterActivity_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterActivityExists(stateResourceName),
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
				Config: testAccTwilioTaskRouterActivity_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterActivityExists(stateResourceName),
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

func testAccCheckTwilioTaskRouterActivityDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

	for _, rs := range s.RootModule().Resources {
		if rs.Type != ActivityResourceName {
			continue
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).Activity(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving activity information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioTaskRouterActivityExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).TaskRouter

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Workspace(rs.Primary.Attributes["workspace_sid"]).Activity(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving activity information %s", err)
		}

		return nil
	}
}

func testAccTwilioTaskRouterActivityImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Workspaces/%s/Activities/%s", rs.Primary.Attributes["workspace_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioTaskRouterActivity_basic(friendlyName string) string {
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
`, friendlyName, friendlyName)
}
