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

var proxyServiceResourceName = "twilio_proxy_phone_number"

func TestAccTwilioStudio_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.phone_number", proxyServiceResourceName)

	testData := acceptance.TestAccData
	uniqueName := acctest.RandString(10)
	isReserved := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioProxyPhoneNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyPhoneNumber_basic(testData, uniqueName, isReserved),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_reserved", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "iso_country"),
					resource.TestCheckResourceAttrSet(stateResourceName, "in_use"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioProxyPhoneNumber_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.phone_number", proxyServiceResourceName)

	testData := acceptance.TestAccData
	uniqueName := acctest.RandString(10)
	isReserved := true
	newIsReserved := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioProxyPhoneNumberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioProxyPhoneNumber_basic(testData, uniqueName, isReserved),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_reserved", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "iso_country"),
					resource.TestCheckResourceAttrSet(stateResourceName, "in_use"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioProxyPhoneNumber_basic(acceptance.TestAccData, uniqueName, newIsReserved),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioProxyPhoneNumberExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sid", testData.PhoneNumberSid),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_reserved", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "phone_number"),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "iso_country"),
					resource.TestCheckResourceAttrSet(stateResourceName, "in_use"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioProxyPhoneNumberDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Proxy

	for _, rs := range s.RootModule().Resources {
		if rs.Type != proxyServiceResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.ID).Get(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving proxy phone number information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioProxyPhoneNumberExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Proxy

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.ID).Get(); err != nil {
			return fmt.Errorf("Error occurred when retrieving proxy phone number information %s", err)
		}

		return nil
	}
}

func testAccTwilioProxyPhoneNumber_basic(testData *acceptance.TestData, uniqueName string, isReserved bool) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
	unique_name = "%s"
}
  
resource "twilio_proxy_phone_number" "phone_number" {
	service_sid = twilio_proxy_service.service.sid
	sid         = "%s"
	is_reserved = %v
}`, uniqueName, testData.PhoneNumberSid, isReserved)
}
