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

var serviceResourceName = "twilio_messaging_service"

func TestAccTwilioMessagingService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioMessagingServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "area_code_geomatch", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_to_long_code", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "mms_converter", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "smart_encoding", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "sticky_sender", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "validity_period", "14400"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioMessagingServiceImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioMessagingService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioMessagingServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "area_code_geomatch", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_to_long_code", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "mms_converter", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "smart_encoding", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "sticky_sender", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "validity_period", "14400"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
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
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "area_code_geomatch", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_to_long_code", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "mms_converter", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "smart_encoding", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "sticky_sender", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "validity_period", "14400"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_blankFriendlyName(t *testing.T) {
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_basic(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected \"friendly_name\" to not be an empty string, got `),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_fallback(friendlyName, fallbackMethod, fallbackURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_method", fallbackMethod),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", fallbackURL),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_url", ""),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_fallback(friendlyName, fallbackMethod, fallbackURL),
				ExpectError: regexp.MustCompile(`(?s)expected fallback_method to be one of \[POST GET\], got test`),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_fallback(friendlyName, fallbackMethod, fallbackURL),
				ExpectError: regexp.MustCompile(`(?s)expected "fallback_url" to have a host, got fallback`),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_inbound(friendlyName, inboundMethod, inboundURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_method", inboundMethod),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", inboundURL),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "inbound_request_url", ""),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_inbound(friendlyName, inboundMethod, inboundURL),
				ExpectError: regexp.MustCompile(`(?s)expected inbound_method to be one of \[POST GET\], got test`),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_inbound(friendlyName, inboundMethod, inboundURL),
				ExpectError: regexp.MustCompile(`(?s)expected "inbound_request_url" to have a host, got inbound`),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_statusCallback(friendlyName, statusCallbackURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", statusCallbackURL),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
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
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_statusCallback(friendlyName, statusCallbackURL),
				ExpectError: regexp.MustCompile(`(?s)expected "status_callback_url" to have a host, got status`),
			},
		},
	})
}

func TestAccTwilioMessagingService_validityPeriod(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)
	validityPeriod := 14400
	newValidityPeriod := 1

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_validityPeriod(friendlyName, validityPeriod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "validity_period", "14400"),
				),
			},
			{
				Config: testAccTwilioMessagingService_validityPeriod(friendlyName, newValidityPeriod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "validity_period", "1"),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "validity_period", "14400"),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_invalidValidityPeriodOf0(t *testing.T) {
	friendlyName := acctest.RandString(10)
	validityPeriod := 0

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_validityPeriod(friendlyName, validityPeriod),
				ExpectError: regexp.MustCompile(`(?s)expected validity_period to be in the range \(1 - 14400\), got 0`),
			},
		},
	})
}

func TestAccTwilioMessagingService_invalidValidityPeriodOf14401(t *testing.T) {
	friendlyName := acctest.RandString(10)
	validityPeriod := 14401

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioMessagingService_validityPeriod(friendlyName, validityPeriod),
				ExpectError: regexp.MustCompile(`(?s)expected validity_period to be in the range \(1 - 14400\), got 14401`),
			},
		},
	})
}

func TestAccTwilioMessagingService_stickySender(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_stickySenderFalse(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sticky_sender", "false"),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sticky_sender", "true"),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_smartEncoding(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_smartEncodingFalse(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "smart_encoding", "false"),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "smart_encoding", "true"),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_mmsConverter(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_mmsConverterFalse(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "mms_converter", "false"),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "mms_converter", "true"),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_fallbackToLongCode(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_fallbackToLongCodeFalse(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_to_long_code", "false"),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "fallback_to_long_code", "true"),
				),
			},
		},
	})
}

func TestAccTwilioMessagingService_areaCodeGeomatch(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioMessagingService_areaCodeGeomatchFalse(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "area_code_geomatch", "false"),
				),
			},
			{
				Config: testAccTwilioMessagingService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioMessagingServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "area_code_geomatch", "true"),
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
			return fmt.Errorf("Error occurred when retrieving service information %s", err.Error())
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
			return fmt.Errorf("Error occurred when retrieving service information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioMessagingServiceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s", rs.Primary.Attributes["sid"]), nil
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

func testAccTwilioMessagingService_validityPeriod(friendlyName string, validityPeriod int) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "%s"
  validity_period = %d
}
`, friendlyName, validityPeriod)
}

func testAccTwilioMessagingService_stickySenderFalse(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "%s"
  sticky_sender = false
}
`, friendlyName)
}

func testAccTwilioMessagingService_smartEncodingFalse(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "%s"
  smart_encoding = false
}
`, friendlyName)
}

func testAccTwilioMessagingService_mmsConverterFalse(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "%s"
  mms_converter = false
}
`, friendlyName)
}

func testAccTwilioMessagingService_fallbackToLongCodeFalse(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "%s"
  fallback_to_long_code = false
}
`, friendlyName)
}

func testAccTwilioMessagingService_areaCodeGeomatchFalse(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_messaging_service" "service" {
  friendly_name = "%s"
  area_code_geomatch = false
}
`, friendlyName)
}
