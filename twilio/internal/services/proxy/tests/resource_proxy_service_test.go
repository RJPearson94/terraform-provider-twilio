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
					resource.TestCheckResourceAttr(stateResourceName, "chat_service_sid", ""),
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
	newUniqueName := acctest.RandString(10)

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
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "chat_instance_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "chat_service_sid", ""),
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
				Config: testAccTwilioProxyService_basic(newUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "chat_instance_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "chat_service_sid", ""),
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
