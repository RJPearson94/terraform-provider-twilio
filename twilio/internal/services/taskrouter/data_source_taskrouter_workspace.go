package taskrouter

import (
	"context"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterWorkspaceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
				Description:  "The SID of the workspace to retrieve",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this workspace",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the workspace",
			},
			"event_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to call when an event is fired in the workspace",
			},
			"event_filters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of event types the workspace is subscribed to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"multi_task_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether multi-tasking is enabled for the workspace",
			},
			"prioritize_queue_order": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The order in which task queues are prioritized",
			},
			"default_activity_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the default activity for the workspace",
			},
			"default_activity_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default activity for the workspace",
			},
			"timeout_activity_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the timeout activity for the workspace",
			},
			"timeout_activity_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the timeout activity for the workspace",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the workspace was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the workspace was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the workspace resource",
			},
		},
	}
}

func dataSourceTaskRouterWorkspaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	sid := d.Get("sid").(string)
	getResponse, err := client.Workspace(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("TaskRouter workspace with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read taskrouter workspace: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("event_callback_url", getResponse.EventCallbackURL)

	if getResponse.EventsFilter != nil && *getResponse.EventsFilter != "" {
		d.Set("event_filters", strings.Split(*getResponse.EventsFilter, ","))
	}

	d.Set("default_activity_name", getResponse.DefaultActivityName)
	d.Set("default_activity_sid", getResponse.DefaultActivitySid)
	d.Set("multi_task_enabled", getResponse.MultiTaskEnabled)
	d.Set("prioritize_queue_order", getResponse.PrioritizeQueueOrder)
	d.Set("timeout_activity_name", getResponse.TimeoutActivityName)
	d.Set("timeout_activity_sid", getResponse.TimeoutActivitySid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
