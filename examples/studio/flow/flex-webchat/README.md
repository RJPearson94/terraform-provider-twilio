# Studio Flow - Flex Webchat

This example provisions the following resources:

- studio flow

This example uses the following Flow Widgets

- trigger
- send-to-flex

## Requirements

| Name      | Version |
| --------- | ------- |
| terraform | >= 0.12 |

## Providers

| Name   | Version |
| ------ | ------- |
| twilio | n/a     |

## Inputs

| Name         | Description              | Type     | Default | Required |
| ------------ | ------------------------ | -------- | ------- | :------: |
| channel_sid  | Task Router Channel sid  | `string` | n/a     |   yes    |
| workflow_sid | Task Router Workflow sid | `string` | n/a     |   yes    |

## Outputs

| Name | Description               |
| ---- | ------------------------- |
| flow | The Generated Studio Flow |
