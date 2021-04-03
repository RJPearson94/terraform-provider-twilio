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

func dataSourceSIPCredential() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPCredentialRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: helper.CredentialSidValidation(),
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"credential_list_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.CredentialListSidValidation(),
			},
			"username": {
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

func dataSourceSIPCredentialRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	credentialListSid := d.Get("credential_list_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Account(accountSid).Sip.CredentialList(credentialListSid).Credential(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("SIP credential with sid (%s) was not found for account with sid (%s) and credential list with sid (%s)", sid, accountSid, credentialListSid)
		}
		return diag.Errorf("Failed to read SIP credential: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("credential_list_sid", getResponse.CredentialListSid)
	d.Set("username", getResponse.Username)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
