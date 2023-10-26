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

const trunkResourceName = "twilio_sip_trunking_trunk"

func TestAccTwilioSIPTrunkingTrunk_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttr(stateResourceName, "cnam_lookup_enabled", "false"),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_method", ""),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "domain_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "recording.#"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.mode", "do-not-record"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.trim", "do-not-trim"),
					resource.TestCheckResourceAttrSet(stateResourceName, "secure"),
					resource.TestCheckResourceAttr(stateResourceName, "transfer_mode", "disable-all"),
					resource.TestCheckResourceAttr(stateResourceName, "auth_type", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "auth_type_set.#"),
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
				ImportStateIdFunc: testAccTwilioSIPTrunkingTrunkImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "cnam_lookup_enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_method", ""),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "domain_name", ""),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "recording.#"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.mode", "do-not-record"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.trim", "do-not-trim"),
					resource.TestCheckResourceAttrSet(stateResourceName, "secure"),
					resource.TestCheckResourceAttr(stateResourceName, "transfer_mode", "disable-all"),
					resource.TestCheckResourceAttr(stateResourceName, "auth_type", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "auth_type_set.#"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_friendlyName(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "cnam_lookup_enabled"),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_method", ""),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_url", ""),
					resource.TestCheckResourceAttr(stateResourceName, "domain_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "recording.#"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.mode", "do-not-record"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.trim", "do-not-trim"),
					resource.TestCheckResourceAttrSet(stateResourceName, "secure"),
					resource.TestCheckResourceAttr(stateResourceName, "transfer_mode", "disable-all"),
					resource.TestCheckResourceAttr(stateResourceName, "auth_type", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "auth_type_set.#"),
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

func TestAccTwilioSIPTrunkingTrunk_recording(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	mode := "record-from-answer"
	trim := "trim-silence"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "recording.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.mode", "do-not-record"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.trim", "do-not-trim"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_recording(mode, trim),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "recording.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.mode", mode),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.trim", trim),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_emptyRecordingBlock(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "recording.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.mode", "do-not-record"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.trim", "do-not-trim"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "recording.#", "1"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.mode", "do-not-record"),
					resource.TestCheckResourceAttr(stateResourceName, "recording.0.trim", "do-not-trim"),
				),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_invalidRecordingMode(t *testing.T) {
	mode := "record-from-answer"
	trim := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingTrunk_recording(mode, trim),
				ExpectError: regexp.MustCompile(`(?s)expected recording.0.trim to be one of \["trim-silence" "do-not-trim"\], got test`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_invalidRecordingTrim(t *testing.T) {
	mode := "test"
	trim := "trim-silence"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingTrunk_recording(mode, trim),
				ExpectError: regexp.MustCompile(`(?s)expected recording.0.mode to be one of \["do-not-record" "record-from-ringing" "record-from-answer" "record-from-ringing-dual" "record-from-answer-dual"\], got test`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_disasterRecovery(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	method := "POST"
	url := "https://test.com/disaster-recovery"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_disasterRecovery(method, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_method", method),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_url", url),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_method", ""),
					resource.TestCheckResourceAttr(stateResourceName, "disaster_recovery_url", ""),
				),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_invalidDisasterRecoveryMethod(t *testing.T) {
	method := "DELETE"
	url := "http://localhost/disaster-recovery"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingTrunk_disasterRecovery(method, url),
				ExpectError: regexp.MustCompile(`(?s)expected disaster_recovery_method to be one of \["GET" "POST"\], got DELETE`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_invalidDisasterRecoveryURL(t *testing.T) {
	method := "POST"
	url := "testURL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingTrunk_disasterRecovery(method, url),
				ExpectError: regexp.MustCompile(`(?s)expected "disaster_recovery_url" to have a host, got testURL`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	friendlyName := ""
	newFriendlyName := acctest.RandString(64)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_friendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_friendlyName(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
				),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_invalidFriendlyNameWithLengthOf65(t *testing.T) {
	friendlyName := "7y80krlx0npe98jtdhahyvx8jvfz09x21x226uxj8gowkun6dgl2p1xj819qjzgtt"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingTrunk_friendlyName(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(0 - 64\), got 7y80krlx0npe98jtdhahyvx8jvfz09x21x226uxj8gowkun6dgl2p1xj819qjzgtt`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_secure(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "secure", "false"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_secureTrue(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "secure", "true"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "secure", "false"),
				),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_transferMode(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	newTransferMode := "sip-only"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "transfer_mode", "disable-all"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_transferMode(newTransferMode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "transfer_mode", newTransferMode),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "transfer_mode", "disable-all"),
				),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_invalidTransferMode(t *testing.T) {
	transferMode := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioSIPTrunkingTrunk_transferMode(transferMode),
				ExpectError: regexp.MustCompile(`(?s)expected transfer_mode to be one of \["enable-all" "sip-only" "disable-all"\], got test`),
			},
		},
	})
}

func TestAccTwilioSIPTrunkingTrunk_cnamLookupEnabled(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.trunk", trunkResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioSIPTrunkingTrunkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "cnam_lookup_enabled", "false"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_cnamLookupEnabledTrue(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "cnam_lookup_enabled", "true"),
				),
			},
			{
				Config: testAccTwilioSIPTrunkingTrunk_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioSIPTrunkingTrunkExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "cnam_lookup_enabled", "false"),
				),
			},
		},
	})
}

func testAccCheckTwilioSIPTrunkingTrunkDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

	for _, rs := range s.RootModule().Resources {
		if rs.Type != trunkResourceName {
			continue
		}

		if _, err := client.Trunk(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving trunk information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioSIPTrunkingTrunkExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).SIPTrunking

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Trunk(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving trunk information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioSIPTrunkingTrunkImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Trunks/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioSIPTrunkingTrunk_basic() string {
	return `
resource "twilio_sip_trunking_trunk" "trunk" {}
`
}

func testAccTwilioSIPTrunkingTrunk_recording(mode string, trim string) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {
  recording {
    mode = "%s"
    trim = "%s"
  }
}
`, mode, trim)
}

func testAccTwilioSIPTrunkingTrunk_emptyRecordingBlock() string {
	return `
resource "twilio_sip_trunking_trunk" "trunk" {
  recording {}
}
`
}

func testAccTwilioSIPTrunkingTrunk_disasterRecovery(method string, url string) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {
  disaster_recovery_method = "%s"
  disaster_recovery_url    = "%s"
}
`, method, url)
}

func testAccTwilioSIPTrunkingTrunk_friendlyName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {
  friendly_name = "%s"
}
`, friendlyName)
}

func testAccTwilioSIPTrunkingTrunk_secureTrue() string {
	return `
resource "twilio_sip_trunking_trunk" "trunk" {
  secure = true
}
`
}

func testAccTwilioSIPTrunkingTrunk_transferMode(transferMode string) string {
	return fmt.Sprintf(`
resource "twilio_sip_trunking_trunk" "trunk" {
  transfer_mode = "%s"
}
`, transferMode)
}

func testAccTwilioSIPTrunkingTrunk_cnamLookupEnabledTrue() string {
	return `
resource "twilio_sip_trunking_trunk" "trunk" {
  cnam_lookup_enabled = true
}
`
}
