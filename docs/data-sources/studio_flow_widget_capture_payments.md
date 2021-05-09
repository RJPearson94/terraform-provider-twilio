---
page_title: "Twilio Studio Flow Widget - Capture payments"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the capture payments widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `payment_token_type` - (Optional) The payment type. Valid values include: `one-time` or `reusable`
- `payment_connector` - (Optional) The unique name of a Payment Gateway Connector
- `payment_amount` - (Optional) The amount to charge
- `description` - (Optional) A description of the the payment
- `language` - (Optional) The language to use when asking the caller for card details
- `min_postal_code_length` - (Optional) The minimum length of an acceptable postal code
- `timeout` - (Optional) The time in seconds to wait for a caller to press a digit on there phone before validating the digits which have been captured
- `max_attempts` - (Optional) The maximum number of attempts at capturing the card details
- `security_code` - (Optional) Whether to ask the caller for the security code of their card
- `currency` - (Optional) The currency to use when charging the card
- `postal_code` - (Optional) The postal code to use when not prompting the caller for their postal code
- `payment_method` - (Optional) The method of payment. Valid values include: `ACH_DEBIT` or `CREDIT_CARD`
- `bank_account_type` - (Optional) The type of bank account which is being charged. Valid values include: `COMMERCIAL_CHECKING`, `COMMERCIAL_SAVINGS`, `CONSUMER_CHECKING` or `CONSUMER_SAVINGS`
- `valid_card_types` - (Optional) The list of cards which are supported. Valid values for items in the list include: `amex`, `diners-club`, `discover`, `enroute`, `jcb`, `maestro`, `master-card`, `optima` or `visa`
- `parameters` - (Optional) A list of `parameter` blocks as documented below

~> Due to data type and validation restrictions liquid templates are not supported for the `payment_token_type`, `min_postal_code_length`, `timeout`, `max_attempts`, `security_code`, `payment_method`, `bank_account_type` and `valid_card_types` (items) arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `parameter` block supports the following:

- `key` - (Mandatory) The parameter name/ key to pass with the payment
- `value` - (Mandatory) The value of the parameter to pass with the payment

---

A `transitions` block supports the following:

- `hangup` - (Optional) The widget to transition to when the caller hangs up
- `max_failed_attempts` - (Optional) The widget to transition to when the maximum number of failed attempts is reached
- `pay_interrupted` - (Optional) The widget to transition to when the payment is interrupted/ stopped by the caller pressing `*`
- `provider_error` - (Optional) The widget to transition to when an error occurs calling the payment provider
- `success` - (Optional) The widget to transition to when the payment is successfully processed
- `validation_error` - (Optional) The widget to transition to when invalid input is received

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the capture payments widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the capture payments widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the capture payments widget
- `json` - The JSON state definition for the capture payments widget
