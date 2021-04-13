package autopilot

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAutopilotFieldType() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAutopilotFieldTypeCreate,
		ReadContext:   resourceAutopilotFieldTypeRead,
		UpdateContext: resourceAutopilotFieldTypeUpdate,
		DeleteContext: resourceAutopilotFieldTypeDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Assistants/(.*)/FieldTypes/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("assistant_sid", match[1])
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 255),
			},
			"unique_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
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

func resourceAutopilotFieldTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	createInput := &field_types.CreateFieldTypeInput{
		UniqueName:   d.Get("unique_name").(string),
		FriendlyName: utils.OptionalStringWithEmptyStringDefault(d, "friendly_name"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).FieldTypes.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create autopilot field type: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAutopilotFieldTypeRead(ctx, d, meta)
}

func resourceAutopilotFieldTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read autopilot field type: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)
	return nil
}

func resourceAutopilotFieldTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	updateInput := &field_type.UpdateFieldTypeInput{
		UniqueName:   utils.OptionalString(d, "unique_name"),
		FriendlyName: utils.OptionalStringWithEmptyStringDefault(d, "friendly_name"),
	}

	updateResp, err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update autopilot field type: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotFieldTypeRead(ctx, d, meta)
}

func resourceAutopilotFieldTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	if err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete autopilot field type: %s", err.Error())
	}
	d.SetId("")
	return nil
}
