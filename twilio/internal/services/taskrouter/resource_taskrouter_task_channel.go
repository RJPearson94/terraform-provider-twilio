package taskrouter

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channels"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTaskRouterTaskChannel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTaskRouterTaskChannelCreate,
		ReadContext:   resourceTaskRouterTaskChannelRead,
		UpdateContext: resourceTaskRouterTaskChannelUpdate,
		DeleteContext: resourceTaskRouterTaskChannelDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Workspaces/(.*)/TaskChannels/(.*)"
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this task channel by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this task channel",
			},
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
				Description:  "The SID of the TaskRouter workspace. Changing this forces a new resource",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A human-readable label for the task channel",
			},
			"unique_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A unique identifier for the task channel. Changing this forces a new resource",
			},
			"channel_optimized_routing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the task channel should prioritize optimized routing. Defaults to `false`",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the task channel was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the task channel was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the task channel resource",
			},
		},
	}
}

func resourceTaskRouterTaskChannelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &task_channels.CreateTaskChannelInput{
		FriendlyName:            d.Get("friendly_name").(string),
		UniqueName:              d.Get("unique_name").(string),
		ChannelOptimizedRouting: utils.OptionalBool(d, "channel_optimized_routing"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannels.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create task channel: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterTaskChannelRead(ctx, d, meta)
}

func resourceTaskRouterTaskChannelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannel(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read task channel: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("channel_optimized_routing", getResponse.ChannelOptimizedRouting)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceTaskRouterTaskChannelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &task_channel.UpdateTaskChannelInput{
		FriendlyName:            utils.OptionalString(d, "friendly_name"),
		ChannelOptimizedRouting: utils.OptionalBool(d, "channel_optimized_routing"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannel(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update task channel: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterTaskChannelRead(ctx, d, meta)
}

func resourceTaskRouterTaskChannelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannel(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete task channel: %s", err.Error())
	}
	d.SetId("")
	return nil
}
