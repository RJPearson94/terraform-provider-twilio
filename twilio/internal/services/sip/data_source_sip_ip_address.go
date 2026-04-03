package sip

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPIPAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPIPAddressRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPIPAddressSidValidation(),
				Description:  "The SID of the SIP IP address to look up",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this SIP IP address",
			},
			"ip_access_control_list_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPIPAccessControlListSidValidation(),
				Description:  "The SID of the IP access control list that this IP address belongs to",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the SIP IP address",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address in the access control list",
			},
			"cidr_length_prefix": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The CIDR prefix length for the IP address range",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP IP address was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP IP address was last updated, in RFC 3339 format",
			},
		},
	}
}

func dataSourceSIPIPAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	ipAccessControlListSid := d.Get("ip_access_control_list_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Account(accountSid).Sip.IpAccessControlList(ipAccessControlListSid).IpAddress(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("SIP IP address with sid (%s) was not found for account with sid (%s) and IP access control list with sid (%s)", sid, accountSid, ipAccessControlListSid)
		}
		return diag.Errorf("Failed to read SIP IP address: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("ip_access_control_list_sid", getResponse.IpAccessControlListSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("ip_address", getResponse.IpAddress)
	d.Set("cidr_length_prefix", getResponse.CidrPrefixLength)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
