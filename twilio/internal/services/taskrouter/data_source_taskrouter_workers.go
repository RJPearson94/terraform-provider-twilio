package taskrouter

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceTaskRouterWorkers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRouterWorkersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"workspace_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"workers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"activity_sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceTaskRouterWorkersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	workspaceSid := d.Get("workspace_sid").(string)
	paginator := client.Workspace(workspaceSid).Workers.NewWorkersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No workers were found for taskrouter workspace with sid (%s)", workspaceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter worker: %s", err.Error())
	}

	d.SetId(workspaceSid)
	d.Set("workspace_sid", workspaceSid)

	workers := make([]interface{}, 0)

	for _, worker := range paginator.Workers {
		d.Set("account_sid", worker.AccountSid)

		workersMap := make(map[string]interface{})

		workersMap["sid"] = worker.Sid
		workersMap["friendly_name"] = worker.FriendlyName
		workersMap["activity_sid"] = worker.ActivitySid
		workersMap["attributes"] = worker.Attributes
		workersMap["activity_name"] = worker.ActivityName
		workersMap["available"] = worker.Available
		workersMap["date_created"] = worker.DateCreated.Format(time.RFC3339)

		if worker.DateUpdated != nil {
			workersMap["date_updated"] = worker.DateUpdated.Format(time.RFC3339)
		}

		if worker.DateStatusChanged != nil {
			workersMap["date_status_changed"] = worker.DateStatusChanged.Format(time.RFC3339)
		}

		workersMap["url"] = worker.URL

		workers = append(workers, workersMap)
	}

	d.Set("workers", &workers)

	return nil
}
