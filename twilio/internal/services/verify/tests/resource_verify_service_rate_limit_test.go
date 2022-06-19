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

var rateLimitResourceName = "twilio_verify_service_rate_limit"

func TestAccTwilioVerifyRateLimit_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.rate_limit", rateLimitResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyRateLimit_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyRateLimitExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "description", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioVerifyRateLimitImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioVerifyRateLimit_blankUniqueName(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyRateLimit_blankUniqueName(),
				ExpectError: regexp.MustCompile(`(?s)expected \"unique_name\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioVerifyRateLimit_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyRateLimit_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccCheckTwilioVerifyRateLimitDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

	for _, rs := range s.RootModule().Resources {
		if rs.Type != rateLimitResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).RateLimit(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving rate limit information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioVerifyRateLimitExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).RateLimit(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving rate limit information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioVerifyRateLimitImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/RateLimits/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioVerifyRateLimit_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = twilio_verify_service.service.sid
  unique_name = "%[1]s"
}
`, uniqueName)
}

func testAccTwilioVerifyRateLimit_blankUniqueName() string {
	return `
resource "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name = ""
}
`
}

func testAccTwilioVerifyRateLimit_invalidServiceSid() string {
	return `
resource "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = "service_sid"
  unique_name = "invalid service sid"
}
`
}
