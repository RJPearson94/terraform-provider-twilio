# Messaging Short Code

This example provisions the following resources:

- messaging service
- messaging short code resource

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

| Name           | Description                                                  | Type     | Default | Required |
| -------------- | ------------------------------------------------------------ | -------- | ------- | :------: |
| short_code_sid | SID of Twilio Short Code to associate with Messaging Service | `string` | n/a     |   yes    |

## Outputs

| Name       | Description                                          |
| ---------- | ---------------------------------------------------- |
| short_code | The Short Code associated with the Messaging Service |
| service    | The Generated Messaging Service                      |
