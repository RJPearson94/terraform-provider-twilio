package sip_trunking

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip_trunking/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/credential_lists"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSIPTrunkingCredentialList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPTrunkingCredentialListCreate,
		ReadContext:   resourceSIPTrunkingCredentialListRead,
		DeleteContext: resourceSIPTrunkingCredentialListDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Trunks/(.*)/CredentialLists/(.*)"
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
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: helper.TrunkSidValidation(),
			},
			"credential_list_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.CredentialListSidValidation(),
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
	}
}

func resourceSIPTrunkingCredentialListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	createInput := &credential_lists.CreateCredentialListInput{
		CredentialListSid: d.Get("credential_list_sid").(string),
	}

	createResult, err := client.Trunk(d.Get("trunk_sid").(string)).CredentialLists.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP trunk credential list: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSIPTrunkingCredentialListRead(ctx, d, meta)
}

func resourceSIPTrunkingCredentialListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	getResponse, err := client.Trunk(d.Get("trunk_sid").(string)).CredentialList(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP trunk credential list: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("trunk_sid", getResponse.TrunkSid)
	d.Set("credential_list_sid", getResponse.Sid) // The CredentialListSid is stored as the resource sid
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}

func resourceSIPTrunkingCredentialListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	if err := client.Trunk(d.Get("trunk_sid").(string)).CredentialList(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP trunk credential list: %s", err.Error())
	}
	d.SetId("")
	return nil
}
