package sip

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPIPAddresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPIPAddressesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_access_control_list_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_addresses": {
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
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cidr_length_prefix": {
							Type:     schema.TypeInt,
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

func dataSourceSIPIPAddressesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	ipAccessControlListSid := d.Get("ip_access_control_list_sid").(string)
	paginator := client.Account(accountSid).Sip.IpAccessControlList(ipAccessControlListSid).IpAddresses.NewIpAddressesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No SIP IP addresses were found for account with sid (%s) and IP access control list with sid (%s)", accountSid, ipAccessControlListSid)
		}
		return diag.Errorf("Failed to list SIP IP addresses: %s", err.Error())
	}

	d.SetId(accountSid + "/" + ipAccessControlListSid)
	d.Set("account_sid", accountSid)
	d.Set("ip_access_control_list_sid", ipAccessControlListSid)

	ipAddresses := make([]interface{}, 0)

	for _, ipAddress := range paginator.IpAddresses {
		ipAddressMap := make(map[string]interface{})

		ipAddressMap["sid"] = ipAddress.Sid
		ipAddressMap["friendly_name"] = ipAddress.FriendlyName
		ipAddressMap["ip_address"] = ipAddress.IpAddress
		ipAddressMap["cidr_length_prefix"] = ipAddress.CidrPrefixLength
		ipAddressMap["date_created"] = ipAddress.DateCreated.Time.Format(time.RFC3339)

		if ipAddress.DateUpdated != nil {
			ipAddressMap["date_updated"] = ipAddress.DateUpdated.Time.Format(time.RFC3339)
		}

		ipAddresses = append(ipAddresses, ipAddressMap)
	}

	d.Set("ip_addresses", &ipAddresses)

	return nil
}
