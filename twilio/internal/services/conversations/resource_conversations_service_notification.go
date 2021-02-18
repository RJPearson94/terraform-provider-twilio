package conversations

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/conversations/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/configuration/notification"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConversationsServiceNotification() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConversationsServiceNotificationCreate,
		ReadContext:   resourceConversationsServiceNotificationRead,
		UpdateContext: resourceConversationsServiceNotificationUpdate,
		DeleteContext: resourceConversationsServiceNotificationDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"new_message": {
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
						"badge_count_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"added_to_conversation": {
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
			"removed_from_conversation": {
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
			"log_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceConversationsServiceNotificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Conversations service notification already exists so updating the notification
	return resourceConversationsServiceNotificationUpdate(ctx, d, meta)
}

func resourceConversationsServiceNotificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Service(d.Get("service_sid").(string)).Configuration().Notification().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read conversations service notification: %s", err.Error())
	}

	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("log_enabled", getResponse.LogEnabled)
	d.Set("new_message", helper.FlattenNotificationsNewMessage(getResponse.NewMessage))
	d.Set("added_to_conversation", helper.FlattenNotificationsAction(getResponse.AddedToConversation))
	d.Set("removed_from_conversation", helper.FlattenNotificationsAction(getResponse.RemovedFromConversation))
	d.Set("url", getResponse.URL)

	return nil
}

func resourceConversationsServiceNotificationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	updateInput := &notification.UpdateNotificationInput{
		LogEnabled: utils.OptionalBool(d, "log_enabled"),
	}

	if _, ok := d.GetOk("new_message"); ok {
		updateInput.NewMessage = &notification.UpdateNotificationNewMessageInput{
			Enabled:           utils.OptionalBool(d, "new_message.0.enabled"),
			Template:          utils.OptionalString(d, "new_message.0.template"),
			Sound:             utils.OptionalString(d, "new_message.0.sound"),
			BadgeCountEnabled: utils.OptionalBool(d, "new_message.0.badge_count_enabled"),
		}
	}

	if _, ok := d.GetOk("added_to_conversation"); ok {
		updateInput.AddedToConversation = &notification.UpdateNotificationConversationActionInput{
			Enabled:  utils.OptionalBool(d, "added_to_conversation.0.enabled"),
			Template: utils.OptionalString(d, "added_to_conversation.0.template"),
			Sound:    utils.OptionalString(d, "added_to_conversation.0.sound"),
		}
	}

	if _, ok := d.GetOk("removed_from_conversation"); ok {
		updateInput.RemovedFromConversation = &notification.UpdateNotificationConversationActionInput{
			Enabled:  utils.OptionalBool(d, "removed_from_conversation.0.enabled"),
			Template: utils.OptionalString(d, "removed_from_conversation.0.template"),
			Sound:    utils.OptionalString(d, "removed_from_conversation.0.sound"),
		}
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Configuration().Notification().UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update conversations service notification: %s", err.Error())
	}

	d.SetId(updateResp.ChatServiceSid)
	return resourceConversationsServiceNotificationRead(ctx, d, meta)
}

func resourceConversationsServiceNotificationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Conversations service notification cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}
