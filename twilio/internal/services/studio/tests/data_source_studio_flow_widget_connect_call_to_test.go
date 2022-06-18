package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetConnectCallTo_client(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_call_to.connect_call_to"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_client(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectCallTo","properties":{"caller_id":"{{contact.channel.address}}","noun":"client","to":"test"},"transitions":[{"event":"callCompleted"},{"event":"hangup"}],"type":"connect-call-to"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetConnectCallTo_conference(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_call_to.connect_call_to"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_conference(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectCallTo","properties":{"caller_id":"{{contact.channel.address}}","noun":"conference","to":"CFaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"callCompleted"},{"event":"hangup"}],"type":"connect-call-to"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetConnectCallTo_number(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_call_to.connect_call_to"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_number(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectCallTo","properties":{"caller_id":"{{contact.channel.address}}","noun":"number","to":"+441234567890"},"transitions":[{"event":"callCompleted"},{"event":"hangup"}],"type":"connect-call-to"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetConnectCallTo_numberMulti(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_call_to.connect_call_to"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_numberMulti(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectCallTo","properties":{"caller_id":"{{contact.channel.address}}","noun":"number-multi","to":"+441234567890,+441234567891"},"transitions":[{"event":"callCompleted"},{"event":"hangup"}],"type":"connect-call-to"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetConnectCallTo_sim(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_call_to.connect_call_to"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_sim(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectCallTo","properties":{"caller_id":"{{contact.channel.address}}","noun":"sim","to":"DEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},"transitions":[{"event":"callCompleted"},{"event":"hangup"}],"type":"connect-call-to"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetConnectCallTo_sip(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_call_to.connect_call_to"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_sip(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectCallTo","properties":{"caller_id":"{{contact.channel.address}}","noun":"sip","sip_endpoint":"sip:test@test.com"},"transitions":[{"event":"callCompleted"},{"event":"hangup"}],"type":"connect-call-to"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetConnectCallTo_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_connect_call_to.connect_call_to"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"ConnectCallTo","properties":{"caller_id":"{{contact.channel.address}}","noun":"sip","offset":{"x":10,"y":20},"record":true,"sip_endpoint":"sip:test@test.com","sip_password":"test2","sip_username":"test","timeout":30},"transitions":[{"event":"callCompleted","next":"ConnectCallTo"},{"event":"hangup","next":"ConnectCallTo"}],"type":"connect-call-to"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_client() string {
	return `
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "client"
  to   = "test"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_conference() string {
	return `
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "conference"
  to   = "CFaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_number() string {
	return `
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "number"
  to   = "+441234567890"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_numberMulti() string {
	return `
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "number-multi"
  to   = "+441234567890,+441234567891"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_sim() string {
	return `
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "sim"
  to   = "DEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_sip() string {
	return `
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name         = "ConnectCallTo"
  noun         = "sip"
  sip_endpoint = "sip:test@test.com"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetConnectCallTo_complete() string {
	return `
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"

  transitions {
    call_completed = "ConnectCallTo"
    hangup         = "ConnectCallTo"
  }

  caller_id    = "{{contact.channel.address}}"
  record       = true
  noun         = "sip"
  timeout      = 30
  sip_username = "test"
  sip_password = "test2"
  sip_endpoint = "sip:test@test.com"

  offset {
    x = 10
    y = 20
  }
}
`
}
