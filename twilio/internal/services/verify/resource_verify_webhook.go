package verify

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/webhooks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVerifyWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVerifyWebhookCreate,
		ReadContext:   resourceVerifyWebhookRead,
		UpdateContext: resourceVerifyWebhookUpdate,
		DeleteContext: resourceVerifyWebhookDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Webhooks/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this webhook by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this webhook",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
				Description:  "The SID of the Verify service. Changing this forces a new resource",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A human-readable label for the webhook",
			},
			"event_types": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The list of events that trigger the webhook. Valid values: `*`, `factor.created`, `factor.verified`, `factor.deleted`, `challenge.approved`, `challenge.denied`",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						// All
						"*",

						// Factor
						"factor.created",
						"factor.verified",
						"factor.deleted",

						// Challenge
						"challenge.approved",
						"challenge.denied",
					}, false),
				},
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"enabled",
					"disabled",
				}, false),
				Default:     "enabled",
				Description: "The status of the webhook. Valid values: `enabled`, `disabled`. Defaults to `enabled`",
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"v1",
					"v2",
				}, false),
				Default:     "v2",
				Description: "The webhook version. Valid values: `v1`, `v2`. Defaults to `v2`",
			},
			"webhook_url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsURLWithHTTPS,
				Description:  "The HTTPS URL that Twilio calls when an event occurs",
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

func resourceVerifyWebhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	createInput := &webhooks.CreateWebhookInput{
		FriendlyName: d.Get("friendly_name").(string),
		WebhookURL:   d.Get("webhook_url").(string),
		EventTypes:   utils.ConvertToStringSlice(d.Get("event_types").([]interface{})),
		Status:       utils.OptionalString(d, "status"),
		Version:      utils.OptionalString(d, "version"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Webhooks.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create webhook: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceVerifyWebhookRead(ctx, d, meta)
}

func resourceVerifyWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	getResponse, err := client.Service(d.Get("service_sid").(string)).Webhook(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read webhook: %s", err.Error())
	}

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

func resourceVerifyWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	updateInput := &webhook.UpdateWebhookInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		WebhookURL:   utils.OptionalString(d, "webhook_url"),
		EventTypes:   utils.OptionalStringSlice(d, "event_types"),
		Status:       utils.OptionalString(d, "status"),
		Version:      utils.OptionalString(d, "version"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Webhook(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update webhook: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceVerifyWebhookRead(ctx, d, meta)
}

func resourceVerifyWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	if err := client.Service(d.Get("service_sid").(string)).Webhook(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete webhook: %s", err.Error())
	}

	d.SetId("")
	return nil
}
