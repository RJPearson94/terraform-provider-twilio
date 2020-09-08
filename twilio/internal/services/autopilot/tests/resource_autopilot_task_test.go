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

var taskResourceName = "twilio_autopilot_task"

func TestAccTwilioAutopilotTask_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAutopilotTaskImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAutopilotTask_actions(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_actions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTask_updateActions(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_actions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotTask_updatedActions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"New World\"}}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTask_actionsURL(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/action"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTask_updateActionsURL(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/action"
	newURL := "http://localhost/action2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, newURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions_url", newURL),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action2\"}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTask_changeActionsToUseURL(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/action"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_actions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTask_changeActionsURLToActionsJSON(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)
	url := "http://localhost/action"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotTask_actions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioAutopilotTask_invalidActionURL(t *testing.T) {
	uniqueName := acctest.RandString(10)
	url := "actionURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
				ExpectError: regexp.MustCompile(`(?s)expected "actions_url" to have a host, got actionURL`),
			},
		},
	})
}

func TestAccTwilioAutopilotTask_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
	uniqueName := acctest.RandString(10)
	newUniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAutopilotTask_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioAutopilotTask_basic(newUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "actions"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioAutopilotTaskDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

	for _, rs := range s.RootModule().Resources {
		if rs.Type != taskResourceName {
			continue
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving task information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAutopilotTaskExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Autopilot

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Assistant(rs.Primary.Attributes["assistant_sid"]).Task(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving task information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAutopilotTaskImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Assistants/%s/Tasks/%s", rs.Primary.Attributes["assistant_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAutopilotTask_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
}
`, uniqueName, uniqueName)
}

func testAccTwilioAutopilotTask_actions(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
  actions       = <<EOF
{
	"actions": [
		{
			"say": {
				"speech": "Hello World"
			}
		}
	]
}
EOF
}
`, uniqueName, uniqueName)
}

func testAccTwilioAutopilotTask_updatedActions(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
  actions       = <<EOF
{
	"actions": [
		{
			"say": {
				"speech": "New World"
			}
		}
	]
}
EOF
}
`, uniqueName, uniqueName)
}

func testAccTwilioAutopilotTask_actionsURL(uniqueName string, url string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%s"
  actions_url   = "%s"
}
`, uniqueName, uniqueName, url)
}
