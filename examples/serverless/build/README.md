# Serverless Build

This example provisions the following resources:

- serverless service
- serverless function
- serverless function version
- serverless build

**Note:** if a function or asset is removed, then function or asset resource could be deleted before the model build has been deleted. This deletion may fail as the function or asset is still part of an active build. If this happens just re-run the apply and the resource will be removed otherwise you might be able to use the `depends_on` to change the way Terraform walks the dependency tree.

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 0.13  |
| twilio    | >= 0.8.1 |

## Providers

| Name   | Version  |
| ------ | -------- |
| random | n/a      |
| twilio | >= 0.8.1 |

## Modules

No Modules.

## Resources

| Name                                                                                                                              |
| --------------------------------------------------------------------------------------------------------------------------------- |
| [random_string](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string)                            |
| [twilio_serverless_build](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.1/docs/resources/serverless_build)       |
| [twilio_serverless_function](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.1/docs/resources/serverless_function) |
| [twilio_serverless_service](https://registry.terraform.io/providers/RJPearson94/twilio/0.8.1/docs/resources/serverless_service)   |

## Inputs

No input.

## Outputs

| Name     | Description                       |
| -------- | --------------------------------- |
| build    | The Generated Serverless Build    |
| function | The Generated Serverless Function |
| service  | The Generated Serverless Service  |
