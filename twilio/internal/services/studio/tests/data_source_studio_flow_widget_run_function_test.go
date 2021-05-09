package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetRunFunction_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_run_function.run_function"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetRunFunction_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"RunFunction","properties":{"url":"https://test-function.twil.io/test-function"},"transitions":[{"event":"fail"},{"event":"success"}],"type":"run-function"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetRunFunction_legacyService(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_run_function.run_function"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetRunFunction_legacyService(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"RunFunction","properties":{"service_sid":"default","url":"https://test-function.twil.io/test-function"},"transitions":[{"event":"fail"},{"event":"success"}],"type":"run-function"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetRunFunction_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_run_function.run_function"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetRunFunction_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"RunFunction","properties":{"environment_sid":"ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","function_sid":"ZHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","offset":{"x":10,"y":20},"parameters":[{"key":"key","value":"value"},{"key":"key2","value":"value2"}],"service_sid":"ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","url":"https://test-function.twil.io/test-function"},"transitions":[{"event":"fail","next":"RunFunction"},{"event":"success","next":"RunFunction"}],"type":"run-function"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetRunFunction_basic() string {
	return `
data "twilio_studio_flow_widget_run_function" "run_function" {
	name = "RunFunction"
	url = "https://test-function.twil.io/test-function"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetRunFunction_legacyService() string {
	return `
data "twilio_studio_flow_widget_run_function" "run_function" {
	name = "RunFunction"
	
	service_sid = "default"
    url = "https://test-function.twil.io/test-function"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetRunFunction_complete() string {
	return `
data "twilio_studio_flow_widget_run_function" "run_function" {
	name = "RunFunction"
	
    transitions {
		fail = "RunFunction"
		success = "RunFunction"
	}

    function_sid = "ZHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	environment_sid = "ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	service_sid = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    parameters {
		key = "key"
        value = "value"
    }
	parameters {
		key = "key2"
        value = "value2"
    }
    url = "https://test-function.twil.io/test-function"

	offset {
		x = 10
		y = 20
	}
}
`
}
