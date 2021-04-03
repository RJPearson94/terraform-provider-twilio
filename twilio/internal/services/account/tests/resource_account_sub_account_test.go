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

var subAccountResourceName = "twilio_account_sub_account"

func TestAccTwilioAccountSubAccount_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.sub_account", subAccountResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountSubAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountSubAccount_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountSubAccountExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "owner_account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "status", "active"),
					resource.TestCheckResourceAttrSet(stateResourceName, "type"),
					resource.TestCheckResourceAttrSet(stateResourceName, "auth_token"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAccountSubAccountImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAccountSubAccount_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.sub_account", subAccountResourceName)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountSubAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountSubAccount_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountSubAccountExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "friendly_name"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "owner_account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "status", "active"),
					resource.TestCheckResourceAttrSet(stateResourceName, "type"),
					resource.TestCheckResourceAttrSet(stateResourceName, "auth_token"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioAccountSubAccount_friendlyName(newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountSubAccountExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "owner_account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "status", "active"),
					resource.TestCheckResourceAttrSet(stateResourceName, "type"),
					resource.TestCheckResourceAttrSet(stateResourceName, "auth_token"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioAccountSubAccount_status(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.sub_account", subAccountResourceName)
	newStatus := "suspended"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountSubAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountSubAccount_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountSubAccountExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status", "active"),
				),
			},
			{
				Config: testAccTwilioAccountSubAccount_status(newStatus),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountSubAccountExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status", newStatus),
				),
			},
			{
				Config: testAccTwilioAccountSubAccount_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountSubAccountExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "status", "active"),
				),
			},
		},
	})
}

func TestAccTwilioAccountSubAccount_invalidStatus(t *testing.T) {
	status := "test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountSubAccount_status(status),
				ExpectError: regexp.MustCompile(`(?s)expected status to be one of \[closed suspended active\], got test`),
			},
		},
	})
}

func testAccCheckTwilioAccountSubAccountDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != subAccountResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving account information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioAccountSubAccountExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving account information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioAccountSubAccountImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAccountSubAccount_basic() string {
	return `
resource "twilio_account_sub_account" "sub_account" {}
`
}

func testAccTwilioAccountSubAccount_friendlyName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_account_sub_account" "sub_account" {
  friendly_name = "%s"
}
`, friendlyName)
}

func testAccTwilioAccountSubAccount_status(status string) string {
	return fmt.Sprintf(`
resource "twilio_account_sub_account" "sub_account" {
	status = "%s"
}
`, status)
}
