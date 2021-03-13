# Serverless Environment - Multiple Environments

This example provisions the following resources:

- serverless service
- serverless environment x2 (staging, prod)

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 0.13  |
| twilio    | >= 0.2.0 |

## Providers

| Name   | Version  |
| ------ | -------- |
| random | n/a      |
| twilio | >= 0.2.0 |

## Modules

No Modules.

## Resources

| Name                                                                                                                                    |
| --------------------------------------------------------------------------------------------------------------------------------------- |
| [random_string](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string)                                  |
| [twilio_serverless_environment](https://registry.terraform.io/providers/RJPearson94/twilio/0.2.0/docs/resources/serverless_environment) |
| [twilio_serverless_service](https://registry.terraform.io/providers/RJPearson94/twilio/0.2.0/docs/resources/serverless_service)         |

## Inputs

No input.

## Outputs

| Name                | Description                                  |
| ------------------- | -------------------------------------------- |
| prod_environment    | The Generated Prod Serverless Environment    |
| service             | The Generated Serverless Service             |
| staging_environment | The Generated Staging Serverless Environment |
