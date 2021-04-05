package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var webhookDataSourceName = "twilio_conversations_webhook"

func TestAccDataSourceTwilioConversationsWebhook_basic(t *testing.T) {
	stateDataSource := fmt.Sprintf("data.%s.webhook", webhookDataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioConversationsWebhook_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSource, "id"),
					resource.TestCheckResourceAttrSet(stateDataSource, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSource, "target"),
					resource.TestCheckResourceAttrSet(stateDataSource, "method"),
					resource.TestCheckResourceAttrSet(stateDataSource, "pre_webhook_url"),
					resource.TestCheckResourceAttrSet(stateDataSource, "post_webhook_url"),
					resource.TestCheckResourceAttrSet(stateDataSource, "filters.#"),
					resource.TestCheckResourceAttrSet(stateDataSource, "url"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioConversationsWebhook_basic() string {
	return `
data "twilio_conversations_webhook" "webhook" {}
`
}
