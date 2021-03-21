# Serverless Deployment

This example provisions the following resources:

- serverless service
- serverless function
- serverless function version
- serverless asset
- serverless asset version
- serverless build
- serverless environment
- serverless deployment

**Note:** if a function or asset is removed, then function or asset resource could be deleted before the model build has been deleted. This deletion may fail as the function or asset is still part of an active build. If this happens just re-run the apply and the resource will be removed otherwise you might be able to use the `depends_on` to change the way Terraform walks the dependency tree.

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 0.13  |
| twilio    | >= 0.8.2 |

## Providers

| Name   | Version  |
| ------ | -------- |
| random | n/a      |
| twilio | >= 0.8.2 |

## Modules

No Modules.

## Resources

| Name                                                                                                                                    |
| --------------------------------------------------------------------------------------------------------------------------------------- |
| [random_string](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string)                                  |
| [twilio_serverless_asset](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.2/docs/resources/serverless_asset)             |
| [twilio_serverless_build](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.2/docs/resources/serverless_build)             |
| [twilio_serverless_deployment](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.2/docs/resources/serverless_deployment)   |
| [twilio_serverless_environment](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.2/docs/resources/serverless_environment) |
| [twilio_serverless_function](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.2/docs/resources/serverless_function)       |
| [twilio_serverless_service](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.2/docs/resources/serverless_service)         |

## Inputs

No input.

## Outputs

| Name        | Description                          |
| ----------- | ------------------------------------ |
| asset       | The Generated Serverless Asset       |
| build       | The Generated Serverless Build       |
| deployment  | The Generated Serverless Deployment  |
| environment | The Generated Serverless Environment |
| function    | The Generated Serverless Function    |
| service     | The Generated Serverless Service     |
