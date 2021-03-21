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

var compositionHookResourceName = "twilio_video_composition_hook"

func TestAccTwilioVideoCompositionHook_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.composition_hook", compositionHookResourceName)
	friendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "mp4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVideoCompositionHook_basic(friendlyName, audioSource, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoCompositionHookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.0", audioSource),
					resource.TestCheckResourceAttr(stateResourceName, "format", format),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources_excluded.#", "0"),
					resource.TestCheckResourceAttrSet(stateResourceName, "enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "resolution"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "status_callback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trim"),
					resource.TestCheckResourceAttr(stateResourceName, "video_layout", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioVideoCompositionHookImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioVideoCompositionHook_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.composition_hook", compositionHookResourceName)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "mp4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVideoCompositionHook_basic(friendlyName, audioSource, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoCompositionHookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.0", audioSource),
					resource.TestCheckResourceAttr(stateResourceName, "format", format),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources_excluded.#", "0"),
					resource.TestCheckResourceAttrSet(stateResourceName, "enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "resolution"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "status_callback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trim"),
					resource.TestCheckResourceAttr(stateResourceName, "video_layout", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioVideoCompositionHook_basic(newFriendlyName, audioSource, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoCompositionHookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.0", audioSource),
					resource.TestCheckResourceAttr(stateResourceName, "format", format),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources_excluded.#", "0"),
					resource.TestCheckResourceAttrSet(stateResourceName, "enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "resolution"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "status_callback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trim"),
					resource.TestCheckResourceAttr(stateResourceName, "video_layout", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioVideoCompositionHook_videoLayout(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.composition_hook", compositionHookResourceName)
	friendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "mp4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVideoCompositionHook_videoLayout(friendlyName, audioSource, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoCompositionHookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.0", audioSource),
					resource.TestCheckResourceAttr(stateResourceName, "format", format),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources_excluded.#", "0"),
					resource.TestCheckResourceAttrSet(stateResourceName, "enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "resolution"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "status_callback_method"),
					resource.TestCheckResourceAttrSet(stateResourceName, "trim"),
					resource.TestCheckResourceAttr(stateResourceName, "video_layout", "{\"grid\":{\"cells_excluded\":[],\"height\":null,\"max_columns\":null,\"max_rows\":null,\"reuse\":\"show_oldest\",\"video_sources\":[\"*\"],\"video_sources_excluded\":[],\"width\":null,\"x_pos\":0,\"y_pos\":0,\"z_pos\":0}}"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioVideoCompositionHook_invalidVideoLayout(t *testing.T) {
	friendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "mp4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVideoCompositionHook_invalidVideoLayout(friendlyName, audioSource, format),
				ExpectError: regexp.MustCompile(`(?s)"video_layout" contains an invalid JSON`),
			},
		},
	})
}

func TestAccTwilioVideoCompositionHook_invalidFormat(t *testing.T) {
	friendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVideoCompositionHook_basic(friendlyName, audioSource, format),
				ExpectError: regexp.MustCompile(`(?s)expected format to be one of \[mp4 webm\], got test`),
			},
		},
	})
}

func TestAccTwilioVideoCompositionHook_statusCallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.composition_hook", compositionHookResourceName)
	friendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "mp4"
	callbackURL := "https://test.com/callback"
	callbackMethod := "POST"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVideoCompositionHook_statusCallback(friendlyName, audioSource, format, callbackURL, callbackMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVideoCompositionHookExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources.0", audioSource),
					resource.TestCheckResourceAttr(stateResourceName, "format", format),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "audio_sources_excluded.#", "0"),
					resource.TestCheckResourceAttrSet(stateResourceName, "enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "resolution"),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_url", callbackURL),
					resource.TestCheckResourceAttr(stateResourceName, "status_callback_method", callbackMethod),
					resource.TestCheckResourceAttrSet(stateResourceName, "trim"),
					resource.TestCheckResourceAttr(stateResourceName, "video_layout", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioVideoCompositionHook_invalidStatusCallbackURL(t *testing.T) {
	friendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "mp4"
	callbackURL := "test"
	callbackMethod := "POST"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVideoCompositionHook_statusCallback(friendlyName, audioSource, format, callbackURL, callbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected "status_callback_url" to have a host, got test`),
			},
		},
	})
}

func TestAccTwilioVideoCompositionHook_invalidStatusCallbackMethod(t *testing.T) {
	friendlyName := acctest.RandString(10)
	audioSource := "*"
	format := "mp4"
	callbackURL := "https://test.com/callback"
	callbackMethod := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVideoCompositionHookDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVideoCompositionHook_statusCallback(friendlyName, audioSource, format, callbackURL, callbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected status_callback_method to be one of \[GET POST\], got test`),
			},
		},
	})
}

func testAccCheckTwilioVideoCompositionHookDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Video

	for _, rs := range s.RootModule().Resources {
		if rs.Type != compositionHookResourceName {
			continue
		}

		if _, err := client.CompositionHook(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving composition hook information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioVideoCompositionHookExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Video

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.CompositionHook(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving composition hook information %s", err)
		}

		return nil
	}
}

func testAccTwilioVideoCompositionHookImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/CompositionHooks/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioVideoCompositionHook_basic(friendlyName string, audio_source string, format string) string {
	return fmt.Sprintf(`
resource "twilio_video_composition_hook" "composition_hook" {
  friendly_name = "%s"
  audio_sources = ["%s"]
  format        = "%s"
}
`, friendlyName, audio_source, format)
}

func testAccTwilioVideoCompositionHook_videoLayout(friendlyName string, audio_source string, format string) string {
	return fmt.Sprintf(`
resource "twilio_video_composition_hook" "composition_hook" {
  friendly_name = "%s"
  audio_sources = ["%s"]
  format        = "%s"
  video_layout = jsonencode({
    "grid" : {
      "cells_excluded" : [],
      "height" : null,
      "max_columns" : null,
      "max_rows" : null,
      "reuse" : "show_oldest",
      "video_sources" : ["*"],
      "video_sources_excluded" : [],
      "width" : null,
      "x_pos" : 0,
      "y_pos" : 0,
      "z_pos" : 0
    }
  })
}
`, friendlyName, audio_source, format)
}

func testAccTwilioVideoCompositionHook_invalidVideoLayout(friendlyName string, audio_source string, format string) string {
	return fmt.Sprintf(`
resource "twilio_video_composition_hook" "composition_hook" {
  friendly_name = "%s"
  audio_sources = ["%s"]
  format        = "%s"
  video_layout  = "test"
}
`, friendlyName, audio_source, format)
}

func testAccTwilioVideoCompositionHook_statusCallback(friendlyName string, audio_source string, format string, callbackURL string, callbackMethod string) string {
	return fmt.Sprintf(`
resource "twilio_video_composition_hook" "composition_hook" {
  friendly_name          = "%s"
  audio_sources          = ["%s"]
  format                 = "%s"
  status_callback_url    = "%s"
  status_callback_method = "%s"
}
`, friendlyName, audio_source, format, callbackURL, callbackMethod)
}
