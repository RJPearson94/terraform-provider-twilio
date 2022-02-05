package sip

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain/auth/registrations/credential_list_mappings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSIPDomainRegistrationCredentialListMapping() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPDomainRegistrationCredentialListMappingCreate,
		ReadContext:   resourceSIPDomainRegistrationCredentialListMappingRead,
		DeleteContext: resourceSIPDomainRegistrationCredentialListMappingDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/SIP/Domains/(.*)/Auth/Registrations/CredentialListMappings/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("account_sid", match[1])
				d.Set("domain_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"domain_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.SIPDomainSidValidation(),
			},
			"credential_list_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.SIPCredentialListSidValidation(),
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
	}
}

func resourceSIPDomainRegistrationCredentialListMappingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &credential_list_mappings.CreateCredentialListMappingInput{
		CredentialListSid: d.Get("credential_list_sid").(string),
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).Sip.Domain(d.Get("domain_sid").(string)).Auth.Registrations.CredentialListMappings.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP domain registration credential list mapping: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSIPDomainRegistrationCredentialListMappingRead(ctx, d, meta)
}

func resourceSIPDomainRegistrationCredentialListMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).Sip.Domain(d.Get("domain_sid").(string)).Auth.Registrations.CredentialListMapping(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP domain registration credential list mapping: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("credential_list_sid", getResponse.Sid) // The CredentialListSid is stored as the resource sid
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}

func resourceSIPDomainRegistrationCredentialListMappingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	if err := client.Account(d.Get("account_sid").(string)).Sip.Domain(d.Get("domain_sid").(string)).Auth.Registrations.CredentialListMapping(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP domain registration credential list mapping: %s", err.Error())
	}
	d.SetId("")
	return nil
}
