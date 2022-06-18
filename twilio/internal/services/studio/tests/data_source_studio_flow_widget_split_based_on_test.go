package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetSplitBasedOn_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_split_based_on.split_based_on"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSplitBasedOn_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SplitBasedOn","properties":{"input":"{{contact.channel.address}}"},"transitions":[{"event":"noMatch"},{"event":"match","next":"SplitBasedOn","conditions":[{"arguments":["{{contact.channel.address}}"],"friendly_name":"If value equal_to test","type":"equal_to","value":"test"}]}],"type":"split-based-on"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetSplitBasedOn_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_split_based_on.split_based_on"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetSplitBasedOn_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"SplitBasedOn","properties":{"input":"{{contact.channel.address}}","offset":{"x":10,"y":20}},"transitions":[{"event":"noMatch","next":"SplitBasedOn"},{"event":"match","next":"SplitBasedOn","conditions":[{"arguments":["{{contact.channel.address}}"],"friendly_name":"If value equal_to test","type":"equal_to","value":"test"}]},{"event":"match","next":"SplitBasedOn","conditions":[{"arguments":["{{contact.channel.address}}"],"friendly_name":"If value not_equal_to test 2","type":"not_equal_to","value":"test 2"}]}],"type":"split-based-on"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetSplitBasedOn_basic() string {
	return `
data "twilio_studio_flow_widget_split_based_on" "split_based_on" {
  name = "SplitBasedOn"

  transitions {
    matches {
      next = "SplitBasedOn"
      conditions {
        arguments     = ["{{contact.channel.address}}"]
        friendly_name = "If value equal_to test"
        type          = "equal_to"
        value         = "test"
      }
    }
  }

  input = "{{contact.channel.address}}"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetSplitBasedOn_complete() string {
	return `
data "twilio_studio_flow_widget_split_based_on" "split_based_on" {
  name = "SplitBasedOn"

  transitions {
    no_match = "SplitBasedOn"
    matches {
      next = "SplitBasedOn"
      conditions {
        arguments     = ["{{contact.channel.address}}"]
        friendly_name = "If value equal_to test"
        type          = "equal_to"
        value         = "test"
      }
    }

    matches {
      next = "SplitBasedOn"
      conditions {
        arguments     = ["{{contact.channel.address}}"]
        friendly_name = "If value not_equal_to test 2"
        type          = "not_equal_to"
        value         = "test 2"
      }
    }
  }

  input = "{{contact.channel.address}}"

  offset {
    x = 10
    y = 20
  }
}
`
}
