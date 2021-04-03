package taskrouter

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/taskrouter/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTaskRouterWorker() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTaskRouterWorkerCreate,
		ReadContext:   resourceTaskRouterWorkerRead,
		UpdateContext: resourceTaskRouterWorkerUpdate,
		DeleteContext: resourceTaskRouterWorkerDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Workspaces/(.*)/Workers/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("workspace_sid", match[1])
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
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: helper.WorkspaceSidValidation(),
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"activity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: helper.ActivitySidValidation(),
			},
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "{}",
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available": {
				Type:     schema.TypeBool,
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
			"date_status_changed": {
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

func resourceTaskRouterWorkerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &workers.CreateWorkerInput{
		FriendlyName: d.Get("friendly_name").(string),
		ActivitySid:  utils.OptionalString(d, "activity_sid"),
		Attributes:   utils.OptionalJSONString(d, "attributes"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Workers.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create taskrouter worker: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkerRead(ctx, d, meta)
}

func resourceTaskRouterWorkerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read taskrouter worker: %s", err.Error())
	}

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

func resourceTaskRouterWorkerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	workerActivitySid, workerActivitySidErr := optionalWorkerActivitySid(ctx, d, meta)
	if workerActivitySidErr != nil {
		return workerActivitySidErr
	}

	updateInput := &worker.UpdateWorkerInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		ActivitySid:  workerActivitySid,
		Attributes:   utils.OptionalJSONString(d, "attributes"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update taskrouter worker: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkerRead(ctx, d, meta)
}

func resourceTaskRouterWorkerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete taskrouter worker: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func optionalWorkerActivitySid(ctx context.Context, d *schema.ResourceData, meta interface{}) (*string, diag.Diagnostics) {
	activitySidSchemaKey := "activity_sid"
	if v, ok := d.GetOk(activitySidSchemaKey); ok {
		return sdkUtils.String(v.(string)), nil
	}
	if ok := d.HasChange(activitySidSchemaKey); ok {
		getResponse, err := meta.(*common.TwilioClient).TaskRouter.Workspace(d.Get("workspace_sid").(string)).FetchWithContext(ctx)
		if err != nil {
			return nil, diag.Errorf("Failed to read taskrouter workspace to get default activity sid for worker: %s", err.Error())
		}
		return sdkUtils.String(getResponse.DefaultActivitySid), nil
	}
	return nil, nil
}
