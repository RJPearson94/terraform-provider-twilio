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

var environmentResourceName = "twilio_serverless_environment"

func TestAccTwilioServerlessEnvironment_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.environment", environmentResourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioServerlessEnvironmentDestroy,
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
			return fmt.Errorf("Error occurred when retrieving environment information %s", err)
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
			return fmt.Errorf("Error occurred when retrieving environment information %s", err)
		}

		return nil
	}
}

func testAccTwilioServerlessEnvironment_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
	unique_name   = "service-%s"
	friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
	service_sid   = twilio_serverless_service.service.sid
	unique_name = "%s"
}`, uniqueName, uniqueName)
}
