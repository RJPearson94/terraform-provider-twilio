# SIP Trunk Phone Number

This example provisions the following resources:

- SIP Trunk
- SIP Trunk Phone Number

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 0.13  |
| twilio    | >= 0.4.0 |

## Providers

| Name   | Version  |
| ------ | -------- |
| twilio | >= 0.4.0 |

## Inputs

| Name         | Description                                 | Type     | Default | Required |
| ------------ | ------------------------------------------- | -------- | ------- | :------: |
| phone_number | The phone number to assign to the SIP trunk | `string` | n/a     |   yes    |

## Outputs

| Name         | Description                |
| ------------ | -------------------------- |
| phone_number | The SIP Trunk Phone Number |
| trunk        | The SIP Trunk              |
