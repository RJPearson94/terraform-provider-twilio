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

var deploymentResourceName = "twilio_serverless_deployment"

func TestAccTwilioServerlessDeployment_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.deployment", deploymentResourceName)
	uniqueName := acctest.RandString(10)

	// Run tests in parallel as I got rate limited when they ran in parallel
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessDeployment_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessDeploymentExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "environment_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioServerlessDeploymentImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioServerlessDeployment_createBeforeDestroy(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.deployment", deploymentResourceName)
	uniqueName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessDeployment_createBeforeDestroy(uniqueName, "Hello World"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessDeploymentExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "environment_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_latest_deployment", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioServerlessDeployment_createBeforeDestroy(uniqueName, "New Response"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessDeploymentExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "environment_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "is_latest_deployment"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				// apply to refresh state to verify the deployment is still latest
				Config: testAccTwilioServerlessDeployment_createBeforeDestroy(uniqueName, "New Response"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessDeploymentExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "is_latest_deployment", "true"),
				),
			},
		},
	})
}

func TestAccTwilioServerlessDeployment_removeBuildAndDeployment(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.deployment", deploymentResourceName)
	uniqueName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessDeploymentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessDeployment_createBeforeDestroy(uniqueName, "Hello World"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessDeploymentExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "environment_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "is_latest_deployment", "true"),
					resource.TestCheckResourceAttrSet(stateResourceName, "build_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckNoResourceAttr(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioServerlessDeployment_removeBuildAndDeployment(uniqueName, "Hello World"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildIsDestroyed,
				),
			},
		},
	})
}

func TestAccTwilioServerlessDeployment_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessDeployment_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioServerlessDeployment_invalidEnvironmentSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessDeployment_invalidEnvironmentSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of environment_sid to match regular expression "\^ZE\[0-9a-fA-F\]\{32\}\$", got environment_sid`),
			},
		},
	})
}

func TestAccTwilioServerlessDeployment_invalidBuildSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessDeployment_invalidBuildSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of build_sid to match regular expression "\^ZB\[0-9a-fA-F\]\{32\}\$", got build_sid`),
			},
		},
	})
}

func testAccCheckTwilioServerlessDeploymentDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != deploymentResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.Attributes["environment_sid"]).Deployment(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving deployment information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioServerlessBuildIsDestroyed(s *terraform.State) error {
	rs, ok := s.RootModule().Resources["twilio_serverless_build.build"]
	if !ok {
		return nil
	}

	return fmt.Errorf("Build resource with sid (%s) still exists after a destroy. This should have been deleted", rs.Primary.ID)
}

func testAccCheckTwilioServerlessDeploymentExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Environment(rs.Primary.Attributes["environment_sid"]).Deployment(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving deployment information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioServerlessDeploymentImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Environments/%s/Deployments/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["environment_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioServerlessDeployment_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = <<EOF
exports.handler = function (context, event, callback) {
	callback(null, "Hello World");
};
EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }
  polling {
    enabled = true
  }
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%[1]s"
}

resource "twilio_serverless_deployment" "deployment" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  build_sid       = twilio_serverless_build.build.sid
}
`, uniqueName)
}

func testAccTwilioServerlessDeployment_createBeforeDestroy(uniqueName string, greetingMessage string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = <<EOF
exports.handler = function (context, event, callback) {
	callback(null, "%[2]s");
};
EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }
  polling {
    enabled = true
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%[1]s"
}

resource "twilio_serverless_deployment" "deployment" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  build_sid       = twilio_serverless_build.build.sid

  lifecycle {
    create_before_destroy = true
  }
}
`, uniqueName, greetingMessage)
}

func testAccTwilioServerlessDeployment_removeBuildAndDeployment(uniqueName string, greetingMessage string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = <<EOF
exports.handler = function (context, event, callback) {
	callback(null, "%[2]s");
};
EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "private"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "%[1]s"
}
`, uniqueName, greetingMessage)
}

func testAccTwilioServerlessDeployment_invalidServiceSid() string {
	return `
resource "twilio_serverless_deployment" "deployment" {
  service_sid     = "service_sid"
  environment_sid = "ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  build_sid       = "ZBaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioServerlessDeployment_invalidEnvironmentSid() string {
	return `
resource "twilio_serverless_deployment" "deployment" {
  service_sid     = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  environment_sid = "environment_sid"
  build_sid       = "ZBaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccTwilioServerlessDeployment_invalidBuildSid() string {
	return `
resource "twilio_serverless_deployment" "deployment" {
  service_sid     = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  environment_sid = "ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  build_sid       = "build_sid"
}
`
}
