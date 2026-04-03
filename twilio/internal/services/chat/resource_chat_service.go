package chat

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/chat/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/services"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChatService() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		CreateContext: resourceChatServiceCreate,
		ReadContext:   resourceChatServiceRead,
		UpdateContext: resourceChatServiceUpdate,
		DeleteContext: resourceChatServiceDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this chat service by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this chat service",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 256),
				Description:  "A human-readable label for the chat service",
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
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The limits configuration for the chat service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_members": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      100,
							ValidateFunc: validation.IntBetween(1, 1000),
							Description:  "The maximum number of members allowed in a channel. Defaults to `100`",
						},
						"user_channels": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      250,
							ValidateFunc: validation.IntBetween(1, 1000),
							Description:  "The maximum number of channels a user can join. Defaults to `250`",
						},
					},
				},
			},
			"media": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The media configuration for the chat service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compatibility_message": {
							Type:        schema.TypeString,
							Optional:    true,
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
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The notification configuration for the chat service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether notification logging is enabled. Defaults to `false`",
						},
						"new_message": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The notification settings for new messages",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether new message notifications are enabled. Defaults to `false`",
									},
									"template": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The notification template for new messages",
									},
									"sound": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The sound to play for new message notifications",
									},
									"badge_count_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether the badge count is enabled for new message notifications. Defaults to `false`",
									},
								},
							},
						},
						"added_to_channel": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "The notification settings for being added to a channel",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether added-to-channel notifications are enabled. Defaults to `false`",
									},
									"template": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The notification template for being added to a channel",
									},
									"sound": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The sound to play for added-to-channel notifications",
									},
								},
							},
						},
						"removed_from_channel": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "The notification settings for being removed from a channel",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether removed-from-channel notifications are enabled. Defaults to `false`",
									},
									"template": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The notification template for being removed from a channel",
									},
									"sound": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The sound to play for removed-from-channel notifications",
									},
								},
							},
						},
						"invited_to_channel": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "The notification settings for being invited to a channel",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether invited-to-channel notifications are enabled. Defaults to `false`",
									},
									"template": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The notification template for being invited to a channel",
									},
									"sound": {
										Type:        schema.TypeString,
										Optional:    true,
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
				Optional:    true,
				Default:     0,
				Description: "The number of retry attempts for post-event webhook requests. Defaults to `0`",
			},
			"post_webhook_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL for post-event webhook requests",
			},
			"pre_webhook_retry_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The number of retry attempts for pre-event webhook requests. Defaults to `0`",
			},
			"pre_webhook_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL for pre-event webhook requests",
			},
			"webhook_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of webhook event triggers to subscribe to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"webhook_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "POST",
				ValidateFunc: validation.StringInSlice([]string{
					"POST",
					"GET",
				}, false),
				Description: "The HTTP method used for webhook requests. Valid values are `POST` or `GET`. Defaults to `POST`",
			},
			"reachability_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the reachability indicator is enabled for the chat service. Defaults to `false`",
			},
			"read_status_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the message read status feature is enabled. Defaults to `true`",
			},
			"typing_indicator_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
				Description: "The duration in seconds after which a typing indicator times out. Defaults to `5`",
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

func resourceChatServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	createInput := &services.CreateServiceInput{
		FriendlyName: d.Get("friendly_name").(string),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create chat service: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	log.Println("[INFO] Only the friendly name can be set on the creation of a chat service so updating resource to add the additional config")
	return resourceChatServiceUpdate(ctx, d, meta)
}

func resourceChatServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read chat service: %s", err.Error())
	}

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

func resourceChatServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	updateInput := &service.UpdateServiceInput{
		FriendlyName:           utils.OptionalString(d, "friendly_name"),
		ReadStatusEnabled:      utils.OptionalBool(d, "read_status_enabled"),
		ReachabilityEnabled:    utils.OptionalBool(d, "reachability_enabled"),
		TypingIndicatorTimeout: utils.OptionalInt(d, "typing_indicator_timeout"),
		PostWebhookURL:         utils.OptionalStringWithEmptyStringOnChange(d, "post_webhook_url"),
		PostWebhookRetryCount:  utils.OptionalIntWith0OnChange(d, "post_webhook_retry_count"),
		PreWebhookURL:          utils.OptionalStringWithEmptyStringOnChange(d, "pre_webhook_url"),
		PreWebhookRetryCount:   utils.OptionalIntWith0OnChange(d, "pre_webhook_retry_count"),
		WebhookMethod:          utils.OptionalString(d, "webhook_method"),
	}

	if _, ok := d.GetOk("notifications"); ok {
		notifications := &service.UpdateServiceNotificationsInput{
			LogEnabled: utils.OptionalBool(d, "notifications.0.log_enabled"),
		}

		if _, ok := d.GetOk("notifications.0.new_message"); ok {
			notifications.NewMessage = &service.UpdateServiceNotificationsNewMessageInput{
				Enabled:           utils.OptionalBool(d, "notifications.0.new_message.0.enabled"),
				Template:          utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.new_message.0.template"),
				Sound:             utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.new_message.0.sound"),
				BadgeCountEnabled: utils.OptionalBool(d, "notifications.0.new_message.0.badge_count_enabled"),
			}
		}

		if _, ok := d.GetOk("notifications.0.added_to_channel"); ok {
			notifications.AddedToChannel = &service.UpdateServiceNotificationsActionInput{
				Enabled:  utils.OptionalBool(d, "notifications.0.added_to_channel.0.enabled"),
				Template: utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.added_to_channel.0.template"),
				Sound:    utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.added_to_channel.0.sound"),
			}
		}

		if _, ok := d.GetOk("notifications.0.removed_from_channel"); ok {
			notifications.RemovedFromChannel = &service.UpdateServiceNotificationsActionInput{
				Enabled:  utils.OptionalBool(d, "notifications.0.removed_from_channel.0.enabled"),
				Template: utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.removed_from_channel.0.template"),
				Sound:    utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.removed_from_channel.0.sound"),
			}
		}

		if _, ok := d.GetOk("notifications.0.invited_to_channel"); ok {
			notifications.InvitedToChannel = &service.UpdateServiceNotificationsActionInput{
				Enabled:  utils.OptionalBool(d, "notifications.0.invited_to_channel.0.enabled"),
				Template: utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.invited_to_channel.0.template"),
				Sound:    utils.OptionalStringWithEmptyStringOnChange(d, "notifications.0.invited_to_channel.0.sound"),
			}
		}

		updateInput.Notifications = notifications
	}

	if _, ok := d.GetOk("limits"); ok {
		updateInput.Limits = &service.UpdateServiceLimitsInput{
			ChannelMembers: utils.OptionalInt(d, "limits.0.channel_members"),
			UserChannels:   utils.OptionalInt(d, "limits.0.user_channels"),
		}
	}

	if _, ok := d.GetOk("media"); ok {
		updateInput.Media = &service.UpdateServiceMediaInput{
			CompatibilityMessage: utils.OptionalString(d, "media.0.compatibility_message"),
		}
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update chat service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatServiceRead(ctx, d, meta)
}

func resourceChatServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete chat service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
