package sip_trunking

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPTrunkingIPAccessControlLists() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPTrunkingIPAccessControlListsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk to retrieve IP access control lists for",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the SIP trunk IP access control lists",
			},
			"ip_access_control_lists": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of IP access control lists associated with the SIP trunk",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this SIP trunk IP access control list by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the SIP trunk IP access control list",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the SIP trunk IP access control list was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the SIP trunk IP access control list was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the SIP trunk IP access control list resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceSIPTrunkingIPAccessControlListsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	trunkSid := d.Get("trunk_sid").(string)
	paginator := client.Trunk(trunkSid).IpAccessControlLists.NewIpAccessControlListsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No IP access control lists were found for SIP trunk with sid (%s)", trunkSid)
		}
		return diag.Errorf("Failed to list SIP trunk IP access control lists: %s", err.Error())
	}

	d.SetId(trunkSid)
	d.Set("trunk_sid", trunkSid)

	IpAccessControlLists := make([]interface{}, 0)

	for _, IpAccessControlList := range paginator.IpAccessControlLists {
		d.Set("account_sid", IpAccessControlList.AccountSid)

		IpAccessControlListMap := make(map[string]interface{})

		IpAccessControlListMap["sid"] = IpAccessControlList.Sid
		IpAccessControlListMap["friendly_name"] = IpAccessControlList.FriendlyName
		IpAccessControlListMap["date_created"] = IpAccessControlList.DateCreated.Format(time.RFC3339)

		if IpAccessControlList.DateUpdated != nil {
			IpAccessControlListMap["date_updated"] = IpAccessControlList.DateUpdated.Format(time.RFC3339)
		}

		IpAccessControlListMap["url"] = IpAccessControlList.URL

		IpAccessControlLists = append(IpAccessControlLists, IpAccessControlListMap)
	}

	d.Set("ip_access_control_lists", &IpAccessControlLists)

	return nil
}
