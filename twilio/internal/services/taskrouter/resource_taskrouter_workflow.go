package taskrouter

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflow"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workflows"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceTaskRouterWorkflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterWorkflowCreate,
		Read:   resourceTaskRouterWorkflowRead,
		Update: resourceTaskRouterWorkflowUpdate,
		Delete: resourceTaskRouterWorkflowDelete,
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
			"workspace_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fallback_assignment_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assignment_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_reservation_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

func resourceTaskRouterWorkflowCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &workflows.CreateWorkflowInput{
		FriendlyName:                  d.Get("friendly_name").(string),
		Configuration:                 d.Get("configuration").(string),
		AssignmentCallbackURL:         utils.OptionalString(d, "assignment_callback_url"),
		FallbackAssignmentCallbackURL: utils.OptionalString(d, "fallback_assignment_callback_url"),
		TaskReservationTimeout:        utils.OptionalInt(d, "task_reservation_timeout"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Workflows.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create workflow: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkflowRead(d, meta)
}

func resourceTaskRouterWorkflowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Workflow(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read workflow: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("fallback_assignment_callback_url", getResponse.FallbackAssignmentCallbackURL)
	d.Set("assignment_callback_url", getResponse.AssignmentCallbackURL)
	d.Set("task_reservation_timeout", getResponse.TaskReservationTimeout)
	d.Set("document_content_type", getResponse.DocumentContentType)
	d.Set("configuration", getResponse.Configuration.(string))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceTaskRouterWorkflowUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &workflow.UpdateWorkflowInput{
		FriendlyName:                  utils.OptionalString(d, "friendly_name"),
		Configuration:                 utils.OptionalString(d, "configuration"),
		AssignmentCallbackURL:         utils.OptionalString(d, "assignment_callback_url"),
		FallbackAssignmentCallbackURL: utils.OptionalString(d, "fallback_assignment_callback_url"),
		TaskReservationTimeout:        utils.OptionalInt(d, "task_reservation_timeout"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Workflow(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update workflow: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkflowRead(d, meta)
}

func resourceTaskRouterWorkflowDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).Workflow(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete workflow: %s", err.Error())
	}
	d.SetId("")
	return nil
}
