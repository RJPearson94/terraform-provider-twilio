package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsWebhook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsWebhookRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this webhook",
			},
			"target": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target of the webhook",
			},
			"method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP method for the webhook",
			},
			"pre_webhook_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL called before an event is sent to the webhook",
			},
			"post_webhook_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL called after an event is sent to the webhook",
			},
			"filters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of webhook event triggers subscribed to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the webhook resource",
			},
		},
	}
}

func dataSourceConversationsWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Webhook().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Conversation webhook was not found")
		}
		return diag.Errorf("Failed to read conversation webhook: %s", err.Error())
	}

	d.SetId(getResponse.AccountSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("target", getResponse.Target)
	d.Set("pre_webhook_url", getResponse.PreWebhookURL)
	d.Set("post_webhook_url", getResponse.PostWebhookURL)
	d.Set("filters", getResponse.Filters)
	d.Set("method", getResponse.Method)
	d.Set("url", getResponse.URL)

	return nil
}
