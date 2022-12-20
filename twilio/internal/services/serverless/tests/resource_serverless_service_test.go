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

var serviceResourceName = "twilio_serverless_service"

func TestAccTwilioServerlessService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "include_credentials", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "ui_editable", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "domain_base"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioServerlessServiceImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioServerlessService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "include_credentials", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "ui_editable", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "domain_base"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttr(stateResourceName, "include_credentials", "true"),
					resource.TestCheckResourceAttr(stateResourceName, "ui_editable", "false"),
					resource.TestCheckResourceAttrSet(stateResourceName, "domain_base"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioServerlessService_uiEditable(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "ui_editable", "false"),
				),
			},
			{
				Config: testAccTwilioServerlessService_uiEditableTrue(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "ui_editable", "true"),
				),
			},
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "ui_editable", "false"),
				),
			},
		},
	})
}

func TestAccTwilioServerlessService_includeCredentials(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "include_credentials", "true"),
				),
			},
			{
				Config: testAccTwilioServerlessService_includeCredentialsFalse(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "include_credentials", "false"),
				),
			},
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "include_credentials", "true"),
				),
			},
		},
	})
}

func TestAccTwilioServerlessService_uniqueName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	uniqueName := acctest.RandString(1)
	newUniqueName := acctest.RandString(50)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
				),
			},
			{
				Config: testAccTwilioServerlessService_basic(newUniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", newUniqueName),
				),
			},
		},
	})
}

func TestAccTwilioServerlessService_invalidUniqueNameWith0Characters(t *testing.T) {
	uniqueName := ""
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 50\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessService_invalidUniqueNameWith51Characters(t *testing.T) {
	uniqueName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of unique_name to be in the range \(1 - 50\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioServerlessService_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(255)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
		},
	})
}

func TestAccTwilioServerlessService_invalidFriendlyNameWith0Characters(t *testing.T) {
	uniqueName := acctest.RandString(10)
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 255\), got `),
			},
		},
	})
}

func TestAccTwilioServerlessService_invalidFriendlyNameWith256Characters(t *testing.T) {
	uniqueName := acctest.RandString(10)
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 255\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func testAccCheckTwilioServerlessServiceDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

	for _, rs := range s.RootModule().Resources {
		if rs.Type != serviceResourceName {
			continue
		}

		if _, err := client.Service(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving service information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioServerlessServiceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Serverless

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving service information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioServerlessServiceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Services/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioServerlessService_basic(uniqueName string, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%s"
  friendly_name = "%s"
}
`, uniqueName, friendlyName)
}

func testAccTwilioServerlessService_uiEditableTrue(uniqueName string, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%s"
  friendly_name = "%s"
  ui_editable   = true
}
`, uniqueName, friendlyName)
}

func testAccTwilioServerlessService_includeCredentialsFalse(uniqueName string, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name         = "%s"
  friendly_name       = "%s"
  include_credentials = false
}
`, uniqueName, friendlyName)
}
