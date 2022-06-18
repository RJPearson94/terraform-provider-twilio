package tests

import (
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio/tests/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTwilioStudioFlowWidgetCapturePayments_basic(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_capture_payments.capture_payments"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetCapturePayments_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"CapturePayments","properties":{"payment_token_type":"reusable"},"transitions":[{"event":"hangup"},{"event":"maxFailedAttempts"},{"event":"payInterrupted"},{"event":"providerError"},{"event":"success"},{"event":"validationError"}],"type":"capture-payments"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioStudioFlowWidgetCapturePayments_complete(t *testing.T) {
	stateDataSourceName := "data.twilio_studio_flow_widget_capture_payments.capture_payments"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioStudioFlowWidgetCapturePayments_complete(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(stateDataSourceName, "json", `{"name":"CapturePayments","properties":{"currency":"usd","description":"Pay Bill","language":"en-GB","max_attempts":2,"min_postal_code_length":3,"offset":{"x":10,"y":20},"parameters":[{"key":"key","value":"value"},{"key":"key2","value":"value2"}],"payment_amount":"10.99","payment_connector":"stripe","payment_method":"ACH_DEBIT","payment_token_type":"reusable","postal_code":"false","security_code":true,"timeout":5,"valid_card_types":["visa","amex"]},"transitions":[{"event":"hangup","next":"CapturePayments"},{"event":"maxFailedAttempts","next":"CapturePayments"},{"event":"payInterrupted","next":"CapturePayments"},{"event":"providerError","next":"CapturePayments"},{"event":"success","next":"CapturePayments"},{"event":"validationError","next":"CapturePayments"}],"type":"capture-payments"}`),
					helper.ValidateFlowWidget(stateDataSourceName),
				),
			},
		},
	})
}

func testAccDataSourceTwilioStudioFlowWidgetCapturePayments_basic() string {
	return `
data "twilio_studio_flow_widget_capture_payments" "capture_payments" {
  name               = "CapturePayments"
  payment_token_type = "reusable"
}
`
}

func testAccDataSourceTwilioStudioFlowWidgetCapturePayments_complete() string {
	return `
data "twilio_studio_flow_widget_capture_payments" "capture_payments" {
  name = "CapturePayments"

  transitions {
    hangup              = "CapturePayments"
    max_failed_attempts = "CapturePayments"
    pay_interrupted     = "CapturePayments"
    provider_error      = "CapturePayments"
    success             = "CapturePayments"
    validation_error    = "CapturePayments"
  }

  currency               = "usd"
  description            = "Pay Bill"
  language               = "en-GB"
  max_attempts           = 2
  min_postal_code_length = 3
  parameters {
    key   = "key"
    value = "value"
  }
  parameters {
    key   = "key2"
    value = "value2"
  }
  payment_amount     = "10.99"
  payment_connector  = "stripe"
  payment_method     = "ACH_DEBIT"
  payment_token_type = "reusable"
  postal_code        = "false"
  security_code      = true
  timeout            = 5
  valid_card_types = [
    "visa",
    "amex"
  ]

  offset {
    x = 10
    y = 20
  }
}
`
}
