package autopilot

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_values"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAutopilotFieldValue() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutopilotFieldValueCreate,
		Read:   resourceAutopilotFieldValueRead,
		Delete: resourceAutopilotFieldValueDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			"assistant_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"field_type_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"language": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"synonym_of": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceAutopilotFieldValueCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &field_values.CreateFieldValueInput{
		Language:  d.Get("language").(string),
		Value:     d.Get("value").(string),
		SynonymOf: utils.OptionalString(d, "synonym_of"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Get("field_type_sid").(string)).FieldValues.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create autopilot field value: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAutopilotFieldValueRead(d, meta)
}

func resourceAutopilotFieldValueRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Get("field_type_sid").(string)).FieldValue(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot field value: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("field_type_sid", getResponse.FieldTypeSid)
	d.Set("language", getResponse.Language)
	d.Set("value", getResponse.Value)
	d.Set("synonym_of", getResponse.SynonymOf)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)
	return nil
}

func resourceAutopilotFieldValueDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Get("field_type_sid").(string)).FieldValue(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete autopilot field value: %s", err.Error())
	}
	d.SetId("")
	return nil
}
