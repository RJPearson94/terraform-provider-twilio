package phone_number

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/phone_number/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePhoneNumber() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePhoneNumberRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_requirements": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"beta": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"capabilities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sms": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"emergency_address_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"emergency_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"messaging": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fallback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"method": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"trunk_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"voice": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"caller_id_lookup": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fallback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fax": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fallback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"identity_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bundle_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_callback_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"origin": {
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
		},
	}
}

func dataSourcePhoneNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Account(accountSid).IncomingPhoneNumber(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Phone number with sid (%s) was not found in account (%s)", sid, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to read phone number: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("address_sid", getResponse.AddressSid)
	d.Set("address_requirements", getResponse.AddressRequirements)
	d.Set("beta", getResponse.Beta)
	d.Set("bundle_sid", getResponse.BundleSid)
	d.Set("capabilities", helper.FlattenCapabilities(&getResponse.Capabilities))
	d.Set("emergency_address_sid", getResponse.EmergencyAddressSid)
	d.Set("emergency_status", getResponse.EmergencyStatus)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("identity_sid", getResponse.IdentitySid)
	d.Set("messaging", helper.FlattenMessaging(getResponse))
	d.Set("origin", getResponse.Origin)
	d.Set("phone_number", getResponse.PhoneNumber)
	d.Set("status", getResponse.Status)
	d.Set("status_callback_url", getResponse.StatusCallback)
	d.Set("status_callback_method", getResponse.StatusCallbackMethod)
	d.Set("trunk_sid", getResponse.TrunkSid)

	if helper.IsVoiceReceiveMode(getResponse.VoiceReceiveMode) {
		d.Set("voice", helper.FlattenVoice(getResponse))
	} else {
		d.Set("fax", helper.FlattenFax(getResponse))
	}

	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	return nil
}
