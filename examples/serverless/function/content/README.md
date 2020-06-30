# Serverless Function - Content

This example provisions the following resources:

- serverless service
- serverless function
- serverless function version (using the content input)

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

| Name             | Description                               |
| ---------------- | ----------------------------------------- |
| function         | The Generated Serverless Function         |
| function_version | The Generated Serverless Function Version |
| service          | The Generated Serverless Service          |
