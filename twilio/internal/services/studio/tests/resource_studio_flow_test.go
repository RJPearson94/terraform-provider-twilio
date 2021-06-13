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

const resourceName = "twilio_studio_flow"

func TestAccTwilioStudio_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.flow", resourceName)
	friendlyName := acctest.RandString(10)
	status := "draft"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioStudioFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioStudioFlow_basic(friendlyName, status),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioStudioFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "status", status),
					resource.TestCheckResourceAttr(stateResourceName, "validate", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "definition"),
					resource.TestCheckResourceAttr(stateResourceName, "commit_message", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "revision"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "webhook_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "valid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:            stateResourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccTwilioStudioFlowImportStateIdFunc(stateResourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"validate"},
			},
		},
	})
}

func TestAccTwilioStudioFlow_invalidStatus(t *testing.T) {
	friendlyName := acctest.RandString(10)
	status := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioStudioFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioStudioFlow_basic(friendlyName, status),
				ExpectError: regexp.MustCompile(`(?s)expected status to be one of \[draft published\], got test`),
			},
		},
	})
}

func TestAccTwilioStudioFlow_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.flow", resourceName)

	friendlyName := acctest.RandString(10)
	status := "draft"
	newStatus := "published"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioStudioFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioStudioFlow_basic(friendlyName, status),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioStudioFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "status", status),
					resource.TestCheckResourceAttr(stateResourceName, "validate", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "definition"),
					resource.TestCheckResourceAttr(stateResourceName, "commit_message", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "revision"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "webhook_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "valid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioStudioFlow_basic(friendlyName, newStatus),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioStudioFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "status", newStatus),
					resource.TestCheckResourceAttr(stateResourceName, "validate", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "definition"),
					resource.TestCheckResourceAttr(stateResourceName, "commit_message", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "revision"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "webhook_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "valid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioStudioFlow_blankFriendlyName(t *testing.T) {
	friendlyName := ""
	status := "draft"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioStudioFlow_basic(friendlyName, status),
				ExpectError: regexp.MustCompile(`(?s)expected \"friendly_name\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioStudioFlow_withWidgets(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.flow", resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioStudioFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioStudioFlow_withWidgets(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioStudioFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", "With widgets"),
					resource.TestCheckResourceAttr(stateResourceName, "status", "published"),
					resource.TestCheckResourceAttr(stateResourceName, "validate", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "definition", `{"description":"Flow with trigger widget","flags":{"allow_concurrent_calls":true},"initial_state":"Trigger","states":[{"name":"Trigger","properties":{"offset":{"x":200,"y":0}},"transitions":[{"event":"incomingCall"},{"event":"incomingMessage"},{"event":"incomingRequest"}],"type":"trigger"}]}`),
				),
			},
		},
	})
}

func TestAccTwilioStudioFlow_commitMessage(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.flow", resourceName)

	friendlyName := acctest.RandString(10)
	status := "published"
	commitMessage := "Test commit message"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioStudioFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioStudioFlow_withCommitMessage(friendlyName, status, commitMessage),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioStudioFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "commit_message", commitMessage),
				),
			},
			{
				Config: testAccTwilioStudioFlow_basic(friendlyName, status),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioStudioFlowExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "commit_message", ""),
				),
			},
		},
	})
}

func TestAccTwilioStudioFlow_withInvalidFlow(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioStudioFlow_invalidFlow(),
				ExpectError: regexp.MustCompile("(?s)Failed to validate studio flow: Flow definition validation failed, check `details` for more information. Details are .*"),
			},
		},
	})
}

func testAccCheckTwilioStudioFlowDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Studio

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourceName {
			continue
		}

		if _, err := client.Flow(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving flow information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioStudioFlowExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Studio

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Flow(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving flow information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioStudioFlowImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Flows/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioStudioFlow_basic(friendlyName string, status string) string {
	return fmt.Sprintf(`
resource "twilio_studio_flow" "flow" {
  friendly_name = "%s"
  status        = "%s"
  definition = jsonencode({
    "description" : "A New Flow",
    "flags" : {
      "allow_concurrent_calls" : true
    },
    "initial_state" : "Trigger",
    "states" : [
      {
        "name" : "Trigger",
        "properties" : {
          "offset" : {
            "x" : 0,
            "y" : 0
          }
        },
        "transitions" : [],
        "type" : "trigger"
      }
    ]
  })
}
`, friendlyName, status)
}

func testAccTwilioStudioFlow_withCommitMessage(friendlyName string, status string, commitMessage string) string {
	return fmt.Sprintf(`
resource "twilio_studio_flow" "flow" {
  friendly_name  = "%s"
  status         = "%s"
  commit_message = "%s"
  definition = jsonencode({
    "description" : "A New Flow",
    "flags" : {
      "allow_concurrent_calls" : true
    },
    "initial_state" : "Trigger",
    "states" : [
      {
        "name" : "Trigger",
        "properties" : {
          "offset" : {
            "x" : 0,
            "y" : 0
          }
        },
        "transitions" : [],
        "type" : "trigger"
      }
    ]
  })
}
`, friendlyName, status, commitMessage)
}

func testAccTwilioStudioFlow_withWidgets() string {
	return `
data "twilio_studio_flow_widget_trigger" "trigger" {
	name = "Trigger"
	
	offset {
		x = 200
		y = 0
	}
}
	
data "twilio_studio_flow_definition" "definition" {
	description   = "Flow with trigger widget"
	initial_state = data.twilio_studio_flow_widget_trigger.trigger.name
	
	flags {
		allow_concurrent_calls = true
	}
	
	states {
		json = data.twilio_studio_flow_widget_trigger.trigger.json
	}
}
	
resource "twilio_studio_flow" "flow" {
	friendly_name = "With widgets"
	status        = "published"
	definition    = data.twilio_studio_flow_definition.definition.json
	validate      = true
}
`
}

func testAccTwilioStudioFlow_invalidFlow() string {
	return `
resource "twilio_studio_flow" "flow" {
  friendly_name = "invalid flow"
  status        = "draft"
  validate      = true
  definition = jsonencode({
    "description" : "Invalid flow"
  })
}
`
}
