package sip_trunking

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip_trunking/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPTrunkingPhoneNumber() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPTrunkingPhoneNumberRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.PhoneNumberSidValidation(),
				Description:  "The SID of the SIP trunk phone number to look up",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this SIP trunk phone number",
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk the phone number belongs to",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the SIP trunk phone number",
			},
			"phone_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The phone number in E.164 format",
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
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk phone number was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk phone number was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the SIP trunk phone number resource",
			},
		},
	}
}

func dataSourceSIPTrunkingPhoneNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Trunk(trunkSid).PhoneNumber(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("SIP trunk phone number with sid (%s) was not found for SIP trunk with sid (%s)", sid, trunkSid)
		}
		return diag.Errorf("Failed to read SIP trunk phone number: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("address_requirements", getResponse.AddressRequirements)
	d.Set("beta", getResponse.Beta)
	d.Set("capabilities", helper.FlattenCapabilities(&getResponse.Capabilities))
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("messaging", helper.FlattenMessaging(getResponse))
	d.Set("phone_number", getResponse.PhoneNumber)
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

	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
