# Verify Service Rate Limit Bucket

This example provisions the following resources:

- Service
- Service Rate Limit
- Service Rate Limit Bucket

## Requirements

| Name                                                                     | Version   |
| ------------------------------------------------------------------------ | --------- |
| <a name="requirement_terraform"></a> [terraform](#requirement_terraform) | >= 0.13   |
| <a name="requirement_twilio"></a> [twilio](#requirement_twilio)          | >= 0.18.0 |

## Providers

| Name                                                      | Version   |
| --------------------------------------------------------- | --------- |
| <a name="provider_twilio"></a> [twilio](#provider_twilio) | >= 0.18.0 |

## Modules

No modules.

## Resources

| Name                                                                                                                                                                           | Type     |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| [twilio_verify_service.service](https://registry.terraform.io/providers/RJPearson94/twilio/latest/docs/resources/verify_service)                                               | resource |
| [twilio_verify_service_rate_limit.rate_limit](https://registry.terraform.io/providers/RJPearson94/twilio/latest/docs/resources/verify_service_rate_limit)                      | resource |
| [twilio_verify_service_rate_limit_bucket.rate_limit_bucket](https://registry.terraform.io/providers/RJPearson94/twilio/latest/docs/resources/verify_service_rate_limit_bucket) | resource |

## Inputs

No inputs.

## Outputs

| Name                                                                                   | Description                                    |
| -------------------------------------------------------------------------------------- | ---------------------------------------------- |
| <a name="output_rate_limit"></a> [rate_limit](#output_rate_limit)                      | The Generated Verify Service Rate Limit        |
| <a name="output_rate_limit_bucket"></a> [rate_limit_bucket](#output_rate_limit_bucket) | The Generated Verify Service Rate Limit Bucket |
| <a name="output_service"></a> [service](#output_service)                               | The Generated Verify service                   |
