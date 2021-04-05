package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var configurationDataSourceName = "twilio_conversations_configuration"

func TestAccDataSourceTwilioConversationsConfiguration_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.configuration", configurationDataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsConfiguration_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_closed_timer"),
					resource.TestCheckResourceAttr(stateDataSourceName, "default_inactive_timer", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsConfiguration_basic() string {
	return `
data "twilio_conversations_configuration" "configuration" {}
`
}
