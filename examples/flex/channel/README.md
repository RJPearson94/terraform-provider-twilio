# Flex Flow

This example provisions the following resources:

- flex flow
- flex channel

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

| Name             | Description                                            | Type     | Default | Required |
| ---------------- | ------------------------------------------------------ | -------- | ------- | :------: |
| chat_service_sid | The SID of the Chat Service to associate the Flow with | `string` | n/a     |   yes    |

## Outputs

| Name    | Description                |
| ------- | -------------------------- |
| channel | The Generated Flex Channel |
| flow    | The Generated Flex Flow    |
