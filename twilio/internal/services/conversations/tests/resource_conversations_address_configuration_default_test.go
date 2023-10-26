package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var conversationAddressConfigurationDefaultResourceName = "twilio_conversations_address_configuration_default"

func TestAccTwilioConversationsAddressConfigurationDefault_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_default", conversationAddressConfigurationDefaultResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationDefaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationDefault_basic(address, addressType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationDefaultExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "default"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "service_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioConversationsAddressConfigurationDefaultImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationDefault_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_default", conversationAddressConfigurationDefaultResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"
	enabled := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationDefaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationDefault_enabled(address, addressType, enabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationDefaultExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "default"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "service_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsAddressConfigurationDefault_basic(address, addressType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationDefaultExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "default"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "service_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationDefault_invalidConversationServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationDefault_invalidConversationServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationDefault_invalidType(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationDefault_invalidType(),
				ExpectError: regexp.MustCompile(`(?s)expected type to be one of \["sms" "whatsapp"\], got type`),
			},
		},
	})
}

func testAccCheckTwilioConversationsAddressConfigurationDefaultDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != conversationAddressConfigurationDefaultResourceName {
			continue
		}

		if _, err := client.Configuration().Address(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving address configuration information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioConversationsAddressConfigurationDefaultExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Configuration().Address(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving address configuration information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsAddressConfigurationDefaultImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Configuration/Addresses/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsAddressConfigurationDefault_basic(address string, addressType string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_address_configuration_default" "address_configuration_default" {
  address = "%[1]s"
  type    = "%[2]s"
}
`, address, addressType)
}

func testAccTwilioConversationsAddressConfigurationDefault_enabled(address string, addressType string, enabled bool) string {
	return fmt.Sprintf(`
resource "twilio_conversations_address_configuration_default" "address_configuration_default" {
  address = "%[1]s"
  type    = "%[2]s"
  enabled = "%[3]v"
}
`, address, addressType, enabled)
}

func testAccTwilioConversationsAddressConfigurationDefault_invalidConversationServiceSid() string {
	return `
resource "twilio_conversations_address_configuration_default" "address_configuration_default" {
  address 								 = "+4471234567890"
  type    								 = "sms"
  service_sid = "service_sid"
}
`
}

func testAccTwilioConversationsAddressConfigurationDefault_invalidType() string {
	return `
resource "twilio_conversations_address_configuration_default" "address_configuration_default" {
  address = "+4471234567890"
  type 		= "type"
}
`
}
