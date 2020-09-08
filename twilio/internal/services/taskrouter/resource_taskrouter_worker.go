package taskrouter

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/worker"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceTaskRouterWorker() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterWorkerCreate,
		Read:   resourceTaskRouterWorkerRead,
		Update: resourceTaskRouterWorkerUpdate,
		Delete: resourceTaskRouterWorkerDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Workspaces/(.*)/Workers/(.*)"
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
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"workspace_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"activity_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attributes": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"activity_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"available": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_status_changed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTaskRouterWorkerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &workers.CreateWorkerInput{
		FriendlyName: d.Get("friendly_name").(string),
		ActivitySid:  utils.OptionalString(d, "activity_sid"),
		Attributes:   utils.OptionalJSONString(d, "attributes"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).Workers.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create taskrouter worker: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterWorkerRead(d, meta)
}

func resourceTaskRouterWorkerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter worker: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("activity_sid", getResponse.ActivitySid)
	d.Set("attributes", getResponse.Attributes)
	d.Set("activity_name", getResponse.ActivityName)
	d.Set("available", getResponse.Available)
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
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &worker.UpdateWorkerInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		ActivitySid:  utils.OptionalString(d, "activity_sid"),
		Attributes:   utils.OptionalJSONString(d, "attributes"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update taskrouter worker: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterWorkerRead(d, meta)
}

func resourceTaskRouterWorkerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Workspace(d.Get("workspace_sid").(string)).Worker(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete taskrouter worker: %s", err.Error())
	}
	d.SetId("")
	return nil
}
