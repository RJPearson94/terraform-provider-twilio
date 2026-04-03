package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterWorkflowRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkflowSidValidation(),
				Description:  "The SID of the workflow to retrieve",
			},
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
				Description:  "The SID of the TaskRouter workspace",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this workflow",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the workflow",
			},
			"fallback_assignment_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to call when a task assignment event is not handled by the primary callback",
			},
			"assignment_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to call when a task is assigned to a worker",
			},
			"task_reservation_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The timeout in seconds for a task reservation",
			},
			"document_content_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The MIME type of the workflow document",
			},
			"configuration": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string of the workflow configuration",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the workflow was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the workflow was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the workflow resource",
			},
		},
	}
}

func dataSourceTaskRouterWorkflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Workspace(workspaceSid).Workflow(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Workflow with sid (%s) was not found for taskrouter workspace with sid (%s)", sid, workspaceSid)
		}
		return diag.Errorf("Failed to read workflow: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("fallback_assignment_callback_url", getResponse.FallbackAssignmentCallbackURL)
	d.Set("assignment_callback_url", getResponse.AssignmentCallbackURL)
	d.Set("task_reservation_timeout", getResponse.TaskReservationTimeout)
	d.Set("document_content_type", getResponse.DocumentContentType)
	d.Set("configuration", getResponse.Configuration)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
