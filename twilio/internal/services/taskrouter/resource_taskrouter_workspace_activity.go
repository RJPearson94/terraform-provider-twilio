package taskrouter

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activities"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activity"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTaskRouterWorkspaceActivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterWorkspaceActivityCreate,
		Read:   resourceTaskRouterWorkspaceActivityRead,
		Update: resourceTaskRouterWorkspaceActivityUpdate,
		Delete: resourceTaskRouterWorkspaceActivityDelete,
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

func resourceTaskRouterWorkspaceActivityCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &activities.CreateActivityInput{
		FriendlyName: d.Get("friendly_name").(string),
		Available:    sdkUtils.Bool(d.Get("available").(bool)),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Activities.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create taskrouter workspace activity: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkspaceActivityRead(d, meta)
}

func resourceTaskRouterWorkspaceActivityRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter workspace activity: %s", err)
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

func resourceTaskRouterWorkspaceActivityUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &activity.UpdateActivityInput{
		FriendlyName: d.Get("friendly_name").(string),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update taskrouter workspace activity: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkspaceActivityRead(d, meta)
}

func resourceTaskRouterWorkspaceActivityDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete taskrouter workspace activity: %s", err.Error())
	}
	d.SetId("")
	return nil
}
