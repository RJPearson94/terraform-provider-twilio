package twiml

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/twiml/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/application"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/applications"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTwimlApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTwimlAppCreate,
		ReadContext:   resourceTwimlAppRead,
		UpdateContext: resourceTwimlAppUpdate,
		DeleteContext: resourceTwimlAppDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/Applications/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("account_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"messaging": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_callback_url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"fallback_url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
					},
				},
			},
			"voice": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"caller_id_lookup": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"fallback_url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"status_callback_url": {
							Type:         schema.TypeString,
							Optional:     true,
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
					},
				},
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTwimlAppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &applications.CreateApplicationInput{
		FriendlyName: utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
	}

	if _, ok := d.GetOk("messaging"); ok {
		createInput.MessageStatusCallback = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.status_callback_url")
		createInput.SmsFallbackMethod = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.fallback_method")
		createInput.SmsFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.fallback_url")
		createInput.SmsMethod = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.method")
		createInput.SmsURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.url")
	}

	if _, ok := d.GetOk("voice"); ok {
		createInput.VoiceCallerIDLookup = utils.OptionalBool(d, "voice.0.caller_id_lookup")
		createInput.VoiceFallbackMethod = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_method")
		createInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_url")
		createInput.VoiceMethod = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.method")
		createInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.url")
		createInput.StatusCallback = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.status_callback_url")
		createInput.StatusCallbackMethod = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.status_callback_method")
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).Applications.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create application: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTwimlAppRead(ctx, d, meta)
}

func resourceTwimlAppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).Application(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read application: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("messaging", helper.FlattenMessaging(getResponse))
	d.Set("voice", helper.FlattenVoice(getResponse))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	return nil
}

func resourceTwimlAppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	updateInput := &application.UpdateApplicationInput{
		FriendlyName: utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
	}

	if _, ok := d.GetOk("messaging"); ok {
		// This is necessary as there is currently an issue with the Applications API where setting MessageStatusCallback an empty string
		// is causing the value to be retained, the only way to clear down the value is to use SmsStatusCallback
		if v, ok := d.GetOk("messaging.0.status_callback_url"); ok {
			updateInput.MessageStatusCallback = sdkUtils.String(v.(string))
		} else if d.HasChange("messaging.0.status_callback_url") {
			updateInput.SmsStatusCallback = sdkUtils.String("")
		}
		updateInput.SmsFallbackMethod = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.fallback_method")
		updateInput.SmsFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.fallback_url")
		updateInput.SmsMethod = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.method")
		updateInput.SmsURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.url")
	}

	if _, ok := d.GetOk("voice"); ok {
		updateInput.VoiceCallerIDLookup = utils.OptionalBool(d, "voice.0.caller_id_lookup")
		updateInput.VoiceFallbackMethod = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_method")
		updateInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_url")
		updateInput.VoiceMethod = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.method")
		updateInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.url")
		updateInput.StatusCallback = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.status_callback_url")
		updateInput.StatusCallbackMethod = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.status_callback_method")
	}

	updateResp, err := client.Account(d.Get("account_sid").(string)).Application(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update application: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTwimlAppRead(ctx, d, meta)
}

func resourceTwimlAppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	if err := client.Account(d.Get("account_sid").(string)).Application(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete application: %s", err.Error())
	}
	d.SetId("")
	return nil
}
