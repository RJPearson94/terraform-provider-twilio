package autopilot

import (
	"context"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
)

func dataSourceAutopilotAssistant() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotAssistantRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_model_build_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"callback_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_queries": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"development_stage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"needs_model_build": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"defaults": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stylesheet": {
				Type:     schema.TypeString,
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
	}
}

func dataSourceAutopilotAssistantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	sid := d.Get("sid").(string)
	getResponse, err := client.Assistant(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Assistant with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read autopilot assistant: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("latest_model_build_sid", getResponse.LatestModelBuildSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)

	if getResponse.CallbackEvents != nil {
		d.Set("callback_events", strings.Split(*getResponse.CallbackEvents, " "))
	}

	d.Set("callback_url", getResponse.CallbackURL)
	d.Set("log_queries", getResponse.LogQueries)
	d.Set("development_stage", getResponse.DevelopmentStage)
	d.Set("needs_model_build", getResponse.NeedsModelBuild)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	getDefaultsResponse, err := client.Assistant(d.Id()).Defaults().FetchWithContext(ctx)
	if err != nil {
		return diag.Errorf("Failed to read autopilot assistant defaults: %s", err.Error())
	}
	defaultsJSON, err := structure.FlattenJsonToString(getDefaultsResponse.Data)
	if err != nil {
		return diag.Errorf("Unable to flatten defaults json to string: %s", err.Error())
	}
	d.Set("defaults", defaultsJSON)

	getStyleSheetResponse, err := client.Assistant(d.Id()).StyleSheet().FetchWithContext(ctx)
	if err != nil {
		return diag.Errorf("Failed to read autopilot assistant stylesheet: %s", err.Error())
	}
	styleSheetJSON, err := structure.FlattenJsonToString(getStyleSheetResponse.Data)
	if err != nil {
		return diag.Errorf("Unable to flatten stylesheet json to string: %s", err.Error())
	}
	d.Set("stylesheet", styleSheetJSON)

	return nil
}
