package studio

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceStudioFlow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStudioFlowCreate,
		ReadContext:   resourceStudioFlowRead,
		UpdateContext: resourceStudioFlowUpdate,
		DeleteContext: resourceStudioFlowDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Flows/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"draft",
					"published",
				}, false),
			},
			"definition": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"commit_message": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Updated via Terraform",
			},
			"validate": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"valid": {
				Type:     schema.TypeBool,
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
			"webhook_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceStudioFlowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Studio

	if err := validateRequest(ctx, d, meta); err != nil {
		return err
	}

	definitionJSONString, _ := structure.NormalizeJsonString(d.Get("definition").(string))
	createInput := &flows.CreateFlowInput{
		FriendlyName:  d.Get("friendly_name").(string),
		Status:        d.Get("status").(string),
		Definition:    definitionJSONString,
		CommitMessage: utils.OptionalString(d, "commit_message"),
	}

	createResult, err := client.Flows.CreateWithContext(ctx, createInput)
	if err != nil {
		return handleError("Failed to create studio flow", err)
	}

	d.SetId(createResult.Sid)

	// This is a duplicate of resourceStudioFlowRead as the API is eventually consistent and was causing issues with the provider
	d.Set("sid", createResult.Sid)
	d.Set("account_sid", createResult.AccountSid)
	d.Set("friendly_name", createResult.FriendlyName)

	json, err := structure.FlattenJsonToString(createResult.Definition)
	if err != nil {
		return diag.Errorf("Unable to flatten definition json to string")
	}
	d.Set("definition", json)
	d.Set("status", createResult.Status)
	d.Set("revision", createResult.Revision)
	d.Set("commit_message", createResult.CommitMessage)
	d.Set("valid", createResult.Valid)
	d.Set("date_created", createResult.DateCreated.Format(time.RFC3339))

	if createResult.DateUpdated != nil {
		d.Set("date_updated", createResult.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", createResult.URL)
	d.Set("webhook_url", createResult.WebhookURL)
	return nil
}

func resourceStudioFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Studio

	getResponse, err := client.Flow(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read studio flow: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)

	json, err := structure.FlattenJsonToString(getResponse.Definition)
	if err != nil {
		return diag.Errorf("Unable to flatten definition json to string")
	}
	d.Set("definition", json)
	d.Set("status", getResponse.Status)
	d.Set("revision", getResponse.Revision)
	d.Set("commit_message", getResponse.CommitMessage)
	d.Set("valid", getResponse.Valid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)
	d.Set("webhook_url", getResponse.WebhookURL)
	return nil
}

func resourceStudioFlowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Studio

	if err := validateRequest(ctx, d, meta); err != nil {
		return err
	}

	updateInput := &flow.UpdateFlowInput{
		FriendlyName:  utils.OptionalString(d, "friendly_name"),
		Status:        d.Get("status").(string),
		Definition:    utils.OptionalJSONString(d, "definition"),
		CommitMessage: utils.OptionalString(d, "commit_message"),
	}

	updateResp, err := client.Flow(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return handleError("Failed to update studio flow", err)
	}

	d.SetId(updateResp.Sid)

	// This is a duplicate of resourceStudioFlowRead as the API is eventually consistent and was causing issues with the provider
	d.Set("sid", updateResp.Sid)
	d.Set("account_sid", updateResp.AccountSid)
	d.Set("friendly_name", updateResp.FriendlyName)

	json, err := structure.FlattenJsonToString(updateResp.Definition)
	if err != nil {
		return diag.Errorf("Unable to flatten definition json to string")
	}
	d.Set("definition", json)
	d.Set("status", updateResp.Status)
	d.Set("revision", updateResp.Revision)
	d.Set("commit_message", updateResp.CommitMessage)
	d.Set("valid", updateResp.Valid)
	d.Set("date_created", updateResp.DateCreated.Format(time.RFC3339))

	if updateResp.DateUpdated != nil {
		d.Set("date_updated", updateResp.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", updateResp.URL)
	d.Set("webhook_url", updateResp.WebhookURL)
	return nil
}

func resourceStudioFlowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Studio

	if err := client.Flow(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete studio flow: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func validateRequest(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.Get("validate").(bool) {
		client := meta.(*common.TwilioClient).Studio

		definitionJSONString, _ := structure.NormalizeJsonString(d.Get("definition").(string))
		validateInput := &flow_validation.ValidateFlowInput{
			FriendlyName:  d.Get("friendly_name").(string),
			Status:        d.Get("status").(string),
			Definition:    definitionJSONString,
			CommitMessage: utils.OptionalString(d, "commit_message"),
		}

		resp, err := client.FlowValidation.ValidateWithContext(ctx, validateInput)
		if err != nil {
			return handleError("Failed to validate studio flow", err)
		}
		if !resp.Valid {
			return diag.Errorf("The template is invalid")
		}
	}
	return nil
}

func handleError(errorPrefix string, err error) diag.Diagnostics {
	if twilioErr, ok := err.(*sdkUtils.TwilioError); ok && twilioErr.Details != nil {
		errDetails, _ := structure.FlattenJsonToString(*twilioErr.Details)
		return diag.Errorf("%s: %s. Details are %s", errorPrefix, err.Error(), errDetails)
	}
	return diag.Errorf("%s: %s", errorPrefix, err.Error())
}
