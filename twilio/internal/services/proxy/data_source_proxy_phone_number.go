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

func dataSourceProxyPhoneNumber() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxyPhoneNumberRead,

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
			"phone_number": &schema.Schema{
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
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_country": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"in_use": &schema.Schema{
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

func dataSourceProxyPhoneNumberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).PhoneNumber(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Phone number with sid (%s) was not found for proxy service with sid (%s)", sid, serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read proxy phone number resource: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("phone_number", getResponse.PhoneNumber)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("iso_country", getResponse.IsoCountry)
	d.Set("is_reserved", getResponse.IsReserved)
	d.Set("capabilities", helper.FlattenPhoneNumberCapabilities(getResponse.Capabilities))
	d.Set("in_use", getResponse.InUse)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
