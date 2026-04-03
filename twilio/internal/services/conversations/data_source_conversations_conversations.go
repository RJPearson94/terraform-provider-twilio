package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsConversations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsConversationsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these conversations",
			},
			"conversations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of conversations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this conversation by Twilio",
						},
						"unique_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique, developer-assigned name for the conversation",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the conversation",
						},
						"attributes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A JSON string of attributes associated with the conversation",
						},
						"messaging_service_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the messaging service associated with the conversation",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the conversation",
						},
						"timers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Timer settings for the conversation",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"date_closed": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The date and time the conversation will be closed, in RFC 3339 format",
									},
									"date_inactive": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The date and time the conversation will be marked as inactive, in RFC 3339 format",
									},
								},
							},
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the conversation was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the conversation was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the conversation resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceConversationsConversationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Conversations.NewConversationsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No conversations were found for conversations service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list conversations: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	conversations := make([]interface{}, 0)

	for _, conversation := range paginator.Conversations {
		d.Set("account_sid", conversation.AccountSid)

		conversationMap := make(map[string]interface{})

		conversationMap["sid"] = conversation.Sid
		conversationMap["unique_name"] = conversation.UniqueName
		conversationMap["friendly_name"] = conversation.FriendlyName
		conversationMap["attributes"] = conversation.Attributes
		conversationMap["messaging_service_sid"] = conversation.MessagingServiceSid
		conversationMap["state"] = conversation.State

		timerMap := make(map[string]interface{}, 0)

		if conversation.Timers.DateClosed != nil {
			timerMap["date_closed"] = conversation.Timers.DateClosed.Format(time.RFC3339)
		}

		if conversation.Timers.DateInactive != nil {
			timerMap["date_inactive"] = conversation.Timers.DateInactive.Format(time.RFC3339)
		}

		conversationMap["timers"] = &[]interface{}{
			timerMap,
		}
		conversationMap["date_created"] = conversation.DateCreated.Format(time.RFC3339)

		if conversation.DateUpdated != nil {
			conversationMap["date_updated"] = conversation.DateUpdated.Format(time.RFC3339)
		}

		conversationMap["url"] = conversation.URL

		conversations = append(conversations, conversationMap)
	}

	d.Set("conversations", &conversations)

	return nil
}
