package autopilot

import (
	"context"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAutopilotWebhooks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotWebhooksRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
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
						"unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"events": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"webhook_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"webhook_method": {
							Type:     schema.TypeString,
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
				},
			},
		},
	}
}

func dataSourceAutopilotWebhooksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	assistantSid := d.Get("assistant_sid").(string)
	paginator := client.Assistant(assistantSid).Webhooks.NewWebhooksPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No webhooks were found for assistant with sid (%s)", assistantSid)
		}
		return diag.Errorf("Failed to list autopilot webhooks: %s", err.Error())
	}

	d.SetId(assistantSid)
	d.Set("assistant_sid", assistantSid)

	webhooks := make([]interface{}, 0)

	for _, webhook := range paginator.Webhooks {
		d.Set("account_sid", webhook.AccountSid)

		webhookMap := make(map[string]interface{})

		webhookMap["sid"] = webhook.Sid
		webhookMap["unique_name"] = webhook.UniqueName
		webhookMap["webhook_url"] = webhook.WebhookURL
		webhookMap["webhook_method"] = webhook.WebhookMethod
		webhookMap["events"] = strings.Split(webhook.Events, " ")
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
