package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var serviceConfigurationResourceName = "twilio_conversations_service_configuration"

func TestAccTwilioConversationsServiceConfiguration_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.service_configuration", serviceConfigurationResourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsServiceConfiguration_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsServiceConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_chat_service_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_conversation_creator_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_conversation_role_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "reachability_enabled"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioConversationsServiceConfigurationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Service(rs.Primary.ID).Configuration().Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving service configuration information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsServiceConfiguration_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_service" "service" {
  friendly_name = "%s"
}

resource "twilio_conversations_service_configuration" "service_configuration" {
  service_sid = twilio_conversations_service.service.sid
}
`, friendlyName)
}
