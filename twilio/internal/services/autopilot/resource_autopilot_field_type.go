package autopilot

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_types"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAutopilotFieldType() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutopilotFieldTypeCreate,
		Read:   resourceAutopilotFieldTypeRead,
		Update: resourceAutopilotFieldTypeUpdate,
		Delete: resourceAutopilotFieldTypeDelete,
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
			"assistant_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unique_name": {
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

func resourceAutopilotFieldTypeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	createInput := &field_types.CreateFieldTypeInput{
		UniqueName:   d.Get("unique_name").(string),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).FieldTypes.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create autopilot field type: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAutopilotFieldTypeRead(d, meta)
}

func resourceAutopilotFieldTypeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot field type: %s", err.Error())
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

func resourceAutopilotFieldTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	updateInput := &field_type.UpdateFieldTypeInput{
		UniqueName:   utils.OptionalString(d, "unique_name"),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	updateResp, err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update autopilot field type: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotFieldTypeRead(d, meta)
}

func resourceAutopilotFieldTypeDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	if err := client.Assistant(d.Get("assistant_sid").(string)).FieldType(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete autopilot field type: %s", err.Error())
	}
	d.SetId("")
	return nil
}
