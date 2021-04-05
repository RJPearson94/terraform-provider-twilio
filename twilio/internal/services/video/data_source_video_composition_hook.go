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
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"audio_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"audio_sources_excluded": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resolution": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_callback_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trim": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"video_layout": {
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
