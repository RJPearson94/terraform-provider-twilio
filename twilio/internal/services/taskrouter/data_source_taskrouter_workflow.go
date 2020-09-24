package taskrouter

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceTaskRouterWorkflow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRouterWorkflowRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
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
	}
}

func dataSourceTaskRouterWorkflowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Workspace(workspaceSid).Workflow(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Workflow with sid (%s) was not found for taskrouter workspace with sid (%s)", sid, workspaceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read workflow: %s", err.Error())
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
