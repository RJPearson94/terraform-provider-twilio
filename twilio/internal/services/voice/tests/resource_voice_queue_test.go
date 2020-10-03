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

var queueResourceName = "twilio_voice_queue"

func TestAccTwilioAccountQueue_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.queue", queueResourceName)
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		Providers:         acceptance.TestAccProviders,
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAccountQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountQueue_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "max_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "average_wait_time"),
					resource.TestCheckResourceAttrSet(stateResourceName, "current_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAccountQueueImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAccountQueue_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.queue", queueResourceName)
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckTwilioAccountQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountQueue_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "max_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "average_wait_time"),
					resource.TestCheckResourceAttrSet(stateResourceName, "current_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioAccountQueue_basic(testData, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "max_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "average_wait_time"),
					resource.TestCheckResourceAttrSet(stateResourceName, "current_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccCheckTwilioAccountQueueDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != queueResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Queue(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving queue information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioAccountQueueExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Queue(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving queue information %s", err)
		}

		return nil
	}
}

func testAccTwilioAccountQueueImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/Queues/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAccountQueue_basic(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_voice_queue" "queue" {
  account_sid   = "%s"
  friendly_name = "%s"
}
`, testData.AccountSid, friendlyName)
}
