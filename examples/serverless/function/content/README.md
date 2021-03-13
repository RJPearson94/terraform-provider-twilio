# Serverless Function - Content

This example provisions the following resources:

- serverless service
- serverless function
- serverless function version (using the content input)

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

| Name                                                                                                                              |
| --------------------------------------------------------------------------------------------------------------------------------- |
| [random_string](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string)                            |
| [twilio_serverless_function](https://registry.terraform.io/providers/RJPearson94/twilio/0.2.0/docs/resources/serverless_function) |
| [twilio_serverless_service](https://registry.terraform.io/providers/RJPearson94/twilio/0.2.0/docs/resources/serverless_service)   |

## Inputs

No input.

## Outputs

| Name     | Description                       |
| -------- | --------------------------------- |
| function | The Generated Serverless Function |
| service  | The Generated Serverless Service  |
