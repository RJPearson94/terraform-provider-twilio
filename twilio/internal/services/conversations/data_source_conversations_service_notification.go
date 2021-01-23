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
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
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
			"added_to_conversation": {
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
			"removed_from_conversation": {
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
			"log_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
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
