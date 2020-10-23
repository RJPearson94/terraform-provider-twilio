# Phone Number

This example provisions the following resources:

- Phone number

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 0.13  |
| twilio    | >= 0.4.0 |

## Providers

| Name   | Version  |
| ------ | -------- |
| twilio | >= 0.4.0 |

## Inputs

| Name          | Description                                        | Type     | Default | Required |
| ------------- | -------------------------------------------------- | -------- | ------- | :------: |
| account_sid  | The account SID to associate the phone number with | `string` | n/a     |   yes    |
| phone_number | The phone number to purchase                       | `string` | n/a     |   yes    |

## Outputs

| Name          | Description                |
| ------------- | -------------------------- |
| phone_number | The Purchased Phone Number |
