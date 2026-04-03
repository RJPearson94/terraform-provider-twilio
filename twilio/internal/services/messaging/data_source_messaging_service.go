package messaging

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMessagingService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMessagingServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
				Description:  "The SID of the messaging service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this messaging service",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the messaging service",
			},
			"area_code_geomatch": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether area code geomatch is enabled on the messaging service",
			},
			"fallback_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP method used to call the fallback URL",
			},
			"fallback_to_long_code": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether fallback to long code is enabled for the messaging service",
			},
			"fallback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to call when an inbound message error is received",
			},
			"inbound_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP method used to call the inbound request URL",
			},
			"inbound_request_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to call when a message is received by any phone number or short code in the messaging service",
			},
			"mms_converter": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the MMS converter is enabled for the messaging service",
			},
			"smart_encoding": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether smart encoding is enabled for the messaging service",
			},
			"status_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to call when a message status change event is received",
			},
			"sticky_sender": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether sticky sender is enabled on the messaging service",
			},
			"use_inbound_webhook_on_number": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the inbound webhook on the phone number is used for incoming messages",
			},
			"validity_period": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of seconds that the messaging service will keep a message in the sending queue before it is considered failed",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the messaging service was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the messaging service was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the messaging service resource",
			},
		},
	}
}

func dataSourceMessagingServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Messaging service with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read messaging service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("area_code_geomatch", getResponse.AreaCodeGeomatch)
	d.Set("fallback_method", getResponse.FallbackMethod)
	d.Set("fallback_to_long_code", getResponse.FallbackToLongCode)
	d.Set("fallback_url", getResponse.FallbackURL)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("inbound_method", getResponse.InboundMethod)
	d.Set("inbound_request_url", getResponse.InboundRequestURL)
	d.Set("mms_converter", getResponse.MmsConverter)
	d.Set("smart_encoding", getResponse.SmartEncoding)
	d.Set("status_callback_url", getResponse.StatusCallback)
	d.Set("sticky_sender", getResponse.StickySender)
	d.Set("use_inbound_webhook_on_number", getResponse.UseInboundWebhookOnNumber)
	d.Set("validity_period", getResponse.ValidityPeriod)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
