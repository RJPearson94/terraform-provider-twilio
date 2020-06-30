# Proxy Phone Number

This example provisions the following resources:

- proxy service
- proxy phone number

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

| Name             | Description                                     | Type     | Default | Required |
| ---------------- | ----------------------------------------------- | -------- | ------- | :------: |
| phone_number_sid | Twilio Phone Number SID to associate with proxy | `string` | n/a     |   yes    |

## Outputs

| Name    | Description                 |
| ------- | --------------------------- |
| service | The Generated Proxy Service |
