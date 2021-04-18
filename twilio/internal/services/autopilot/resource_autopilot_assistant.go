package autopilot

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistants"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const callbackEventsSeperator = " "

func resourceAutopilotAssistant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAutopilotAssistantCreate,
		ReadContext:   resourceAutopilotAssistantRead,
		UpdateContext: resourceAutopilotAssistantUpdate,
		DeleteContext: resourceAutopilotAssistantDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Assistants/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
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
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 255),
			},
			"unique_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},
			"callback_events": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"log_queries": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"development_stage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"in-development",
					"in-production",
				}, false),
			},
			"needs_model_build": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"defaults": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"stylesheet": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
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

func resourceAutopilotAssistantCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	createInput := &assistants.CreateAssistantInput{
		FriendlyName:   utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		UniqueName:     utils.OptionalStringWithEmptyStringOnChange(d, "unique_name"),
		LogQueries:     utils.OptionalBool(d, "log_queries"),
		CallbackURL:    utils.OptionalStringWithEmptyStringOnChange(d, "callback_url"),
		CallbackEvents: utils.OptionalSeperatedStringWithEmptyStringOnChange(d, "callback_events", callbackEventsSeperator),
		Defaults:       utils.OptionalJSONString(d, "defaults"),
		StyleSheet:     utils.OptionalJSONString(d, "stylesheet"),
	}

	createResult, err := client.Assistants.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create autopilot assistant: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if d.Get("development_stage") != nil && d.Get("development_stage").(string) != "in-development" {
		return resourceAutopilotAssistantUpdate(ctx, d, meta)
	}
	return resourceAutopilotAssistantRead(ctx, d, meta)
}

func resourceAutopilotAssistantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	getResponse, err := client.Assistant(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read autopilot assistant: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("latest_model_build_sid", getResponse.LatestModelBuildSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)

	if getResponse.CallbackEvents != nil && *getResponse.CallbackEvents != "" {
		d.Set("callback_events", strings.Split(*getResponse.CallbackEvents, callbackEventsSeperator))
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
	defaultsJSONString, err := structure.FlattenJsonToString(getDefaultsResponse.Data)
	if err != nil {
		return diag.Errorf("Unable to flatten defaults json to string: %s", err.Error())
	}
	d.Set("defaults", defaultsJSONString)

	getStyleSheetResponse, err := client.Assistant(d.Id()).StyleSheet().FetchWithContext(ctx)
	if err != nil {
		return diag.Errorf("Failed to read autopilot assistant stylesheet: %s", err.Error())
	}
	styleSheetJSONString, err := structure.FlattenJsonToString(getStyleSheetResponse.Data)
	if err != nil {
		return diag.Errorf("Unable to flatten stylesheet json to string: %s", err.Error())
	}
	d.Set("stylesheet", styleSheetJSONString)

	return nil
}

func resourceAutopilotAssistantUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	updateInput := &assistant.UpdateAssistantInput{
		FriendlyName:     utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		UniqueName:       utils.OptionalString(d, "unique_name"),
		LogQueries:       utils.OptionalBool(d, "log_queries"),
		CallbackURL:      utils.OptionalStringWithEmptyStringOnChange(d, "callback_url"),
		CallbackEvents:   utils.OptionalSeperatedStringWithEmptyStringOnChange(d, "callback_events", callbackEventsSeperator),
		DevelopmentStage: utils.OptionalString(d, "development_stage"),
		Defaults:         utils.OptionalJSONString(d, "defaults"),
		StyleSheet:       utils.OptionalJSONString(d, "stylesheet"),
	}

	updateResp, err := client.Assistant(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update autopilot assistant: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotAssistantRead(ctx, d, meta)
}

func resourceAutopilotAssistantDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	if err := client.Assistant(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete autopilot assistant: %s", err.Error())
	}
	d.SetId("")
	return nil
}
