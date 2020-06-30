# Serverless Environment - Multiple Environments

This example provisions the following resources:

- serverless service
- serverless environment x2 (staging, prod)

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

No input.

## Outputs

| Name                | Description                                  |
| ------------------- | -------------------------------------------- |
| prod_environment    | The Generated Prod Serverless Environment    |
| service             | The Generated Serverless Service             |
| staging_environment | The Generated Staging Serverless Environment |
