package sip_trunking

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPTrunkingCredentialLists() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPTrunkingCredentialListsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credential_lists": {
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
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSIPTrunkingCredentialListsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	trunkSid := d.Get("trunk_sid").(string)
	paginator := client.Trunk(trunkSid).CredentialLists.NewCredentialListsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No credential lists were found for SIP trunk with sid (%s)", trunkSid)
		}
		return diag.Errorf("Failed to list SIP trunk credential lists: %s", err.Error())
	}

	d.SetId(trunkSid)
	d.Set("trunk_sid", trunkSid)

	credentialLists := make([]interface{}, 0)

	for _, credentialList := range paginator.CredentialLists {
		d.Set("account_sid", credentialList.AccountSid)

		credentialListMap := make(map[string]interface{})

		credentialListMap["sid"] = credentialList.Sid
		credentialListMap["friendly_name"] = credentialList.FriendlyName
		credentialListMap["date_created"] = credentialList.DateCreated.Format(time.RFC3339)

		if credentialList.DateUpdated != nil {
			credentialListMap["date_updated"] = credentialList.DateUpdated.Format(time.RFC3339)
		}

		credentialListMap["url"] = credentialList.URL

		credentialLists = append(credentialLists, credentialListMap)
	}

	d.Set("credential_lists", &credentialLists)

	return nil
}
