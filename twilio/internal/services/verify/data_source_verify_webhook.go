package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVerifyWebhook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyWebhookRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyWebhookSidValidation(),
				Description:  "The SID of the webhook to fetch",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
				Description:  "The SID of the Verify service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this webhook",
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
	}
}

func dataSourceVerifyWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	sid := d.Get("sid").(string)
	serviceSid := d.Get("service_sid").(string)
	getResponse, err := client.Service(serviceSid).Webhook(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Webhook with sid (%s) was not found for Verify service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read Verify webhook: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("event_types", getResponse.EventTypes)
	d.Set("status", getResponse.Status)
	d.Set("version", getResponse.Version)
	d.Set("webhook_url", getResponse.WebhookURL)
	d.Set("webhook_method", getResponse.WebhookMethod)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
