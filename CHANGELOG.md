## v0.12.0 (2021-06-13)

FIXES

- The `account_sid` attribute for `twilio_voice_queue` will force a new resource when the value is changed. This implementation now matches the documented behaviour

FEATURES

- **New Data Source:** `twilio_twiml_app` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/twiml_app.md)
- - **New Resource:** `twilio_twiml_app` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/twiml_app.md)

## v0.11.0 (2021-05-09)

FIXES

- Previously when removing a commit message from a `twilio_studio_flow` resource an error was thrown because an empty string was being passed up to the Twilio API for the commit message. This was incorrect, the commit message parameter should not have been sent to the API as a new version is created without the commit message, hence removing it from the resource
- Previously the provider acceptable a `reachability_enabled` boolean on the `twilio_chat_service` resource however this value was not passed to the Twilio API and always caused a drift. This has been corrected

FEATURES

- **New Data Source:** `twilio_studio_flow_definition` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_definition.md)
- **New Data Source:** `twilio_studio_flow_widget_add_twiml_redirect` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_add_twiml_redirect.md)
- **New Data Source:** `twilio_studio_flow_widget_capture_payments` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_capture_payments.md)
- **New Data Source:** `twilio_studio_flow_widget_connect_call_to` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_connect_call_to.md)
- **New Data Source:** `twilio_studio_flow_widget_connect_virtual_agent` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_connect_virtual_agent.md)
- **New Data Source:** `twilio_studio_flow_widget_enqueue_call` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_enqueue_call.md)
- **New Data Source:** `twilio_studio_flow_widget_fork_stream` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_fork_stream.md)
- **New Data Source:** `twilio_studio_flow_widget_gather_input_on_call` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_gather_input_on_call.md)
- **New Data Source:** `twilio_studio_flow_widget_make_http_request` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_make_http_request.md)
- **New Data Source:** `twilio_studio_flow_widget_make_outgoing_call` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_make_outgoing_call.md)
- **New Data Source:** `twilio_studio_flow_widget_record_call` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_record_call.md)
- **New Data Source:** `twilio_studio_flow_widget_record_voicemail` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_record_voicemail.md)
- **New Data Source:** `twilio_studio_flow_widget_run_function` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_run_function.md)
- **New Data Source:** `twilio_studio_flow_widget_say_play` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_say_play.md)
- **New Data Source:** `twilio_studio_flow_widget_send_and_wait_for_reply` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_send_and_wait_for_reply.md)
- **New Data Source:** `twilio_studio_flow_widget_send_message` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_send_message.md)
- **New Data Source:** `twilio_studio_flow_widget_send_to_autopilot` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_send_to_autopilot.md)
- **New Data Source:** `twilio_studio_flow_widget_send_to_flex` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_send_to_flex.md)
- **New Data Source:** `twilio_studio_flow_widget_set_variables` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_set_variables.md)
- **New Data Source:** `twilio_studio_flow_widget_split_based_on` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_split_based_on.md)
- **New Data Source:** `twilio_studio_flow_widget_state` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_state.md)
- **New Data Source:** `twilio_studio_flow_widget_trigger` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/studio_flow_widget_trigger.md)

- Show the validation errors which are returned in the API response when you create, update or validate a studio flow. This improved error logging on `twilio_studio_flow` resource can aid with debugging issues with your Studio Flow definition

## v0.10.0 (2021-04-25)

### FEATURES

- Make the `sid` argument optional on the `twilio_account_details` data source to allow the account SID configured on the provider to be used. This means the data source can be used instead of having to supply the account SID twice, once in the provider configuration and for any resources/ data sources that have the account sid as an argument
- Add the ability to search for a `local`, `mobile` or `toll_free` phone number on the `twilio_phone_number` resource. This allows you to search for and purchase a phone number in a single resource. This gets around the issue with data sources returns a new available phone numbers list on each plan which caused a new phone number to be purchased when being supplied as the input to the `twilio_phone_number` resource

### FIXES

- Previously when creating a serverless build with multiple functions or asset versions that were created simultaneously, the Twilio API could return these in any order. As the order may not match the order the asset or function versions were supplied, Terraform will detect this as a drift and the build is re-created on the next deployment. A diff suppression has been added to the `sid` argument in the `asset_version` and `function_version` block to check if the SID was present in the list, just in a different position

### NOTES

- The following data sources `twilio_phone_number_available_local_numbers`, `twilio_phone_number_available_mobile_numbers` and `twilio_phone_number_available_toll_free_numbers` have been deprecated. The `search_criteria` block on the `twilio_phone_number` resource should be used instead
- The `function_version` and `asset_version` lists on the `twilio_serverless_build` resource, `twilio_serverless_build` data source and `twilio_serverless_builds` data source are now sorted by SID. This is to ensure the outputs are consistent despite the API returning the values in a random order

## v0.9.0 (2021-04-18)

NOTES

**This change is a major overhaul of the Terraform provider and contains several breaking changes.**

The provider schema now includes validation on several arguments (see documentation for any new limits/ constraints), the introduction of argument defaults. These default values were included to allow me to address an issue when the provider was not clearing down an argument value in Twilio when it was removed from the Terraform configuration. This is a change in the behaviour of the provider **breaking change**

Due to the number of changes it is recommended you run a Terraform plan and make any necessary configuration changes to your Terraform configuration before applying the changes when upgrading the provider version. Once this has been completed all the new defaults should be applied and you can any Terraform configuration changes

- The deprecated `sid` argument on the `twilio_sip_trunking_phone_number` resources has now been removed and the `phone_number_sid` argument has now become mandatory **breaking change**

Defaults have been applied to the following resources:

- twilio_account_address - emergency_enabled
- twilio_account_sub_account - status
- twilio_autopilot_assistant - log_queries
- twilio_autopilot_task_sample - source_channel
- twilio_autopilot_webhook - webhook_method
- twilio_chat_channel - attributes, type
- twilio_chat_channel_member - attributes
- twilio_chat_channel_studio_webhook - retry_count
- twilio_chat_channel_trigger_webhook - retry_count, method
- twilio_chat_channel_webhook - retry_count, method
- twilio_chat_service - post_webhook_retry_count, pre_webhook_retry_count, webhook_method, reachability_enabled, read_status_enabled, typing_indicator_timeout, limits.channel_members, limits.user_channels, notifications.log_enabled, new_message.enabled, new_message.badge_count_enabled, added_to_channel.enabled, invited_to_channel.enabled, removed_from_channel.enabled
- twilio_chat_user - attributes
- twilio_conversations_conversation - attributes, state
- twilio_conversations_conversation_trigger_webhook - method
- twilio_conversations_conversation_webhook - method
- twilio_conversations_push_credential_apn - sandbox
- twilio_conversations_service_configuration - reachability_enabled
- twilio_conversations_service_notification - log_enabled, new_message.enabled, new_message.badge_count_enabled, added_to_conversation.enabled, removed_from_conversation.enabled
- twilio_conversations_user - attributes
- twilio_flex_flow - enabled, retry_count, janitor_enabled, long_lived
- twilio_messaging_service - area_code_geomatch, fallback_method, fallback_to_long_code, inbound_method, mms_converter, smart_encoding, sticky_sender, validity_period
- twilio_phone_number - status_callback_method, messaging.method, messaging.fallback_method, voice.method, voice.fallback_method, voice.caller_id_lookup, fax.method, fax.fallback_method
- twilio_proxy_service - default_ttl, geo_match_level, number_selection_behavior
- twilio_serverless_service - include_credentials, ui_editable
- twilio_sip_ip_address - cidr_length_prefix
- twilio_sip_domain - sip_registration, secure, emergency_calling_enabled, voice_method, voice_fallback_method, voice_status_callback_method
- twilio_sip_trunking_trunk - cnam_lookup_enabled, recording.mode, recording.trim, secure and transfer_mode
- twilio_taskrouter_workspace - template, prioritize_queue_order
- twilio_taskrouter_workflow - task_reservation_timeout
- twilio_taskrouter_worker - attributes
- twilio_taskrouter_task queue - max_reserved_workers, task_order, target workers
- twilio_taskrouter_task channel - channel_optimized_routing
- twilio_video_composition_hook - status_callback_method, resolution, trim, format and enabled
- twilio_voice_queue - max size

For the new default values, please see the relevant documentation.

FIXES

- The TaskRouter workspace EventsFilter will now be null when an empty string is returned from the API on the `twilio_taskrouter_workspace` resource and data source. Thanks to @bobtfish for the fix
- Remove `event_callback_url` attribute from`taskrouter_task_queue` resource and data source as the value is not returned from the API **breaking change**
- Remove `chat_service_sid` attribute from`twilio_proxy_service` resource and data source as the value is not returned from the API **breaking change**
- The Integration block on the `twilio_flex_flow` resource is now mandatory **breaking change**
- A bug in the voice block in the `twilio_sip_domain` resource which caused the urls and methods to not be updated has been fixed. This bug was caused by the code referencing old argument names. The argument names have been updated and the appropriate test coverage has been added.
- The following attributes: `password` (`twilio_sip_credential`), `secret` (`twilio_conversations_push_credential_fcm`), `private_key` (`twilio_conversations_push_credential_apn`), `auth_token` (`provider`) and `api_secret` (`provider`) have now been marked sensitive

FEATURES

- **New Data Source:** `twilio_video_composition_hook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/video_composition_hook.md)
- **New Resource:** `twilio_video_composition_hook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/video_composition_hook.md)

## v0.8.2 (2021-03-21)

FIXES

- When destroying a deployment resource. The `build_sid` was not being superseded correctly, this causes the following error `Failed to delete serverless build: Cannot delete Build because it's part of an active Deployment.` to be thrown when attempting to destroy a build.
  This error was caused because the Terraform code was incorrectly comparing if the pointers for the environment `build_sid` and the deployment `build_sid` (which was created from the value that was held in Terraform state) were the same. This has since been updated to dereference the pointer for the environment build_sid, so the 2 values could be compared.
  This was not caught by the pre-existing acceptance tests because the tests tore down all resources, the environment was being deleted before the build so no active deployments were preventing the build from being deleted. Additional acceptance tests have been added to ensure this functionality works as expected.

## v0.8.1 (2021-03-13)

FIXES

- When using the lifecycle meta-argument `create_before_destroy` on the `twilio_serverless_deployment` resource and Terraform detected a change that forced a new resource to be created. Terraform created the new deployment and when the old deployment was being destroyed the provider was incorrectly removing the new build. Causing no active build to be deployed and causing an outage on the API. This has now been fixed by checking the environment to see if the build sid matches the active build on the environment. If they match the build will be superseded otherwise the state will be removed (as deployments can't be deleted)

## v0.8.0 (2021-02-19)

NOTES

- Add deprecation warning to Programmable Chat resources and data sources
- Add deprecation warning to `twilio_sip_trunking_phone_number` `sid` argument
- Improved error messages when account or API Key credential validation failures occur (this is part of Go SDK v0.14.1 release)

FIXES

- Programmable Chat Channel Member `LastConsumedTimestamp` was being set as a time.Time, this data type is incompatible with Terraform, so setting the value as a RFC3339 string if the value is not nil

FEATURES

- **New Data Source:** `twilio_sip_trunking_credential_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_credential_list.md)
- **New Data Source:** `twilio_sip_trunking_credential_lists` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_credential_lists.md)
- **New Data Source:** `twilio_sip_trunking_ip_access_control_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_ip_access_control_list.md)
- **New Data Source:** `twilio_sip_trunking_ip_access_control_lists` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_trunking_ip_access_control_lists.md)
- **New Resource:** `twilio_sip_trunking_credential_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_trunking_credential_list.md)
- **New Resource:** `twilio_sip_trunking_ip_access_control_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_trunking_ip_access_control_list.md)
- **New Data Source:** `twilio_sip_credential` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_credential.md)
- **New Data Source:** `twilio_sip_credentials` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_credentials.md)
- **New Data Source:** `twilio_sip_credential_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_credential_list.md)
- **New Data Source:** `twilio_sip_domain` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_domain.md)
- **New Data Source:** `twilio_sip_domain_ip_access_control_list_mapping` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_domain_ip_access_control_list_mapping.md)
- **New Data Source:** `twilio_sip_domain_ip_access_control_list_mappings` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_domain_ip_access_control_list_mappings.md)
- **New Data Source:** `twilio_sip_domain_credential_list_mapping` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_domain_credential_list_mapping.md)
- **New Data Source:** `twilio_sip_domain_credential_list_mappings` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_domain_credential_list_mappings.md)
- **New Data Source:** `twilio_sip_ip_access_control_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_ip_access_control_list.md)
- **New Data Source:** `twilio_sip_ip_address` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_ip_address.md)
- **New Data Source:** `twilio_sip_ip_addresses` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/sip_ip_addresses.md)
- **New Resource:** `twilio_sip_credential` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_credential.md)
- **New Resource:** `twilio_sip_credential_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_credential_list.md)
- **New Resource:** `twilio_sip_domain` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_domain.md)
- **New Resource:** `twilio_sip_domain_ip_access_control_list_mapping` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_domain_ip_access_control_list_mapping.md)
- **New Resource:** `twilio_sip_domain_credential_list_mapping` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_domain_credential_list_mapping.md)
- **New Resource:** `twilio_sip_ip_access_control_list` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_ip_access_control_list.md)
- **New Resource:** `twilio_sip_ip_address` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/sip_ip_address.md)
- **New Data Source:** `twilio_flex_plugin` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/flex_plugin.md)
- **New Data Source:** `twilio_flex_plugin_configuration` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/flex_plugin_configuration.md)
- **New Data Source:** `twilio_flex_plugin_release` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/flex_plugin_release.md)
- **New Resource:** `twilio_flex_plugin` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/flex_plugin.md)
- **New Resource:** `twilio_flex_plugin_configuration` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/flex_plugin_configuration.md)
- **New Resource:** `twilio_flex_plugin_release` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/flex_plugin_release.md)
- **Updated Resource:** `twilio_sip_trunking_phone_number` Add `phone_number_sid` argument and `phone_number_sid` attribute (to superseed `sid` argument)
- **Updated Resource:** `twilio_serverless_build` Add `runtime` argument and `runtime` attribute
- **Updated Data Source:** `twilio_serverless_build` Add `runtime` attribute
- **Updated Data Source:** `twilio_serverless_builds` Add `runtime` attribute

## v0.7.1 (2021-02-08)

Update conversation documentation to remove typos and improve the readability of the documentation

## v0.7.0 (2021-02-08)

FEATURES

- **New Data Source:** `twilio_conversations_configuration` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_configuration.md)
- **New Data Source:** `twilio_conversations_conversation_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_conversation_webhook.md)
- **New Data Source:** `twilio_conversations_conversation_webhooks` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_conversation_webhooks.md)
- **New Data Source:** `twilio_conversations_conversation` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_conversation.md)
- **New Data Source:** `twilio_conversations_conversations` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_conversations.md)
- **New Data Source:** `twilio_conversations_role` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_role.md)
- **New Data Source:** `twilio_conversations_roles` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_roles.md)
- **New Data Source:** `twilio_conversations_service_configuration` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_service_configuration.md)
- **New Data Source:** `twilio_conversations_service_notification` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_service_notification.md)
- **New Data Source:** `twilio_conversations_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_service.md)
- **New Data Source:** `twilio_conversations_user` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_user.md)
- **New Data Source:** `twilio_conversations_users` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_users.md)
- **New Data Source:** `twilio_conversations_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/data-sources/conversations_webhook.md)
- **New Resource:** `twilio_conversations_configuration` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_configuration.md)
- **New Resource:** `twilio_conversations_conversation_studio_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_conversation_studio_webhook.md)
- **New Resource:** `twilio_conversations_conversation_trigger_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_conversation_trigger_webhook.md)
- **New Resource:** `twilio_conversations_conversation_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_conversation_webhook.md)
- **New Resource:** `twilio_conversations_conversation` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_conversation.md)
- **New Resource:** `twilio_conversations_push_credential_apn` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_push_credential_apn.md)
- **New Resource:** `twilio_conversations_push_credential_fcm` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_push_credential_fcm.md)
- **New Resource:** `twilio_conversations_role` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_role.md)
- **New Resource:** `twilio_conversations_service_configuration` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_service_configuration.md)
- **New Resource:** `twilio_conversations_service_notification` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_service_notification.md)
- **New Resource:** `twilio_conversations_service` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_service.md)
- **New Resource:** `twilio_conversations_user` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_user.md)
- **New Resource:** `twilio_conversations_webhook` [docs](https://github.com/RJPearson94/terraform-provider-twilio/blob/main/docs/resources/conversations_webhook.md)

## v0.6.2 (2021-01-22)

Correct bug when setting VoiceReceiveMode on twilio phone number resource. The voice receive mode was incorrectly being set to the voice and fax state instead of the corresponding string value on both create and update. This has now been corrected

Correct some documentation on the terraform registry

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
