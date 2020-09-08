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
  unique_name   = "service-%s"
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
  unique_name   = "service-%s"
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

resource "twilio_serverless_asset" "asset2" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "%s-2"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset-2"
  visibility        = "%s"
}
`, uniqueName, friendlyName, visibility, friendlyName, visibility)
}
