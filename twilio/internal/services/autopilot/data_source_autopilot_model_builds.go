package autopilot

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAutopilotModelBuilds() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutopilotModelBuildsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model_builds": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Required: true,
						},
						"unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"build_duration": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": {
							Type:     schema.TypeInt,
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

func dataSourceAutopilotModelBuildsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	assistantSid := d.Get("assistant_sid").(string)
	paginator := client.Assistant(assistantSid).ModelBuilds.NewModelBuildsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No model builds were found for assistant with sid (%s)", assistantSid)
		}
		return fmt.Errorf("[ERROR] Failed to list autopilot model builds: %s", err.Error())
	}

	d.SetId(assistantSid)
	d.Set("assistant_sid", assistantSid)

	modelBuilds := make([]interface{}, len(paginator.ModelBuilds)-1)

	for _, modelBuild := range paginator.ModelBuilds {
		d.Set("account_sid", modelBuild.AccountSid)

		modelBuildMap := make(map[string]interface{})

		modelBuildMap["sid"] = modelBuild.Sid
		modelBuildMap["unique_name"] = modelBuild.UniqueName
		modelBuildMap["build_duration"] = modelBuild.BuildDuration
		modelBuildMap["status"] = modelBuild.Status
		modelBuildMap["error_code"] = modelBuild.ErrorCode
		modelBuildMap["date_created"] = modelBuild.DateCreated.Format(time.RFC3339)

		if modelBuild.DateUpdated != nil {
			modelBuildMap["date_updated"] = modelBuild.DateUpdated.Format(time.RFC3339)
		}

		modelBuildMap["url"] = modelBuild.URL

		modelBuilds = append(modelBuilds, modelBuildMap)
	}

	d.Set("model_builds", &modelBuilds)

	return nil
}
