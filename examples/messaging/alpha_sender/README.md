# Messaging Alpha Sender

This example provisions the following resources:

- messaging service
- messaging alpha sender resource

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

| Name         | Description                                      | Type     | Default | Required |
| ------------ | ------------------------------------------------ | -------- | ------- | :------: |
| alpha_sender | Alpha Sender to associate with Messaging Service | `string` | n/a     |   yes    |

## Outputs

| Name         | Description                                            |
| ------------ | ------------------------------------------------------ |
| alpha_sender | The Alpha Sender associated with the Messaging Service |
| service      | The Generated Messaging Service                        |