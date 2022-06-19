package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const rateLimitsDataSourceName = "twilio_verify_service_rate_limits"

func TestAccDataSourceTwilioVerifyRateLimits_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.rate_limits", rateLimitsDataSourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyRateLimits_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "rate_limits.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limits.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "rate_limits.0.unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateDataSourceName, "rate_limits.0.description", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limits.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limits.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limits.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyRateLimits_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyRateLimits_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyRateLimits_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = twilio_verify_service.service.sid
  unique_name = "%[1]s"
}

data "twilio_verify_service_rate_limits" "rate_limits" {
  service_sid = twilio_verify_service_rate_limit.rate_limit.service_sid
}
`, uniqueName)
}

func testAccDataSourceTwilioVerifyRateLimits_invalidServiceSid() string {
	return `
data "twilio_verify_service_rate_limits" "rate_limits" {
  service_sid = "service_sid"
}
`
}
