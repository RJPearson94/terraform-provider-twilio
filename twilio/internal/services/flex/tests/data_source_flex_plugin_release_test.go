package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var pluginReleaseDataSourceName = "twilio_flex_plugin_release"

func TestAccDataSourceTwilioFlexPluginRelease_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.plugin_release", pluginReleaseDataSourceName)

	name := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioFlexPluginRelease_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "configuration_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioFlexPluginRelease_basic(name string) string {
	return fmt.Sprintf(`
resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "%s"
}

resource "twilio_flex_plugin_release" "plugin_release" {
  configuration_sid = twilio_flex_plugin_configuration.plugin_configuration.sid
}

data "twilio_flex_plugin_release" "plugin_release" {
  sid = twilio_flex_plugin_release.plugin_release.sid
}
`, name)
}
