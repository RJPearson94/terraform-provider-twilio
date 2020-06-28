package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var dataSourceName = "twilio_studio_flow"

func TestAccDataSourceTwilioStudioFlow_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("%s.flow", dataSourceName)
	friendlyName := acctest.RandString(10)
	status := "draft"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioStudioFlowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioStudioFlow_complete(friendlyName, status),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "sid", regexp.MustCompile(`^FW(.+)$`)),
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
