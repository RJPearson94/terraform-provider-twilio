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

var serviceResourceName = "twilio_sync_service"

func TestAccTwilioSyncService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSyncServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSyncService_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "acl_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_window", "5000"),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_webhooks_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "webhooks_from_rest_enabled", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSyncServiceImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSyncService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSyncServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSyncService_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "acl_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_window", "5000"),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_webhooks_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "webhooks_from_rest_enabled", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioSyncService_friendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "acl_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_window", "5000"),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_webhooks_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "webhooks_from_rest_enabled", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioSyncService_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSyncServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSyncService_friendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioSyncService_friendlyName(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
			{
				Config: testAccTwilioSyncService_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioSyncService_reachabilityDebouncingWindow(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSyncServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSyncService_reachabilityDebouncingWindow(1000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_window", "1000"),
				),
			},
			{
				Config: testAccTwilioSyncService_reachabilityDebouncingWindow(30000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_window", "30000"),
				),
			},
			{
				Config: testAccTwilioSyncService_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "reachability_debouncing_window", "5000"),
				),
			},
		},
	})
}

func TestAccTwilioSyncService_invalidReachabilityDebouncingWindow(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSyncService_reachabilityDebouncingWindow(1),
				ExpectError: regexp.MustCompile(`(?s)expected reachability_debouncing_window to be in the range \(1000 - 30000\), got 1`),
			},
		},
	})
}

func TestAccTwilioSyncService_webhookUrl(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	webhookUrl := "http://localhost.com/webhookUrl"
	webhookUrlSecure := "https://localhost.com/webhookUrl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSyncServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSyncService_webhookUrl(webhookUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrl),
				),
			},
			{
				Config: testAccTwilioSyncService_webhookUrl(webhookUrlSecure),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", webhookUrlSecure),
				),
			},
			{
				Config: testAccTwilioSyncService_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSyncServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "webhook_url", ""),
				),
			},
		},
	})
}

func TestAccTwilioConversationsWebhook_invalidWebhookUrl(t *testing.T) {
	webhookUrl := "webhookUrl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSyncService_webhookUrl(webhookUrl),
				ExpectError: regexp.MustCompile(`(?s)expected "webhook_url" to have a host, got webhookUrl`),
			},
		},
	})
}

func testAccCheckTwilioSyncServiceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Sync

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

func testAccCheckTwilioSyncServiceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Sync

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

func testAccTwilioSyncServiceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSyncService_basic() string {
	return `resource "twilio_sync_service" "service" {}`
}

func testAccTwilioSyncService_friendlyName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_sync_service" "service" {
  friendly_name = "%[1]s"
}
`, friendlyName)
}

func testAccTwilioSyncService_reachabilityDebouncingWindow(reachabilityDebouncingWindow int) string {
	return fmt.Sprintf(`
resource "twilio_sync_service" "service" {
  reachability_debouncing_window = %[1]d
}
`, reachabilityDebouncingWindow)
}

func testAccTwilioSyncService_webhookUrl(webhookUrl string) string {
	return fmt.Sprintf(`
resource "twilio_sync_service" "service" {
  webhook_url = "%[1]s"
}
`, webhookUrl)
}
