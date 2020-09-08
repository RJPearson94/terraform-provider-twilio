package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMessagingService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMessagingServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"area_code_geomatch": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fallback_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fallback_to_long_code": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fallback_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbound_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbound_request_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mms_converter": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"smart_encoding": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status_callback_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sticky_sender": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"validity_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMessagingServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Messaging service with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read messaging service: %s", err)
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
	d.Set("validity_period", getResponse.ValidityPeriod)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
