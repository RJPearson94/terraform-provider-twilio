package autopilot

import (
	"encoding/json"
	"fmt"
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
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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

func resourceAutopilotAssistantCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	createInput := &assistants.CreateAssistantInput{
		FriendlyName:   utils.OptionalString(d, "friendly_name"),
		UniqueName:     utils.OptionalString(d, "unique_name"),
		LogQueries:     utils.OptionalBool(d, "log_queries"),
		CallbackURL:    utils.OptionalString(d, "callback_url"),
		CallbackEvents: utils.OptionalSeperatedString(d, "callback_events", callbackEventsSeperator),
		Defaults:       utils.OptionalJSONString(d, "defaults"),
		StyleSheet:     utils.OptionalJSONString(d, "stylesheet"),
	}

	createResult, err := client.Assistants.Create(createInput)
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

	getResponse, err := client.Assistant(d.Id()).Get()
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

	getDefaultsResponse, err := client.Assistant(d.Id()).Defaults().Get()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to read autopilot assistant defaults: %s", err.Error())
	}
	if err := marshalJSONData(d, "defaults", getDefaultsResponse.Data); err != nil {
		return err
	}

	getStyleSheetResponse, err := client.Assistant(d.Id()).StyleSheet().Get()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to read autopilot assistant stylesheet: %s", err.Error())
	}
	if err := marshalJSONData(d, "stylesheet", getStyleSheetResponse.Data); err != nil {
		return err
	}

	return nil
}

func resourceAutopilotAssistantUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

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

	updateResp, err := client.Assistant(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update autopilot assistant: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotAssistantRead(d, meta)
}

func resourceAutopilotAssistantDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	if err := client.Assistant(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete autopilot assistant: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func marshalJSONData(d *schema.ResourceData, id string, data interface{}) error {
	if data != nil {
		jsonByteArray, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to marshal %s data to string: %s", id, err.Error())
		}
		d.Set(id, string(jsonByteArray))
	}
	return nil
}
