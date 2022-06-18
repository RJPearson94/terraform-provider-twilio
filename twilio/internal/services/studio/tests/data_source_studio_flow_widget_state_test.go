package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetState_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_state.state"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetState_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"State","properties":{"digits":"123"},"transitions":[{"event":"audioComplete"}],"type":"say-play"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetState_withNext(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_state.state"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetState_withNext(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"State","properties":{"digits":"123"},"transitions":[{"event":"audioComplete","next":"State"}],"type":"say-play"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetState_withNextAndCondition(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_state.state"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetState_withNextAndCondition(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"State","properties":{"channel":"TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","input":"{{contact.channel.address}}","workflow":"WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"test","next":"testTransition","conditions":[{"arguments":["{{contact.channel.address}}"],"friendly_name":"If value equal_to test","type":"equal_to","value":"test"}]}],"type":"test-type"}`),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetState_withMultipleConditions(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_state.state"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetState_withMultipleConditions(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"State","properties":{"channel":"TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","input":"{{contact.channel.address}}","workflow":"WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"test","conditions":[{"arguments":["{{contact.channel.address}}"],"friendly_name":"If value equal_to test","type":"equal_to","value":"test"},{"arguments":["{{contact.channel.address}}"],"friendly_name":"If value equal_to test 2","type":"equal_to","value":"test 2"}]}],"type":"test-type"}`),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetState_withMultipleTransitions(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_state.state"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetState_multipleTransitions(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"State","properties":{"input":"{{contact.channel.address}}"},"transitions":[{"event":"test","next":"testTransition","conditions":[{"arguments":["{{contact.channel.address}}"],"friendly_name":"If value equal_to test","type":"equal_to","value":"test"}]},{"event":"noMatch"}],"type":"test-type"}`),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetState_basic() string {
	return `
data "twilio_studio_flow_widget_state" "state" {
  name = "State"
  type = "say-play"

  transitions {
    event = "audioComplete"
  }

  properties = {
    "digits" : "123"
  }
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetState_withNext() string {
	return `
data "twilio_studio_flow_widget_state" "state" {
  name = "State"
  type = "say-play"

  transitions {
    event = "audioComplete"
    next  = "State"
  }

  properties = {
    "digits" : "123"
  }
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetState_withNextAndCondition() string {
	return `
data "twilio_studio_flow_widget_state" "state" {
  name = "State"
  type = "test-type"

  transitions {
    event = "test"
    next  = "testTransition"
    conditions {
      arguments     = ["{{contact.channel.address}}"]
      friendly_name = "If value equal_to test"
      type          = "equal_to"
      value         = "test"
    }
  }

  properties = {
    "input" : "{{contact.channel.address}}",
    "workflow" : "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "channel" : "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  }
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetState_withMultipleConditions() string {
	return `
data "twilio_studio_flow_widget_state" "state" {
  name = "State"
  type = "test-type"

  transitions {
    event = "test"
    conditions {
      arguments     = ["{{contact.channel.address}}"]
      friendly_name = "If value equal_to test"
      type          = "equal_to"
      value         = "test"
    }
    conditions {
      arguments     = ["{{contact.channel.address}}"]
      friendly_name = "If value equal_to test 2"
      type          = "equal_to"
      value         = "test 2"
    }
  }

  properties = {
    "input" : "{{contact.channel.address}}",
    "workflow" : "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "channel" : "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  }
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetState_multipleTransitions() string {
	return `
data "twilio_studio_flow_widget_state" "state" {
  name = "State"
  type = "test-type"

  transitions {
    event = "test"
    next  = "testTransition"
    conditions {
      arguments     = ["{{contact.channel.address}}"]
      friendly_name = "If value equal_to test"
      type          = "equal_to"
      value         = "test"
    }
  }

  transitions {
    event = "noMatch"
  }

  properties = {
    "input" : "{{contact.channel.address}}",
  }
}
`
}
