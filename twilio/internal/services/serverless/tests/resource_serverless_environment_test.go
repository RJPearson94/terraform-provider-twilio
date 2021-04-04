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

var environmentResourceName = "twilio_serverless_environment"

func TestAccTwilioServerlessEnvironment_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.environment", environmentResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessEnvironment_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessEnvironmentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "domain_suffix", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "build_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioServerlessEnvironmentImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioServerlessEnvironment_domainSuffix(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.environment", environmentResourceName)
	uniqueName := acctest.RandString(10)
	domainSuffix := acctest.RandString(1)
	newDomainSuffix := acctest.RandString(16)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessEnvironment_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessEnvironmentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "domain_suffix", ""),
				),
			},
			{
				Config: testAccTwilioServerlessEnvironment_domainSuffix(uniqueName, domainSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessEnvironmentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "domain_suffix", domainSuffix),
				),
			},
			{
				Config: testAccTwilioServerlessEnvironment_domainSuffix(uniqueName, newDomainSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessEnvironmentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "domain_suffix", newDomainSuffix),
				),
			},
			{
				Config: testAccTwilioServerlessEnvironment_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessEnvironmentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "domain_suffix", ""),
				),
			},
		},
	})
}

func TestAccTwilioServerlessEnvironment_invalidDomainSuffixWith0Characters(t *testing.T) {
	domainSuffix := ""
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessEnvironment_domainSuffixWithStubbedServiceSid(domainSuffix),
				ExpectError: regexp.MustCompile(`(?s)expected length of domain_suffix to be in the range \(1 - 16\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessEnvironment_invalidDomainSuffixNameWith17Characters(t *testing.T) {
	domainSuffix := "aaaaaaaaaaaaaaaaa"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessEnvironment_domainSuffixWithStubbedServiceSid(domainSuffix),
				ExpectError: regexp.MustCompile(`(?s)expected length of domain_suffix to be in the range \(1 - 16\), got aaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessEnvironment_uniqueName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.environment", environmentResourceName)
	uniqueName := acctest.RandString(1)
	newUniqueName := acctest.RandString(50)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessEnvironment_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessEnvironmentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
				),
			},
			{
				Config: testAccTwilioServerlessEnvironment_basic(newUniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessEnvironmentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
				),
			},
		},
	})
}

func TestAccTwilioServerlessEnvironment_invalidUniqueNameWith0Characters(t *testing.T) {
	uniqueName := ""
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessEnvironment_uniqueNameWithStubbedServiceSid(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 100\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessEnvironment_invalidUniqueNameWith101Characters(t *testing.T) {
	uniqueName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessEnvironment_uniqueNameWithStubbedServiceSid(uniqueName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 100\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessEnvironment_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessEnvironment_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func testAccCheckTwilioServerlessEnvironmentDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != environmentResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving environment information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioServerlessEnvironmentExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving environment information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioServerlessEnvironmentImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Environments/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioServerlessEnvironment_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%[1]s"
}
`, uniqueName)
}

func testAccTwilioServerlessEnvironment_domainSuffix(uniqueName string, domainSuffix string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid   = twilio_serverless_service.service.sid
  unique_name   = "%[1]s"
  domain_suffix = "%[2]s"
}
`, uniqueName, domainSuffix)
}

func testAccTwilioServerlessEnvironment_uniqueNameWithStubbedServiceSid(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_environment" "environment" {
  service_sid = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name = "%s"
}
`, uniqueName)
}

func testAccTwilioServerlessEnvironment_domainSuffixWithStubbedServiceSid(domainSuffix string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_environment" "environment" {
  service_sid   = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  unique_name   = "test"
  domain_suffix = "%s"
}
`, domainSuffix)
}

func testAccTwilioServerlessEnvironment_invalidServiceSid() string {
	return `
resource "twilio_serverless_environment" "environment" {
  service_sid = "service_sid"
  unique_name = "invalid_service_sid"
}
`
}
