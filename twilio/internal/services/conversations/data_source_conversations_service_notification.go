package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/conversations/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsServiceNotification() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsServiceNotificationRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this service notification",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service",
			},
			"new_message": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Notification settings for new messages",
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
							Description: "The template for new message notifications",
						},
						"sound": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The sound to play for new message notifications",
						},
						"badge_count_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether badge count is enabled for new message notifications",
						},
					},
				},
			},
			"added_to_conversation": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Notification settings for when a user is added to a conversation",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether added-to-conversation notifications are enabled",
						},
						"template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The template for added-to-conversation notifications",
						},
						"sound": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The sound to play for added-to-conversation notifications",
						},
					},
				},
			},
			"removed_from_conversation": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Notification settings for when a user is removed from a conversation",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether removed-from-conversation notifications are enabled",
						},
						"template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The template for removed-from-conversation notifications",
						},
						"sound": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The sound to play for removed-from-conversation notifications",
						},
					},
				},
			},
			"log_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether notification logging is enabled for the service",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the service notification resource",
			},
		},
	}
}

func dataSourceConversationsServiceNotificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	getResponse, err := client.Service(serviceSid).Configuration().Notification().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Conversation notification was not found for service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to read conversations service notification: %s", err.Error())
	}

	d.SetId(getResponse.ChatServiceSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("log_enabled", getResponse.LogEnabled)
	d.Set("new_message", helper.FlattenNotificationsNewMessage(getResponse.NewMessage))
	d.Set("added_to_conversation", helper.FlattenNotificationsAction(getResponse.AddedToConversation))
	d.Set("removed_from_conversation", helper.FlattenNotificationsAction(getResponse.RemovedFromConversation))
	d.Set("url", getResponse.URL)

	return nil
}
