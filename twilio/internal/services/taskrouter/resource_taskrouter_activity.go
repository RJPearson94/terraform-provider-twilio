package taskrouter

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activities"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/activity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTaskRouterActivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterActivityCreate,
		Read:   resourceTaskRouterActivityRead,
		Update: resourceTaskRouterActivityUpdate,
		Delete: resourceTaskRouterActivityDelete,
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

func resourceTaskRouterActivityCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &activities.CreateActivityInput{
		FriendlyName: d.Get("friendly_name").(string),
		Available:    utils.OptionalBool(d, "available"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Activities.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create taskrouter activity: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterActivityRead(d, meta)
}

func resourceTaskRouterActivityRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter activity: %s", err)
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

func resourceTaskRouterActivityUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &activity.UpdateActivityInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update taskrouter activity: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterActivityRead(d, meta)
}

func resourceTaskRouterActivityDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).Activity(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete taskrouter activity: %s", err.Error())
	}
	d.SetId("")
	return nil
}
