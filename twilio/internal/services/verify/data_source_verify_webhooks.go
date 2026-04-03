package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhooks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVerifyWebhooks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyWebhooksRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
				Description:  "The SID of the Verify service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these webhooks",
			},
			"webhooks": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of webhooks for the Verify service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this webhook by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the webhook",
						},
						"event_types": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The list of events that trigger the webhook",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the webhook",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The webhook version",
						},
						"webhook_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTPS URL that Twilio calls when an event occurs",
						},
						"webhook_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used when calling the webhook URL",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the webhook was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the webhook was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the webhook resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceVerifyWebhooksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.Verify

	options := &webhooks.WebhooksPageOptions{}

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Webhooks.NewWebhooksPaginatorWithOptions(options)
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No webhooks were found for Verify service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list Verify webhooks: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("account_sid", twilioClient.AccountSid)
	d.Set("service_sid", serviceSid)

	webhooks := make([]interface{}, 0)

	for _, webhook := range paginator.Webhooks {
		webhookMap := make(map[string]interface{})

		webhookMap["sid"] = webhook.Sid
		webhookMap["friendly_name"] = webhook.FriendlyName
		webhookMap["event_types"] = webhook.EventTypes
		webhookMap["status"] = webhook.Status
		webhookMap["version"] = webhook.Version
		webhookMap["webhook_url"] = webhook.WebhookURL
		webhookMap["webhook_method"] = webhook.WebhookMethod
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
