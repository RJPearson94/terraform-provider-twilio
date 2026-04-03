package credentials

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/aws_credential"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials/aws_credentials"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCredentialsAWS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAWSCreate,
		ReadContext:   resourceAWSRead,
		UpdateContext: resourceAWSUpdate,
		DeleteContext: resourceAWSDelete,

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
				Description: "The unique SID assigned to this AWS credential by Twilio",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this AWS credential. Changing this forces a new resource",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A human-readable label for the AWS credential",
			},
			"aws_access_key_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The AWS access key ID used to authenticate with AWS services. Changing this forces a new resource",
			},
			"aws_secret_access_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "The AWS secret access key used to authenticate with AWS services. Sensitive -- will not be shown in logs or plans. Changing this forces a new resource",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the AWS credential was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the AWS credential was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the AWS credential resource",
			},
		},
	}
}

func resourceAWSCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	createInput := &aws_credentials.CreateAWSCredentialInput{
		Credentials:  fmt.Sprintf("%s:%s", d.Get("aws_access_key_id").(string), d.Get("aws_secret_access_key").(string)),
		FriendlyName: utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		AccountSid:   utils.OptionalString(d, "account_sid"),
	}

	createResult, err := client.Credentials.AWSCredentials.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create aws credential: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAWSRead(ctx, d, meta)
}

func resourceAWSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	getResponse, err := client.Credentials.AWSCredential(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read aws credential: %s", err.Error())
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

func resourceAWSUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	updateInput := &aws_credential.UpdateAWSCredentialInput{
		FriendlyName: utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
	}

	updateResp, err := client.Credentials.AWSCredential(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update aws credential: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAWSRead(ctx, d, meta)
}

func resourceAWSDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Accounts

	if err := client.Credentials.AWSCredential(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete aws credential: %s", err.Error())
	}

	d.SetId("")
	return nil
}
