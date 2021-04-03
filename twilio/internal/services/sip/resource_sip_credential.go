package sip

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credential"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credentials"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSIPCredential() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPCredentialCreate,
		ReadContext:   resourceSIPCredentialRead,
		UpdateContext: resourceSIPCredentialUpdate,
		DeleteContext: resourceSIPCredentialDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/SIP/CredentialLists/(.*)/Credentials/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("account_sid", match[1])
				d.Set("credential_list_sid", match[2])
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
			"credential_list_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.CredentialListSidValidation(),
			},
			"username": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile("^.{12,}$"), "Must contain at least 12 characters"),
					validation.StringMatch(regexp.MustCompile("^.*[A-Z].*$"), "Must contain a uppercase letter"),
					validation.StringMatch(regexp.MustCompile("^.*[a-z].*$"), "Must contain a lowercase letter"),
					validation.StringMatch(regexp.MustCompile("^.*[0-9].*$"), "Must contain a number"),
				),
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

func resourceSIPCredentialCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &credentials.CreateCredentialInput{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).Sip.CredentialList(d.Get("credential_list_sid").(string)).Credentials.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP credential: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSIPCredentialRead(ctx, d, meta)
}

func resourceSIPCredentialRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).Sip.CredentialList(d.Get("credential_list_sid").(string)).Credential(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP credential: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("credential_list_sid", getResponse.CredentialListSid)
	d.Set("username", getResponse.Username)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}

func resourceSIPCredentialUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	updateInput := &credential.UpdateCredentialInput{
		Password: d.Get("password").(string),
	}

	updateResult, err := client.Account(d.Get("account_sid").(string)).Sip.CredentialList(d.Get("credential_list_sid").(string)).Credential(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update SIP credential: %s", err.Error())
	}

	d.SetId(updateResult.Sid)
	return resourceSIPCredentialRead(ctx, d, meta)
}

func resourceSIPCredentialDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	if err := client.Account(d.Get("account_sid").(string)).Sip.CredentialList(d.Get("credential_list_sid").(string)).Credential(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP credential: %s", err.Error())
	}
	d.SetId("")
	return nil
}
