# Messaging Phone Number

This example provisions the following resources:

- messaging service
- messaging phone number resource

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

| Name             | Description                                                    | Type     | Default | Required |
| ---------------- | -------------------------------------------------------------- | -------- | ------- | :------: |
| phone_number_sid | SID of Twilio Phone Number to associate with Messaging Service | `string` | n/a     |   yes    |

## Outputs

| Name         | Description                                            |
| ------------ | ------------------------------------------------------ |
| phone_number | The Phone Number associated with the Messaging Service |
| service      | The Generated Messaging Service                        |
