package sip

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPDomainIPAccessControlListMappings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPDomainIPAccessControlListMappingsRead,

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
			"ip_access_control_list_mappings": {
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

func dataSourceSIPDomainIPAccessControlListMappingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	domainSid := d.Get("domain_sid").(string)
	paginator := client.Account(accountSid).Sip.Domain(domainSid).Auth.Calls.IpAccessControlListMappings.NewIpAccessControlListMappingsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No SIP domain IP access control list mappings were found for account with sid (%s) and domain with sid (%s)", accountSid, domainSid)
		}
		return diag.Errorf("Failed to list SIP domain IP access control list mappings: %s", err.Error())
	}

	d.SetId(accountSid + "/" + domainSid)
	d.Set("account_sid", accountSid)
	d.Set("domain_sid", domainSid)

	ipAccessControlListMappings := make([]interface{}, 0)

	for _, ipAccessControlListMapping := range paginator.IpAccessControlListMappings {
		ipAccessControlListMappingMap := make(map[string]interface{})

		ipAccessControlListMappingMap["sid"] = ipAccessControlListMapping.Sid
		ipAccessControlListMappingMap["friendly_name"] = ipAccessControlListMapping.FriendlyName
		ipAccessControlListMappingMap["date_created"] = ipAccessControlListMapping.DateCreated.Time.Format(time.RFC3339)

		if ipAccessControlListMapping.DateUpdated != nil {
			ipAccessControlListMappingMap["date_updated"] = ipAccessControlListMapping.DateUpdated.Time.Format(time.RFC3339)
		}

		ipAccessControlListMappings = append(ipAccessControlListMappings, ipAccessControlListMappingMap)
	}

	d.Set("ip_access_control_list_mappings", &ipAccessControlListMappings)

	return nil
}
