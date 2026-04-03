package sip

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPDomain() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPDomainRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPDomainSidValidation(),
				Description:  "The SID of the SIP domain to look up",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this SIP domain",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fully qualified domain name for the SIP domain",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the SIP domain",
			},
			"voice": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The voice settings for the SIP domain",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_callback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL called for voice status callback events",
						},
						"status_callback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice status callback URL",
						},
						"fallback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice fallback URL",
						},
						"fallback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when an error occurs retrieving or executing the TwiML for voice calls",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice URL",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL called when the SIP domain receives an incoming voice call",
						},
					},
				},
			},
			"emergency": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The emergency calling settings for the SIP domain",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"calling_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether emergency calling is enabled for this SIP domain",
						},
						"caller_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the phone number used as the emergency caller ID",
						},
					},
				},
			},
			"byoc_trunk_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the BYOC trunk associated with this SIP domain",
			},
			"secure": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether secure SIP (SIPS) is enabled for the domain",
			},
			"sip_registration": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether SIP registration is allowed for the domain",
			},
			"auth_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication type configured for the SIP domain",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP domain was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP domain was last updated, in RFC 3339 format",
			},
		},
	}
}

func dataSourceSIPDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Account(accountSid).Sip.Domain(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("SIP domain with sid (%s) was not found for account with sid (%s)", sid, accountSid)
		}
		return diag.Errorf("Failed to read SIP domain: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("auth_type", getResponse.AuthType)
	d.Set("byoc_trunk_sid", getResponse.ByocTrunkSid)
	d.Set("domain_name", getResponse.DomainName)
	d.Set("emergency", helper.FlattenEmergency(getResponse))
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("secure", getResponse.Secure)
	d.Set("sip_registration", getResponse.SipRegistration)
	d.Set("voice", helper.FlattenVoice(getResponse))
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
