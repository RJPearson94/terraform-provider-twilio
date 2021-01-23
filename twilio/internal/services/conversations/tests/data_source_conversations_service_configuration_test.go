package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var serviceConfigurationDataSourceName = "twilio_conversations_service_configuration"

func TestAccDataSourceTwilioConversationsServiceConfiguration_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.service_configuration", serviceConfigurationDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsServiceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_chat_service_role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_conversation_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "default_conversation_role_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "reachability_enabled"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsServiceConfiguration_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

data "twilio_conversations_service_configuration" "service_configuration" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}
