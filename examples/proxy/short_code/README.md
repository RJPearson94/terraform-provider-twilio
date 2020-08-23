# Proxy Short Code

This example provisions the following resources:

- proxy service
- proxy short code

## Requirements

| Name      | Version |
| --------- | ------- |
| terraform | >= 0.12 |

## Providers

| Name   | Version |
| ------ | ------- |
| random | n/a     |
| twilio | n/a     |

## Inputs

| Name           | Description                                   | Type     | Default | Required |
| -------------- | --------------------------------------------- | -------- | ------- | :------: |
| short_code_sid | Twilio Short Code SID to associate with proxy | `string` | n/a     |   yes    |

## Outputs

| Name       | Description                                      |
| ---------- | ------------------------------------------------ |
| service    | The Generated Proxy Service                      |
| short_code | The Short Code associated with the Proxy Service |
