package credentials

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/public_key"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/public_keys"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCredentialsPublicKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePublicKeyCreate,
		ReadContext:   resourcePublicKeyRead,
		UpdateContext: resourcePublicKeyUpdate,
		DeleteContext: resourcePublicKeyDelete,

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
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourcePublicKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	createInput := &public_keys.CreatePublicKeyInput{
		PublicKey:    d.Get("public_key").(string),
		FriendlyName: utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		AccountSid:   utils.OptionalString(d, "account_sid"),
	}

	createResult, err := client.Credentials.PublicKeys.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create public key: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourcePublicKeyRead(ctx, d, meta)
}

func resourcePublicKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	getResponse, err := client.Credentials.PublicKey(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read public key: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}

func resourcePublicKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	updateInput := &public_key.UpdatePublicKeyInput{
		FriendlyName: utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
	}

	updateResp, err := client.Credentials.PublicKey(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update public key: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourcePublicKeyRead(ctx, d, meta)
}

func resourcePublicKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	if err := client.Credentials.PublicKey(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete public key: %s", err.Error())
	}

	d.SetId("")
	return nil
}
