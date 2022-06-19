package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const rateLimitBucketsDataSourceName = "twilio_verify_service_rate_limit_buckets"

func TestAccDataSourceTwilioVerifyRateLimitBuckets_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.rate_limit_buckets", rateLimitBucketsDataSourceName)
	uniqueName := acctest.RandString(10)
	max := 10
	interval := 2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyRateLimitBuckets_basic(uniqueName, max, interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limit_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "rate_limit_buckets.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limit_buckets.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "rate_limit_buckets.0.max", "10"),
					resource.TestCheckResourceAttr(stateDataSourceName, "rate_limit_buckets.0.interval", "2"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limit_buckets.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limit_buckets.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limit_buckets.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyRateLimitBuckets_invalidRateLimitSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyRateLimitBuckets_invalidRateLimitSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of rate_limit_sid to match regular expression "\^RK\[0-9a-fA-F\]\{32\}\$", got rate_limit_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyRateLimitBuckets_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyRateLimitBuckets_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyRateLimitBuckets_basic(uniqueName string, max int, interval int) string {
	return fmt.Sprintf(`
resource "twilio_verify_service" "service" {
  friendly_name = "%[1]s"
}

resource "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = twilio_verify_service.service.sid
  unique_name = "%[1]s"
}

resource "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  service_sid    = twilio_verify_service_rate_limit.rate_limit.service_sid
  rate_limit_sid = twilio_verify_service_rate_limit.rate_limit.sid
  max            = %[2]d
  interval       = %[3]d
}

data "twilio_verify_service_rate_limit_buckets" "rate_limit_buckets" {
  service_sid    = twilio_verify_service_rate_limit_bucket.rate_limit_bucket.service_sid
  rate_limit_sid = twilio_verify_service_rate_limit_bucket.rate_limit_bucket.rate_limit_sid
}
`, uniqueName, max, interval)
}

func testAccDataSourceTwilioVerifyRateLimitBuckets_invalidRateLimitSid() string {
	return `
data "twilio_verify_service_rate_limit_buckets" "rate_limit_buckets" {
  service_sid    = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  rate_limit_sid = "rate_limit_sid"
}
`
}

func testAccDataSourceTwilioVerifyRateLimitBuckets_invalidServiceSid() string {
	return `
data "twilio_verify_service_rate_limit_buckets" "rate_limit_buckets" {
  service_sid    = "service_sid"
  rate_limit_sid = "RKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}
