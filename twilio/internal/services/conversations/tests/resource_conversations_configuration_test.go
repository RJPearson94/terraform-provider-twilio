package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var configurationResourceName = "twilio_conversations_configuration"

func TestAccTwilioConversationsConfiguration_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.configuration", configurationResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConfiguration_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_closed_timer"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_inactive_timer"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsConfiguration_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.configuration", configurationResourceName)
	defaultClosedTimer := "PT10M"
	newDefaultClosedTimer := "PT15M"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsConfiguration_defaultClosedTimer(defaultClosedTimer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "default_closed_timer", defaultClosedTimer),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_inactive_timer"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsConfiguration_defaultClosedTimer(newDefaultClosedTimer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsConfigurationExists(stateResourceName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_service_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "default_closed_timer", newDefaultClosedTimer),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_inactive_timer"),
					resource.TestCheckResourceAttrSet(stateResourceName, "default_messaging_service_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioConversationsConfigurationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Configuration().Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving configuration information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsConfiguration_basic() string {
	return `
resource "twilio_conversations_configuration" "configuration" {}
`
}

func testAccTwilioConversationsConfiguration_defaultClosedTimer(defaultClosedTimer string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_configuration" "configuration" {
  default_closed_timer = "%s"
}
`, defaultClosedTimer)
}
