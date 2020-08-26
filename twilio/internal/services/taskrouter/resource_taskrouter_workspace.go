package taskrouter

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const workspaceEventsSeperator = ","

func resourceTaskRouterWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterWorkspaceCreate,
		Read:   resourceTaskRouterWorkspaceRead,
		Update: resourceTaskRouterWorkspaceUpdate,
		Delete: resourceTaskRouterWorkspaceDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Workspaces/(.*)"
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
				Type:     schema.TypeString,
				Required: true,
			},
			"event_callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"event_filters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"multi_task_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prioritize_queue_order": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"LIFO",
					"FIFO",
				}, false),
			},
			"default_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_activity_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_activity_sid": {
				Type:     schema.TypeString,
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
		},
	}
}

func resourceTaskRouterWorkspaceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &workspaces.CreateWorkspaceInput{
		FriendlyName:         d.Get("friendly_name").(string),
		EventCallbackURL:     utils.OptionalString(d, "event_callback_url"),
		EventsFilter:         utils.OptionalSeperatedString(d, "event_filters", workspaceEventsSeperator),
		MultiTaskEnabled:     utils.OptionalBool(d, "multi_task_enabled"),
		Template:             utils.OptionalString(d, "template"),
		PrioritizeQueueOrder: utils.OptionalString(d, "prioritize_queue_order"),
	}

	createResult, err := client.Workspaces.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create taskrouter workspace: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkspaceRead(d, meta)
}

func resourceTaskRouterWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Workspace(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter workspace: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("event_callback_url", getResponse.EventCallbackURL)

	if getResponse.EventsFilter != nil {
		d.Set("event_filters", strings.Split(*getResponse.EventsFilter, workspaceEventsSeperator))
	}

	d.Set("default_activity_name", getResponse.DefaultActivityName)
	d.Set("default_activity_sid", getResponse.DefaultActivitySid)
	d.Set("multi_task_enabled", getResponse.MultiTaskEnabled)
	d.Set("prioritize_queue_order", getResponse.PrioritizeQueueOrder)
	d.Set("timeout_activity_name", getResponse.TimeoutActivityName)
	d.Set("timeout_activity_sid", getResponse.TimeoutActivitySid)

	if value, ok := d.GetOk("template"); ok {
		d.Set("template", value.(string))
	}

	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceTaskRouterWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &workspace.UpdateWorkspaceInput{
		FriendlyName:         utils.OptionalString(d, "friendly_name"),
		EventCallbackURL:     utils.OptionalString(d, "event_callback_url"),
		EventsFilter:         utils.OptionalSeperatedString(d, "event_filters", workspaceEventsSeperator),
		MultiTaskEnabled:     utils.OptionalBool(d, "multi_task_enabled"),
		Template:             utils.OptionalString(d, "template"),
		PrioritizeQueueOrder: utils.OptionalString(d, "prioritize_queue_order"),
	}

	updateResp, err := client.Workspace(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update taskrouter workspace: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkspaceRead(d, meta)
}

func resourceTaskRouterWorkspaceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Workspace(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete taskrouter workspace: %s", err.Error())
	}
	d.SetId("")
	return nil
}
