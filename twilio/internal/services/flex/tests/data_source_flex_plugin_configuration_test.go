package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var pluginConfigurationDataSourceName = "twilio_flex_plugin_configuration"

func TestAccDataSourceTwilioFlexPluginConfiguration_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.plugin_configuration", pluginConfigurationDataSourceName)

	name := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioFlexPluginConfiguration_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "name", name),
					resource.TestCheckResourceAttr(stateDataSourceName, "description", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "plugins.#", "0"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioFlexPluginConfiguration_withPlugins(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.plugin_configuration", pluginConfigurationDataSourceName)

	name := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioFlexPluginConfiguration_withPlugins(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "name", name),
					resource.TestCheckResourceAttr(stateDataSourceName, "description", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "plugins.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "plugins.0.plugin_version_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "plugins.0.plugin_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "plugins.0.plugin_url", "https://example.com"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "plugins.0.phase"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "plugins.0.private"),
					resource.TestCheckResourceAttr(stateDataSourceName, "plugins.0.unique_name", name),
					resource.TestCheckResourceAttr(stateDataSourceName, "plugins.0.version", "1.0.0"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "plugins.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "plugins.0.url"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioFlexPluginConfiguration_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioFlexPluginConfiguration_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^FJ\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioFlexPluginConfiguration_basic(name string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "%s"
}

data "twilio_flex_plugin_configuration" "plugin_configuration" {
  sid = twilio_flex_plugin_configuration.plugin_configuration.sid
}
`, name)
}

func testAccDataSourceTwilioFlexPluginConfiguration_withPlugins(name string) string {
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

data "twilio_flex_plugin_configuration" "plugin_configuration" {
  sid = twilio_flex_plugin_configuration.plugin_configuration.sid
}
`, name)
}

func testAccDataSourceTwilioFlexPluginConfiguration_invalidSid() string {
	return `
data "twilio_flex_plugin_configuration" "plugin_configuration" {
  sid = "sid"
}
`
}
