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

var buildResourceName = "twilio_serverless_build"

func TestAccTwilioServerlessBuild_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.build", buildResourceName)
	assetResourceName := "twilio_serverless_asset.asset"
	functionResourceName := "twilio_serverless_function.function"
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessBuildDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessBuild_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "asset_version.#", "1"),
					resource.TestCheckResourceAttrPair(stateResourceName, "asset_version.0.sid", assetResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "function_version.#", "1"),
					resource.TestCheckResourceAttrPair(stateResourceName, "function_version.0.sid", functionResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "dependencies.%"),
					resource.TestCheckResourceAttrSet(stateResourceName, "runtime"),
					resource.TestCheckResourceAttrSet(stateResourceName, "status"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:            stateResourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccTwilioServerlessBuildImportStateIdFunc(stateResourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"polling"},
			},
		},
	})
}

func TestAccTwilioServerlessBuild_functions(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.build", buildResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessBuildDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessBuild_multipleFunctions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "function_version.#", "4"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.0.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.1.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.2.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.3.sid"),
				),
			},
			{
				Config: testAccTwilioServerlessBuild_oneFunction(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "function_version.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.0.sid"),
				),
			},
			{
				Config: testAccTwilioServerlessBuild_multipleFunctions(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "function_version.#", "4"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.0.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.1.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.2.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "function_version.3.sid"),
				),
			},
		},
	})
}

func TestAccTwilioServerlessBuild_assets(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.build", buildResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessBuildDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessBuild_multipleAssets(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "asset_version.#", "4"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.0.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.1.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.2.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.3.sid"),
				),
			},
			{
				Config: testAccTwilioServerlessBuild_oneAsset(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "asset_version.#", "1"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.0.sid"),
				),
			},
			{
				Config: testAccTwilioServerlessBuild_multipleAssets(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "asset_version.#", "4"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.0.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.1.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.2.sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "asset_version.3.sid"),
				),
			},
		},
	})
}

func TestAccTwilioServerlessBuild_dependenciesAndRuntime(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.build", buildResourceName)
	uniqueName := acctest.RandString(10)
	version := "3.6.2"
	runtime := "node16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessBuildDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessBuild_dependenciesAndRuntime(uniqueName, version, runtime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessBuildExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "dependencies.%", "6"),
					resource.TestCheckResourceAttr(stateResourceName, "dependencies.twilio", version),
					resource.TestCheckResourceAttr(stateResourceName, "runtime", runtime),
				),
			},
		},
	})
}

func TestAccTwilioServerlessBuild_invalidRuntime(t *testing.T) {
	uniqueName := acctest.RandString(10)
	version := "3.6.2"
	runtime := "python2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessBuild_dependenciesAndRuntime(uniqueName, version, runtime),
				ExpectError: regexp.MustCompile(`(?s)expected runtime to be one of \["node14" "node16"\], got python2`),
			},
		},
	})
}

func TestAccTwilioServerlessBuild_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessBuild_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioServerlessBuild_invalidFunctionVersionSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessBuild_invalidFunctionVersionSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of function_version.0.sid to match regular expression "\^ZN\[0-9a-fA-F\]\{32\}\$", got function_version_sid`),
			},
		},
	})
}

func TestAccTwilioServerlessBuild_invalidAssetVersionSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessBuild_invalidAssetVersionSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of asset_version.0.sid to match regular expression "\^ZN\[0-9a-fA-F\]\{32\}\$", got asset_version_sid`),
			},
		},
	})
}

func testAccCheckTwilioServerlessBuildDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != buildResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Build(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving build information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioServerlessBuildExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Build(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving build information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioServerlessBuildImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Builds/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioServerlessBuild_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
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

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }
  asset_version {
    sid = twilio_serverless_asset.asset.latest_version_sid
  }
  polling {
    enabled = true
  }
}
`, uniqueName)
}

func testAccTwilioServerlessBuild_oneFunction(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
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
`, uniqueName)
}

func testAccTwilioServerlessBuild_multipleFunctions(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
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

resource "twilio_serverless_function" "function2" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = <<EOF
  exports.handler = function (context, event, callback) {
	  callback(null, "Hello World 2");
  };
  EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld2.js"
  path              = "/test-function2"
  visibility        = "private"
}

resource "twilio_serverless_function" "function3" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = <<EOF
  exports.handler = function (context, event, callback) {
	  callback(null, "Hello World 3");
  };
  EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld3.js"
  path              = "/test-function3"
  visibility        = "private"
}

resource "twilio_serverless_function" "function4" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = <<EOF
  exports.handler = function (context, event, callback) {
	  callback(null, "Hello World 3");
  };
  EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld4.js"
  path              = "/test-function4"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }
  function_version {
    sid = twilio_serverless_function.function2.latest_version_sid
  }
  function_version {
    sid = twilio_serverless_function.function3.latest_version_sid
  }
  function_version {
    sid = twilio_serverless_function.function4.latest_version_sid
  }
  polling {
    enabled = true
  }
}
`, uniqueName)
}

func testAccTwilioServerlessBuild_oneAsset(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  asset_version {
    sid = twilio_serverless_asset.asset.latest_version_sid
  }
  polling {
    enabled = true
  }
}
`, uniqueName)
}

func testAccTwilioServerlessBuild_multipleAssets(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "private"
}

resource "twilio_serverless_asset" "asset2" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test2"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset2"
  visibility        = "private"
}

resource "twilio_serverless_asset" "asset3" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test3"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset3"
  visibility        = "private"
}

resource "twilio_serverless_asset" "asset4" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test4"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset4"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  asset_version {
    sid = twilio_serverless_asset.asset.latest_version_sid
  }
  asset_version {
    sid = twilio_serverless_asset.asset2.latest_version_sid
  }
  asset_version {
    sid = twilio_serverless_asset.asset3.latest_version_sid
  }
  asset_version {
    sid = twilio_serverless_asset.asset4.latest_version_sid
  }
  polling {
    enabled = true
  }
}
`, uniqueName)
}

func testAccTwilioServerlessBuild_dependenciesAndRuntime(uniqueName string, twilioVersion string, runtime string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "service-%s"
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

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid
  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }
  asset_version {
    sid = twilio_serverless_asset.asset.latest_version_sid
  }
  dependencies = {
    "twilio"                  = "%s",
    "fs"                      = "0.0.1-security"
    "lodash"                  = "4.17.11"
    "util"                    = "0.11.0"
    "xmldom"                  = "0.1.27"
    "@twilio/runtime-handler" = "1.0.1"
  }
  runtime = "%s"
  polling {
    enabled = true
  }
}
`, uniqueName, twilioVersion, runtime)
}

func testAccTwilioServerlessBuild_invalidServiceSid() string {
	return `
resource "twilio_serverless_build" "build" {
  service_sid = "service_sid"
  function_version {
    sid = "ZHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  }
  dependencies = {
    "twilio"                  = "3.6.2",
    "fs"                      = "0.0.1-security"
    "lodash"                  = "4.17.11"
    "util"                    = "0.11.0"
    "xmldom"                  = "0.1.27"
    "@twilio/runtime-handler" = "1.0.1"
  }
  runtime = "node16"
}
`
}

func testAccTwilioServerlessBuild_invalidFunctionVersionSid() string {
	return `
resource "twilio_serverless_build" "build" {
  service_sid = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  function_version {
    sid = "function_version_sid"
  }
  dependencies = {
    "twilio"                  = "3.6.2",
    "fs"                      = "0.0.1-security"
    "lodash"                  = "4.17.11"
    "util"                    = "0.11.0"
    "xmldom"                  = "0.1.27"
    "@twilio/runtime-handler" = "1.0.1"
  }
  runtime = "node16"
}
`
}

func testAccTwilioServerlessBuild_invalidAssetVersionSid() string {
	return `
resource "twilio_serverless_build" "build" {
  service_sid = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  asset_version {
    sid = "asset_version_sid"
  }
  dependencies = {
    "twilio"                  = "3.6.2",
    "fs"                      = "0.0.1-security"
    "lodash"                  = "4.17.11"
    "util"                    = "0.11.0"
    "xmldom"                  = "0.1.27"
    "@twilio/runtime-handler" = "1.0.1"
  }
  runtime = "node16"
}
`
}
