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

var variableResourceName = "twilio_serverless_variable"

func TestAccTwilioServerlessVariable_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.variable", variableResourceName)
	uniqueName := acctest.RandString(10)
	key := "test-key"
	value := "test-value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessVariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessVariable_basic(uniqueName, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessVariableExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "key", key),
					resource.TestCheckResourceAttr(stateResourceName, "value", value),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "environment_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioServerlessVariableImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioServerlessVariable_key(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.variable", variableResourceName)

	uniqueName := acctest.RandString(10)
	key := acctest.RandString(1)
	newKey := acctest.RandString(128)
	value := "test-value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessVariable_basic(uniqueName, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessVariableExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "key", key),
				),
			},
			{
				Config: testAccTwilioServerlessVariable_basic(uniqueName, newKey, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessVariableExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "key", newKey),
				),
			},
		},
	})
}

func TestAccTwilioServerlessVariable_invalidKeyWith0Characters(t *testing.T) {
	uniqueName := acctest.RandString(10)
	key := ""
	value := "test-value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessVariable_basic(uniqueName, key, value),
				ExpectError: regexp.MustCompile(`(?s)expected length of key to be in the range \(1 - 128\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessVariable_invalidKeyWith129Characters(t *testing.T) {
	uniqueName := acctest.RandString(10)
	key := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	value := "test-value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessVariable_basic(uniqueName, key, value),
				ExpectError: regexp.MustCompile(`(?s)expected length of key to be in the range \(1 - 128\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessVariable_value(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.variable", variableResourceName)

	uniqueName := acctest.RandString(10)
	key := "test-key"
	value := acctest.RandString(1)
	newValue := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessVariable_basic(uniqueName, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessVariableExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "value", value),
				),
			},
			{
				Config: testAccTwilioServerlessVariable_basic(uniqueName, key, newValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessVariableExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "value", newValue),
				),
			},
		},
	})
}

func TestAccTwilioServerlessFunction_invalidValue(t *testing.T) {
	uniqueName := acctest.RandString(10)
	key := "test-key"
	value := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessVariable_basic(uniqueName, key, value),
				ExpectError: regexp.MustCompile(`(?s)expected \"value\" to not be an empty string, got `),
			},
		},
	})
}

func TestAccTwilioServerlessVariable_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessVariable_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioServerlessVariable_invalidEnvironmentSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessVariable_invalidEnvironmentSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of environment_sid to match regular expression "\^ZE\[0-9a-fA-F\]\{32\}\$", got environment_sid`),
			},
		},
	})
}

func testAccCheckTwilioServerlessVariableDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != variableResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.Attributes["environment_sid"]).Variable(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving variable information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioServerlessVariableExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.Attributes["environment_sid"]).Variable(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving variable information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioServerlessVariableImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Environments/%s/Variables/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["environment_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioServerlessVariable_basic(uniqueName string, key string, value string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%s"
}

resource "twilio_serverless_variable" "variable" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  key             = "%s"
  value           = "%s"
}
`, uniqueName, uniqueName, key, value)
}

func testAccTwilioServerlessVariable_invalidServiceSid() string {
	return `
resource "twilio_serverless_variable" "variable" {
  service_sid     = "service_sid"
  environment_sid = "ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  key             = "test"
  value           = "test"
}
`
}

func testAccTwilioServerlessVariable_invalidEnvironmentSid() string {
	return `
resource "twilio_serverless_variable" "variable" {
  service_sid     = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  environment_sid = "environment_sid"
  key             = "test"
  value           = "test"
}
`
}
