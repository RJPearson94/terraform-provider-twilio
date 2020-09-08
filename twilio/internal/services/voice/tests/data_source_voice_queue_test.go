package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var accountQueueDataSourceName = "twilio_voice_queue"

func TestAccDataSourceTwilioAccountQueue_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.queue", accountQueueDataSourceName)
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountQueue_complete(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "max_size"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "average_wait_time"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "current_size"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccTwilioAccountQueue_complete(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_voice_queue" "queue" {
  account_sid   = "%s"
  friendly_name = "%s"
}

data "twilio_voice_queue" "queue" {
  account_sid = twilio_voice_queue.queue.account_sid
  sid         = twilio_voice_queue.queue.sid
}
`, testData.AccountSid, friendlyName)
}
