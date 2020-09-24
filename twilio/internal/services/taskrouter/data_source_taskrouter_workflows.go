package taskrouter

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceTaskRouterWorkflows() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRouterWorkflowsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"workspace_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"workflows": {
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
						"fallback_assignment_callback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"assignment_callback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"task_reservation_timeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"document_content_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"configuration": {
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

func dataSourceTaskRouterWorkflowsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	workspaceSid := d.Get("workspace_sid").(string)
	paginator := client.Workspace(workspaceSid).Workflows.NewWorkflowsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No workflows were found for taskrouter workspace with sid (%s)", workspaceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read workflow: %s", err.Error())
	}

	d.SetId(workspaceSid)
	d.Set("workspace_sid", workspaceSid)

	workflows := make([]interface{}, 0)

	for _, workflow := range paginator.Workflows {
		d.Set("account_sid", workflow.AccountSid)

		workflowsMap := make(map[string]interface{})

		workflowsMap["sid"] = workflow.Sid
		workflowsMap["friendly_name"] = workflow.FriendlyName
		workflowsMap["fallback_assignment_callback_url"] = workflow.FallbackAssignmentCallbackURL
		workflowsMap["assignment_callback_url"] = workflow.AssignmentCallbackURL
		workflowsMap["task_reservation_timeout"] = workflow.TaskReservationTimeout
		workflowsMap["document_content_type"] = workflow.DocumentContentType
		workflowsMap["configuration"] = workflow.Configuration
		workflowsMap["date_created"] = workflow.DateCreated.Format(time.RFC3339)

		if workflow.DateUpdated != nil {
			workflowsMap["date_updated"] = workflow.DateUpdated.Format(time.RFC3339)
		}

		workflowsMap["url"] = workflow.URL

		workflows = append(workflows, workflowsMap)
	}

	d.Set("workflows", &workflows)

	return nil
}
