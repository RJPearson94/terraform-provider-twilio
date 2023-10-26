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

var domainResourceName = "twilio_sip_domain"

func TestAccTwilioSIPDomain_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "domain_name", domainName),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "auth_type", ""),
					resource.TestCheckResourceAttr(stateResourceName, "byoc_trunk_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.calling_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.caller_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "secure"),
					resource.TestCheckResourceAttr(stateResourceName, "sip_registration", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioSIPDomainImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSIPDomain_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	newDomainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "domain_name", domainName),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "auth_type", ""),
					resource.TestCheckResourceAttr(stateResourceName, "byoc_trunk_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.calling_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.caller_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "secure"),
					resource.TestCheckResourceAttr(stateResourceName, "sip_registration", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_basic(testData, newDomainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "domain_name", newDomainName),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "auth_type", ""),
					resource.TestCheckResourceAttr(stateResourceName, "byoc_trunk_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.calling_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.caller_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "secure"),
					resource.TestCheckResourceAttr(stateResourceName, "sip_registration", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioSIPDomain_voiceURLAndMethod(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "https://demo.twilio.com/welcome/voice/"
	method := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_voiceURLAndMethod(testData, domainName, url, method),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", url),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", method),
				),
			},
			{
				Config: testAccTwilioSIPDomain_emptyVoiceBlock(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidVoiceURL(t *testing.T) {
	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "test"
	method := "POST"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_voiceURLAndMethod(testData, domainName, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected "voice.0.url" to have a host, got test`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidVoiceMethod(t *testing.T) {
	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "https://demo.twilio.com/welcome/voice/"
	method := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_voiceURLAndMethod(testData, domainName, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected voice.0.method to be one of \["GET" "POST"\], got test`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_voiceFallbackURLAndMethod(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "https://demo.twilio.com/welcome/voice/"
	method := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_voiceFallbackURLAndMethod(testData, domainName, url, method),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", method),
				),
			},
			{
				Config: testAccTwilioSIPDomain_emptyVoiceBlock(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.fallback_method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidVoiceFallbackURL(t *testing.T) {
	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "test"
	method := "POST"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_voiceFallbackURLAndMethod(testData, domainName, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected "voice.0.fallback_url" to have a host, got test`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidVoiceFallbackMethod(t *testing.T) {
	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "https://demo.twilio.com/welcome/voice/"
	method := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_voiceFallbackURLAndMethod(testData, domainName, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected voice.0.fallback_method to be one of \["GET" "POST"\], got test`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_voiceStatusCallbackURLAndMethod(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "https://demo.twilio.com/welcome/voice/"
	method := "GET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_voiceStatusCallbackURLAndMethod(testData, domainName, url, method),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", url),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", method),
				),
			},
			{
				Config: testAccTwilioSIPDomain_emptyVoiceBlock(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "voice.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "voice.0.status_callback_method", "POST"),
				),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidVoiceStatusCallbackURL(t *testing.T) {
	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "test"
	method := "POST"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_voiceStatusCallbackURLAndMethod(testData, domainName, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected "voice.0.status_callback_url" to have a host, got test`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidVoiceStatusCallbackMethod(t *testing.T) {
	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	url := "https://demo.twilio.com/welcome/voice/"
	method := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_voiceStatusCallbackURLAndMethod(testData, domainName, url, method),
				ExpectError: regexp.MustCompile(`(?s)expected voice.0.status_callback_method to be one of \["GET" "POST"\], got test`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_emergency(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"
	emergencyCallingEnabled := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_emergency(testData, domainName, emergencyCallingEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.calling_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "emergency.0.caller_sid", ""),
				),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidEmergencyCallerSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_invalidEmergencyCallerSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of emergency.0.caller_sid to match regular expression "\^PN\[0-9a-fA-F\]\{32\}\$", got caller_sid`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidByocTrunkSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_invalidByocTrunkSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of byoc_trunk_sid to match regular expression "\^BY\[0-9a-fA-F\]\{32\}\$", got byoc_trunk_sid`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_secure(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "secure", "false"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_secureTrue(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "secure", "true"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "secure", "false"),
				),
			},
		},
	})
}

func TestAccTwilioSIPDomain_sipRegistration(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.domain", domainResourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sip_registration", "false"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_sipRegistrationTrue(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sip_registration", "true"),
				),
			},
			{
				Config: testAccTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPDomainExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "sip_registration", "false"),
				),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidDomainName(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_invalidDomainName(),
				ExpectError: regexp.MustCompile(`(?s)expected value of domain_name to match regular expression "\^\[a\-zA\-Z0\-9\-\.\]\+\\\\.sip\\\\.twilio\\\\.com\$", got domain_name`),
			},
		},
	})
}

func TestAccTwilioSIPDomain_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPDomain_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccCheckTwilioSIPDomainDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != domainResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.Domain(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving domain information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPDomainExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Sip.Domain(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving domain information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPDomainImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/SIP/Domains/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPDomain_basic(testData *acceptance.TestData, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
}
`, testData.AccountSid, domainName)
}

func testAccTwilioSIPDomain_voiceURLAndMethod(testData *acceptance.TestData, domainName string, url string, method string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
  voice {
    url    = "%s"
    method = "%s"
  }
}
`, testData.AccountSid, domainName, url, method)
}

func testAccTwilioSIPDomain_voiceFallbackURLAndMethod(testData *acceptance.TestData, domainName string, url string, method string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
  voice {
    fallback_url    = "%s"
    fallback_method = "%s"
  }
}
`, testData.AccountSid, domainName, url, method)
}

func testAccTwilioSIPDomain_voiceStatusCallbackURLAndMethod(testData *acceptance.TestData, domainName string, url string, method string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
  voice {
    status_callback_url    = "%s"
    status_callback_method = "%s"
  }
}
`, testData.AccountSid, domainName, url, method)
}

func testAccTwilioSIPDomain_emptyVoiceBlock(testData *acceptance.TestData, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
  voice {}
}
`, testData.AccountSid, domainName)
}

func testAccTwilioSIPDomain_emergency(testData *acceptance.TestData, domainName string, emergencyCallingEnabled bool) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
  emergency {
    calling_enabled = %t
  }
}
`, testData.AccountSid, domainName, emergencyCallingEnabled)
}

func testAccTwilioSIPDomain_invalidEmergencyCallerSid() string {
	return `
resource "twilio_sip_domain" "domain" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_name = "invalid-caller-sid.sip.twilio.com"
  emergency {
    calling_enabled = false
    caller_sid      = "caller_sid"
  }
}
`
}

func testAccTwilioSIPDomain_invalidByocTrunkSid() string {
	return `
resource "twilio_sip_domain" "domain" {
  account_sid    = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_name    = "invalid-byoc-trunk-sid.sip.twilio.com"
  byoc_trunk_sid = "byoc_trunk_sid"
}
`
}

func testAccTwilioSIPDomain_secureTrue(testData *acceptance.TestData, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
  secure      = true
}
`, testData.AccountSid, domainName)
}

func testAccTwilioSIPDomain_sipRegistrationTrue(testData *acceptance.TestData, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid      = "%s"
  domain_name      = "%s"
  sip_registration = true
}
`, testData.AccountSid, domainName)
}

func testAccTwilioSIPDomain_invalidAccountSid() string {
	return `
resource "twilio_sip_domain" "domain" {
  account_sid    = "account_sid"
  domain_name    = "invalid-account-sid.sip.twilio.com"
  byoc_trunk_sid = "BYaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioSIPDomain_invalidDomainName() string {
	return `
resource "twilio_sip_domain" "domain" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  domain_name = "domain_name"
}
`
}
