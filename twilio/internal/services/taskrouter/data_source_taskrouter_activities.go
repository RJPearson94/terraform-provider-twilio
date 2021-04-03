package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterActivities() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterActivitiesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"activities": {
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

func dataSourceTaskRouterActivitiesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	workspaceSid := d.Get("workspace_sid").(string)
	paginator := client.Workspace(workspaceSid).Activities.NewActivitiesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No activities were found for taskrouter workspace with sid (%s)", workspaceSid)
		}
		return diag.Errorf("Failed to read taskrouter activity: %s", err.Error())
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
