# Studio Flow - With Conditional Widget

This example provisions a Studio Flow using the Studio Widget Data Sources. This example shows how conditionals can be used to dynamically generate the Studio Flow definition JSON with and without the Autopilot Assistant widget

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

| Name                                                                                                                                                                    |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [twilio_studio_flow](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/resources/studio_flow)                                                      |
| [twilio_studio_flow_definition](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/data-sources/studio_flow_definition)                             |
| [twilio_studio_flow_widget_send_to_autopilot](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/data-sources/studio_flow_widget_send_to_autopilot) |
| [twilio_studio_flow_widget_send_to_flex](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/data-sources/studio_flow_widget_send_to_flex)           |
| [twilio_studio_flow_widget_trigger](https://registry.terraform.io/providers/RJPearson94/twilio/0.11.0/docs/data-sources/studio_flow_widget_trigger)                     |

## Inputs

| Name                    | Description             | Type     | Default | Required |
| ----------------------- | ----------------------- | -------- | ------- | :------: |
| autopilot_assistant_sid | Autopilot Assistant sid | `string` | `null`  |    no    |
| channel_sid             | Flex Channel sid        | `string` | n/a     |   yes    |
| workflow_sid            | TaskRouter Workflow sid | `string` | n/a     |   yes    |

## Outputs

| Name | Description               |
| ---- | ------------------------- |
| flow | The Generated Studio Flow |
