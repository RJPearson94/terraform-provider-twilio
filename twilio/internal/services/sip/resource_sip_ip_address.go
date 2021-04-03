package sip

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_address"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list/ip_addresses"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSIPIPAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPIPAddressCreate,
		ReadContext:   resourceSIPIPAddressRead,
		UpdateContext: resourceSIPIPAddressUpdate,
		DeleteContext: resourceSIPIPAddressDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/SIP/IpAccessControlLists/(.*)/IpAddresses/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("account_sid", match[1])
				d.Set("ip_access_control_list_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"ip_access_control_list_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.SIPIPAccessControlListSidValidation(),
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"ip_address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"cidr_length_prefix": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
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

func resourceSIPIPAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &ip_addresses.CreateIpAddressInput{
		FriendlyName:     d.Get("friendly_name").(string),
		IpAddress:        d.Get("ip_address").(string),
		CidrPrefixLength: utils.OptionalInt(d, "cidr_length_prefix"),
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).Sip.IpAccessControlList(d.Get("ip_access_control_list_sid").(string)).IpAddresses.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP IP address resource: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSIPIPAddressRead(ctx, d, meta)
}

func resourceSIPIPAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).Sip.IpAccessControlList(d.Get("ip_access_control_list_sid").(string)).IpAddress(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP IP address: %s", err.Error())
	}

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

func resourceSIPIPAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	updateInput := &ip_address.UpdateIpAddressInput{
		FriendlyName:     utils.OptionalString(d, "friendly_name"),
		IpAddress:        utils.OptionalString(d, "ip_address"),
		CidrPrefixLength: utils.OptionalInt(d, "cidr_length_prefix"),
	}

	updateResult, err := client.Account(d.Get("account_sid").(string)).Sip.IpAccessControlList(d.Get("ip_access_control_list_sid").(string)).IpAddress(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update SIP IP address: %s", err.Error())
	}

	d.SetId(updateResult.Sid)
	return resourceSIPIPAddressRead(ctx, d, meta)
}

func resourceSIPIPAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	if err := client.Account(d.Get("account_sid").(string)).Sip.IpAccessControlList(d.Get("ip_access_control_list_sid").(string)).IpAddress(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP IP address: %s", err.Error())
	}
	d.SetId("")
	return nil
}
