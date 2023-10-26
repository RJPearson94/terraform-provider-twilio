package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var proxyServiceResourceName = "twilio_proxy_service"

func TestAccTwilioProxyService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "chat_instance_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "default_ttl", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "geo_match_level", "country"),
					resource.TestCheckResourceAttr(stateResourceName, "number_selection_behavior", "prefer-sticky"),
					resource.TestCheckResourceAttr(stateResourceName, "intercept_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "out_of_session_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioProxyServiceImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioProxyService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)

	uniqueName := acctest.RandString(10)
	defaultTTL := 10

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyService_defaultTTL(uniqueName, defaultTTL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "chat_instance_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "default_ttl", "10"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "geo_match_level", "country"),
					resource.TestCheckResourceAttr(stateResourceName, "number_selection_behavior", "prefer-sticky"),
					resource.TestCheckResourceAttr(stateResourceName, "intercept_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "out_of_session_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "chat_instance_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "default_ttl", "0"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "geo_match_level", "country"),
					resource.TestCheckResourceAttr(stateResourceName, "number_selection_behavior", "prefer-sticky"),
					resource.TestCheckResourceAttr(stateResourceName, "intercept_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "out_of_session_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioProxyService_callbacks(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)

	uniqueName := acctest.RandString(10)
	url := "https://test.com/callbackURL"
	interceptURL := "https://test.com/interceptURL"
	outOfSessionURL := "https://test.com/outOfSessionURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyService_callbacks(uniqueName, url, interceptURL, outOfSessionURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "intercept_callback_url", interceptURL),
					resource.TestCheckResourceAttr(stateResourceName, "out_of_session_callback_url", outOfSessionURL),
				),
			},
			{
				Config: testAccTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "intercept_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "out_of_session_callback_url", ""),
				),
			},
		},
	})
}

func TestAccTwilioProxyService_invalidCallbackURL(t *testing.T) {
	uniqueName := acctest.RandString(10)
	url := "callbackURL"
	interceptURL := "https://test.com/interceptURL"
	outOfSessionURL := "https://test.com/outOfSessionURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_callbacks(uniqueName, url, interceptURL, outOfSessionURL),
				ExpectError: regexp.MustCompile(`(?s)expected "callback_url" to have a host, got callbackURL`)},
		},
	})
}

func TestAccTwilioProxyService_invalidInterceptCallbackURL(t *testing.T) {
	uniqueName := acctest.RandString(10)
	url := "https://test.com/callbackURL"
	interceptURL := "interceptURL"
	outOfSessionURL := "https://test.com/outOfSessionURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_callbacks(uniqueName, url, interceptURL, outOfSessionURL),
				ExpectError: regexp.MustCompile(`(?s)expected "intercept_callback_url" to have a host, got interceptURL`),
			},
		},
	})
}

func TestAccTwilioProxyService_invalidOutOfSessionCallbackURL(t *testing.T) {
	uniqueName := acctest.RandString(10)
	url := "https://test.com/callbackURL"
	interceptURL := "https://test.com/interceptURL"
	outOfSessionURL := "outOfSessionURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_callbacks(uniqueName, url, interceptURL, outOfSessionURL),
				ExpectError: regexp.MustCompile(`(?s)expected "out_of_session_callback_url" to have a host, got outOfSessionURL`),
			},
		},
	})
}

func TestAccTwilioProxyService_geoMatchLevel(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)

	uniqueName := acctest.RandString(10)
	geoMatchLevel := "area-code"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyService_geoMatchLevel(uniqueName, geoMatchLevel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "geo_match_level", geoMatchLevel),
				),
			},
			{
				Config: testAccTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "geo_match_level", "country"),
				),
			},
		},
	})
}

func TestAccTwilioProxyService_invalidGeoMatchLevel(t *testing.T) {
	uniqueName := acctest.RandString(10)
	geoMatchLevel := "geo_match_level"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_geoMatchLevel(uniqueName, geoMatchLevel),
				ExpectError: regexp.MustCompile(`(?s)expected geo_match_level to be one of \["area-code" "country" "extended-area-code"\], got geo_match_level`),
			},
		},
	})
}

func TestAccTwilioProxyService_numberSelectionBehaviour(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)

	uniqueName := acctest.RandString(10)
	numberSelectionBehavior := "avoid-sticky"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyService_numberSelectionBehaviour(uniqueName, numberSelectionBehavior),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "number_selection_behavior", numberSelectionBehavior),
				),
			},
			{
				Config: testAccTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "number_selection_behavior", "prefer-sticky"),
				),
			},
		},
	})
}

func TestAccTwilioProxyService_invalidNumberSelectionBehaviour(t *testing.T) {
	uniqueName := acctest.RandString(10)
	numberSelectionBehavior := "number_selection_behavior"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_numberSelectionBehaviour(uniqueName, numberSelectionBehavior),
				ExpectError: regexp.MustCompile(`(?s)expected number_selection_behavior to be one of \["avoid-sticky" "prefer-sticky"\], got number_selection_behavior`),
			},
		},
	})
}

func TestAccTwilioProxyService_uniqueName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)

	uniqueName := acctest.RandString(1)
	newUniqueName := acctest.RandString(191)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
				),
			},
			{
				Config: testAccTwilioProxyService_basic(newUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
				),
			},
		},
	})
}

func TestAccTwilioProxyService_invalidUniqueNameWith0Characters(t *testing.T) {
	uniqueName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_basic(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 191\), got `),
			},
		},
	})
}

func TestAccTwilioProxyService_invalidUniqueNameWith192Characters(t *testing.T) {
	uniqueName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_basic(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 191\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioProxyService_chatInstanceSid(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)

	uniqueName := acctest.RandString(10)
	chatInstanceSid := "ISaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioProxyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyService_chatInstanceSid(uniqueName, chatInstanceSid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "chat_instance_sid", chatInstanceSid),
				),
			},
			{
				Config: testAccTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "chat_instance_sid", ""),
				),
			},
		},
	})
}

func TestAccTwilioProxyService_invalidChatInstanceSid(t *testing.T) {
	uniqueName := "invalid_chat_instance_sid"
	chatInstanceSid := "chat_instance_sid"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioProxyService_chatInstanceSid(uniqueName, chatInstanceSid),
				ExpectError: regexp.MustCompile(`(?s)expected value of chat_instance_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got chat_instance_sid`),
			},
		},
	})
}

func testAccCheckTwilioProxyServiceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Proxy

	for _, rs := range s.RootModule().Resources {
		if rs.Type != proxyServiceResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.ID).Fetch(); err != nil {
			if _, ok := err.(*sdkUtils.TwilioError); ok {
				// currently proxy returns a 400 error not 404
				if err.(*sdkUtils.TwilioError).Error() == "Invalid Service Sid" {
					return nil
				}
			}
			return fmt.Errorf("Error occurred when retrieving service information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioProxyServiceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Proxy

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

func testAccTwilioProxyServiceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioProxyService_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name = "%s"
}
`, uniqueName)
}

func testAccTwilioProxyService_callbacks(uniqueName string, url string, interceptURL string, outOfSessionURL string) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name                 = "%s"
  callback_url                = "%s"
  intercept_callback_url      = "%s"
  out_of_session_callback_url = "%s"
}
`, uniqueName, url, interceptURL, outOfSessionURL)
}

func testAccTwilioProxyService_geoMatchLevel(uniqueName string, geoMatchLevel string) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name     = "%s"
  geo_match_level = "%s"
}
`, uniqueName, geoMatchLevel)
}

func testAccTwilioProxyService_numberSelectionBehaviour(uniqueName string, numberSelectionBehaviour string) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name               = "%s"
  number_selection_behavior = "%s"
}
`, uniqueName, numberSelectionBehaviour)
}

func testAccTwilioProxyService_defaultTTL(uniqueName string, defaultTTL int) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name = "%s"
  default_ttl = %d
}
`, uniqueName, defaultTTL)
}

func testAccTwilioProxyService_chatInstanceSid(uniqueName string, chatInstanceSid string) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name       = "%s"
  chat_instance_sid = "%s"
}
`, uniqueName, chatInstanceSid)
}
