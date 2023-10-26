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

const appResourceName = "twilio_twiml_app"

func TestAccTwilioTwimlApp_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTwimlApp_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.caller_id_lookup", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioTwimlAppImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioTwimlApp_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTwimlApp_basic(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.caller_id_lookup", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioTwimlApp_friendlyName(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.caller_id_lookup", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioTwimlApp_messagingStatusCallbackUrl(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

	testData := acceptance.TestAccData
	statusCallbackUrl := "http://localhost.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTwimlApp_messagingStatusCallbackUrl(testData, statusCallbackUrl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.status_callback_url", statusCallbackUrl),
				),
			},
			{
				Config: testAccTwilioTwimlApp_defaultMessaging(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.status_callback_url", ""),
				),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidMessagingStatusCallbackUrl(t *testing.T) {
	testData := acceptance.TestAccData
	statusCallbackUrl := "statusCallbackUrl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_messagingStatusCallbackUrl(testData, statusCallbackUrl),
				ExpectError: regexp.MustCompile(`(?s)expected "messaging.0.status_callback_url" to have a host, got statusCallbackUrl`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_messagingFallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

	testData := acceptance.TestAccData
	fallbackUrl := "http://localhost.com"
	fallbackMethod := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTwimlApp_messagingFallback(testData, fallbackUrl, fallbackMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_url", fallbackUrl),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_method", fallbackMethod),
				),
			},
			{
				Config: testAccTwilioTwimlApp_messaging(testData, "http://localhost.com/messaging", "POST"), // Can't use defaultMessaging due to the url bug which is showing up as drift
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.fallback_method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidMessagingFallbackMethod(t *testing.T) {
	testData := acceptance.TestAccData
	fallbackUrl := "http://localhost.com"
	fallbackMethod := "DELETE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_messagingFallback(testData, fallbackUrl, fallbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected messaging.0.fallback_method to be one of \["GET" "POST"\], got DELETE`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidMessagingFallbackUrl(t *testing.T) {
	testData := acceptance.TestAccData
	fallbackUrl := "fallbackUrl"
	fallbackMethod := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_messagingFallback(testData, fallbackUrl, fallbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected "messaging.0.fallback_url" to have a host, got fallbackUrl`),
			},
		},
	})
}

// The Twilio API is not currently allowing this to be replaced with an empty string
// func TestAccTwilioTwimlApp_messaging(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

// 	testData := acceptance.TestAccData
// 	url := "http://localhost.com"
// 	method := "GET"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioTwimlApp_messaging(testData, url, method),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioTwimlAppExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
// 					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.url", url),
// 					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.method", method),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioTwimlApp_defaultMessaging(testData),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioTwimlAppExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "messaging.#", "1"),
// 					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.url", ""),
// 					resource.TestCheckResourceAttr(stateResourceName, "messaging.0.method", "POST"),
// 				),
// 			},
// 		},
// 	})
// }

func TestAccTwilioTwimlApp_invalidMessagingMethod(t *testing.T) {
	testData := acceptance.TestAccData
	url := "http://localhost.com"
	method := "DELETE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_messaging(testData, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected messaging.0.method to be one of \["GET" "POST"\], got DELETE`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidMessagingUrl(t *testing.T) {
	testData := acceptance.TestAccData
	url := "url"
	method := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_messaging(testData, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected "messaging.0.url" to have a host, got url`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_voiceCallerIdLookup(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

	testData := acceptance.TestAccData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTwimlApp_voiceCallerIdLookupTrue(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.caller_id_lookup", "true"),
				),
			},
			{
				Config: testAccTwilioTwimlApp_defaultVoice(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.caller_id_lookup", "false"),
				),
			},
		},
	})
}

func TestAccTwilioTwimlApp_voiceStatusCallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

	testData := acceptance.TestAccData
	statusCallbackUrl := "http://localhost.com"
	statusCallbackMethod := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTwimlApp_voiceStatusCallback(testData, statusCallbackUrl, statusCallbackMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", statusCallbackUrl),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", statusCallbackMethod),
				),
			},
			{
				Config: testAccTwilioTwimlApp_defaultVoice(testData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidVoiceStatusCallbackMethod(t *testing.T) {
	testData := acceptance.TestAccData
	statusCallbackUrl := "http://localhost.com"
	statusCallbackMethod := "DELETE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_voiceStatusCallback(testData, statusCallbackUrl, statusCallbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected voice.0.status_callback_method to be one of \["GET" "POST"\], got DELETE`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidVoiceStatusCallbackUrl(t *testing.T) {
	testData := acceptance.TestAccData
	statusCallbackUrl := "fallbackUrl"
	statusCallbackMethod := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_voiceStatusCallback(testData, statusCallbackUrl, statusCallbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected "voice.0.status_callback_url" to have a host, got fallbackUrl`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_voiceFallback(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

	testData := acceptance.TestAccData
	fallbackUrl := "http://localhost.com"
	fallbackMethod := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTwimlApp_voiceFallback(testData, fallbackUrl, fallbackMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", fallbackUrl),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", fallbackMethod),
				),
			},
			{
				Config: testAccTwilioTwimlApp_voice(testData, "http://localhost.com/voice", "POST"), // Can't use defaultVoice due to the url bug which is showing up as drift
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTwimlAppExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidVoiceFallbackMethod(t *testing.T) {
	testData := acceptance.TestAccData
	fallbackUrl := "http://localhost.com"
	fallbackMethod := "DELETE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_voiceFallback(testData, fallbackUrl, fallbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected voice.0.fallback_method to be one of \["GET" "POST"\], got DELETE`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidVoiceFallbackUrl(t *testing.T) {
	testData := acceptance.TestAccData
	fallbackUrl := "fallbackUrl"
	fallbackMethod := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_voiceFallback(testData, fallbackUrl, fallbackMethod),
				ExpectError: regexp.MustCompile(`(?s)expected "voice.0.fallback_url" to have a host, got fallbackUrl`),
			},
		},
	})
}

// The Twilio API is not currently allowing this to be replaced with an empty string
// func TestAccTwilioTwimlApp_voice(t *testing.T) {
// 	stateResourceName := fmt.Sprintf("%s.app", appResourceName)

// 	testData := acceptance.TestAccData
// 	url := "http://localhost.com"
// 	method := "GET"

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:          func() { acceptance.PreCheck(t) },
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		CheckDestroy:      testAccCheckTwilioTwimlAppDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTwilioTwimlApp_voice(testData, url, method),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioTwimlAppExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
// 					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", url),
// 					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", method),
// 				),
// 			},
// 			{
// 				Config: testAccTwilioTwimlApp_defaultVoice(testData),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckTwilioTwimlAppExists(stateResourceName),
// 					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
// 					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
// 					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
// 				),
// 			},
// 		},
// 	})
// }

func TestAccTwilioTwimlApp_invalidVoiceMethod(t *testing.T) {
	testData := acceptance.TestAccData
	url := "http://localhost.com"
	method := "DELETE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_voice(testData, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected voice.0.method to be one of \["GET" "POST"\], got DELETE`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidVoiceUrl(t *testing.T) {
	testData := acceptance.TestAccData
	url := "url"
	method := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_voice(testData, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected "voice.0.url" to have a host, got url`),
			},
		},
	})
}

func TestAccTwilioTwimlApp_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioTwimlApp_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccCheckTwilioTwimlAppDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != appResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Application(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving app information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioTwimlAppExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Application(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving app information %s", err)
		}

		return nil
	}
}

func testAccTwilioTwimlAppImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/Applications/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioTwimlApp_basic(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
}
`, testData.AccountSid)
}

func testAccTwilioTwimlApp_friendlyName(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid   = "%s"
  friendly_name = "%s"
}
`, testData.AccountSid, friendlyName)
}

func testAccTwilioTwimlApp_messagingStatusCallbackUrl(testData *acceptance.TestData, statusCallbackUrl string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  messaging {
    status_callback_url = "%s"
  }
}
`, testData.AccountSid, statusCallbackUrl)
}

func testAccTwilioTwimlApp_messagingFallback(testData *acceptance.TestData, fallbackUrl string, fallbackMethod string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  messaging {
    url             = "http://localhost.com/messaging"
    method          = "POST"
    fallback_url    = "%s"
    fallback_method = "%s"
  }
}
`, testData.AccountSid, fallbackUrl, fallbackMethod)
}

func testAccTwilioTwimlApp_messaging(testData *acceptance.TestData, url string, method string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  messaging {
    url    = "%s"
    method = "%s"
  }
}
`, testData.AccountSid, url, method)
}

func testAccTwilioTwimlApp_defaultMessaging(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  messaging {}
}
`, testData.AccountSid)
}

func testAccTwilioTwimlApp_voiceStatusCallback(testData *acceptance.TestData, statusCallbackUrl string, statusCallbackMethod string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  voice {
    status_callback_url    = "%s"
    status_callback_method = "%s"
  }
}
`, testData.AccountSid, statusCallbackUrl, statusCallbackMethod)
}

func testAccTwilioTwimlApp_voiceFallback(testData *acceptance.TestData, fallbackUrl string, fallbackMethod string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  voice {
    url             = "http://localhost.com/voice"
    method          = "POST"
    fallback_url    = "%s"
    fallback_method = "%s"
  }
}
`, testData.AccountSid, fallbackUrl, fallbackMethod)
}

func testAccTwilioTwimlApp_voice(testData *acceptance.TestData, url string, method string) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  voice {
    url    = "%s"
    method = "%s"
  }
}
`, testData.AccountSid, url, method)
}

func testAccTwilioTwimlApp_voiceCallerIdLookupTrue(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  voice {
    caller_id_lookup = true
  }
}
`, testData.AccountSid)
}

func testAccTwilioTwimlApp_defaultVoice(testData *acceptance.TestData) string {
	return fmt.Sprintf(`
resource "twilio_twiml_app" "app" {
  account_sid = "%s"
  voice {}
}
`, testData.AccountSid)
}

func testAccTwilioTwimlApp_invalidAccountSid() string {
	return `
resource "twilio_twiml_app" "app" {
  account_sid = "account_sid"
}
`
}
