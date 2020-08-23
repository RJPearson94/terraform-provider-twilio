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

var serviceResourceName = "twilio_messaging_service"

func TestAccTwilioMessagingService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioMessagingServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "area_code_geomatch"),
					resource.TestCheckResourceAttrSet(stateResourceName, "fallback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "fallback_to_long_code"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "inbound_method"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "mms_converter"),
					resource.TestCheckResourceAttrSet(stateResourceName, "smart_encoding"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "sticky_sender"),
					resource.TestCheckResourceAttrSet(stateResourceName, "validity_period"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioMessagingServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "area_code_geomatch"),
					resource.TestCheckResourceAttrSet(stateResourceName, "fallback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "fallback_to_long_code"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "inbound_method"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "mms_converter"),
					resource.TestCheckResourceAttrSet(stateResourceName, "smart_encoding"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "sticky_sender"),
					resource.TestCheckResourceAttrSet(stateResourceName, "validity_period"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "area_code_geomatch"),
					resource.TestCheckResourceAttrSet(stateResourceName, "fallback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "fallback_to_long_code"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "inbound_method"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "mms_converter"),
					resource.TestCheckResourceAttrSet(stateResourceName, "smart_encoding"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "sticky_sender"),
					resource.TestCheckResourceAttrSet(stateResourceName, "validity_period"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioMessagingServiceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Messaging

	for _, rs := range s.RootModule().Resources {
		if rs.Type != serviceResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving service information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioMessagingServiceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Messaging

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving service information %s", err)
		}

		return nil
	}
}

func testAccTwilioMessagingService_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
	friendly_name = "%s"
}`, friendlyName)
}
