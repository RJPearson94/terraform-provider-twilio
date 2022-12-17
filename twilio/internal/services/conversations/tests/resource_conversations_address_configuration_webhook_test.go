package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var conversationAddressConfigurationWebhookResourceName = "twilio_conversations_address_configuration_webhook"

func TestAccTwilioConversationsAddressConfigurationWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_webhook", conversationAddressConfigurationWebhookResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"
	webhookUrl := "https://localhost/webhook"
	webhookFilters := []string{"onMessageAdded", "onMessageUpdated", "onMessageRemoved"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationWebhook_basic(address, addressType, webhookUrl, webhookFilters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.#", "3"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.0", "onMessageAdded"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.1", "onMessageUpdated"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.2", "onMessageRemoved"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrl),
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
				ImportStateIdFunc: testAccTwilioConversationsAddressConfigurationWebhookImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_webhook", conversationAddressConfigurationWebhookResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"
	enabled := false
	webhookUrl := "https://localhost/webhook"
	webhookFilters := []string{"onMessageAdded"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationWebhook_enabled(address, addressType, enabled, webhookUrl, webhookFilters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.0", "onMessageAdded"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrl),
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
				Config: testAccTwilioConversationsAddressConfigurationWebhook_basic(address, addressType, webhookUrl, webhookFilters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "address", address),
					resource.TestCheckResourceAttr(stateResourceName, "type", addressType),
					resource.TestCheckResourceAttr(stateResourceName, "integration_type", "webhook"),
					resource.TestCheckResourceAttr(stateResourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.0", "onMessageAdded"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrl),
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

func TestAccTwilioConversationsAddressConfigurationWebhook_webhookMethod(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_webhook", conversationAddressConfigurationWebhookResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"
	webhookMethod := "GET"
	webhookUrl := "https://localhost/webhook"
	webhookFilters := []string{"onMessageAdded"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationWebhook_webhookMethod(address, addressType, webhookUrl, webhookFilters, webhookMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", webhookMethod),
				),
			},
			{
				Config: testAccTwilioConversationsAddressConfigurationWebhook_basic(address, addressType, webhookUrl, webhookFilters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationWebhook_webhookFilters(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.address_configuration_webhook", conversationAddressConfigurationWebhookResourceName)
	address := acceptance.TestAccData.PhoneNumber
	addressType := "sms"
	webhookUrl := "https://localhost/webhook"
	webhookFilters := []string{"onMessageAdded", "onMessageUpdated"}
	newWebhookFilters := []string{"onMessageAdded"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsAddressConfigurationWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsAddressConfigurationWebhook_basic(address, addressType, webhookUrl, webhookFilters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.0", "onMessageAdded"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.1", "onMessageUpdated"),
				),
			},
			{
				Config: testAccTwilioConversationsAddressConfigurationWebhook_basic(address, addressType, webhookUrl, newWebhookFilters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsAddressConfigurationWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_filters.0", "onMessageAdded"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationWebhook_invalidConversationServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationWebhook_invalidConversationServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^IS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationWebhook_invalidType(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationWebhook_invalidType(),
				ExpectError: regexp.MustCompile(`(?s)expected type to be one of \[sms whatsapp\], got type`),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookFilter(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookFilter(),
				ExpectError: regexp.MustCompile(`(?s)expected webhook_filters.0 to be one of \[onMessageAdded onMessageUpdated onMessageRemoved onConversationUpdated onConversationStateUpdated onConversationRemoved onParticipantAdded onParticipantUpdated onParticipantRemoved onDeliveryUpdated\], got unknown`),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookUrl(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookUrl(),
				ExpectError: regexp.MustCompile(`(?s)expected "webhook_url" to have a host, got webhook`),
			},
		},
	})
}

func TestAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookMethod(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookMethod(),
				ExpectError: regexp.MustCompile(`(?s)expected webhook_method to be one of \[GET POST\], got webhook_method`),
			},
		},
	})
}

func testAccCheckTwilioConversationsAddressConfigurationWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != conversationAddressConfigurationWebhookResourceName {
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

func testAccCheckTwilioConversationsAddressConfigurationWebhookExists(name string) resource.TestCheckFunc {
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

func testAccTwilioConversationsAddressConfigurationWebhookImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Configuration/Addresses/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsAddressConfigurationWebhook_basic(address string, addressType string, webhookUrl string, webhookFilters []string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "%[1]s"
  type            = "%[2]s"
  webhook_filters = %[3]s
  webhook_url     = "%[4]s"
}
`, address, addressType, `["`+strings.Join(webhookFilters, `","`)+`"]`, webhookUrl)
}

func testAccTwilioConversationsAddressConfigurationWebhook_enabled(address string, addressType string, enabled bool, webhookUrl string, webhookFilters []string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "%[1]s"
  type            = "%[2]s"
  enabled         = "%[3]v"
  webhook_filters = %[4]s
  webhook_url     = "%[5]s"
}
`, address, addressType, enabled, `["`+strings.Join(webhookFilters, `","`)+`"]`, webhookUrl)
}

func testAccTwilioConversationsAddressConfigurationWebhook_webhookMethod(address string, addressType string, webhookUrl string, webhookFilters []string, method string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "%[1]s"
  type            = "%[2]s"
  webhook_filters = %[3]s
  webhook_url     = "%[4]s"
  webhook_method  = "%[5]s"
}
`, address, addressType, `["`+strings.Join(webhookFilters, `","`)+`"]`, webhookUrl, method)
}

func testAccTwilioConversationsAddressConfigurationWebhook_invalidConversationServiceSid() string {
	return `
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address 								 = "+4471234567890"
  type    								 = "sms"
  service_sid = "service_sid"
	webhook_filters          = ["onMessageAdded"]
	webhook_url              = "https://localhost/webhook"
}
`
}

func testAccTwilioConversationsAddressConfigurationWebhook_invalidType() string {
	return `
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "+4471234567890"
  type 		        = "type"
	webhook_filters = ["onMessageAdded"]
	webhook_url     = "https://localhost/webhook"
}
`
}

func testAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookFilter() string {
	return `
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "+4471234567890"
  type 		        = "sms"
	webhook_filters = ["unknown"]
	webhook_url     = "https://localhost/webhook"
}
`
}

func testAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookUrl() string {
	return `
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "+4471234567890"
  type 		        = "sms"
	webhook_filters = ["onMessageAdded"]
	webhook_url     = "webhook"
}
`
}

func testAccTwilioConversationsAddressConfigurationWebhook_invalidWebhookMethod() string {
	return `
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "+4471234567890"
  type 		        = "sms"
	webhook_filters = ["onMessageAdded"]
	webhook_url     = "https://localhost/webhook"
	webhook_method  = "webhook_method"
}
`
}
