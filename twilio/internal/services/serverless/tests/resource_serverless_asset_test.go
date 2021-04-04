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

var assetResourceName = "twilio_serverless_asset"

func TestAccTwilioServerlessAsset_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.asset", assetResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessAssetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessAsset_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "{}"),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/json"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "test.json"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-asset"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", visibility),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "latest_version_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:            stateResourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccTwilioServerlessAssetImportStateIdFunc(stateResourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"content", "content_file_name", "content_type", "source_hash"},
			},
		},
	})
}

func TestAccTwilioServerlessAsset_multipleAssets(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.asset", assetResourceName)
	stateResourceName2 := fmt.Sprintf("%s.asset2", assetResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessAssetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessAsset_multipleAssets(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					testAccCheckTwilioServerlessAssetExists(stateResourceName2),
				),
			},
		},
	})
}

func TestAccTwilioServerlessAssetVersion_invalidVisibility(t *testing.T) {
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	visibility := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessAsset_basic(uniqueName, friendlyName, visibility),
				ExpectError: regexp.MustCompile(`(?s)expected visibility to be one of \[public protected private\], got test`),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.asset", assetResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)
	visibility := "private"
	newVisibility := "protected"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessAssetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessAsset_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "{}"),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/json"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "test.json"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-asset"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", visibility),
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
				Config: testAccTwilioServerlessAsset_basic(uniqueName, newFriendlyName, newVisibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "content", "{}"),
					resource.TestCheckResourceAttr(stateResourceName, "content_type", "application/json"),
					resource.TestCheckResourceAttr(stateResourceName, "content_file_name", "test.json"),
					resource.TestCheckResourceAttr(stateResourceName, "path", "/test-asset"),
					resource.TestCheckResourceAttr(stateResourceName, "visibility", newVisibility),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_invalidServiceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessAsset_invalidServiceSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of service_sid to match regular expression "\^ZS\[0-9a-fA-F\]\{32\}\$", got service_sid`),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.asset", assetResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(255)
	visibility := "private"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessAssetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessAsset_basic(uniqueName, friendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioServerlessAsset_basic(uniqueName, newFriendlyName, visibility),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_invalidFriendlyNameWith0Characters(t *testing.T) {
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessAsset_friendlyNameWithStubbedServiceSid(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 255\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_invalidFriendlyNameWith256Characters(t *testing.T) {
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessAsset_friendlyNameWithStubbedServiceSid(friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 255\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_path(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.asset", assetResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(1)
	visibility := "private"
	path := "/a"
	newPath := "/" + acctest.RandString(254)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessAssetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessAsset_path(uniqueName, friendlyName, visibility, path),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "path", path),
				),
			},
			{
				Config: testAccTwilioServerlessAsset_path(uniqueName, friendlyName, visibility, newPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessAssetExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "path", newPath),
				),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_invalidPathWith0Characters(t *testing.T) {
	path := ""
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessAsset_pathWithWithStubbedServiceSid(path),
				ExpectError: regexp.MustCompile(`(?s)expected length of path to be in the range \(1 - 255\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_invalidPathWith256Characters(t *testing.T) {
	path := "/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessAsset_pathWithWithStubbedServiceSid(path),
				ExpectError: regexp.MustCompile(`(?s)expected length of path to be in the range \(1 - 255\), got /aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessAsset_invalidContentType(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessAsset_blankContentType(),
				ExpectError: regexp.MustCompile(`(?s)expected \"content_type\" to not be an empty string, got `),
			},
		},
	})
}

func testAccCheckTwilioServerlessAssetDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != assetResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Asset(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving asset information %s", err.Error())
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Asset(rs.Primary.ID).Version(rs.Primary.Attributes["latest_version_sid"]).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving asset version information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioServerlessAssetExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Asset(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving asset information %s", err.Error())
		}

		if _, err := client.Service(rs.Primary.Attributes["service_sid"]).Asset(rs.Primary.ID).Version(rs.Primary.Attributes["latest_version_sid"]).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving asset version information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioServerlessAssetImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s/Assets/%s", rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioServerlessAsset_basic(uniqueName string, friendlyName string, visibility string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%s"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "%s"
}
`, uniqueName, friendlyName, visibility)
}

func testAccTwilioServerlessAsset_multipleAssets(uniqueName string, friendlyName string, visibility string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%[1]s"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%[2]s"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "%[3]s"
}

resource "twilio_serverless_asset" "asset2" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%[2]s-2"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset-2"
  visibility        = "%[3]s"
}
`, uniqueName, friendlyName, visibility)
}

func testAccTwilioServerlessAsset_path(uniqueName string, friendlyName string, visibility string, path string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%s"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "%s"
  visibility        = "%s"
}
`, uniqueName, friendlyName, path, visibility)
}

func testAccTwilioServerlessAsset_invalidServiceSid() string {
	return `
resource "twilio_serverless_asset" "asset2" {
  service_sid       = "service_sid"
  friendly_name     = "invalid_service_sid"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/invalid-service-sid"
  visibility        = "private"
}
`
}

func testAccTwilioServerlessAsset_pathWithWithStubbedServiceSid(path string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_asset" "asset" {
  service_sid       = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name     = "invalid_path"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "%s"
  visibility        = "private"
}
`, path)
}

func testAccTwilioServerlessAsset_friendlyNameWithStubbedServiceSid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_asset" "asset" {
  service_sid       = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name     = "%s"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/friendly-ame"
  visibility        = "private"
}
`, friendlyName)
}

func testAccTwilioServerlessAsset_blankContentType() string {
	return `
resource "twilio_serverless_asset" "asset" {
  service_sid       = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  friendly_name     = "invalid_content_type"
  content           = "{}"
  content_type      = ""
  content_file_name = "test.json"
  path              = "/invalid-content-type"
  visibility        = "private"
}
`
}
