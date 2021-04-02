package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const dataSourceName = "twilio_studio_flow"

func TestAccDataSourceTwilioStudioFlow_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.flow", dataSourceName)
	friendlyName := acctest.RandString(10)
	status := "draft"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlow_complete(friendlyName, status),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "status"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "definition"),
					resource.TestCheckResourceAttr(stateDataSourceName, "commit_message", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "revision"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "webhook_url"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "valid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountQueue_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioStudioFlow_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^FW\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlow_complete(friendlyName string, status string) string {
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

data "twilio_studio_flow" "flow" {
  sid = twilio_studio_flow.flow.sid
}
`, friendlyName, status)
}

func testAccDataSourceTwilioStudioFlow_invalidSid() string {
	return `
data "twilio_studio_flow" "flow" {
  sid = "sid"
}
`
}
