package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var dataSourceName = "twilio_studio_flow"

func TestAccDataSourceTwilioStudioFlow_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.flow", dataSourceName)
	friendlyName := acctest.RandString(10)
	status := "draft"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioStudioFlow_complete(friendlyName, status),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "sid", regexp.MustCompile(`^FW(.+)$`)),
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

func testAccTwilioStudioFlow_complete(friendlyName string, status string) string {
	return fmt.Sprintf(`
resource "twilio_studio_flow" "flow" {
  friendly_name = "%s"
  status        = "%s"
  definition    = <<EOF
{
	"description": "A New Flow",
	"flags": {
		"allow_concurrent_calls": true
	},
	"initial_state": "Trigger",
	"states": [
		{
		"name": "Trigger",
		"properties": {
			"offset": {
			"x": 0,
			"y": 0
			}
		},
		"transitions": [],
		"type": "trigger"
		}
	]
}
EOF
}

data "twilio_studio_flow" "flow" {
  sid = twilio_studio_flow.flow.sid
}
`, friendlyName, status)
}
