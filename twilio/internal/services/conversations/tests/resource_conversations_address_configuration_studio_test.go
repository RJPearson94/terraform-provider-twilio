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

var conversationAddressConfigurationStudioResourceName = "twilio_conversations_address_configuration_studio"

func TestAccTwilioConversationsAddressConfigurationStudio_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_studio", conversationAddressConfigurationStudioResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"
	retryCount := 2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationStudioDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationStudio_basic(address, addressType, retryCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationStudioExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "flow_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "retry_count", "2"),
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
				ImportStateIdFunc: testAccTwilioConversationsAddressConfigurationStudioImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationStudio_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_studio", conversationAddressConfigurationStudioResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"
	enabled := false
	retryCount := 2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationStudioDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationStudio_enabled(address, addressType, enabled, retryCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationStudioExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "flow_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "retry_count", "2"),
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
				Config: testAccTwilioConversationsAddressConfigurationStudio_basic(address, addressType, retryCount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationStudioExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "studio"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "flow_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "retry_count", "2"),
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

func TestAccTwilioConversationsAddressConfigurationStudio_invalidStudioFlowSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationStudio_invalidStudioFlowSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of flow_sid to match regular expression "\^FW\[0-9a-fA-F\]\{32\}\$", got flow_sid`),
			},
		},
	})
}
func TestAccTwilioConversationsAddressConfigurationStudio_invalidConversationServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationStudio_invalidConversationServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationStudio_invalidType(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationStudio_invalidType(),
				ExpectError: regexp.MustCompile(`(?s)expected type to be one of \[sms whatsapp\], got type`),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationStudio_invalidStudioRetryCount(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationStudio_invalidStudioRetryCount(),
				ExpectError: regexp.MustCompile(`(?s)expected retry_count to be in the range \(0 - 3\), got 4`),
			},
		},
	})
}

func testAccCheckTwilioConversationsAddressConfigurationStudioDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != conversationAddressConfigurationStudioResourceName {
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

func testAccCheckTwilioConversationsAddressConfigurationStudioExists(name string) resource.TestCheckFunc {
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

func testAccTwilioConversationsAddressConfigurationStudioImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Configuration/Addresses/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsAddressConfigurationStudio_basic(address string, addressType string, retryCount int) string {
	return fmt.Sprintf(`
resource "twilio_studio_flow" "flow" {
  friendly_name = "Test Flow"
  status        = "published"
  definition = jsonencode({
    "description" : "A New Flow",
    "flags" : {
      "allow_concurrent_calls" : true
    },
    "initial_state" : "Trigger",
    "states" : [
      {
        "name" : "Trigger",
        "properties" : {
          "offset" : {
            "x" : 0,
            "y" : 0
          }
        },
        "transitions" : [],
        "type" : "trigger"
      }
    ]
  })
}

resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address     = "%[1]s"
  type        = "%[2]s"
  flow_sid    = twilio_studio_flow.flow.sid
  retry_count = %[3]d
}
`, address, addressType, retryCount)
}

func testAccTwilioConversationsAddressConfigurationStudio_enabled(address string, addressType string, enabled bool, retryCount int) string {
	return fmt.Sprintf(`
resource "twilio_studio_flow" "flow" {
  friendly_name = "Test Flow"
  status        = "published"
  definition = jsonencode({
    "description" : "A New Flow",
    "flags" : {
      "allow_concurrent_calls" : true
    },
    "initial_state" : "Trigger",
    "states" : [
      {
        "name" : "Trigger",
        "properties" : {
          "offset" : {
            "x" : 0,
            "y" : 0
          }
        },
        "transitions" : [],
        "type" : "trigger"
      }
    ]
  })
}

resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address     = "%[1]s"
  type        = "%[2]s"
  enabled     = "%[3]v"
  flow_sid    = twilio_studio_flow.flow.sid
  retry_count = %[4]d
}
`, address, addressType, enabled, retryCount)
}

func testAccTwilioConversationsAddressConfigurationStudio_invalidStudioFlowSid() string {
	return `
resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address 					 = "+4471234567890"
  type    					 = "sms"
	flow_sid    = "flow_sid"
	retry_count = 1
}
`
}

func testAccTwilioConversationsAddressConfigurationStudio_invalidConversationServiceSid() string {
	return `
resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address 								 = "+4471234567890"
  type    								 = "sms"
  service_sid = "service_sid"
	flow_sid          = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	retry_count       = 1
}
`
}

func testAccTwilioConversationsAddressConfigurationStudio_invalidType() string {
	return `
resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address            = "+4471234567890"
  type 	             = "type"
	flow_sid    = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	retry_count = 1
}
`
}

func testAccTwilioConversationsAddressConfigurationStudio_invalidStudioRetryCount() string {
	return `
resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address            = "+4471234567890"
  type 	             = "sms"
	flow_sid    = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	retry_count = 4
}
`
}
