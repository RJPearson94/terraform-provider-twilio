# Credentials - AWS

This example provisions the following resources:

- aws credentials resource

## Requirements

| Name      | Version   |
| --------- | --------- |
| terraform | >= 0.13   |
| twilio    | >= 0.14.0 |

## Providers

| Name   | Version   |
| ------ | --------- |
| aws    | n/a       |
| twilio | >= 0.14.0 |

## Modules

No Modules.

## Resources

| Name                                                                                                                       |
| -------------------------------------------------------------------------------------------------------------------------- |
| [aws_iam_access_key](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_access_key)           |
| [aws_iam_user](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user)                       |
| [twilio_credentials_aws](https://registry.terraform.io/providers/RJPearson94/twilio/0.14.0/docs/resources/credentials_aws) |

## Inputs

No input.

## Outputs

| Name           | Description                           |
| -------------- | ------------------------------------- |
| aws_credential | The Generated AWS Credential Resource |
