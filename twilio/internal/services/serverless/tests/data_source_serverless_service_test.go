package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var serviceDataSourceName = "twilio_serverless_service"

func TestAccDataSourceTwilioServerlessService_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("%s.service", serviceDataSourceName)
	uniqueName := acctest.RandString(10)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioServerlessService_basic(uniqueName, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioServerlessServiceExists(stateDataSourceName),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "include_credentials"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "ui_editable"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioServerlessService_basic(uniqueName string, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_serverless_service" "service" {
  unique_name   = "%s"
  friendly_name = "%s"
}

data "twilio_serverless_service" "service" {
  sid = twilio_serverless_service.service.sid
}
`, uniqueName, friendlyName)
}
