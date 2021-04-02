package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const accountQueueDataSourceName = "twilio_voice_queue"

func TestAccDataSourceTwilioAccountQueue_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.queue", accountQueueDataSourceName)
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountQueue_complete(testData, friendlyName),
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

func TestAccDataSourceTwilioAccountQueue_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountQueue_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountQueue_invalidSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountQueue_invalidSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of sid to match regular expression "\^QU\[0-9a-fA-F\]\{32\}\$", got sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAccountQueue_complete(testData *acceptance.TestData, friendlyName string) string {
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

func testAccDataSourceTwilioAccountQueue_invalidAccountSid() string {
	return `
data "twilio_voice_queue" "queue" {
  account_sid = "account_sid"
  sid         = "QUaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioAccountQueue_invalidSid() string {
	return `
data "twilio_voice_queue" "queue" {
  account_sid = "ACaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  sid         = "sid"
}
`
}
