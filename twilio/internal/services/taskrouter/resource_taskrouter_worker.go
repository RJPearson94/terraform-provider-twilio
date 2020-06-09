package taskrouter

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTaskRouterWorker() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterWorkerCreate,
		Read:   resourceTaskRouterWorkerRead,
		Update: resourceTaskRouterWorkerUpdate,
		Delete: resourceTaskRouterWorkerDelete,
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
			"activity_sid": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("activity_sid_default") != nil {
						return old == d.Get("activity_sid_default").(string) && new == ""
					}
					return false
				},
			},
			"activity_sid_default": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "{}",
			},
			"activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available": {
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
			"date_status_changed": {
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

func resourceTaskRouterWorkerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &workers.CreateWorkerInput{
		FriendlyName: d.Get("friendly_name").(string),
		ActivitySid:  d.Get("activity_sid").(string),
		Attributes:   d.Get("attributes").(string),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Workers.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create taskrouter worker: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkerRead(d, meta)
}

func resourceTaskRouterWorkerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter worker: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("friendly_name", getResponse.FriendlyName)

	if d.Get("activity_sid").(string) == "" {
		d.Set("activity_sid_default", getResponse.ActivitySid)
	}

	d.Set("activity_sid", getResponse.ActivitySid)
	d.Set("attributes", getResponse.Attributes)
	d.Set("activity_name", getResponse.ActivityName)
	d.Set("available", getResponse.Available)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	if getResponse.DateStatusChanged != nil {
		d.Set("date_status_changed", getResponse.DateStatusChanged.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceTaskRouterWorkerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &worker.UpdateWorkerInput{
		FriendlyName: d.Get("friendly_name").(string),
		ActivitySid:  d.Get("activity_sid").(string),
		Attributes:   d.Get("attributes").(string),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update taskrouter worker: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkerRead(d, meta)
}

func resourceTaskRouterWorkerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete taskrouter worker: %s", err.Error())
	}
	d.SetId("")
	return nil
}
