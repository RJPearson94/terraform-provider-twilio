# Credentials - Public Key

This example provisions the following resources:

- public key resource

## Requirements

| Name      | Version   |
| --------- | --------- |
| terraform | >= 0.13   |
| twilio    | >= 0.14.0 |

## Providers

| Name   | Version   |
| ------ | --------- |
| twilio | >= 0.14.0 |

## Modules

No Modules.

## Resources

| Name                                                                                                                                     |
| ---------------------------------------------------------------------------------------------------------------------------------------- |
| [twilio_credentials_public_key](https://registry.terraform.io/providers/RJPearson94/twilio/0.14.0/docs/resources/credentials_public_key) |

## Inputs

| Name       | Description                                                     | Type     | Default | Required |
| ---------- | --------------------------------------------------------------- | -------- | ------- | :------: |
| public_key | A public key which can be used for public key client validation | `string` | n/a     |   yes    |

## Outputs

| Name       | Description                       |
| ---------- | --------------------------------- |
| public_key | The Generated Public Key Resource |
