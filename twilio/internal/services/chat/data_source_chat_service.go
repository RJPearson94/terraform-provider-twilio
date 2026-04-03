package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/chat/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatService() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatServiceSidValidation(),
				Description:  "The SID of the chat service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this chat service",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the chat service",
			},
			"default_channel_creator_role_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default role assigned to the creator of a channel",
			},
			"default_channel_role_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default role assigned to members of a channel",
			},
			"default_service_role_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default role assigned to users of the chat service",
			},
			"limits": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The limits configuration for the chat service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_members": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of members allowed in a channel",
						},
						"user_channels": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of channels a user can join",
						},
					},
				},
			},
			"media": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The media configuration for the chat service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compatibility_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message to display when a media message has no text fallback",
						},
						"size_limit_mb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum media file size in megabytes",
						},
					},
				},
			},
			"notifications": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The notification configuration for the chat service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether notification logging is enabled",
						},
						"new_message": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The notification settings for new messages",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether new message notifications are enabled",
									},
									"template": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The notification template for new messages",
									},
									"sound": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The sound to play for new message notifications",
									},
									"badge_count_enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the badge count is enabled for new message notifications",
									},
								},
							},
						},
						"added_to_channel": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The notification settings for being added to a channel",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether added-to-channel notifications are enabled",
									},
									"template": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The notification template for being added to a channel",
									},
									"sound": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The sound to play for added-to-channel notifications",
									},
								},
							},
						},
						"removed_from_channel": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The notification settings for being removed from a channel",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether removed-from-channel notifications are enabled",
									},
									"template": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The notification template for being removed from a channel",
									},
									"sound": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The sound to play for removed-from-channel notifications",
									},
								},
							},
						},
						"invited_to_channel": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The notification settings for being invited to a channel",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether invited-to-channel notifications are enabled",
									},
									"template": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The notification template for being invited to a channel",
									},
									"sound": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The sound to play for invited-to-channel notifications",
									},
								},
							},
						},
					},
				},
			},
			"post_webhook_retry_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of retry attempts for post-event webhook requests",
			},
			"post_webhook_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for post-event webhook requests",
			},
			"pre_webhook_retry_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of retry attempts for pre-event webhook requests",
			},
			"pre_webhook_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for pre-event webhook requests",
			},
			"webhook_filters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of webhook event triggers subscribed to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"webhook_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP method used for webhook requests",
			},
			"reachability_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the reachability indicator is enabled for the chat service",
			},
			"read_status_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the message read status feature is enabled",
			},
			"typing_indicator_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The duration in seconds after which a typing indicator times out",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the chat service was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the chat service was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the chat service resource",
			},
		},
	}
}

func dataSourceChatServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Chat service with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read chat service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("default_channel_creator_role_sid", getResponse.DefaultChannelCreatorRoleSid)
	d.Set("default_channel_role_sid", getResponse.DefaultChannelRoleSid)
	d.Set("default_service_role_sid", getResponse.DefaultServiceRoleSid)
	d.Set("limits", helper.FlattenLimits(getResponse.Limits))
	d.Set("media", helper.FlattenMedia(getResponse.Media))
	d.Set("notifications", helper.FlattenNotifications(getResponse.Notifications))
	d.Set("post_webhook_retry_count", getResponse.PostWebhookRetryCount)
	d.Set("post_webhook_url", getResponse.PostWebhookURL)
	d.Set("pre_webhook_retry_count", getResponse.PreWebhookRetryCount)
	d.Set("pre_webhook_url", getResponse.PreWebhookURL)
	d.Set("reachability_enabled", getResponse.ReachabilityEnabled)
	d.Set("read_status_enabled", getResponse.ReadStatusEnabled)
	d.Set("typing_indicator_timeout", getResponse.TypingIndicatorTimeout)
	d.Set("webhook_filters", getResponse.WebhookFilters)
	d.Set("webhook_method", getResponse.WebhookMethod)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
