package sip

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPCredentials() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPCredentialsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
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
			"credentials": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceSIPCredentialsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	credentialListSid := d.Get("credential_list_sid").(string)
	paginator := client.Account(accountSid).Sip.CredentialList(credentialListSid).Credentials.NewCredentialsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No SIP credentials were found for account with sid (%s) and credential list with sid (%s)", accountSid, credentialListSid)
		}
		return diag.Errorf("Failed to list SIP credentials: %s", err.Error())
	}

	d.SetId(accountSid + "/" + credentialListSid)
	d.Set("account_sid", accountSid)
	d.Set("credential_list_sid", credentialListSid)

	credentials := make([]interface{}, 0)

	for _, credential := range paginator.Credentials {
		credentialMap := make(map[string]interface{})

		credentialMap["sid"] = credential.Sid
		credentialMap["username"] = credential.Username
		credentialMap["date_created"] = credential.DateCreated.Time.Format(time.RFC3339)

		if credential.DateUpdated != nil {
			credentialMap["date_updated"] = credential.DateUpdated.Time.Format(time.RFC3339)
		}

		credentials = append(credentials, credentialMap)
	}

	d.Set("credentials", &credentials)

	return nil
}
