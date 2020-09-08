package proxy

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/proxy/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProxyShortCode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProxyShortCodeRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_reserved": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"capabilities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"fax_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_fax_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_mms_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_sms_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_voice_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sip_trunking": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sms_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sms_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"short_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_country": {
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
	}
}

func dataSourceProxyShortCodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).ShortCode(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Short code with sid (%s) was not found for proxy service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read proxy short code resource: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("short_code", getResponse.ShortCode)
	d.Set("iso_country", getResponse.IsoCountry)
	d.Set("is_reserved", getResponse.IsReserved)
	d.Set("capabilities", helper.FlattenShortCodeCapabilities(getResponse.Capabilities))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
