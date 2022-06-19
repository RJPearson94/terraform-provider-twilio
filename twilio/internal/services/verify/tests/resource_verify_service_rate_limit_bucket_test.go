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

var rateLimitBucketResourceName = "twilio_verify_service_rate_limit_bucket"

func TestAccTwilioVerifyRateLimitBucket_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.rate_limit_bucket", rateLimitBucketResourceName)
	uniqueName := acctest.RandString(10)
	max := 10
	interval := 2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyRateLimitBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyRateLimitBucket_basic(uniqueName, max, interval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyRateLimitBucketExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max", "10"),
					resource.TestCheckResourceAttr(stateResourceName, "interval", "2"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "rate_limit_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioVerifyRateLimitBucketImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioVerifyRateLimitBucket_max(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.rate_limit_bucket", rateLimitBucketResourceName)
	uniqueName := acctest.RandString(10)
	max := 10
	newMax := 5
	interval := 2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyRateLimitBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyRateLimitBucket_basic(uniqueName, max, interval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyRateLimitBucketExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max", "10"),
				),
			},
			{
				Config: testAccTwilioVerifyRateLimitBucket_basic(uniqueName, newMax, interval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyRateLimitBucketExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max", "5"),
				),
			},
		},
	})
}

func TestAccTwilioVerifyRateLimitBucket_interval(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.rate_limit_bucket", rateLimitBucketResourceName)
	uniqueName := acctest.RandString(10)
	max := 10
	interval := 2
	newInterval := 3

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioVerifyRateLimitBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioVerifyRateLimitBucket_basic(uniqueName, max, interval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyRateLimitBucketExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "interval", "2"),
				),
			},
			{
				Config: testAccTwilioVerifyRateLimitBucket_basic(uniqueName, max, newInterval),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioVerifyRateLimitBucketExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "interval", "3"),
				),
			},
		},
	})
}

func TestAccTwilioVerifyRateLimitBucket_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyRateLimitBucket_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^VA\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioVerifyRateLimitBucket_invalidRateLimitSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioVerifyRateLimitBucket_invalidRateLimitSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of rate_limit_sid to match regular expression "\^RK\[0-9a-fA-F\]\{32\}\$", got rate_limit_sid`),
			},
		},
	})
}

func testAccCheckTwilioVerifyRateLimitBucketDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

	for _, rs := range s.RootModule().Resources {
		if rs.Type != rateLimitBucketResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).RateLimit(rs.Primary.Attributes["rate_limit_sid"]).Bucket(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving rate limit bucket information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioVerifyRateLimitBucketExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Verify

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).RateLimit(rs.Primary.Attributes["rate_limit_sid"]).Bucket(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving rate limit bucket information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioVerifyRateLimitBucketImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/RateLimits/%s/Buckets/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["rate_limit_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioVerifyRateLimitBucket_basic(uniqueName string, max int, interval int) string {
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
`, uniqueName, max, interval)
}

func testAccTwilioVerifyRateLimitBucket_invalidServiceSid() string {
	return `
resource "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  service_sid    = "service_sid"
  rate_limit_sid = "RKaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  max            = 10
  interval       = 2
}
`
}

func testAccTwilioVerifyRateLimitBucket_invalidRateLimitSid() string {
	return `
resource "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  service_sid    = "VAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  rate_limit_sid = "rate_limit_sid"
  max            = 10
  interval       = 2
}
`
}
