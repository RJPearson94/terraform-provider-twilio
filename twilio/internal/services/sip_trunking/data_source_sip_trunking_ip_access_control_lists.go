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
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_access_control_lists": {
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
