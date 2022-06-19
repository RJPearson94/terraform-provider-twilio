# Verify Webhook

This example provisions the following resources:

- Service
- Webhook

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

| Name                                                                                                                             | Type     |
| -------------------------------------------------------------------------------------------------------------------------------- | -------- |
| [twilio_verify_service.service](https://registry.terraform.io/providers/RJPearson94/twilio/latest/docs/resources/verify_service) | resource |
| [twilio_verify_webhook.webhook](https://registry.terraform.io/providers/RJPearson94/twilio/latest/docs/resources/verify_webhook) | resource |

## Inputs

No inputs.

## Outputs

| Name                                                     | Description                  |
| -------------------------------------------------------- | ---------------------------- |
| <a name="output_service"></a> [service](#output_service) | The Generated Verify service |
| <a name="output_webhook"></a> [webhook](#output_webhook) | The Generated Verify Webhook |
