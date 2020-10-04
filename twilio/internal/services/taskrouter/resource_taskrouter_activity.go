package taskrouter

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activities"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTaskRouterActivity() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTaskRouterActivityCreate,
		ReadContext:   resourceTaskRouterActivityRead,
		UpdateContext: resourceTaskRouterActivityUpdate,
		DeleteContext: resourceTaskRouterActivityDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Workspaces/(.*)/Activities/(.*)"
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"available": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
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

func resourceTaskRouterActivityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &activities.CreateActivityInput{
		FriendlyName: d.Get("friendly_name").(string),
		Available:    utils.OptionalBool(d, "available"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Activities.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create taskrouter activity: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterActivityRead(ctx, d, meta)
}

func resourceTaskRouterActivityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read taskrouter activity: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("available", getResponse.Available)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceTaskRouterActivityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &activity.UpdateActivityInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update taskrouter activity: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterActivityRead(ctx, d, meta)
}

func resourceTaskRouterActivityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete taskrouter activity: %s", err.Error())
	}
	d.SetId("")
	return nil
}
