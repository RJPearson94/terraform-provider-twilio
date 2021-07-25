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
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"area_code_geomatch": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fallback_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fallback_to_long_code": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fallback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbound_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbound_request_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mms_converter": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"smart_encoding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status_callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sticky_sender": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_inbound_webhook_on_number": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"validity_period": {
				Type:     schema.TypeInt,
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
