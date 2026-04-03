package sip_trunking

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_url"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/origination_urls"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSIPTrunkingOriginationURL() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPTrunkingOriginationURLCreate,
		ReadContext:   resourceSIPTrunkingOriginationURLRead,
		UpdateContext: resourceSIPTrunkingOriginationURLUpdate,
		DeleteContext: resourceSIPTrunkingOriginationURLDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Trunks/(.*)/OriginationUrls/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("trunk_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this SIP trunk origination URL by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this SIP trunk origination URL",
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk to add the origination URL to. Changing this forces a new resource",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the origination URL is enabled and available for use",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A human-readable label for the SIP trunk origination URL",
			},
			"priority": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "The priority of the origination URL, from 0 to 65535. Lower values have higher priority",
			},
			"sip_url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^sip:.+$"), ""),
				Description:  "The SIP address to route origination calls to, must start with `sip:`",
			},
			"weight": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "The weight of the origination URL, from 0 to 65535. Used for load balancing among URLs with the same priority",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk origination URL was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk origination URL was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the SIP trunk origination URL resource",
			},
		},
	}
}

func resourceSIPTrunkingOriginationURLCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	createInput := &origination_urls.CreateOriginationURLInput{
		Enabled:      d.Get("enabled").(bool),
		FriendlyName: d.Get("friendly_name").(string),
		Priority:     d.Get("priority").(int),
		SipURL:       d.Get("sip_url").(string),
		Weight:       d.Get("weight").(int),
	}

	createResult, err := client.Trunk(d.Get("trunk_sid").(string)).OriginationURLs.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP trunk origination url: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSIPTrunkingOriginationURLRead(ctx, d, meta)
}

func resourceSIPTrunkingOriginationURLRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	getResponse, err := client.Trunk(d.Get("trunk_sid").(string)).OriginationURL(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP trunk origination url: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("enabled", getResponse.Enabled)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("priority", getResponse.Priority)
	d.Set("sip_url", getResponse.SipURL)
	d.Set("trunk_sid", getResponse.TrunkSid)
	d.Set("weight", getResponse.Weight)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}

func resourceSIPTrunkingOriginationURLUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	updateInput := &origination_url.UpdateOriginationURLInput{
		Enabled:      utils.OptionalBool(d, "enabled"),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		Priority:     utils.OptionalIntWith0OnChange(d, "priority"),
		SipURL:       utils.OptionalString(d, "sip_url"),
		Weight:       utils.OptionalIntWith0OnChange(d, "weight"),
	}

	updateResult, err := client.Trunk(d.Get("trunk_sid").(string)).OriginationURL(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update SIP trunk origination url: %s", err.Error())
	}

	d.SetId(updateResult.Sid)
	return resourceSIPTrunkingOriginationURLRead(ctx, d, meta)
}

func resourceSIPTrunkingOriginationURLDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	if err := client.Trunk(d.Get("trunk_sid").(string)).OriginationURL(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP trunk origination url: %s", err.Error())
	}
	d.SetId("")
	return nil
}
