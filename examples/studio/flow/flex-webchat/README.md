# Studio Flow - Flex Webchat

This example provisions the following resources:

- studio flow

This example uses the following Flow Widgets

- trigger
- send-to-flex

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 0.13  |
| twilio    | >= 0.2.0 |

## Providers

| Name   | Version  |
| ------ | -------- |
| twilio | >= 0.2.0 |

## Modules

No Modules.

## Resources

| Name                                                                                                              |
| ----------------------------------------------------------------------------------------------------------------- |
| [twilio_studio_flow](https://registry.terraform.io/providers/RJPearson94/twilio/0.2.0/docs/resources/studio_flow) |

## Inputs

| Name         | Description              | Type     | Default | Required |
| ------------ | ------------------------ | -------- | ------- | :------: |
| channel_sid  | Task Router Channel sid  | `string` | n/a     |   yes    |
| workflow_sid | Task Router Workflow sid | `string` | n/a     |   yes    |

## Outputs

| Name | Description               |
| ---- | ------------------------- |
| flow | The Generated Studio Flow |
