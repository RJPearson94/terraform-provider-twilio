package autopilot

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAutopilotModelBuilds() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotModelBuildsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
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

func dataSourceAutopilotModelBuildsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	assistantSid := d.Get("assistant_sid").(string)
	paginator := client.Assistant(assistantSid).ModelBuilds.NewModelBuildsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No model builds were found for assistant with sid (%s)", assistantSid)
		}
		return diag.Errorf("Failed to list autopilot model builds: %s", err.Error())
	}

	d.SetId(assistantSid)
	d.Set("assistant_sid", assistantSid)

	modelBuilds := make([]interface{}, 0)

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
