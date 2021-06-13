# Studio Flow - With Widgets

This example provisions a Studio Flow using the Studio Widget Data Sources

## Requirements

| Name      | Version   |
| --------- | --------- |
| terraform | >= 0.13   |
| twilio    | >= 0.11.0 |

## Providers

| Name   | Version   |
| ------ | --------- |
| twilio | >= 0.11.0 |

## Modules

No Modules.

## Resources

| Name                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [twilio_studio_flow](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/resources/studio_flow)                                            |
| [twilio_studio_flow_definition](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/data-sources/studio_flow_definition)                   |
| [twilio_studio_flow_widget_send_to_flex](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/data-sources/studio_flow_widget_send_to_flex) |
| [twilio_studio_flow_widget_trigger](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/data-sources/studio_flow_widget_trigger)           |

## Inputs

| Name         | Description             | Type     | Default | Required |
| ------------ | ----------------------- | -------- | ------- | :------: |
| channel_sid  | Flex Channel sid        | `string` | n/a     |   yes    |
| workflow_sid | TaskRouter Workflow sid | `string` | n/a     |   yes    |

## Outputs

| Name            | Description                          |
| --------------- | ------------------------------------ |
| flow_definition | The Generated Studio Flow Definition |
