package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const accountQueuesDataSourceName = "twilio_voice_queues"

func TestAccDataSourceTwilioAccountQueues_complete(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.queues", accountQueuesDataSourceName)
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioAccountQueues_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(stateDataSourceName, "queues.#", regexp.MustCompile("^[1-9]([0-9]+)?$")),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioAccountQueues_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioAccountQueues_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioAccountQueues_basic(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_voice_queue" "queue" {
  account_sid   = "%s"
  friendly_name = "%s"
}

data "twilio_voice_queues" "queues" {
  account_sid = twilio_voice_queue.queue.account_sid
}
`, testData.AccountSid, friendlyName)
}

func testAccDataSourceTwilioAccountQueues_invalidAccountSid() string {
	return `
data "twilio_voice_queues" "queues" {
  account_sid = "account_sid"
}
`
}
