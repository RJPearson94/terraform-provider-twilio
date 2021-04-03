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
				ValidateFunc: helper.DomainSidValidation(),
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"voice": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_callback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_callback_method": {
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
			"emergency": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"calling_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"caller_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"byoc_trunk_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secure": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"sip_registration": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"auth_type": {
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
