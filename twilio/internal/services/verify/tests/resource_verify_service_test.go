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

var serviceResourceName = "twilio_verify_service"

func TestAccTwilioVerifyService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "custom_code_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "do_not_share_warning_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "dtmf_input_required", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "lookup_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "mailer_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "psd2_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "push.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "push.0.apn_credential_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "push.0.fcm_credential_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "totp.0.issuer"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "30"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "default_template_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "skip_sms_to_landlines", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "tts_name", ""),
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
				ImportStateIdFunc: testAccTwilioVerifyServiceImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioVerifyService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "custom_code_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "do_not_share_warning_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "dtmf_input_required", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "lookup_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "mailer_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "psd2_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "push.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "push.0.apn_credential_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "push.0.fcm_credential_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "totp.0.issuer"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "30"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "default_template_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "skip_sms_to_landlines", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "tts_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioVerifyService_basic(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "custom_code_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "do_not_share_warning_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "dtmf_input_required", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "lookup_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "mailer_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "psd2_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "push.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "push.0.apn_credential_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "push.0.fcm_credential_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "totp.0.issuer"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "30"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "default_template_sid", ""),
					resource.TestCheckResourceAttr(stateResourceName, "skip_sms_to_landlines", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "tts_name", ""),
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

func TestAccTwilioVerifyService_invalidFriendlyNameWith0Characters(t *testing.T) {
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_basic(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 30\), got `),
			},
		},
	})
}

func TestAccTwilioVerifyService_invalidFriendlyNameWith31Characters(t *testing.T) {
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_basic(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 30\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioVerifyService_totp(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	friendlyName := acctest.RandString(1)
	issuer := "Test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyService_totp(friendlyName, issuer, 20, 3, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.issuer", "Test"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "20"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "3"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "0"),
				),
			},
			{
				Config: testAccTwilioVerifyService_totp(friendlyName, issuer, 60, 8, 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.issuer", "Test"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "60"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "8"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "2"),
				),
			},
			{
				Config: testAccTwilioVerifyService_totp(friendlyName, issuer, 20, 3, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.issuer", "Test"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "20"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "3"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "0"),
				),
			},
			{
				Config: testAccTwilioVerifyService_defaultTotp(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "totp.0.issuer"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "30"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "1"),
				),
			},
			{
				Config: testAccTwilioVerifyService_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "totp.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "totp.0.issuer"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.time_step", "30"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.code_length", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "totp.0.skew", "1"),
				),
			},
		},
	})
}

func TestAccTwilioVerifyService_invalidTotpTimestep(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_totp("friendlyName", "test", 0, 3, 0),
				ExpectError: regexp.MustCompile(`(?s)expected totp\.0\.time_step to be in the range \(20 - 60\), got 0`),
			},
		},
	})
}

func TestAccTwilioVerifyService_invalidTotpCodeLength(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_totp("friendlyName", "test", 30, 0, 0),
				ExpectError: regexp.MustCompile(`(?s)expected totp\.0\.code_length to be in the range \(3 - 8\), got 0`),
			},
		},
	})
}

func TestAccTwilioVerifyService_invalidTotpSkew(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_totp("friendlyName", "test", 30, 6, 3),
				ExpectError: regexp.MustCompile(`(?s)expected totp\.0\.skew to be in the range \(0 - 2\), got 3`),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_invalidPushApnCredentialSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_push("friendlyName", "apn_credential_sid", "CRaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
				ExpectError: regexp.MustCompile(`(?s)expected value of push\.0\.apn_credential_sid to match regular expression "\^CR\[0-9a-fA-F\]\{32\}\$", got apn_credential_sid`),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_invalidPushFcmCredentialSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_push("friendlyName", "CRaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "fcm_credential_sid"),
				ExpectError: regexp.MustCompile(`(?s)expected value of push\.0\.fcm_credential_sid to match regular expression "\^CR\[0-9a-fA-F\]\{32\}\$", got fcm_credential_sid`),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_invalidMailerSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_invalidMailerSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of mailer_sid to match regular expression "\^MD\[0-9a-fA-F\]\{32\}\$", got mailer_sid`),
			},
		},
	})
}

func TestAccTwilioVerifyWebhook_invalidDefaultTemplateSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyService_invalidDefaultTemplateSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of default_template_sid to match regular expression "\^HJ\[0-9a-fA-F\]\{32\}\$", got default_template_sid`),
			},
		},
	})
}

func testAccCheckTwilioVerifyServiceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

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

func testAccCheckTwilioVerifyServiceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

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

func testAccTwilioVerifyServiceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioVerifyService_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}
`, friendlyName)
}

func testAccTwilioVerifyService_totp(friendlyName string, issuer string, timeStep int, codeLength int, skew int) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
  totp {
    issuer      = "%[2]s"
    time_step   = %[3]d
    code_length = %[4]d
    skew        = %[5]d
  }
}
`, friendlyName, issuer, timeStep, codeLength, skew)
}

func testAccTwilioVerifyService_defaultTotp(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
  totp {}
}
`, friendlyName)
}

func testAccTwilioVerifyService_push(friendlyName string, apnCredentialSid string, fcmCredentialSid string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
  push {
    apn_credential_sid = "%[2]s"
    fcm_credential_sid = "%[3]s"
  }
}
`, friendlyName, apnCredentialSid, fcmCredentialSid)
}

func testAccTwilioVerifyService_invalidMailerSid() string {
	return `
resource "twilio_verify_service" "service" {
  friendly_name = "invalid mailer sid"
  mailer_sid    = "mailer_sid"
}
`
}

func testAccTwilioVerifyService_invalidDefaultTemplateSid() string {
	return `
resource "twilio_verify_service" "service" {
  friendly_name        = "invalid default template sid"
  default_template_sid = "default_template_sid"
}
`
}
