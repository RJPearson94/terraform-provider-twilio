package studio

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flows"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceStudioFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceStudioFlowCreate,
		Read:   resourceStudioFlowRead,
		Update: resourceStudioFlowUpdate,
		Delete: resourceStudioFlowDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Required: true,
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

func resourceStudioFlowCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio

	if err := validateRequest(d, meta); err != nil {
		return err
	}

	createInput := &flows.CreateFlowInput{
		FriendlyName:  d.Get("friendly_name").(string),
		Status:        d.Get("status").(string),
		Definition:    d.Get("definition").(string),
		CommitMessage: utils.OptionalString(d, "commit_message"),
	}

	createResult, err := client.Flows.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create studio flow: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceStudioFlowRead(d, meta)
}

func resourceStudioFlowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio

	getResponse, err := client.Flow(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read studio flow: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("definition", getResponse.Definition)
	d.Set("status", getResponse.Status)
	d.Set("revision", getResponse.Revision)
	d.Set("commit_message", getResponse.CommitMessage)
	d.Set("valid", getResponse.Valid)
	d.Set("validate", d.Get("validate").(bool))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)
	d.Set("webhook_url", getResponse.WebhookURL)

	return nil
}

func resourceStudioFlowUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio

	if err := validateRequest(d, meta); err != nil {
		return err
	}

	updateInput := &flow.UpdateFlowInput{
		FriendlyName:  utils.OptionalString(d, "friendly_name"),
		Status:        d.Get("status").(string),
		Definition:    utils.OptionalString(d, "definition"),
		CommitMessage: utils.OptionalString(d, "commit_message"),
	}

	updateResp, err := client.Flow(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to Update Studio Flow: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceStudioFlowRead(d, meta)
}

func resourceStudioFlowDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio

	if err := client.Flow(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete studio flow: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func validateRequest(d *schema.ResourceData, meta interface{}) error {
	if d.Get("validate").(bool) {
		client := meta.(*common.TwilioClient).Studio

		validateInput := &flow_validation.ValidateFlowInput{
			FriendlyName:  d.Get("friendly_name").(string),
			Status:        d.Get("status").(string),
			Definition:    d.Get("definition").(string),
			CommitMessage: utils.OptionalString(d, "commit_message"),
		}

		resp, err := client.FlowValidation.Validate(validateInput)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to validate studio flow: %s", err)
		}
		if !resp.Valid {
			return fmt.Errorf("[ERROR] The template is invalid")
		}
	}
	return nil
}
