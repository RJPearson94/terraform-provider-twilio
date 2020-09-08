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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const callbackEventsSeperator = " "

func resourceAutopilotAssistant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutopilotAssistantCreate,
		Read:   resourceAutopilotAssistantRead,
		Update: resourceAutopilotAssistantUpdate,
		Delete: resourceAutopilotAssistantDelete,

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
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_model_build_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unique_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"callback_events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"callback_url": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"log_queries": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"development_stage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"in-development",
					"in-production",
				}, false),
			},
			"needs_model_build": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"defaults": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"stylesheet": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
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

func resourceAutopilotAssistantCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &assistants.CreateAssistantInput{
		FriendlyName:   utils.OptionalString(d, "friendly_name"),
		UniqueName:     utils.OptionalString(d, "unique_name"),
		LogQueries:     utils.OptionalBool(d, "log_queries"),
		CallbackURL:    utils.OptionalString(d, "callback_url"),
		CallbackEvents: utils.OptionalSeperatedString(d, "callback_events", callbackEventsSeperator),
		Defaults:       utils.OptionalJSONString(d, "defaults"),
		StyleSheet:     utils.OptionalJSONString(d, "stylesheet"),
	}

	createResult, err := client.Assistants.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create autopilot assistant: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if d.Get("development_stage") != nil && d.Get("development_stage").(string) != "in-development" {
		return resourceAutopilotAssistantUpdate(d, meta)
	}
	return resourceAutopilotAssistantRead(d, meta)
}

func resourceAutopilotAssistantRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Assistant(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot assistant: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("latest_model_build_sid", getResponse.LatestModelBuildSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)

	if getResponse.CallbackEvents != nil {
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
		return fmt.Errorf("[ERROR] Failed to read autopilot assistant defaults: %s", err.Error())
	}
	defaultsJSONString, err := structure.FlattenJsonToString(getDefaultsResponse.Data)
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to flatten defaults json to string: %s", err.Error())
	}
	d.Set("defaults", defaultsJSONString)

	getStyleSheetResponse, err := client.Assistant(d.Id()).StyleSheet().FetchWithContext(ctx)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to read autopilot assistant stylesheet: %s", err.Error())
	}
	styleSheetJSONString, err := structure.FlattenJsonToString(getStyleSheetResponse.Data)
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to flatten stylesheet json to string: %s", err.Error())
	}
	d.Set("stylesheet", styleSheetJSONString)

	return nil
}

func resourceAutopilotAssistantUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &assistant.UpdateAssistantInput{
		FriendlyName:     utils.OptionalString(d, "friendly_name"),
		UniqueName:       utils.OptionalString(d, "unique_name"),
		LogQueries:       utils.OptionalBool(d, "log_queries"),
		CallbackURL:      utils.OptionalString(d, "callback_url"),
		CallbackEvents:   utils.OptionalSeperatedString(d, "callback_events", callbackEventsSeperator),
		DevelopmentStage: utils.OptionalString(d, "development_stage"),
		Defaults:         utils.OptionalJSONString(d, "defaults"),
		StyleSheet:       utils.OptionalJSONString(d, "stylesheet"),
	}

	updateResp, err := client.Assistant(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update autopilot assistant: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotAssistantRead(d, meta)
}

func resourceAutopilotAssistantDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Assistant(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete autopilot assistant: %s", err.Error())
	}
	d.SetId("")
	return nil
}
