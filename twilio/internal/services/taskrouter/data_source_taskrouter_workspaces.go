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
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"workspaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_callback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_filters": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"multi_task_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"prioritize_queue_order": {
							Type:     schema.TypeString,
							Computed: true,
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
