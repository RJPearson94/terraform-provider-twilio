package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var serviceResourceName = "twilio_serverless_service"

func TestAccTwilioServerlessService_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "include_credentials"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ui_editable"),
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

func TestAccTwilioServerlessService_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service", serviceResourceName)

	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateResourceName, "include_credentials"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ui_editable"),
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
					resource.TestCheckResourceAttrSet(stateResourceName, "include_credentials"),
					resource.TestCheckResourceAttrSet(stateResourceName, "ui_editable"),
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
			return fmt.Errorf("Error occurred when retrieving service information %s", err)
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
			return fmt.Errorf("Error occurred when retrieving service information %s", err)
		}

		return nil
	}
}

func testAccTwilioServerlessService_basic(uniqueName string, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
	unique_name   = "%s"
	friendly_name = "%s"
}`, uniqueName, friendlyName)
}
