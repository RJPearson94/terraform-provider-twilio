package taskrouter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceTaskRouterWorkspace() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRouterWorkspaceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"event_callback_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"event_filters": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"multi_task_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"prioritize_queue_order": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_activity_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_activity_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_activity_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_activity_sid": &schema.Schema{
				Type:     schema.TypeString,
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
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTaskRouterWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	sid := d.Get("sid").(string)
	getResponse, err := client.Workspace(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] TaskRouter workspace with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter workspace: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("event_callback_url", getResponse.EventCallbackURL)

	if getResponse.EventsFilter != nil {
		d.Set("event_filters", strings.Split(*getResponse.EventsFilter, ","))
	}

	d.Set("default_activity_name", getResponse.DefaultActivityName)
	d.Set("default_activity_sid", getResponse.DefaultActivitySid)
	d.Set("multi_task_enabled", getResponse.MultiTaskEnabled)
	d.Set("prioritize_queue_order", getResponse.PrioritizeQueueOrder)
	d.Set("timeout_activity_name", getResponse.TimeoutActivityName)
	d.Set("timeout_activity_sid", getResponse.TimeoutActivitySid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
