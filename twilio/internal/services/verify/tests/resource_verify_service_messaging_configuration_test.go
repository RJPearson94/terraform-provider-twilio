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

var messagingConfigurationResourceName = "twilio_verify_messaging_configuration"

func TestAccTwilioVerifyMessagingConfiguration_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.messaging_configuration", messagingConfigurationResourceName)
	friendlyName := acctest.RandString(10)
	countryCode := "GB"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyMessagingConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyMessagingConfiguration_basic(friendlyName, countryCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyMessagingConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "country_code", countryCode),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioVerifyMessagingConfigurationImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioVerifyMessagingConfiguration_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyMessagingConfiguration_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioVerifyMessagingConfiguration_invalidMessagingServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyMessagingConfiguration_invalidMessagingServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of messaging_service_sid to match regular expression "\^MG\[0-9a-fA-F\]\{32\}\$", got messaging_service_sid`),
			},
		},
	})
}

func testAccCheckTwilioVerifyMessagingConfigurationDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

	for _, rs := range s.RootModule().Resources {
		if rs.Type != messagingConfigurationResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).MessagingConfiguration(rs.Primary.Attributes["sid"]).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving rate limit information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioVerifyMessagingConfigurationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).MessagingConfiguration(rs.Primary.Attributes["sid"]).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving rate limit information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioVerifyMessagingConfigurationImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/MessagingConfigurations/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["country_code"]), nil
	}
}

func testAccTwilioVerifyMessagingConfiguration_basic(friendlyName string, countryCode string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid           = twilio_verify_service.service.sid
  messaging_service_sid = twilio_messaging_service.service.sid
  country_code          = "%[2]s"
}
`, friendlyName, countryCode)
}

func testAccTwilioVerifyMessagingConfiguration_invalidServiceSid() string {
	return `
resource "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid           = "service_sid"
  messaging_service_sid = "MGaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  country_code          = "GB"
}
`
}

func testAccTwilioVerifyMessagingConfiguration_invalidMessagingServiceSid() string {
	return `
resource "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid           = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  messaging_service_sid = "messaging_service_sid"
  country_code          = "GB"
}
`
}
