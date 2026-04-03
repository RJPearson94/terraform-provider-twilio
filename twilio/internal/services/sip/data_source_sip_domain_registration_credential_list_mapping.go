package sip

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPDomainRegistrationCredentialListMapping() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPDomainRegistrationCredentialListMappingRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPCredentialListSidValidation(),
				Description:  "The SID of the SIP domain registration credential list mapping to look up",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this SIP domain registration credential list mapping",
			},
			"domain_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPDomainSidValidation(),
				Description:  "The SID of the SIP domain the credential list is mapped to for registration",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the SIP domain registration credential list mapping",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP domain registration credential list mapping was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP domain registration credential list mapping was last updated, in RFC 3339 format",
			},
		},
	}
}

func dataSourceSIPDomainRegistrationCredentialListMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	domainSid := d.Get("domain_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Account(accountSid).Sip.Domain(domainSid).Auth.Registrations.CredentialListMapping(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("SIP domain regsitration credential list mapping with sid (%s) was not found for account with sid (%s) and domain with sid (%s)", sid, accountSid, domainSid)
		}
		return diag.Errorf("Failed to read SIP domain regsitration credential list mapping: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
