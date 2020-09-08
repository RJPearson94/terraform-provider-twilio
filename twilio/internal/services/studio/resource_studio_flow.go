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
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"draft",
					"published",
				}, false),
			},
			"definition": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"commit_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"validate": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"valid": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceStudioFlowCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	if err := validateRequest(d, meta); err != nil {
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
		return fmt.Errorf("[ERROR] Failed to create studio flow: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceStudioFlowRead(d, meta)
}

func resourceStudioFlowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Flow(d.Id()).FetchWithContext(ctx)
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

	json, err := structure.FlattenJsonToString(getResponse.Definition)
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to flattern definition json to string")
	}
	d.Set("definition", json)
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
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	if err := validateRequest(d, meta); err != nil {
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
		return fmt.Errorf("Failed to Update Studio Flow: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceStudioFlowRead(d, meta)
}

func resourceStudioFlowDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Flow(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete studio flow: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func validateRequest(d *schema.ResourceData, meta interface{}) error {
	if d.Get("validate").(bool) {
		client := meta.(*common.TwilioClient).Studio
		ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
		defer cancel()

		definitionJSONString, _ := structure.NormalizeJsonString(d.Get("definition").(string))
		validateInput := &flow_validation.ValidateFlowInput{
			FriendlyName:  d.Get("friendly_name").(string),
			Status:        d.Get("status").(string),
			Definition:    definitionJSONString,
			CommitMessage: utils.OptionalString(d, "commit_message"),
		}

		resp, err := client.FlowValidation.ValidateWithContext(ctx, validateInput)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to validate studio flow: %s", err.Error())
		}
		if !resp.Valid {
			return fmt.Errorf("[ERROR] The template is invalid")
		}
	}
	return nil
}
