## v0.6.2 (2021-01-22)

Correct bug when setting VoiceReceiveMode on twilio phone number resource. The voice receive mode was incorrectly being set to the voice and fax state instead of the corresponding string value on both create and update. This has now been corrected

## v0.6.1 (2020-12-12)

Correct some incorrect documentation on the terraform registry

## v0.6.0 (2020-12-12)

FEATURES

- **New Data Source:** `twilio_sip_trunking_origination_url` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_origination_url.md)
- **New Data Source:** `twilio_sip_trunking_origination_urls` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_origination_urls.md)
- **New Data Source:** `twilio_sip_trunking_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_phone_number.md)
- **New Data Source:** `twilio_sip_trunking_phone_numbers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_phone_numbers.md)
- **New Data Source:** `twilio_sip_trunking_trunk` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_trunk.md)
- **New Resource:** `twilio_sip_trunking_origination_url` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_trunking_origination_url.md)
- **New Resource:** `twilio_sip_trunking_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_trunking_phone_number.md)
- **New Resource:** `twilio_sip_trunking_trunk` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_trunking_trunk.md)

## v0.5.0 (2020-11-26)

Bump Terraform SDK Plugin and Twilio Go SDK versions
Deprecated searching for fax enabled phone numbers using the `twilio_phone_number_available_local_numbers`, `twilio_phone_number_available_mobile_numbers` and `twilio_phone_number_available_toll_free_numbers` data sources
Deprecated fax block for `twilio_phone_number` resource and `twilio_phone_number` and `twilio_phone_numbers` data sources
Deprecated fax from the capabilities block for `twilio_phone_number` resource and `twilio_phone_number` and `twilio_phone_numbers` data sources
In the next version only the voice block will be supported for `twilio_phone_number` resource and `twilio_phone_number` and `twilio_phone_numbers` data sources

BREAKING CHANGES

The twilio api no longer returns whether a phone number is configured to integrated with services when a call or fax is received. If no value is supplied then the `voice` block will be used for `twilio_phone_number` resource and `twilio_phone_number` and `twilio_phone_numbers` data sources

## v0.4.0 (2020-10-24)

FEATURES

- **New Data Source:** `twilio_phone_number_available_toll_free_numbers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/phone_number_available_toll_free_numbers.md)
- **New Data Source:** `twilio_phone_number_available_mobile_numbers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/phone_number_available_mobile_numbers.md)
- **New Data Source:** `twilio_phone_number_available_local_numbers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/phone_number_available_local_numbers.md)
- **New Data Source:** `twilio_phone_numbers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/phone_number.md)
- **New Data Source:** `twilio_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/phone_number.md)
- **New Resource:** `twilio_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/phone_number.md)

## v0.3.0 (2020-10-04)

Update Terraform SDK Plugin to V2 and add new resources and data source for addresses and queues

BREAKING CHANGES

The terraform sdk plugin has been update to V2, this means that this provider requires terraform 0.12+

FEATURES

- **New Data Source:** `twilio_account_addresses` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/account_addresses.md)
- **New Data Source:** `twilio_account_address` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/account_address.md)
- **New Resource:** `twilio_account_address` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/account_address.md)
- **New Data Source:** `twilio_voice_queues` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/voice_queues.md)
- **New Data Source:** `twilio_voice_queue` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/voice_queue.md)
- **New Resource:** `twilio_voice_queue` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/voice_queue.md)

## v0.2.1 (2020-09-13)

BUG FIXES

`twilio_autopilot_task` Fix issue with action json not being updated. Sending the Action URL superseded the actions json change. The update task function now checks if the actions json has changed and conditionally sets either the actions or actions_url field on the update request

## v0.2.0 (2020-09-12)

FEATURES
`twilio_autopilot_model_build` Add `triggers` to force the recreated of a model build [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_model_build.md)
`twilio_serverless_build` Add `triggers` to force the recreated of a build [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_build.md)
`twilio_serverless_deployment` Add `triggers` to force the recreated of a build and add `is_latest_deployment` to indicate whether the terraform resource is the latest deployment [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_deployment.md)

BUG FIXES

`twilio_autopilot_model_build` A new model build wasn't automatically triggered when dependent resources were updated. [Issue](https://github.com/RJPearson94/terraform-provider-twilio/issues/7) **Breaking Change**
`twilio_serverless_asset` recompute latest_version_sid when the asset was changed [Related Issue](https://github.com/RJPearson94/terraform-provider-twilio/issues/6)
`twilio_serverless_function` recompute latest_version_sid when the function was changed [Related Issue](https://github.com/RJPearson94/terraform-provider-twilio/issues/6)
`twilio_serverless_build` Removed computed flag on asset and function version arguments to allow an artefact to be removed and the build is recreated and force a new resource to be created when when function and/ or asset sid has changed [Related Issue](https://github.com/RJPearson94/terraform-provider-twilio/issues/6) **Breaking Change**

## v0.1.1 (2020-09-06)

Update Documentation with Terraform Registry Information and disclaimer of terraform docs

## v0.1.0 (2020-09-06)

Add first release of the Twilio Terraform provider. The list of supported resources and data sources can be seen in the features section below.

FEATURES

- **New Data Source:** `twilio_account_balance` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/account_balance.md)
- **New Data Source:** `twilio_account_details` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/account_details.md)
- **New Resource:** `twilio_account_sub_account` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/account_sub_account.md)
- **New Data Source:** `twilio_autopilot_assistant` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_assistant.md)
- **New Data Source:** `twilio_autopilot_field_type` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_field_type.md)
- **New Data Source:** `twilio_autopilot_field_types` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_field_types.md)
- **New Data Source:** `twilio_autopilot_field_value` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_field_value.md)
- **New Data Source:** `twilio_autopilot_field_values` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_field_values.md)
- **New Data Source:** `twilio_autopilot_model_build` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_model_build.md)
- **New Data Source:** `twilio_autopilot_model_builds` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_model_builds.md)
- **New Data Source:** `twilio_autopilot_task` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_task.md)
- **New Data Source:** `twilio_autopilot_tasks` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_tasks.md)
- **New Data Source:** `twilio_autopilot_task_field` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_task_field.md)
- **New Data Source:** `twilio_autopilot_task_fields` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_task_fields.md)
- **New Data Source:** `twilio_autopilot_task_sample` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_task_sample.md)
- **New Data Source:** `twilio_autopilot_task_samples` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_task_samples.md)
- **New Data Source:** `twilio_autopilot_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_webhook.md)
- **New Data Source:** `twilio_autopilot_webhooks` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/autopilot_webhooks.md)
- **New Resource:** `twilio_autopilot_assistant` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_assistant.md)
- **New Resource:** `twilio_autopilot_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_webhook.md)
- **New Resource:** `twilio_autopilot_task` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_webhook.md)
- **New Resource:** `twilio_autopilot_task_sample` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_task_sample.md)
- **New Resource:** `twilio_autopilot_task_field` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_task_field.md)
- **New Resource:** `twilio_autopilot_field_type` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_field_type.md)
- **New Resource:** `twilio_autopilot_field_value` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_field_value.md)
- **New Resource:** `twilio_autopilot_model_build` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/autopilot_model_build.md)
- **New Data Source:** `twilio_chat_channel` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_channel.md)
- **New Data Source:** `twilio_chat_channels` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_channels.md)
- **New Data Source:** `twilio_chat_channel_member` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_channel_member.md)
- **New Data Source:** `twilio_chat_channel_members` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_channel_members.md)
- **New Data Source:** `twilio_chat_role` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_role.md)
- **New Data Source:** `twilio_chat_roles` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_roles.md)
- **New Data Source:** `twilio_chat_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_service.md)
- **New Data Source:** `twilio_chat_channel_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_channel_webhook.md)
- **New Data Source:** `twilio_chat_channel_webhooks` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_channel_webhooks.md)
- **New Data Source:** `twilio_chat_user` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_user.md)
- **New Data Source:** `twilio_chat_users` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/chat_users.md)
- **New Resource:** `twilio_chat_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_service.md)
- **New Resource:** `twilio_chat_role` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_role.md)
- **New Resource:** `twilio_chat_channel` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_channel.md)
- **New Resource:** `twilio_chat_channel_member` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_channel_member.md)
- **New Resource:** `twilio_chat_channel_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_channel_webhook.md)
- **New Resource:** `twilio_chat_channel_studio_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_channel_studio_webhook.md)
- **New Resource:** `twilio_chat_channel_trigger_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_channel_trigger_webhook.md)
- **New Resource:** `twilio_chat_user` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/chat_user.md)
- **New Data Source:** `twilio_flex_flow` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/flex_flow.md)
- **New Resource:** `twilio_flex_flow` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/flex_flow.md)
- **New Resource:** `twilio_iam_api_key` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/iam_api_key.md)
- **New Data Source:** `twilio_messaging_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/messaging_service.md)
- **New Data Source:** `twilio_messaging_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/messaging_phone_number.md)
- **New Data Source:** `twilio_messaging_phone_numbers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/messaging_phone_numbers.md)
- **New Data Source:** `twilio_messaging_short_code` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/messaging_short_code.md)
- **New Data Source:** `twilio_messaging_short_codes` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/messaging_short_codes.md)
- **New Data Source:** `twilio_messaging_alpha_sender` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/messaging_alpha_sender.md)
- **New Data Source:** `twilio_messaging_alpha_senders` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/messaging_alpha_senders.md)
- **New Resource:** `twilio_messaging_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/messaging_service.md)
- **New Resource:** `twilio_messaging_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/messaging_phone_number.md)
- **New Resource:** `twilio_messaging_short_code` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/messaging_short_code.md)
- **New Resource:** `twilio_messaging_alpha_sender` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/messaging_alpha_sender.md)
- **New Data Source:** `twilio_proxy_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/proxy_service.md)
- **New Data Source:** `twilio_proxy_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/proxy_phone_number.md)
- **New Data Source:** `twilio_proxy_phone_numbers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/proxy_phone_numbers.md)
- **New Data Source:** `twilio_proxy_short_code` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/proxy_short_code.md)
- **New Data Source:** `twilio_proxy_short_codes` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/proxy_short_codes.md)
- **New Resource:** `twilio_proxy_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/proxy_service.md)
- **New Resource:** `twilio_proxy_phone_number` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/proxy_phone_number.md)
- **New Resource:** `twilio_proxy_short_code` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/proxy_short_code.md)
- **New Data Source:** `twilio_serverless_asset` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_asset.md)
- **New Data Source:** `twilio_serverless_assets` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_assets.md)
- **New Data Source:** `twilio_serverless_build` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_buid.md)
- **New Data Source:** `twilio_serverless_builds` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sererless_builds.md)
- **New Data Source:** `twilio_serverless_deployment` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_deployment.md)
- **New Data Source:** `twilio_serverless_deployments` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_deployments.md)
- **New Data Source:** `twilio_serverless_environment` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_environment.md)
- **New Data Source:** `twilio_serverless_environments` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_environments.md)
- **New Data Source:** `twilio_serverless_function` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/severless_function.md)
- **New Data Source:** `twilio_serverless_functions` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/severless_functions.md)
- **New Data Source:** `twilio_serverless_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_service.md)
- **New Data Source:** `twilio_serverless_variable` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/serverless_variable.md)
- **New Data Source:** `twilio_serverless_variables` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-soures/serverless_variables.md)
- **New Resource:** `twilio_serverless_environment` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/severless_environment.md)
- **New Resource:** `twilio_serverless_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_service.md)
- **New Resource:** `twilio_serverless_variable` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_variable.md)
- **New Resource:** `twilio_serverless_asset` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_asset.md)
- **New Resource:** `twilio_serverless_function` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_function.md)
- **New Resource:** `twilio_serverless_build` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_build.md)
- **New Resource:** `twilio_serverless_deployment` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/serverless_deployment.md)
- **New Data Source:** `twilio_studio_flow` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow.md)
- **New Resource:** `twilio_studio_flow` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/studio_flow.md)
- **New Data Source:** `twilio_taskrouter_activities` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_activities.md)
- **New Data Source:** `twilio_taskrouter_activity` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_activity.md)
- **New Data Source:** `twilio_taskrouter_task_channel` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_task_channel.md)
- **New Data Source:** `twilio_taskrouter_task_channels` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_task_channels.md)
- **New Data Source:** `twilio_taskrouter_task_queue` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_task_queue.md)
- **New Data Source:** `twilio_taskrouter_task_queues` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_task_queues.md)
- **New Data Source:** `twilio_taskrouter_worker` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_worker.md)
- **New Data Source:** `twilio_taskrouter_workers` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_workers.md)
- **New Data Source:** `twilio_taskrouter_workflow` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_workflow.md)
- **New Data Source:** `twilio_taskrouter_workflows` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_workflows.md)
- **New Data Source:** `twilio_taskrouter_workspace` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/taskrouter_workspace.md)
- **New Resource:** `twilio_taskrouter_workspace` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/taskrouter_workspace.md)
- **New Resource:** `twilio_taskrouter_activity` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/taskrouter_activity.md)
- **New Resource:** `twilio_taskrouter_task_queue` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/taskrouter_task_queue.md)
- **New Resource:** `twilio_taskrouter_worker` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/taskrouter_worker.md)
- **New Resource:** `twilio_taskrouter_workflow` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/taskrouter_workflow.md)
- **New Resource:** `twilio_taskrouter_task_channel` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/taskrouter_task_channel.md)
