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
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: helper.TrunkSidValidation(),
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_number": {
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
			"status_callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_callback_method": {
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
