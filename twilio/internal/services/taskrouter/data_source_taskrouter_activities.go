package taskrouter

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceTaskRouterActivities() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTaskRouterActivitiesRead,

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
			"activities": &schema.Schema{
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

func dataSourceTaskRouterActivitiesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	workspaceSid := d.Get("workspace_sid").(string)
	paginator := client.Workspace(workspaceSid).Activities.NewActivitiesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No activities were found for taskrouter workspace with sid (%s)", workspaceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read taskrouter activity: %s", err.Error())
	}

	d.SetId(workspaceSid)
	d.Set("workspace_sid", workspaceSid)

	activities := make([]interface{}, 0)

	for _, activity := range paginator.Activities {
		d.Set("account_sid", activity.AccountSid)

		activitiesMap := make(map[string]interface{})

		activitiesMap["sid"] = activity.Sid
		activitiesMap["friendly_name"] = activity.FriendlyName
		activitiesMap["available"] = activity.Available
		activitiesMap["date_created"] = activity.DateCreated.Format(time.RFC3339)

		if activity.DateUpdated != nil {
			activitiesMap["date_updated"] = activity.DateUpdated.Format(time.RFC3339)
		}

		activitiesMap["url"] = activity.URL

		activities = append(activities, activitiesMap)
	}

	d.Set("activities", &activities)

	return nil
}
