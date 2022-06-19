package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const rateLimitBucketDataSourceName = "twilio_verify_service_rate_limit_bucket"

func TestAccDataSourceTwilioVerifyRateLimitBucket_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.rate_limit_bucket", rateLimitBucketDataSourceName)
	uniqueName := acctest.RandString(10)
	max := 10
	interval := 2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioVerifyRateLimitBucket_basic(uniqueName, max, interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "max", "10"),
					resource.TestCheckResourceAttr(stateDataSourceName, "interval", "2"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "rate_limit_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyRateLimitBucket_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyRateLimitBucket_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^BL\[0-9a-fA-F\]\{32\}\$", got rate_limit_bucket_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyRateLimitBucket_invalidRateLimitSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyRateLimitBucket_invalidRateLimitSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of rate_limit_sid to match regular expression "\^RK\[0-9a-fA-F\]\{32\}\$", got rate_limit_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioVerifyRateLimitBucket_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioVerifyRateLimitBucket_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioVerifyRateLimitBucket_basic(uniqueName string, max int, interval int) string {
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

data "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  sid            = twilio_verify_service_rate_limit_bucket.rate_limit_bucket.sid
  service_sid    = twilio_verify_service_rate_limit_bucket.rate_limit_bucket.service_sid
  rate_limit_sid = twilio_verify_service_rate_limit_bucket.rate_limit_bucket.rate_limit_sid
}
`, uniqueName, max, interval)
}

func testAccDataSourceTwilioVerifyRateLimitBucket_invalidSid() string {
	return `
data "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  sid            = "rate_limit_bucket_sid"
  service_sid    = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  rate_limit_sid = "RKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioVerifyRateLimitBucket_invalidRateLimitSid() string {
	return `
data "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  sid            = "BLaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  service_sid    = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  rate_limit_sid = "rate_limit_sid"
}
`
}

func testAccDataSourceTwilioVerifyRateLimitBucket_invalidServiceSid() string {
	return `
data "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  sid            = "RKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  service_sid    = "service_sid"
  rate_limit_sid = "RKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}
