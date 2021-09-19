package taskrouter

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTaskRouterWorkspaceConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTaskRouterWorkspaceConfigurationCreate,
		ReadContext:   resourceTaskRouterWorkspaceConfigurationRead,
		UpdateContext: resourceTaskRouterWorkspaceConfigurationUpdate,
		DeleteContext: resourceTaskRouterWorkspaceConfigurationDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
			},
			"default_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_activity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.TaskRouterActivitySidValidation(),
			},
			"timeout_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_activity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.TaskRouterActivitySidValidation(),
			},
		},
	}
}

func resourceTaskRouterWorkspaceConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TaskRouter Workspace configuration already exists so updating the configuration
	return resourceTaskRouterWorkspaceConfigurationUpdate(ctx, d, meta)
}

func resourceTaskRouterWorkspaceConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read taskrouter workspace configuration: %s", err.Error())
	}

	d.Set("default_activity_name", getResponse.DefaultActivityName)
	d.Set("default_activity_sid", getResponse.DefaultActivitySid)
	d.Set("timeout_activity_name", getResponse.TimeoutActivityName)
	d.Set("timeout_activity_sid", getResponse.TimeoutActivitySid)

	return nil
}

func resourceTaskRouterWorkspaceConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &workspace.UpdateWorkspaceInput{
		DefaultActivitySid: utils.OptionalString(d, "default_activity_sid"),
		TimeoutActivitySid: utils.OptionalString(d, "timeout_activity_sid"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update taskrouter workspace configuration: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkspaceConfigurationRead(ctx, d, meta)
}

func resourceTaskRouterWorkspaceConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] TaskRouter workspace configuration cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}
