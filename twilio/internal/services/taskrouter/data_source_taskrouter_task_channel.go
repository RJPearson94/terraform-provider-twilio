package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterTaskChannel() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterTaskChannelRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.TaskRouterTaskChannelSidValidation(),
				ExactlyOneOf: []string{"sid", "unique_name"},
				Description:  "The SID of the task channel to retrieve. Exactly one of `sid` or `unique_name` must be specified",
			},
			"unique_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"sid", "unique_name"},
				Description:  "The unique name of the task channel to retrieve. Exactly one of `sid` or `unique_name` must be specified",
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
				Description: "The SID of the account that owns this task channel",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the task channel",
			},
			"channel_optimized_routing": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the task channel prioritizes optimized routing",
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

func dataSourceTaskRouterTaskChannelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	var identifier string

	if v, ok := d.GetOk("sid"); ok {
		identifier = v.(string)
	} else {
		identifier = d.Get("unique_name").(string)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	getResponse, err := client.Workspace(workspaceSid).TaskChannel(identifier).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Task channel with sid/ unique name (%s) was not found for taskrouter workspace with sid (%s)", identifier, workspaceSid)
		}
		return diag.Errorf("Failed to read task channel: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
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
