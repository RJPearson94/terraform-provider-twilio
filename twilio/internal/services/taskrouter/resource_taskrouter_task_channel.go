package taskrouter

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channel"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_channels"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTaskRouterTaskChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterTaskChannelCreate,
		Read:   resourceTaskRouterTaskChannelRead,
		Update: resourceTaskRouterTaskChannelUpdate,
		Delete: resourceTaskRouterTaskChannelDelete,
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
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"channel_optimized_routing": {
				Type:     schema.TypeBool,
				Optional: true,
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

func resourceTaskRouterTaskChannelCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &task_channels.CreateTaskChannelInput{
		FriendlyName:            d.Get("friendly_name").(string),
		UniqueName:              d.Get("unique_name").(string),
		ChannelOptimizedRouting: utils.OptionalBool(d, "channel_optimized_routing"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannels.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create task channel: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterTaskChannelRead(d, meta)
}

func resourceTaskRouterTaskChannelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannel(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read task channel: %s", err)
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

func resourceTaskRouterTaskChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &task_channel.UpdateTaskChannelInput{
		FriendlyName:            utils.OptionalString(d, "friendly_name"),
		ChannelOptimizedRouting: utils.OptionalBool(d, "channel_optimized_routing"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannel(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update task channel: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterTaskChannelRead(d, meta)
}

func resourceTaskRouterTaskChannelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).TaskChannel(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete task channel: %s", err.Error())
	}
	d.SetId("")
	return nil
}
