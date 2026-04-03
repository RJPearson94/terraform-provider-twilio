---
page_title: "twilio_studio_flow_widget_capture_payments Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_capture_payments Data Source

Use this data source to generate the JSON for the Studio Flow capture payments widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/capture-payments) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_capture_payments" "capture_payments" {
  name               = "CapturePayments"
  payment_token_type = "reusable"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_capture_payments" "capture_payments" {
  name = "CapturePayments"

  transitions {
    hangup              = "HangupTransition"
    max_failed_attempts = "MaxFailedAttemptsTransition"
    pay_interrupted     = "PayInterruptedTransition"
    provider_error      = "ProviderErrorTransition"
    success             = "SuccessTransition"
    validation_error    = "ValidationErrorTransition"
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
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `bank_account_type` (String) The bank account type for ACH payments. Valid values: `COMMERCIAL_CHECKING`, `COMMERCIAL_SAVINGS`, `CONSUMER_CHECKING`, `CONSUMER_SAVINGS`
- `currency` (String) The currency for the payment amount as an ISO 4217 code (e.g. `USD`)
- `description` (String) A human-readable description of what is being charged
- `language` (String) The language for payment prompts (e.g. `en-US`)
- `max_attempts` (Number) The maximum number of allowed payment input attempts before the `max_failed_attempts` transition fires
- `min_postal_code_length` (Number) The minimum number of digits required for postal code input
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `parameters` (Block List) Additional key/value parameters to pass to the payment connector (see [below for nested schema](#nestedblock--parameters))
- `payment_amount` (String) The amount to charge, as a string (e.g. `20.00`)
- `payment_connector` (String) The unique name of the payment connector to use for processing
- `payment_method` (String) The payment method to collect. Valid values: `ACH_DEBIT`, `CREDIT_CARD`
- `payment_token_type` (String) Whether to generate a one-time or reusable payment token. Valid values: `one-time`, `reusable`
- `postal_code` (String) Whether to prompt for a postal code. Set to `false` to disable, or provide a pre-filled value
- `security_code` (Boolean) Whether to prompt the caller for the card security code (CVV)
- `timeout` (Number) The timeout in seconds to wait for a digit between inputs
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `valid_card_types` (List of String) The card brands to accept. Valid values: `amex`, `diners-club`, `discover`, `enroute`, `jcb`, `maestro`, `master-card`, `optima`, `visa`

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source

<a id="nestedblock--offset"></a>
### Nested Schema for `offset`

Optional:

- `x` (Number) The x-axis position. Defaults to 0
- `y` (Number) The y-axis position. Defaults to 0


<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `key` (String) The parameter name
- `value` (String) The parameter value


<a id="nestedblock--transitions"></a>
### Nested Schema for `transitions`

Optional:

- `hangup` (String) The name of the next widget when the caller hangs up
- `max_failed_attempts` (String) The name of the next widget when the maximum number of failed payment attempts is reached
- `pay_interrupted` (String) The name of the next widget when the payment capture is interrupted
- `provider_error` (String) The name of the next widget when the payment provider returns an error
- `success` (String) The name of the next widget when payment is successfully captured
- `validation_error` (String) The name of the next widget when a payment input validation error occurs
