package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var compositionHookDataSourceName = "twilio_video_composition_hook"

func TestAccDataSourceTwilioVideoCompositionHook_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.composition_hook", compositionHookDataSourceName)
	friendlyName := acctest.RandString(10)
	audioSource := "*"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountCompositionHook_basic(friendlyName, audioSource),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "audio_sources.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "audio_sources.0", audioSource),
					resource.TestCheckResourceAttr(stateDataSourceName, "format", "webm"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "audio_sources_excluded.#", "0"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "enabled"),
					resource.TestCheckResourceAttr(stateDataSourceName, "resolution", "640x480"),
					resource.TestCheckResourceAttr(stateDataSourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "trim"),
					resource.TestCheckResourceAttr(stateDataSourceName, "video_layout", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountCompositionHook_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountCompositionHook_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^HK\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAccountCompositionHook_basic(friendlyName string, audio_source string) string {
	return fmt.Sprintf(`
resource "twilio_video_composition_hook" "composition_hook" {
  friendly_name = "%s"
  audio_sources = ["%s"]
}

data "twilio_video_composition_hook" "composition_hook" {
  sid = twilio_video_composition_hook.composition_hook.sid
}
`, friendlyName, audio_source)
}

func testAccDataSourceTwilioAccountCompositionHook_invalidSid() string {
	return `
data "twilio_video_composition_hook" "composition_hook" {
  sid = "sid"
}
`
}
