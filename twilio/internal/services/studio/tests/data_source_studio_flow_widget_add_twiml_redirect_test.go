package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetAddTwiMLRedirect_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_add_twiml_redirect.add_twiml_redirect"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetAddTwiMLRedirect_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"AddTwiMLRedirect","properties":{"url":"https://test.com/twiml"},"transitions":[{"event":"fail"},{"event":"return"},{"event":"timeout"}],"type":"add-twiml-redirect"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetAddTwiMLRedirect_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_add_twiml_redirect.add_twiml_redirect"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetAddTwiMLRedirect_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"AddTwiMLRedirect","properties":{"method":"POST","offset":{"x":10,"y":20},"timeout":"100","url":"https://test.com/twiml"},"transitions":[{"event":"fail","next":"AddTwiMLRedirect"},{"event":"return","next":"AddTwiMLRedirect"},{"event":"timeout","next":"AddTwiMLRedirect"}],"type":"add-twiml-redirect"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetAddTwiMLRedirect_basic() string {
	return `
data "twilio_studio_flow_widget_add_twiml_redirect" "add_twiml_redirect" {
  name = "AddTwiMLRedirect"
  url  = "https://test.com/twiml"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetAddTwiMLRedirect_complete() string {
	return `
data "twilio_studio_flow_widget_add_twiml_redirect" "add_twiml_redirect" {
  name = "AddTwiMLRedirect"

  transitions {
    fail    = "AddTwiMLRedirect"
    return  = "AddTwiMLRedirect"
    timeout = "AddTwiMLRedirect"
  }

  method  = "POST"
  timeout = "100"
  url     = "https://test.com/twiml"

  offset {
    x = 10
    y = 20
  }
}
`
}
