package taskrouter

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTaskRouterWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterWorkspaceCreate,
		Read:   resourceTaskRouterWorkspaceRead,
		Update: resourceTaskRouterWorkspaceUpdate,
		Delete: resourceTaskRouterWorkspaceDelete,
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
			"event_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"events_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multi_task_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prioritize_queue_order": {
				Type:     schema.TypeString,
				Optional: true,
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

	createInput := &workspaces.CreateWorkspaceInput{
		FriendlyName:         d.Get("friendly_name").(string),
		EventCallbackUrl:     d.Get("event_callback_url").(string),
		EventsFilter:         d.Get("events_filter").(string),
		MultiTaskEnabled:     sdkUtils.Bool(d.Get("multi_task_enabled").(bool)),
		Template:             d.Get("template").(string),
		PrioritizeQueueOrder: d.Get("prioritize_queue_order").(string),
	}

	createResult, err := client.Workspaces.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create taskrouter workspace: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkspaceRead(d, meta)
}

func resourceTaskRouterWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter workspace: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("event_callback_url", getResponse.EventCallbackURL)
	d.Set("events_filter", getResponse.EventsFilter)
	d.Set("default_activity_name", getResponse.DefaultActivityName)
	d.Set("default_activity_sid", getResponse.DefaultActivitySid)
	d.Set("multi_task_enabled", getResponse.MultiTaskEnabled)
	d.Set("prioritize_queue_order", getResponse.PrioritizeQueueOrder)
	d.Set("timeout_activity_name", getResponse.TimeoutActivityName)
	d.Set("timeout_activity_sid", getResponse.TimeoutActivitySid)
	d.Set("template", d.Get("template").(string))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceTaskRouterWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &workspace.UpdateWorkspaceInput{
		FriendlyName:         d.Get("friendly_name").(string),
		EventCallbackUrl:     d.Get("event_callback_url").(string),
		EventsFilter:         d.Get("events_filter").(string),
		MultiTaskEnabled:     sdkUtils.Bool(d.Get("multi_task_enabled").(bool)),
		Template:             d.Get("template").(string),
		PrioritizeQueueOrder: d.Get("prioritize_queue_order").(string),
	}

	updateResp, err := client.Workspace(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update taskrouter workspace: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkspaceRead(d, meta)
}

func resourceTaskRouterWorkspaceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete taskrouter workspace: %s", err.Error())
	}
	d.SetId("")
	return nil
}
