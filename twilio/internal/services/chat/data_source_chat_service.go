package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/chat/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceChatService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceChatServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_members": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"user_channels": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"media": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compatibility_message": {
							Type:     schema.TypeString,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"new_message": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"badge_count_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"added_to_channel": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"removed_from_channel": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"invited_to_channel": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"template": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sound": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"post_webhook_retry_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"post_webhook_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pre_webhook_retry_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pre_webhook_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhook_filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"webhook_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reachability_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"read_status_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"typing_indicator_timeout": {
				Type:     schema.TypeInt,
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

func dataSourceChatServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Chat service with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read chat service: %s", err)
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
