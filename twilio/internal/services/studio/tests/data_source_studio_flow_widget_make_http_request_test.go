package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetMakeHttpRequest_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_make_http_request.make_http_request"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetMakeHttpRequest_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"MakeHttpRequest","properties":{"content_type":"application/x-www-form-urlencoded;charset=utf-8","method":"GET","url":"https://test.com"},"transitions":[{"event":"failed"},{"event":"success"}],"type":"make-http-request"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetMakeHttpRequest_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_make_http_request.make_http_request"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetMakeHttpRequest_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"MakeHttpRequest","properties":{"body":"{\"say\":\"Hello World\"}","content_type":"application/json;charset=utf-8","method":"POST","offset":{"x":10,"y":20},"parameters":[{"key":"key","value":"value"},{"key":"key2","value":"value2"}],"url":"https://test.com"},"transitions":[{"event":"failed","next":"MakeHttpRequest"},{"event":"success","next":"MakeHttpRequest"}],"type":"make-http-request"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetMakeHttpRequest_basic() string {
	return `
data "twilio_studio_flow_widget_make_http_request" "make_http_request" {
  name = "MakeHttpRequest"

  method       = "GET"
  content_type = "application/x-www-form-urlencoded"
  url          = "https://test.com"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetMakeHttpRequest_complete() string {
	return `
data "twilio_studio_flow_widget_make_http_request" "make_http_request" {
  name = "MakeHttpRequest"

  transitions {
    failed  = "MakeHttpRequest"
    success = "MakeHttpRequest"
  }

  method       = "POST"
  content_type = "application/json"
  url          = "https://test.com"
  body = jsonencode({
    "say" : "Hello World"
  })

  parameters {
    key   = "key"
    value = "value"
  }

  parameters {
    key   = "key2"
    value = "value2"
  }

  offset {
    x = 10
    y = 20
  }
}
`
}
