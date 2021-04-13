package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var taskResourceName = "twilio_autopilot_task"

// func TestAccTwilioAutopilotTask_basic(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(10)

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_basic(uniqueName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
// 					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
// 				),
// 			},
// 			{
// 				ResourceName:      stateResourceName,
// 				ImportState:       true,
// 				ImportStateIdFunc: testAccTwilioAutopilotTaskImportStateIdFunc(stateResourceName),
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_update(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(1)
// 	newUniqueName := acctest.RandString(64)

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_basic(uniqueName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
// 					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioAutopilotTask_basic(newUniqueName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
// 					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "assistant_sid"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_invalidUniqueNameWith0Characters(t *testing.T) {
// 	uniqueName := ""

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccTwilioAutopilotTask_withStubbedAssistantSid(uniqueName),
// 				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got `),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_invalidUniqueNameWith65Characters(t *testing.T) {
// 	uniqueName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccTwilioAutopilotTask_withStubbedAssistantSid(uniqueName),
// 				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 64\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_friendlyName(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(10)
// 	friendlyName := ""
// 	newFriendlyName := acctest.RandString(255)

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_friendlyName(uniqueName, friendlyName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioAutopilotTask_friendlyName(uniqueName, newFriendlyName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioAutopilotTask_basic(uniqueName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_invalidFriendlyNameWith256Characters(t *testing.T) {
// 	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccTwilioAutopilotTask_friendlyNameWithStubbedAssistantSid(friendlyName),
// 				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(0 - 255\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_actions(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(10)

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_actions(uniqueName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
// 				),
// 			},
// 		},
// 	})
// }

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
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
				),
			},
			{
				Config: testAccTwilioAutopilotTask_updatedActions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"New World\"}}]}"),
				),
			},
		},
	})
}

// func TestAccTwilioAutopilotTask_actionsURL(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(10)
// 	url := "http://localhost/action"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_updateActionsURL(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(10)
// 	url := "http://localhost/action"
// 	newURL := "http://localhost/action2"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, newURL),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions_url", newURL),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action2\"}]}"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_changeActionsToUseURL(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(10)
// 	url := "http://localhost/action"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_actions(uniqueName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_changeActionsURLToActionsJSON(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.task", taskResourceName)
// 	uniqueName := acctest.RandString(10)
// 	url := "http://localhost/action"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions_url", url),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"redirect\":\"http://localhost/action\"}]}"),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioAutopilotTask_actions(uniqueName),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioAutopilotTaskExists(stateResourceName),
// 					resource.TestCheckResourceAttrSet(stateResourceName, "actions_url"),
// 					resource.TestCheckResourceAttr(stateResourceName, "actions", "{\"actions\":[{\"say\":{\"speech\":\"Hello World\"}}]}"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_invalidActionURL(t *testing.T) {
// 	uniqueName := acctest.RandString(10)
// 	url := "actionURL"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioAutopilotTaskDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccTwilioAutopilotTask_actionsURL(uniqueName, url),
// 				ExpectError: regexp.MustCompile(`(?s)expected "actions_url" to have a host, got actionURL`),
// 			},
// 		},
// 	})
// }

// func TestAccTwilioAutopilotTask_invalidAssistantSid(t *testing.T) {
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccTwilioAutopilotTask_invalidAssistantSid(),
// 				ExpectError: regexp.MustCompile(`(?s)expected value of assistant_sid to match regular expression "\^UA\[0-9a-fA-F\]\{32\}\$", got assistant_sid`),
// 			},
// 		},
// 	})
// }

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
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
}
`, uniqueName)
}

func testAccTwilioAutopilotTask_friendlyName(uniqueName string, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  friendly_name = "%[2]s"
}
`, uniqueName, friendlyName)
}

func testAccTwilioAutopilotTask_actions(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  actions = jsonencode({
    "actions" : [
      {
        "say" : {
          "speech" : "Hello World"
        }
      }
    ]
  })
}
`, uniqueName)
}

func testAccTwilioAutopilotTask_updatedActions(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  actions = jsonencode({
    "actions" : [
      {
        "say" : {
          "speech" : "New World"
        }
      }
    ]
  })
}
`, uniqueName)
}

func testAccTwilioAutopilotTask_actionsURL(uniqueName string, url string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_assistant" "assistant" {
  unique_name = "%[1]s"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "%[1]s"
  actions_url   = "%[2]s"
}
`, uniqueName, url)
}

func testAccTwilioAutopilotTask_invalidAssistantSid() string {
	return `
resource "twilio_autopilot_task" "task" {
  assistant_sid = "assistant_sid"
  unique_name   = "invalid_assistant_sid"
}
`
}

func testAccTwilioAutopilotTask_withStubbedAssistantSid(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_task" "task" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "%s"
}
`, uniqueName)
}

func testAccTwilioAutopilotTask_friendlyNameWithStubbedAssistantSid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_autopilot_task" "task" {
  assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "friendly_name_with_stubbed_assistant_sid"
  friendly_name = "%s"
}
`, friendlyName)
}
