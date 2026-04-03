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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.PhoneNumberSidValidation(),
				Description:  "The SID of the phone number to look up",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this phone number",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the phone number",
			},
			"phone_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phone number in E.164 format",
			},
			"address_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the address associated with this phone number",
			},
			"address_requirements": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of address required for this phone number",
			},
			"beta": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the phone number is a beta number new to the Twilio platform",
			},
			"capabilities": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The set of boolean capabilities of the phone number",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number supports fax",
						},
						"sms": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number supports SMS",
						},
						"mms": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number supports MMS",
						},
						"voice": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number supports voice",
						},
					},
				},
			},
			"emergency_address_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the emergency address associated with this phone number",
			},
			"emergency_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The emergency calling status of the phone number",
			},
			"messaging": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The messaging settings for the phone number",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the TwiML application that handles incoming messages",
						},
						"fallback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the messaging fallback URL",
						},
						"fallback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming messages",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when the phone number receives an incoming message",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the messaging URL",
						},
					},
				},
			},
			"trunk_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the SIP trunk that handles voice calls for this phone number",
			},
			"voice": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The voice settings for the phone number",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the TwiML application that handles incoming voice calls",
						},
						"caller_id_lookup": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether caller ID lookup is enabled for incoming voice calls",
						},
						"fallback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice fallback URL",
						},
						"fallback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming voice calls",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice URL",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when the phone number receives an incoming voice call",
						},
					},
				},
			},
			"fax": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The fax settings for the phone number",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the TwiML application that handles incoming faxes",
						},
						"fallback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the fax fallback URL",
						},
						"fallback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming faxes",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the fax URL",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when the phone number receives an incoming fax",
						},
					},
				},
			},
			"identity_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the identity resource associated with this phone number",
			},
			"bundle_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the regulatory compliance bundle associated with this phone number",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the phone number",
			},
			"status_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL called for status callback events on the phone number",
			},
			"status_callback_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP method used to call the status callback URL",
			},
			"origin": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The origin of the phone number, such as `twilio` or `hosted`",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the phone number was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the phone number was last updated, in RFC 3339 format",
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
		d.Set("fax", &[]interface{}{})
	} else {
		d.Set("fax", helper.FlattenFax(getResponse))
		d.Set("voice", &[]interface{}{})
	}
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	return nil
}
