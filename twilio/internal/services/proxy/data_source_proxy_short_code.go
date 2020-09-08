package proxy

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/proxy/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceProxyShortCode() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxyShortCodeRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_reserved": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"capabilities": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_inbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"fax_outbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms_inbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms_outbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_fax_domestic": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_mms_domestic": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_sms_domestic": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_voice_domestic": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sip_trunking": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sms_inbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sms_outbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice_inbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice_outbound": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"short_code": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_country": &schema.Schema{
				Type:     schema.TypeString,
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

func dataSourceProxyShortCodeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).ShortCode(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Short code with sid (%s) was not found for proxy service with sid (%s)", sid, serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read proxy short code resource: %s", err.Error())
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
