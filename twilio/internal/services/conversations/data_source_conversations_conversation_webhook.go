package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsConversationWebhook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsConversationWebhookRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationWebhookSidValidation(),
				Description:  "The SID of the conversation webhook to retrieve",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this conversation webhook",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service",
			},
			"conversation_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationSidValidation(),
				Description:  "The SID of the conversation",
			},
			"target": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target of the conversation webhook",
			},
			"configuration": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The configuration of the conversation webhook",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method for the webhook",
						},
						"webhook_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to send webhook requests to",
						},
						"filters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of webhook event filters",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"triggers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of keywords that trigger the webhook",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"replay_after": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The message index to replay messages from",
						},
						"flow_sid": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The SID of the Studio flow",
						},
					},
				},
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the conversation webhook was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the conversation webhook was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the conversation webhook resource",
			},
		},
	}
}

func dataSourceConversationsConversationWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).Conversation(conversationSid).Webhook(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Conversation webhook with sid (%s) was not found for service with sid (%s) and conversation with (%s)", sid, serviceSid, conversationSid)
		}
		return diag.Errorf("Failed to read conversation webhook: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("conversation_sid", getResponse.ConversationSid)
	d.Set("target", getResponse.Target)
	d.Set("configuration", &[]interface{}{
		map[string]interface{}{
			"webhook_url":  getResponse.Configuration.URL,
			"method":       getResponse.Configuration.Method,
			"replay_after": getResponse.Configuration.ReplayAfter,
			"triggers":     getResponse.Configuration.Triggers,
			"flow_sid":     getResponse.Configuration.FlowSid,
			"filters":      getResponse.Configuration.Filters,
		},
	})
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
