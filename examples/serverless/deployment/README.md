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

| Name        | Description                          |
| ----------- | ------------------------------------ |
| asset       | The Generated Serverless Asset       |
| build       | The Generated Serverless Build       |
| deployment  | The Generated Serverless Deployment  |
| environment | The Generated Serverless Environment |
| function    | The Generated Serverless Function    |
| service     | The Generated Serverless Service     |
