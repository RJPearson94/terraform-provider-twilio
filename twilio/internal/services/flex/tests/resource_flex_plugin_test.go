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

var pluginResourceName = "twilio_flex_plugin"

func TestAccTwilioFlexPlugin_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.plugin", pluginResourceName)

	uniqueName := acctest.RandString(10)
	version := "1.0.0"
	pluginURL := "https://example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioFlexPluginDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexPlugin_basic(uniqueName, version, pluginURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexPluginExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttr(stateResourceName, "archived", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "changelog", ""),
					resource.TestCheckResourceAttr(stateResourceName, "version", version),
					resource.TestCheckResourceAttr(stateResourceName, "plugin_url", pluginURL),
					resource.TestCheckResourceAttrSet(stateResourceName, "private"),
					resource.TestCheckResourceAttr(stateResourceName, "version_archived", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioFlexPluginImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioFlexPlugin_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.plugin", pluginResourceName)

	uniqueName := acctest.RandString(10)
	version := "1.0.0"
	pluginURL := "https://example.com"
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioFlexPluginDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexPlugin_basic(uniqueName, version, pluginURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexPluginExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttr(stateResourceName, "archived", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "changelog", ""),
					resource.TestCheckResourceAttr(stateResourceName, "version", version),
					resource.TestCheckResourceAttr(stateResourceName, "plugin_url", pluginURL),
					resource.TestCheckResourceAttrSet(stateResourceName, "private"),
					resource.TestCheckResourceAttr(stateResourceName, "version_archived", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioFlexPlugin_friendlyName(uniqueName, newFriendlyName, version, pluginURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttr(stateResourceName, "archived", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "changelog", ""),
					resource.TestCheckResourceAttr(stateResourceName, "version", version),
					resource.TestCheckResourceAttr(stateResourceName, "plugin_url", pluginURL),
					resource.TestCheckResourceAttrSet(stateResourceName, "private"),
					resource.TestCheckResourceAttr(stateResourceName, "version_archived", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioFlexPlugin_blankUniqueName(t *testing.T) {
	uniqueName := ""
	version := "1.0.0"
	pluginURL := "https://example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioFlexPlugin_basic(uniqueName, version, pluginURL),
				ExpectError: regexp.MustCompile(`(?s)expected \"unique_name\" to not be an empty string, got `),
			},
		},
	})
}

func testAccCheckTwilioFlexPluginDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

	for _, rs := range s.RootModule().Resources {
		if rs.Type != pluginResourceName {
			continue
		}

		if _, err := client.Plugin(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving plugin information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioFlexPluginExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Plugin(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving plugin information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioFlexPluginImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/PluginService/Plugins/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioFlexPlugin_basic(uniqueName string, version string, pluginURL string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin" "plugin" {
  unique_name = "%s"
  version     = "%s"
  plugin_url  = "%s"
}
`, uniqueName, version, pluginURL)
}

func testAccTwilioFlexPlugin_friendlyName(uniqueName string, friendlyName string, version string, pluginURL string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin" "plugin" {
  unique_name   = "%s"
  friendly_name = "%s"
  version       = "%s"
  plugin_url    = "%s"
}
`, uniqueName, friendlyName, version, pluginURL)
}
