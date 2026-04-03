package video

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/video/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVideoCompositionHook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVideoCompositionHookRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VideoCompositionHookSidValidation(),
				Description:  "The SID of the video composition hook to look up",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this video composition hook",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the video composition hook",
			},
			"audio_sources": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of audio source track names included in the composition",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"audio_sources_excluded": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of audio source track names excluded from the composition",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the composition hook is enabled",
			},
			"format": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The file format for the composition",
			},
			"resolution": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resolution of the composition in the format `WIDTHxHEIGHT`",
			},
			"status_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL called for composition status callback events",
			},
			"status_callback_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP method used to call the status callback URL",
			},
			"trim": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether intervals with no media are removed from the composition",
			},
			"video_layout": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string describing the video layout of the composition",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the video composition hook was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the video composition hook was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the video composition hook resource",
			},
		},
	}
}

func dataSourceVideoCompositionHookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Video

	sid := d.Get("sid").(string)
	getResponse, err := client.CompositionHook(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Composition hook with sid (%s) was not found ", sid)
		}
		return diag.Errorf("Failed to read composition hook: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
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
		return diag.Errorf("Unable to flatten video layout json to string. Error %s", err.Error())
	}
	d.Set("video_layout", videoLayout)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
