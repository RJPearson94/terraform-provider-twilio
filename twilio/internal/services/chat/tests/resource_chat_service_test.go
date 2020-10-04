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

var serviceResourceName = "twilio_chat_service"

func TestAccTwilioChatService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_role_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "limits.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "media.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.added_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.removed_from_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.invited_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioChatServiceImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioChatChannelWebhook_invalidPostWebhookURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	webhookURL := "webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatService_postWebhook(friendlyName, webhookURL),
				ExpectError: regexp.MustCompile(`(?s)expected "post_webhook_url" to have a host, got webhook`),
			},
		},
	})
}

func TestAccTwilioChatChannelWebhook_invalidPreWebhookURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	webhookURL := "webhook"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioChatService_preWebhook(friendlyName, webhookURL),
				ExpectError: regexp.MustCompile(`(?s)expected "pre_webhook_url" to have a host, got webhook`),
			},
		},
	})
}

func TestAccTwilioChatService_notifications(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)
	logEnabled := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatService_notifications(friendlyName, logEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatServiceExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_role_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "limits.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "media.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.log_enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.added_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.removed_from_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.invited_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.0.enabled", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioChatService_basicUpdate(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_role_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "limits.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "media.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.added_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.removed_from_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.invited_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioChatService_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatServiceExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_role_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "limits.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "media.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.added_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.removed_from_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.invited_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioChatService_notificationsUpdate(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)
	logEnabled := true
	newLogEnabled := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioChatServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioChatService_notifications(friendlyName, logEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_role_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "limits.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "media.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.log_enabled", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.added_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.removed_from_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.invited_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.0.enabled", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioChatService_notifications(friendlyName, newLogEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioChatServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_channel_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_role_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "limits.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "media.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.log_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.added_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.removed_from_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.invited_to_channel.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "notifications.0.new_message.0.enabled", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioChatServiceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

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

func testAccCheckTwilioChatServiceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Chat

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

func testAccTwilioChatServiceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioChatService_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"
}
`, friendlyName)
}

func testAccTwilioChatService_notifications(friendlyName string, logEnabled bool) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"

  notifications {
    log_enabled = %t

    new_message {
      enabled = true
    }
  }
}
`, friendlyName, logEnabled)
}

func testAccTwilioChatService_postWebhook(friendlyName string, url string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"

  post_webhook_retry_count = 1
  post_webhook_url         = "%s"
}
`, friendlyName, url)
}

func testAccTwilioChatService_preWebhook(friendlyName string, url string) string {
	return fmt.Sprintf(`
resource "twilio_chat_service" "service" {
  friendly_name = "%s"

  pre_webhook_retry_count = 1
  pre_webhook_url         = "%s"
}
`, friendlyName, url)
}
