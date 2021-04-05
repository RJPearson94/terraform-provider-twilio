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

var webhookResourceName = "twilio_conversations_webhook"

func TestAccTwilioConversationsWebhook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsWebhook_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsWebhookExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "target"),
					resource.TestCheckResourceAttrSet(stateResourceName, "method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "pre_webhook_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "post_webhook_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "filters.#"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)

	method := "GET"
	preWebhookUrl := "https://localhost.com/preWebhookUrl"
	postWebhookUrl := "https://localhost.com/postWebhookUrl"
	newMethod := "POST"
	newPreWebhookUrl := "https://localhost.com/newPreWebhookUrl"
	newPostWebhookUrl := "https://localhost.com/newPostWebhookUrl"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsWebhook_withMethodAndUrls(method, preWebhookUrl, postWebhookUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsWebhookExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "target"),
					resource.TestCheckResourceAttr(stateResourceName, "method", method),
					resource.TestCheckResourceAttr(stateResourceName, "pre_webhook_url", preWebhookUrl),
					resource.TestCheckResourceAttr(stateResourceName, "post_webhook_url", postWebhookUrl),
					resource.TestCheckResourceAttrSet(stateResourceName, "filters.#"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsWebhook_withMethodAndUrls(newMethod, newPreWebhookUrl, newPostWebhookUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsWebhookExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "target"),
					resource.TestCheckResourceAttr(stateResourceName, "method", newMethod),
					resource.TestCheckResourceAttr(stateResourceName, "pre_webhook_url", newPreWebhookUrl),
					resource.TestCheckResourceAttr(stateResourceName, "post_webhook_url", newPostWebhookUrl),
					resource.TestCheckResourceAttrSet(stateResourceName, "filters.#"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_invalidMethod(t *testing.T) {
	method := "DELETE"
	preWebhookUrl := "https://localhost.com/preWebhookUrl"
	postWebhookUrl := "https://localhost.com/postWebhookUrl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsWebhook_withMethodAndUrls(method, preWebhookUrl, postWebhookUrl),
				ExpectError: regexp.MustCompile(`(?s)expected method to be one of \[GET POST\], got DELETE`),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_invalidPreWebhookUrl(t *testing.T) {
	method := "GET"
	preWebhookUrl := "preWebhookUrl"
	postWebhookUrl := "https://localhost.com/postWebhookUrl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsWebhook_withMethodAndUrls(method, preWebhookUrl, postWebhookUrl),
				ExpectError: regexp.MustCompile(`(?s)expected "pre_webhook_url" to have a host, got preWebhookUrl`),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_invalidPostWebhookUrl(t *testing.T) {
	method := "GET"
	preWebhookUrl := "https://localhost.com/preWebhookUrl"
	postWebhookUrl := "postWebhookUrl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsWebhook_withMethodAndUrls(method, preWebhookUrl, postWebhookUrl),
				ExpectError: regexp.MustCompile(`(?s)expected "post_webhook_url" to have a host, got postWebhookUrl`),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_invalidTarget(t *testing.T) {
	target := "studio"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsWebhook_withTarget(target),
				ExpectError: regexp.MustCompile(`(?s)expected target to be one of \[webhook flex\], got studio`),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_filters(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.webhook", webhookResourceName)

	filters := []string{"onConversationAdd"}
	newFilters := []string{"onConversationAdd", "onConversationRemove"}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsWebhook_filters(filters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", filters[0]),
				),
			},
			{
				Config: testAccTwilioConversationsWebhook_filters(newFilters),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsWebhookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "filters.#", "2"),
					resource.TestCheckResourceAttr(stateResourceName, "filters.0", newFilters[0]),
					resource.TestCheckResourceAttr(stateResourceName, "filters.1", newFilters[1]),
				),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_invalidFilter(t *testing.T) {
	filters := []string{"test"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsWebhook_filters(filters),
				ExpectError: regexp.MustCompile(`(?s)expected filters.0 to be one of \[onConversationAdd onConversationAdded onConversationRemove onConversationRemoved onConversationStateUpdated onConversationUpdate onConversationUpdated onDeliveryUpdated onMessageAdd onMessageAdded onMessageRemove onMessageRemoved onMessageUpdate onMessageUpdated onParticipantAdd onParticipantAdded onParticipantRemove onParticipantRemoved onParticipantUpdate onParticipantUpdated onUserAdded onUserUpdate onUserUpdated\], got test`),
			},
		},
	})
}

func testAccCheckTwilioConversationsWebhookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != webhookResourceName {
			continue
		}

		if _, err := client.Webhook().Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioConversationsWebhookExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Webhook().Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving webhook information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsWebhook_basic() string {
	return `
resource "twilio_conversations_webhook" "webhook" {}
`
}

func testAccTwilioConversationsWebhook_withTarget(target string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_webhook" "webhook" {
  target = "%s"
}
`, target)
}

func testAccTwilioConversationsWebhook_withMethodAndUrls(method string, preWebhookUrl string, postWebhookUrl string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_webhook" "webhook" {
  method           = "%s"
  pre_webhook_url  = "%s"
  post_webhook_url = "%s"
}
`, method, preWebhookUrl, postWebhookUrl)
}

func testAccTwilioConversationsWebhook_filters(filter []string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_webhook" "webhook" {
  filters = %s
}
`, "[\""+strings.Join(filter, "\",\"")+"\"]")
}
