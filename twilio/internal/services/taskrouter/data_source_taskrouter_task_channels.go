package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterTaskChannels() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterTaskChannelsRead,

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
			"task_channels": {
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
						"unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"channel_optimized_routing": {
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

func dataSourceTaskRouterTaskChannelsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	workspaceSid := d.Get("workspace_sid").(string)
	paginator := client.Workspace(workspaceSid).TaskChannels.NewTaskChannelsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No task channels were found for taskrouter workspace with sid (%s)", workspaceSid)
		}
		return diag.Errorf("Failed to read task channel: %s", err.Error())
	}

	d.SetId(workspaceSid)
	d.Set("workspace_sid", workspaceSid)

	taskChannels := make([]interface{}, 0)

	for _, taskChannel := range paginator.TaskChannels {
		d.Set("account_sid", taskChannel.AccountSid)

		taskChannelsMap := make(map[string]interface{})

		taskChannelsMap["sid"] = taskChannel.Sid
		taskChannelsMap["friendly_name"] = taskChannel.FriendlyName
		taskChannelsMap["unique_name"] = taskChannel.UniqueName
		taskChannelsMap["channel_optimized_routing"] = taskChannel.ChannelOptimizedRouting
		taskChannelsMap["date_created"] = taskChannel.DateCreated.Format(time.RFC3339)

		if taskChannel.DateUpdated != nil {
			taskChannelsMap["date_updated"] = taskChannel.DateUpdated.Format(time.RFC3339)
		}

		taskChannelsMap["url"] = taskChannel.URL

		taskChannels = append(taskChannels, taskChannelsMap)
	}

	d.Set("task_channels", &taskChannels)

	return nil
}
