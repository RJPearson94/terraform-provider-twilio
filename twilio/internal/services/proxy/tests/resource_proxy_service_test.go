package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var proxyServiceResourceName = "twilio_proxy_service"

func TestAccTwilioStudio_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioProxyServiceDestroy,
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
		},
	})
}

func TestAccTwilioProxyService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", proxyServiceResourceName)

	uniqueName := acctest.RandString(10)
	newUniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioProxyServiceDestroy,
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

func testAccCheckTwilioProxyServiceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Proxy

	for _, rs := range s.RootModule().Resources {
		if rs.Type != proxyServiceResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.ID).Get(); err != nil {
			if _, ok := err.(*sdkUtils.TwilioError); ok {
				// currently proxy returns a 400 error not 404
				if err.(*sdkUtils.TwilioError).Error() == "Invalid Service Sid" {
					return nil
				}
			}
			return fmt.Errorf("Error occurred when retrieving service information %s", err)
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

		if _, err := client.Service(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving service information %s", err)
		}

		return nil
	}
}

func testAccTwilioProxyService_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
	unique_name = "%s"
}`, uniqueName)
}
