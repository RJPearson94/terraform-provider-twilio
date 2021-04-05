package video

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/video/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_hook"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/composition_hooks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVideoCompositionHook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVideoCompositionHookCreate,
		ReadContext:   resourceVideoCompositionHookRead,
		UpdateContext: resourceVideoCompositionHookUpdate,
		DeleteContext: resourceVideoCompositionHookDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/CompositionHooks/(.*)"
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
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"audio_sources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"audio_sources_excluded": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "webm",
				ValidateFunc: validation.StringInSlice([]string{
					"mp4",
					"webm",
				}, false),
			},
			"resolution": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "640x480",
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d)+x(\d)+$`), ""),
			},
			"status_callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"status_callback_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "POST",
				ValidateFunc: validation.StringInSlice([]string{
					"GET",
					"POST",
				}, false),
			},
			"trim": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"video_layout": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "{}",
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

func resourceVideoCompositionHookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Video

	createInput := &composition_hooks.CreateCompositionHookInput{
		FriendlyName:         d.Get("friendly_name").(string),
		AudioSources:         utils.OptionalStringSlice(d, "audio_sources"),
		AudioSourcesExcluded: utils.OptionalStringSlice(d, "audio_sources_excluded"),
		Enabled:              utils.OptionalBool(d, "enabled"),
		Format:               utils.OptionalString(d, "format"),
		Resolution:           utils.OptionalString(d, "resolution"),
		StatusCallback:       utils.OptionalStringWithEmptyStringDefault(d, "status_callback_url"),
		StatusCallbackMethod: utils.OptionalString(d, "status_callback_method"),
		Trim:                 utils.OptionalBool(d, "trim"),
		VideoLayout:          utils.OptionalStringWithEmptyStringDefault(d, "video_layout"),
	}

	createResult, err := client.CompositionHooks.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create composition hook: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceVideoCompositionHookRead(ctx, d, meta)
}

func resourceVideoCompositionHookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Video

	getResponse, err := client.CompositionHook(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read composition hook: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("audio_sources", getResponse.AudioSources)
	d.Set("audio_sources_excluded", getResponse.AudioSourcesExcluded)
	d.Set("enabled", getResponse.Enabled)
	d.Set("format", getResponse.Format)
	d.Set("resolution", getResponse.Resolution)
	d.Set("status_callback_url", getResponse.StatusCallback)
	d.Set("status_callback_method", getResponse.StatusCallbackMethod)
	d.Set("trim", getResponse.Trim)

	videoLayout, err := helper.FlattenJsonToStringOrEmptyObjectString(getResponse.VideoLayout)
	if err != nil {
		return diag.Errorf("Unable to flatten video layout json to string. Error ", err.Error())
	}
	d.Set("video_layout", videoLayout)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}

func resourceVideoCompositionHookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Video

	updateInput := &composition_hook.UpdateCompositionHookInput{
		FriendlyName:         d.Get("friendly_name").(string),
		AudioSources:         utils.OptionalStringSlice(d, "audio_sources"),
		AudioSourcesExcluded: utils.OptionalStringSlice(d, "audio_sources_excluded"),
		Enabled:              utils.OptionalBool(d, "enabled"),
		Format:               utils.OptionalString(d, "format"),
		Resolution:           utils.OptionalString(d, "resolution"),
		StatusCallback:       utils.OptionalStringWithEmptyStringDefault(d, "status_callback_url"),
		StatusCallbackMethod: utils.OptionalString(d, "status_callback_method"),
		Trim:                 utils.OptionalBool(d, "trim"),
		VideoLayout:          utils.OptionalStringWithEmptyStringDefault(d, "video_layout"),
	}

	updateResp, err := client.CompositionHook(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update composition hook: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceVideoCompositionHookRead(ctx, d, meta)
}

func resourceVideoCompositionHookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Video

	if err := client.CompositionHook(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete composition hook: %s", err.Error())
	}

	d.SetId("")
	return nil
}
