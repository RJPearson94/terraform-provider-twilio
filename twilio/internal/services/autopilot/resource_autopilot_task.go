package autopilot

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/tasks"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAutopilotTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutopilotTaskCreate,
		Read:   resourceAutopilotTaskRead,
		Update: resourceAutopilotTaskUpdate,
		Delete: resourceAutopilotTaskDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Assistants/(.*)/Tasks/(.*)"
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
				Optional: true,
				Computed: true,
			},
			"actions_url": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validation.IsURLWithHTTPorHTTPS,
				ConflictsWith: []string{"actions"},
			},
			"actions": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
				ConflictsWith:    []string{"actions_url"},
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

func resourceAutopilotTaskCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &tasks.CreateTaskInput{
		UniqueName:   d.Get("unique_name").(string),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		ActionsURL:   utils.OptionalString(d, "actions_url"),
		Actions:      utils.OptionalJSONString(d, "actions"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).Tasks.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create autopilot task: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAutopilotTaskRead(d, meta)
}

func resourceAutopilotTaskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot task: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("actions_url", getResponse.ActionsURL)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	getActionsResponse, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Id()).Actions().FetchWithContext(ctx)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to read autopilot task actions: %s", err.Error())
	}

	actionsJSONString, err := structure.FlattenJsonToString(getActionsResponse.Data)
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to flatten actions json to string: %s", err.Error())
	}
	d.Set("actions", actionsJSONString)

	return nil
}

func resourceAutopilotTaskUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &task.UpdateTaskInput{
		UniqueName:   utils.OptionalString(d, "unique_name"),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	if d.HasChange("actions") {
		updateInput.Actions = utils.OptionalJSONString(d, "actions")
	} else {
		updateInput.ActionsURL = utils.OptionalString(d, "actions_url")
	}

	updateResp, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update autopilot task: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotTaskRead(d, meta)
}

func resourceAutopilotTaskDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete autopilot task: %s", err.Error())
	}
	d.SetId("")
	return nil
}
