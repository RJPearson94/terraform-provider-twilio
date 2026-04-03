package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterWorker() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterWorkerRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkerSidValidation(),
				Description:  "The SID of the worker to retrieve",
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
				Description: "The SID of the account that owns this worker",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the worker",
			},
			"activity_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the worker's current activity",
			},
			"attributes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string of attributes for the worker",
			},
			"activity_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The friendly name of the worker's current activity",
			},
			"available": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the worker is available to receive tasks",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the worker was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the worker was last updated, in RFC 3339 format",
			},
			"date_status_changed": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the worker's activity status last changed, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the worker resource",
			},
		},
	}
}

func dataSourceTaskRouterWorkerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Workspace(workspaceSid).Worker(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Worker with sid (%s) was not found for taskrouter workspace with sid (%s)", sid, workspaceSid)
		}
		return diag.Errorf("Failed to read taskrouter worker: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("activity_sid", getResponse.ActivitySid)
	d.Set("attributes", getResponse.Attributes)
	d.Set("activity_name", getResponse.ActivityName)
	d.Set("available", getResponse.Available)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	if getResponse.DateStatusChanged != nil {
		d.Set("date_status_changed", getResponse.DateStatusChanged.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
