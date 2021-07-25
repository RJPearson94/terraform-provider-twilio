package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var serviceDataSourceName = "twilio_messaging_service"

func TestAccDataSourceTwilioMessagingService_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service", serviceDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "area_code_geomatch", "true"),
					resource.TestCheckResourceAttr(stateDataSourceName, "fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "fallback_to_long_code", "true"),
					resource.TestCheckResourceAttr(stateDataSourceName, "fallback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "inbound_method", "POST"),
					resource.TestCheckResourceAttr(stateDataSourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "mms_converter", "true"),
					resource.TestCheckResourceAttr(stateDataSourceName, "smart_encoding", "true"),
					resource.TestCheckResourceAttr(stateDataSourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "sticky_sender", "true"),
					resource.TestCheckResourceAttr(stateDataSourceName, "use_inbound_webhook_on_number", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "validity_period", "14400"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioMessagingService_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioMessagingService_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^MG\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioMessagingService_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "%s"
}

data "twilio_messaging_service" "service" {
  sid = twilio_messaging_service.service.sid
}
`, friendlyName)
}

func testAccDataSourceTwilioMessagingService_invalidSid() string {
	return `
data "twilio_messaging_service" "service" {
  sid = "sid"
}
`
}
