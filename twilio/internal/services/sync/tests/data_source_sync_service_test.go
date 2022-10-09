package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const serviceDataSourceName = "twilio_sync_service"

func TestAccDataSourceTwilioSyncService_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service", serviceDataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSyncService_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "acl_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "reachability_debouncing_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "reachability_debouncing_window", "5000"),
					resource.TestCheckResourceAttr(stateDataSourceName, "reachability_webhooks_enabled", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhook_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "webhooks_from_rest_enabled", "false"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioSyncService_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioSyncService_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioSyncService_basic() string {
	return `
resource "twilio_sync_service" "service" {}

data "twilio_sync_service" "service" {
  sid = twilio_sync_service.service.sid
}
`
}

func testAccDataSourceTwilioSyncService_invalidSid() string {
	return `
data "twilio_sync_service" "service" {
  sid = "service_sid"
}
`
}
