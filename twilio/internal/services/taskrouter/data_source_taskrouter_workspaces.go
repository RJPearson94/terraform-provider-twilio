package taskrouter

import (
	"context"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterWorkspaces() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterWorkspacesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter workspaces by friendly name",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the workspaces",
			},
			"workspaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of workspaces",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this workspace by Twilio",
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
				},
			},
		},
	}
}

func dataSourceTaskRouterWorkspacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.TaskRouter

	options := &workspaces.WorkspacesPageOptions{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	paginator := client.Workspaces.NewWorkspacesPaginatorWithOptions(options)
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No workspaces were found")
		}
		return diag.Errorf("Failed to read taskrouter workspaces: %s", err.Error())
	}

	d.SetId(twilioClient.AccountSid)
	d.Set("account_sid", twilioClient.AccountSid)

	workspaces := make([]interface{}, 0)

	for _, workspace := range paginator.Workspaces {
		workspacesMap := make(map[string]interface{})

		workspacesMap["sid"] = workspace.Sid
		workspacesMap["friendly_name"] = workspace.FriendlyName
		workspacesMap["event_callback_url"] = workspace.EventCallbackURL

		if workspace.EventsFilter != nil && *workspace.EventsFilter != "" {
			workspacesMap["event_filters"] = strings.Split(*workspace.EventsFilter, ",")
		}

		workspacesMap["default_activity_name"] = workspace.DefaultActivityName
		workspacesMap["default_activity_sid"] = workspace.DefaultActivitySid
		workspacesMap["multi_task_enabled"] = workspace.MultiTaskEnabled
		workspacesMap["prioritize_queue_order"] = workspace.PrioritizeQueueOrder
		workspacesMap["timeout_activity_name"] = workspace.TimeoutActivityName
		workspacesMap["timeout_activity_sid"] = workspace.TimeoutActivitySid
		workspacesMap["date_created"] = workspace.DateCreated.Format(time.RFC3339)

		if workspace.DateUpdated != nil {
			workspacesMap["date_updated"] = workspace.DateUpdated.Format(time.RFC3339)
		}

		workspacesMap["url"] = workspace.URL

		workspaces = append(workspaces, workspacesMap)
	}

	d.Set("workspaces", &workspaces)

	return nil
}
