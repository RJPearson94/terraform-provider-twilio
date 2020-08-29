package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var serviceDataSourceName = "twilio_messaging_service"

func TestAccDataSourceTwilioMessagingService_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service", serviceDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "area_code_geomatch"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "fallback_method"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "fallback_to_long_code"),
					resource.TestCheckResourceAttr(stateDataSourceName, "fallback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "inbound_method"),
					resource.TestCheckResourceAttr(stateDataSourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "mms_converter"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "smart_encoding"),
					resource.TestCheckResourceAttr(stateDataSourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sticky_sender"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "validity_period"),
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
