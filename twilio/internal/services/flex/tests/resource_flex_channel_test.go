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

var resourceFlexChannel = "twilio_flex_channel"

func TestAccTwilioFlexChannel_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.channel", resourceFlexChannel)

	testData := acceptance.TestAccData

	friendlyName := acctest.RandString(10)
	chatFriendlyName := acctest.RandString(10)
	chatUserFriendlyName := acctest.RandString(10)
	identity := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioFlexChannelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexChannel_basic(testData, friendlyName, chatFriendlyName, chatUserFriendlyName, identity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioFlexChannelExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "chat_friendly_name", chatFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "chat_user_friendly_name", chatUserFriendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "identity", identity),
					resource.TestCheckResourceAttrSet(stateResourceName, "flex_flow_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttr(stateResourceName, "chat_unique_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "long_lived"),
					resource.TestCheckResourceAttr(stateResourceName, "pre_engagement_data", ""),
					resource.TestCheckResourceAttr(stateResourceName, "target", ""),
					resource.TestCheckResourceAttr(stateResourceName, "task_attributes", ""),
					resource.TestCheckResourceAttr(stateResourceName, "task_sid", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "user_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func testAccCheckTwilioFlexChannelDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourceFlexChannel {
			continue
		}

		if _, err := client.Channel(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving flex channel information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioFlexChannelExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Flex

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Channel(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving flex channel information %s", err)
		}

		return nil
	}
}

func testAccTwilioFlexChannel_basic(testData *acceptance.TestData, friendlyName string, chatFriendlyName string, chatUserFriendlyName string, identity string) string {
	return fmt.Sprintf(`
resource "twilio_flex_flow" "flow" {
  friendly_name    = "%s"
  chat_service_sid = "%s"
  channel_type     = "web"
  integration_type = "external"
  integration {
    url = "https://test.com/external"
  }
}

resource "twilio_flex_channel" "channel" {
  chat_friendly_name      = "%s"
  chat_user_friendly_name = "%s"
  flex_flow_sid           = twilio_flex_flow.flow.sid
  identity                = "%s"
}
`, friendlyName, testData.FlexChannelServiceSid, chatFriendlyName, chatUserFriendlyName, identity)
}
