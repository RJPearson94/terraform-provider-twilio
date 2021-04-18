package taskrouter

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflow"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflows"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTaskRouterWorkflow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTaskRouterWorkflowCreate,
		ReadContext:   resourceTaskRouterWorkflowRead,
		UpdateContext: resourceTaskRouterWorkflowUpdate,
		DeleteContext: resourceTaskRouterWorkflowDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Workspaces/(.*)/Workflows/(.*)"
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
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"fallback_assignment_callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"assignment_callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"task_reservation_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      120,
				ValidateFunc: validation.IntBetween(1, 86400),
			},
			"document_content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
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

func resourceTaskRouterWorkflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	configurationJSONString, _ := structure.NormalizeJsonString(d.Get("configuration").(string))

	createInput := &workflows.CreateWorkflowInput{
		FriendlyName:                  d.Get("friendly_name").(string),
		Configuration:                 configurationJSONString,
		AssignmentCallbackURL:         utils.OptionalStringWithEmptyStringOnChange(d, "assignment_callback_url"),
		FallbackAssignmentCallbackURL: utils.OptionalStringWithEmptyStringOnChange(d, "fallback_assignment_callback_url"),
		TaskReservationTimeout:        utils.OptionalInt(d, "task_reservation_timeout"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Workflows.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create workflow: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkflowRead(ctx, d, meta)
}

func resourceTaskRouterWorkflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Workflow(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read workflow: %s", err.Error())
	}

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

func resourceTaskRouterWorkflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &workflow.UpdateWorkflowInput{
		FriendlyName:                  utils.OptionalString(d, "friendly_name"),
		Configuration:                 utils.OptionalJSONString(d, "configuration"),
		AssignmentCallbackURL:         utils.OptionalStringWithEmptyStringOnChange(d, "assignment_callback_url"),
		FallbackAssignmentCallbackURL: utils.OptionalStringWithEmptyStringOnChange(d, "fallback_assignment_callback_url"),
		TaskReservationTimeout:        utils.OptionalInt(d, "task_reservation_timeout"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Workflow(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update workflow: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkflowRead(ctx, d, meta)
}

func resourceTaskRouterWorkflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).Workflow(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete workflow: %s", err.Error())
	}
	d.SetId("")
	return nil
}
