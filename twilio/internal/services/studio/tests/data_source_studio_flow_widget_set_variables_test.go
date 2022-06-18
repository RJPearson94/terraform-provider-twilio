package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetSetVariables_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_set_variables.set_variables"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSetVariables_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SetVariables","properties":{},"transitions":[{"event":"next"}],"type":"set-variables"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSetVariables_withVariables(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_set_variables.set_variables"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSetVariables_withVariables(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SetVariables","properties":{"variables":[{"key":"test","value":"testValue"}]},"transitions":[{"event":"next"}],"type":"set-variables"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSetVariables_withMultipleVariables(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_set_variables.set_variables"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSetVariables_withMultipleVariables(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SetVariables","properties":{"variables":[{"key":"test","value":"testValue"},{"key":"test2","value":"testValue2"}]},"transitions":[{"event":"next"}],"type":"set-variables"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSetVariables_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_set_variables.set_variables"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSetVariables_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SetVariables","properties":{"offset":{"x":10,"y":20},"variables":[{"key":"test","value":"testValue"}]},"transitions":[{"event":"next","next":"SetVariables"}],"type":"set-variables"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetSetVariables_basic() string {
	return `
data "twilio_studio_flow_widget_set_variables" "set_variables" {
  name = "SetVariables"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSetVariables_withVariables() string {
	return `
data "twilio_studio_flow_widget_set_variables" "set_variables" {
  name = "SetVariables"

  variables {
    key   = "test"
    value = "testValue"
  }
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSetVariables_withMultipleVariables() string {
	return `
data "twilio_studio_flow_widget_set_variables" "set_variables" {
  name = "SetVariables"

  variables {
    key   = "test"
    value = "testValue"
  }

  variables {
    key   = "test2"
    value = "testValue2"
  }
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSetVariables_complete() string {
	return `
data "twilio_studio_flow_widget_set_variables" "set_variables" {
  name = "SetVariables"

  transitions {
    next = "SetVariables"
  }

  variables {
    key   = "test"
    value = "testValue"
  }

  offset {
    x = 10
    y = 20
  }
}
`
}
