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

func dataSourceProxyPhoneNumber() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProxyPhoneNumberRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.PhoneNumberSidValidation(),
				Description:  "The SID of the Proxy phone number to read",
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
				Description: "The SID of the account that owns this Proxy phone number",
			},
			"phone_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phone number in E.164 format",
			},
			"is_reserved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the phone number is reserved and not assigned to a Proxy session",
			},
			"capabilities": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The capabilities of the Proxy phone number",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive faxes",
						},
						"fax_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can send faxes",
						},
						"mms_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive MMS messages",
						},
						"mms_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can send MMS messages",
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
							Description: "Whether the phone number supports SIP trunking",
						},
						"sms_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive SMS messages",
						},
						"sms_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can send SMS messages",
						},
						"voice_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive voice calls",
						},
						"voice_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can make voice calls",
						},
					},
				},
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the Proxy phone number",
			},
			"iso_country": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ISO country code of the phone number",
			},
			"in_use": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of active Proxy sessions assigned to this phone number",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy phone number was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy phone number was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the Proxy phone number resource",
			},
		},
	}
}

func dataSourceProxyPhoneNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).PhoneNumber(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Phone number with sid (%s) was not found for proxy service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read proxy phone number resource: %s", err.Error())
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
