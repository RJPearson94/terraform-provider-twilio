package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsConversation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsConversationRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationSidValidation(),
				Description:  "The SID of the conversation to retrieve",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this conversation",
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
	}
}

func dataSourceConversationsConversationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).Conversation(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Conversation with sid (%s) was not found for service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read conversations conversation: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("attributes", getResponse.Attributes)
	d.Set("messaging_service_sid", getResponse.MessagingServiceSid)
	d.Set("state", getResponse.State)

	timerMap := make(map[string]interface{}, 0)
	if getResponse.Timers.DateClosed != nil {
		timerMap["date_closed"] = getResponse.Timers.DateClosed.Format(time.RFC3339)
	}
	if getResponse.Timers.DateInactive != nil {
		timerMap["date_inactive"] = getResponse.Timers.DateInactive.Format(time.RFC3339)
	}
	d.Set("timers", &[]interface{}{
		timerMap,
	})

	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
