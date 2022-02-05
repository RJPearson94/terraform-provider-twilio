package sip

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPDomainRegistrationCredentialListMappings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPDomainRegistrationCredentialListMappingsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"domain_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPDomainSidValidation(),
			},
			"credential_list_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": {
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

func dataSourceSIPDomainRegistrationCredentialListMappingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	domainSid := d.Get("domain_sid").(string)
	paginator := client.Account(accountSid).Sip.Domain(domainSid).Auth.Registrations.CredentialListMappings.NewCredentialListMappingsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No SIP domain registration credential list mappings were found for account with sid (%s) and domain with sid (%s)", accountSid, domainSid)
		}
		return diag.Errorf("Failed to list SIP domain registration credential list mappings: %s", err.Error())
	}

	d.SetId(accountSid + "/" + domainSid)
	d.Set("account_sid", accountSid)
	d.Set("domain_sid", domainSid)

	credentialListMappings := make([]interface{}, 0)

	for _, credentialListMapping := range paginator.CredentialListMappings {
		credentialListMappingMap := make(map[string]interface{})

		credentialListMappingMap["sid"] = credentialListMapping.Sid
		credentialListMappingMap["friendly_name"] = credentialListMapping.FriendlyName
		credentialListMappingMap["date_created"] = credentialListMapping.DateCreated.Time.Format(time.RFC3339)

		if credentialListMapping.DateUpdated != nil {
			credentialListMappingMap["date_updated"] = credentialListMapping.DateUpdated.Time.Format(time.RFC3339)
		}

		credentialListMappings = append(credentialListMappings, credentialListMappingMap)
	}

	d.Set("credential_list_mappings", &credentialListMappings)

	return nil
}
