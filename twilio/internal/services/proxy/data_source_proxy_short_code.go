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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ShortCodeSidValidation(),
				Description:  "The SID of the Proxy short code to read",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ProxyServiceSidValidation(),
				Description:  "The SID of the Proxy service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Proxy short code",
			},
			"is_reserved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the short code is reserved and not assigned to a Proxy session",
			},
			"capabilities": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The capabilities of the Proxy short code",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can receive faxes",
						},
						"fax_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can send faxes",
						},
						"mms_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can receive MMS messages",
						},
						"mms_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can send MMS messages",
						},
						"restriction_fax_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether fax is restricted to domestic only",
						},
						"restriction_mms_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether MMS is restricted to domestic only",
						},
						"restriction_sms_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether SMS is restricted to domestic only",
						},
						"restriction_voice_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether voice is restricted to domestic only",
						},
						"sip_trunking": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code supports SIP trunking",
						},
						"sms_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can receive SMS messages",
						},
						"sms_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can send SMS messages",
						},
						"voice_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can receive voice calls",
						},
						"voice_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code can make voice calls",
						},
					},
				},
			},
			"short_code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The short code value",
			},
			"iso_country": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ISO country code of the short code",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy short code was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy short code was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the Proxy short code resource",
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
