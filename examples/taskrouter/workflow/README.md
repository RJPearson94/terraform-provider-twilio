# TaskRouter Workflow

This example provisions the following resources:

- taskrouter workspace
- taskrouter task queue
- taskrouter workflow

## Requirements

| Name      | Version |
| --------- | ------- |
| terraform | >= 0.12 |

## Providers

| Name   | Version |
| ------ | ------- |
| twilio | n/a     |

## Inputs

No input.

## Outputs

| Name        | Description                        |
| ----------- | ---------------------------------- |
| task\_queue | The Generated task queue           |
| workflow    | The Generated workflow             |
| workspace   | The Generated TaskRouter Workspace |
