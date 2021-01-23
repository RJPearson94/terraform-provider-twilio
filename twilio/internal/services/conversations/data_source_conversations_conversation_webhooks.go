package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsConversationWebhooks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsConversationWebhooksRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
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
			"conversation_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"webhooks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"configuration": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"webhook_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"filters": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"triggers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"replay_after": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"flow_sid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
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
				},
			},
		},
	}
}

func dataSourceConversationsConversationWebhooksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	paginator := client.Service(serviceSid).Conversation(conversationSid).Webhooks.NewWebhooksPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No conversation webhooks were found for service with sid (%s) and conversation with (%s)", serviceSid, conversationSid)
		}
		return diag.Errorf("Failed to read conversation webhook: %s", err.Error())
	}

	d.SetId(serviceSid + "/" + conversationSid)
	d.Set("service_sid", serviceSid)
	d.Set("conversation_sid", conversationSid)

	webhooks := make([]interface{}, 0)

	for _, webhook := range paginator.Webhooks {
		d.Set("account_sid", webhook.AccountSid)

		webhookMap := make(map[string]interface{})

		webhookMap["sid"] = webhook.Sid
		webhookMap["target"] = webhook.Target
		webhookMap["configuration"] = &[]interface{}{
			map[string]interface{}{
				"webhook_url":  webhook.Configuration.URL,
				"method":       webhook.Configuration.Method,
				"replay_after": webhook.Configuration.ReplayAfter,
				"triggers":     webhook.Configuration.Triggers,
				"flow_sid":     webhook.Configuration.FlowSid,
				"filters":      webhook.Configuration.Filters,
			},
		}

		webhookMap["date_created"] = webhook.DateCreated.Format(time.RFC3339)

		if webhook.DateUpdated != nil {
			webhookMap["date_updated"] = webhook.DateUpdated.Format(time.RFC3339)
		}

		webhookMap["url"] = webhook.URL

		webhooks = append(webhooks, webhookMap)
	}

	d.Set("webhooks", &webhooks)

	return nil
}
