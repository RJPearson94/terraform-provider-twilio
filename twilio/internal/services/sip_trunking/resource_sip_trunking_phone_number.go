package sip_trunking

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip_trunking/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/phone_numbers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSIPTrunkingPhoneNumber() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPTrunkingPhoneNumberCreate,
		ReadContext:   resourceSIPTrunkingPhoneNumberRead,
		DeleteContext: resourceSIPTrunkingPhoneNumberDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Trunks/(.*)/PhoneNumbers/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("trunk_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this SIP trunk phone number by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this SIP trunk phone number",
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk to associate the phone number with. Changing this forces a new resource",
			},
			"phone_number_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.PhoneNumberSidValidation(),
				Description:  "The SID of the incoming phone number to associate with the trunk. Changing this forces a new resource",
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

func resourceSIPTrunkingPhoneNumberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	createInput := &phone_numbers.CreatePhoneNumberInput{
		PhoneNumberSid: d.Get("phone_number_sid").(string),
	}

	createResult, err := client.Trunk(d.Get("trunk_sid").(string)).PhoneNumbers.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP trunk phone number: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSIPTrunkingPhoneNumberRead(ctx, d, meta)
}

func resourceSIPTrunkingPhoneNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	getResponse, err := client.Trunk(d.Get("trunk_sid").(string)).PhoneNumber(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP trunk phone number: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("address_requirements", getResponse.AddressRequirements)
	d.Set("phone_number_sid", getResponse.Sid) // The PhoneNumberSid is stored as the resource sid
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

func resourceSIPTrunkingPhoneNumberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	if err := client.Trunk(d.Get("trunk_sid").(string)).PhoneNumber(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP trunk phone number: %s", err.Error())
	}
	d.SetId("")
	return nil
}
