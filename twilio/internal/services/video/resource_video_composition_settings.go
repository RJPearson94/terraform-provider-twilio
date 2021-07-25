package video

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_settings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVideoCompositionSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVideoCompositionSettingsCreate,
		ReadContext:   resourceVideoCompositionSettingsRead,
		UpdateContext: resourceVideoCompositionSettingsUpdate,
		DeleteContext: resourceVideoCompositionSettingsDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_credentials_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.CredentialSidValidation(),
			},
			"aws_s3_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"aws_storage_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"encryption_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"encryption_key_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.CredentialSidValidation(),
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVideoCompositionSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Video composition settings already exists so updating the default settings
	return resourceVideoCompositionSettingsUpdate(ctx, d, meta)
}

func resourceVideoCompositionSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Video

	getResponse, err := client.CompositionSettings().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read composition settings: %s", err.Error())
	}

	d.Set("account_sid", getResponse.AccountSid)
	d.Set("aws_credentials_sid", getResponse.AWSCredentialSid)
	d.Set("aws_s3_url", getResponse.AWSS3URL)
	d.Set("aws_storage_enabled", getResponse.AWSStorageEnabled)
	d.Set("encryption_enabled", getResponse.EncryptionEnabled)
	d.Set("encryption_key_sid", getResponse.EncryptionKeySid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("url", getResponse.URL)

	return nil
}

func resourceVideoCompositionSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Video

	updateInput := &composition_settings.UpdateCompositionSettingsInput{
		AWSCredentialSid:  utils.OptionalStringWithEmptyStringOnChange(d, "aws_credentials_sid"),
		AWSS3URL:          utils.OptionalStringWithEmptyStringOnChange(d, "aws_s3_url"),
		AWSStorageEnabled: utils.OptionalBool(d, "aws_storage_enabled"),
		EncryptionEnabled: utils.OptionalBool(d, "encryption_enabled"),
		EncryptionKeySid:  utils.OptionalStringWithEmptyStringOnChange(d, "encryption_key_sid"),
		FriendlyName:      d.Get("friendly_name").(string),
	}

	updateResp, err := client.CompositionSettings().UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update composition settings: %s", err.Error())
	}

	d.SetId(updateResp.AccountSid)
	return resourceVideoCompositionSettingsRead(ctx, d, meta)
}

func resourceVideoCompositionSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Video composition settings cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}
