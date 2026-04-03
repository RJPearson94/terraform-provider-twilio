package account

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccountDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountDetailsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account to look up. If not specified, the provider's account SID is used",
			},
			"owner_account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the parent account that owns this account",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the account",
			},
			"auth_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The authorization token for the account. Sensitive -- will not be shown in logs or plans",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the account. Valid values are `active`, `suspended`, or `closed`",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the account (e.g., `Trial` or `Full`)",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the account was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the account was last updated, in RFC 3339 format",
			},
		},
	}
}

func dataSourceAccountDetailsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.API

	var sid string
	if v, ok := d.GetOk("sid"); ok {
		sid = v.(string)
	} else {
		sid = twilioClient.AccountSid
	}

	getResponse, err := client.Account(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Account with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read account details: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("owner_account_sid", getResponse.OwnerAccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("status", getResponse.Status)
	d.Set("type", getResponse.Type)
	d.Set("auth_token", getResponse.AuthToken)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
