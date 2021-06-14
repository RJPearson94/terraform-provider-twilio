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

var pluginConfigurationResourceName = "twilio_flex_plugin_configuration"

func TestAccTwilioFlexPluginConfiguration_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.plugin_configuration", pluginConfigurationResourceName)

	name := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioFlexPluginConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexPluginConfiguration_basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexPluginConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "name", name),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttr(stateResourceName, "plugins.#", "0"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioFlexPluginConfigurationImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioFlexPluginConfiguration_createBeforeRelease(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.plugin_configuration", pluginConfigurationResourceName)

	name := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioFlexPluginConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexPluginConfiguration_withPluginCreateBeforeRelease(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexPluginConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "name", name),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttr(stateResourceName, "plugins.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioFlexPluginConfiguration_createBeforeRelease(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexPluginConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "name", name),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttr(stateResourceName, "plugins.#", "0"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioFlexPluginConfiguration_blankName(t *testing.T) {
	name := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioFlexPluginConfiguration_basic(name),
				ExpectError: regexp.MustCompile(`(?s)expected \"name\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioFlexPluginConfiguration_withPlugins(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.plugin_configuration", pluginConfigurationResourceName)
	pluginStateResourceName := "twilio_flex_plugin.plugin"

	name := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioFlexPluginConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexPluginConfiguration_withPlugins(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexPluginConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "name", name),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttr(stateResourceName, "plugins.#", "1"),
					resource.TestCheckResourceAttrPair(stateResourceName, "plugins.0.plugin_version_sid", pluginStateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "plugins.0.plugin_sid", pluginStateResourceName, "sid"),
					resource.TestCheckResourceAttrPair(stateResourceName, "plugins.0.plugin_url", pluginStateResourceName, "plugin_url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "plugins.0.phase"),
					resource.TestCheckResourceAttrPair(stateResourceName, "plugins.0.private", pluginStateResourceName, "private"),
					resource.TestCheckResourceAttrPair(stateResourceName, "plugins.0.unique_name", pluginStateResourceName, "unique_name"),
					resource.TestCheckResourceAttrPair(stateResourceName, "plugins.0.version", pluginStateResourceName, "version"),
					resource.TestCheckResourceAttrSet(stateResourceName, "plugins.0.date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "plugins.0.url"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioFlexPluginConfiguration_invalidPluginVersionSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioFlexPluginConfiguration_withInvalidPluginVersionSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of plugins.0.plugin_version_sid to match regular expression "\^FV\[0-9a-fA-F\]\{32\}\$", got plugin_version_sid`),
			},
		},
	})
}

func testAccCheckTwilioFlexPluginConfigurationDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

	for _, rs := range s.RootModule().Resources {
		if rs.Type != pluginConfigurationResourceName {
			continue
		}

		if _, err := client.PluginConfiguration(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving plugin configuration information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioFlexPluginConfigurationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.PluginConfiguration(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving plugin configuration information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioFlexPluginConfigurationImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/PluginService/Configurations/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioFlexPluginConfiguration_basic(name string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "%s"
}
`, name)
}

func testAccTwilioFlexPluginConfiguration_withPlugins(name string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin" "plugin" {
  unique_name = "%[1]s"
  version     = "1.0.0"
  plugin_url  = "https://example.com"
}

resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "%[1]s"
  plugins {
    plugin_version_sid = twilio_flex_plugin.plugin.latest_version_sid
  }
}
`, name)
}

func testAccTwilioFlexPluginConfiguration_createBeforeRelease(name string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "%[1]s"

	lifecycle {
		create_before_destroy = true
	}
}
`, name)
}

func testAccTwilioFlexPluginConfiguration_withPluginCreateBeforeRelease(name string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin" "plugin" {
  unique_name = "%[1]s"
  version     = "1.0.0"
  plugin_url  = "https://example.com"
}

resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "%[1]s"
  plugins {
    plugin_version_sid = twilio_flex_plugin.plugin.latest_version_sid
  }

	lifecycle {
		create_before_destroy = true
	}
}
`, name)
}

func testAccTwilioFlexPluginConfiguration_withInvalidPluginVersionSid() string {
	return `
resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "invalid_plugin_version_sid"
  plugins {
    plugin_version_sid = "plugin_version_sid"
  }
}
`
}
