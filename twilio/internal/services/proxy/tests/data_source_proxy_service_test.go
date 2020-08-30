package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var proxyServiceDataSourceName = "twilio_proxy_service"

func TestAccDataSourceTwilioProxyService_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service", proxyServiceDataSourceName)
	uniqueName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioProxyService_basic(uniqueName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "unique_name", uniqueName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "chat_instance_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "chat_service_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "default_ttl", "0"),
					resource.TestCheckResourceAttr(stateDataSourceName, "callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "geo_match_level", "country"),
					resource.TestCheckResourceAttr(stateDataSourceName, "number_selection_behavior", "prefer-sticky"),
					resource.TestCheckResourceAttr(stateDataSourceName, "intercept_callback_url", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "out_of_session_callback_url", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioProxyService_basic(uniqueName string) string {
	return fmt.Sprintf(`
resource "twilio_proxy_service" "service" {
  unique_name = "%s"
}

data "twilio_proxy_service" "service" {
  sid = twilio_proxy_service.service.sid
}
`, uniqueName)
}
