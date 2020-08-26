package tests

import (
	"fmt"
	"regexp"
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
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioMessagingServiceDestroy,
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
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
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
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioMessagingServiceDestroy,
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
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
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
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
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

func TestAccTwilioMessagingService_fallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)
	fallbackMethod := "GET"
	fallbackURL := "https://test.com/fallback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_fallback(friendlyName, fallbackMethod, fallbackURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_method", fallbackMethod),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", fallbackURL),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_invalidFallbackMethod(t *testing.T) {
	friendlyName := acctest.RandString(10)
	fallbackMethod := "test"
	fallbackURL := "https://test.com/fallback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_fallback(friendlyName, fallbackMethod, fallbackURL),
				ExpectError: regexp.MustCompile("config is invalid: expected fallback_method to be one of \\[POST GET\\], got test"),
			},
		},
	})
}

func TestAccTwilioMessagingService_invalidFallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	fallbackMethod := "GET"
	fallbackURL := "fallback"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_fallback(friendlyName, fallbackMethod, fallbackURL),
				ExpectError: regexp.MustCompile("config is invalid: expected \"fallback_url\" to have a host, got fallback"),
			},
		},
	})
}

func TestAccTwilioMessagingService_inbound(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)
	inboundMethod := "GET"
	inboundURL := "https://test.com/inbound"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_inbound(friendlyName, inboundMethod, inboundURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_method", inboundMethod),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", inboundURL),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_invalidInboundMethod(t *testing.T) {
	friendlyName := acctest.RandString(10)
	inboundMethod := "test"
	inboundURL := "https://test.com/inbound"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_inbound(friendlyName, inboundMethod, inboundURL),
				ExpectError: regexp.MustCompile("config is invalid: expected inbound_method to be one of \\[POST GET\\], got test"),
			},
		},
	})
}

func TestAccTwilioMessagingService_invalidInboundURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	inboundMethod := "GET"
	inboundURL := "inbound"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_inbound(friendlyName, inboundMethod, inboundURL),
				ExpectError: regexp.MustCompile("config is invalid: expected \"inbound_request_url\" to have a host, got inbound"),
			},
		},
	})
}

func TestAccTwilioMessagingService_statusCallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)
	statusCallbackURL := "https://test.com/status"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_statusCallback(friendlyName, statusCallbackURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", statusCallbackURL),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_invalidStatusCallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	statusCallbackURL := "status"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_statusCallback(friendlyName, statusCallbackURL),
				ExpectError: regexp.MustCompile("config is invalid: expected \"status_callback_url\" to have a host, got status"),
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
}
`, friendlyName)
}

func testAccTwilioMessagingService_fallback(friendlyName string, method string, url string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name   = "%s"
  fallback_method = "%s"
  fallback_url    = "%s"
}
`, friendlyName, method, url)
}

func testAccTwilioMessagingService_inbound(friendlyName string, method string, url string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name       = "%s"
  inbound_method      = "%s"
  inbound_request_url = "%s"
}
`, friendlyName, method, url)
}

func testAccTwilioMessagingService_statusCallback(friendlyName string, url string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name       = "%s"
  status_callback_url = "%s"
}
`, friendlyName, url)
}
