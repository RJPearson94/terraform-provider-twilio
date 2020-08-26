package chat

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/services"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceChatService() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatServiceCreate,
		Read:   resourceChatServiceRead,
		Update: resourceChatServiceUpdate,
		Delete: resourceChatServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_channel_creator_role_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_channel_role_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_service_role_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"limits": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_members": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"user_channels": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"media": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compatibility_message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"size_limit_mb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"notifications": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"new_message": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"badge_count_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"added_to_channel": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"removed_from_channel": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"invited_to_channel": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"post_webhook_retry_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"post_webhook_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_webhook_retry_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pre_webhook_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_filters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"webhook_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"POST",
					"GET",
				}, false),
			},
			"reachability_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"read_status_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"typing_indicator_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceChatServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &services.CreateServiceInput{
		FriendlyName: d.Get("friendly_name").(string),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create chat service: %s", err)
	}

	d.SetId(createResult.Sid)

	log.Println("[INFO] Only the friendly name can be set on the creation of a chat service so updating resource to add the additional config")
	return resourceChatServiceUpdate(d, meta)
}

func resourceChatServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read chat service: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("default_channel_creator_role_sid", getResponse.DefaultChannelCreatorRoleSid)
	d.Set("default_channel_role_sid", getResponse.DefaultChannelRoleSid)
	d.Set("default_service_role_sid", getResponse.DefaultServiceRoleSid)
	d.Set("limits", flatternLimits(getResponse.Limits))
	d.Set("media", flatternMedia(getResponse.Media))
	d.Set("notifications", flatternNotifications(getResponse.Notifications))
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

func resourceChatServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &service.UpdateServiceInput{
		FriendlyName:           utils.OptionalString(d, "friendly_name"),
		ReadStatusEnabled:      utils.OptionalBool(d, "read_status_enabled"),
		TypingIndicatorTimeout: utils.OptionalInt(d, "typing_indicator_timeout"),
		PostWebhookURL:         utils.OptionalString(d, "post_webhook_url"),
		PostWebhookRetryCount:  utils.OptionalInt(d, "post_webhook_retry_count"),
		PreWebhookURL:          utils.OptionalString(d, "pre_webhook_url"),
		PreWebhookRetryCount:   utils.OptionalInt(d, "pre_webhook_retry_count"),
		WebhookMethod:          utils.OptionalString(d, "webhook_method"),
	}

	if _, ok := d.GetOk("notifications"); ok {
		updateInput.NotificationsLogEnabled = utils.OptionalBool(d, "notifications.0.log_enabled")

		if _, ok := d.GetOk("notifications.0.new_message"); ok {
			updateInput.NotificationsNewMessageEnabled = utils.OptionalBool(d, "notifications.0.new_message.0.enabled")
			updateInput.NotificationsNewMessageTemplate = utils.OptionalString(d, "notifications.0.new_message.0.template")
			updateInput.NotificationsNewMessageSound = utils.OptionalString(d, "notifications.0.new_message.0.sound")
			updateInput.NotificationsNewMessageBadgeCountEnabled = utils.OptionalBool(d, "notifications.0.new_message.0.badge_count_enabled")
		}

		if _, ok := d.GetOk("notifications.0.added_to_channel"); ok {
			updateInput.NotificationsAddedToChannelEnabled = utils.OptionalBool(d, "notifications.0.added_to_channel.0.enabled")
			updateInput.NotificationsAddedToChannelTemplate = utils.OptionalString(d, "notifications.0.added_to_channel.0.template")
			updateInput.NotificationsAddedToChannelSound = utils.OptionalString(d, "notifications.0.added_to_channel.0.sound")
		}

		if _, ok := d.GetOk("notifications.0.removed_from_channel"); ok {
			updateInput.NotificationsRemovedFromChannelEnabled = utils.OptionalBool(d, "notifications.0.removed_from_channel.0.enabled")
			updateInput.NotificationsRemovedFromChannelTemplate = utils.OptionalString(d, "notifications.0.removed_from_channel.0.template")
			updateInput.NotificationsRemovedFromChannelSound = utils.OptionalString(d, "notifications.0.removed_from_channel.0.sound")
		}

		if _, ok := d.GetOk("notifications.0.invited_to_channel"); ok {
			updateInput.NotificationsInvitedToChannelEnabled = utils.OptionalBool(d, "notifications.0.invited_to_channel.0.enabled")
			updateInput.NotificationsInvitedToChannelTemplate = utils.OptionalString(d, "notifications.0.invited_to_channel.0.template")
			updateInput.NotificationsInvitedToChannelSound = utils.OptionalString(d, "notifications.0.invited_to_channel.0.sound")
		}
	}

	if _, ok := d.GetOk("limits"); ok {
		updateInput.LimitsChannelMembers = utils.OptionalInt(d, "limits.0.channel_members")
		updateInput.LimitsUserChannels = utils.OptionalInt(d, "limits.0.user_channels")
	}

	if _, ok := d.GetOk("media"); ok {
		updateInput.MediaCompatibilityMessage = utils.OptionalString(d, "media.0.compatibility_message")
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update chat service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatServiceRead(d, meta)
}

func resourceChatServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete chat service: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func flatternNotifications(input map[string]interface{}) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["added_to_channel"] = flatternChannelUserModification(input["added_to_channel"].(map[string]interface{}))
	result["invited_to_channel"] = flatternChannelUserModification(input["invited_to_channel"].(map[string]interface{}))
	result["removed_from_channel"] = flatternChannelUserModification(input["removed_from_channel"].(map[string]interface{}))
	result["new_message"] = flatternNewMesage(input["new_message"].(map[string]interface{}))
	result["log_enabled"] = input["log_enabled"]
	results = append(results, result)

	return &results
}

func flatternChannelUserModification(input map[string]interface{}) *[]interface{} {
	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["enabled"] = input["enabled"]
	result["template"] = input["template"]
	result["sound"] = input["sound"]

	results = append(results, result)
	return &results
}

func flatternNewMesage(input map[string]interface{}) *[]interface{} {
	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["enabled"] = input["enabled"]
	result["template"] = input["template"]
	result["sound"] = input["sound"]
	result["badge_count_enabled"] = input["badge_count_enabled"]

	results = append(results, result)
	return &results
}

func flatternLimits(input map[string]interface{}) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["user_channels"] = input["user_channels"]
	result["channel_members"] = input["channel_members"]

	results = append(results, result)
	return &results
}

func flatternMedia(input map[string]interface{}) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["compatibility_message"] = input["compatibility_message"]
	result["size_limit_mb"] = input["size_limit_mb"]

	results = append(results, result)
	return &results
}
