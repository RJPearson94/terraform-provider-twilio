package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var pluginDataSourceName = "twilio_flex_plugin"

func TestAccDataSourceTwilioFlexPlugin_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.plugin", pluginDataSourceName)

	uniqueName := acctest.RandString(10)
	version := "1.0.0"
	pluginURL := "https://example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioFlexPlugin_basic(uniqueName, version, pluginURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "friendly_name"),
					resource.TestCheckResourceAttr(stateDataSourceName, "description", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "archived", "false"),
					resource.TestCheckResourceAttr(stateDataSourceName, "changelog", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "version", version),
					resource.TestCheckResourceAttr(stateDataSourceName, "plugin_url", pluginURL),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "private"),
					resource.TestCheckResourceAttr(stateDataSourceName, "version_archived", "false"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioFlexPlugin_basic(uniqueName string, version string, pluginURL string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin" "plugin" {
  unique_name = "%s"
  version     = "%s"
  plugin_url  = "%s"
}

data "twilio_flex_plugin" "plugin" {
  sid = twilio_flex_plugin.plugin.sid
}
`, uniqueName, version, pluginURL)
}
