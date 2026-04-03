# SIP Domain

This example provisions the following resources:

- SIP Credential List
- SIP Credential
- SIP Domain (with voice URL and status callback)
- SIP Domain Credential List Mapping

## Requirements

| Name      | Version   |
| --------- | --------- |
| terraform | >= 0.13   |
| twilio    | >= 0.27.0 |

## Providers

| Name   | Version   |
| ------ | --------- |
| twilio | >= 0.27.0 |

## Inputs

| Name                      | Description                                              | Type     | Required |
| ------------------------- | -------------------------------------------------------- | -------- | -------- |
| account_sid               | The SID of the Twilio account                            | `string` | yes      |
| domain_name               | The unique SIP domain name (must end in .sip.twilio.com) | `string` | yes      |
| voice_url                 | The URL Twilio calls on inbound SIP calls                | `string` | yes      |
| voice_status_callback_url | The URL Twilio calls with call status updates            | `string` | yes      |
| sip_username              | The username for the SIP credential                      | `string` | yes      |
| sip_password              | The password for the SIP credential                      | `string` | yes      |

## Outputs

| Name                    | Description                             |
| ----------------------- | --------------------------------------- |
| credential_list         | The SIP credential list                 |
| credential              | The SIP credential (sensitive)          |
| domain                  | The SIP domain                          |
| credential_list_mapping | The SIP domain credential list mapping  |
