# Account Address

This example provisions the following resources:

- Address

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 0.13  |
| twilio    | >= 0.3.0 |

## Providers

| Name   | Version  |
| ------ | -------- |
| twilio | >= 0.3.0 |

## Inputs

| Name          | Description                                   | Type     | Default | Required |
| ------------- | --------------------------------------------- | -------- | ------- | :------: |
| account_sid   | The account SID to associate the address with | `string` | n/a     |   yes    |
| city          | The address city                              | `string` | n/a     |   yes    |
| customer_name | Your name/ business name                      | `string` | n/a     |   yes    |
| iso_country   | The address ISO country                       | `string` | n/a     |   yes    |
| postal_code   | The address postal code                       | `string` | n/a     |   yes    |
| region        | The address region                            | `string` | n/a     |   yes    |
| street        | The address street                            | `string` | n/a     |   yes    |

## Outputs

| Name    | Description           |
| ------- | --------------------- |
| address | The Generated Address |
