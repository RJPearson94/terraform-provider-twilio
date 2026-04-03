package proxy

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/services"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceProxyService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProxyServiceCreate,
		ReadContext:   resourceProxyServiceRead,
		UpdateContext: resourceProxyServiceUpdate,
		DeleteContext: resourceProxyServiceDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)"
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this Proxy service by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Proxy service",
			},
			"chat_instance_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ChatServiceSidValidation(),
				Description:  "The SID of the Chat service instance to associate with the Proxy service",
			},
			"unique_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 191),
				Description:  "A unique, developer-assigned name for the Proxy service. Must be between 1 and 191 characters",
			},
			"default_ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The default time-to-live (TTL) for sessions in the Proxy service, in seconds. Defaults to `0`",
			},
			"callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL to receive callback events for the Proxy service",
			},
			"geo_match_level": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "country",
				ValidateFunc: validation.StringInSlice([]string{
					"area-code",
					"country",
					"extended-area-code",
				}, false),
				Description: "The geographic area for matching proxy numbers. Valid values are `area-code`, `country`, or `extended-area-code`. Defaults to `country`",
			},
			"number_selection_behavior": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "prefer-sticky",
				ValidateFunc: validation.StringInSlice([]string{
					"avoid-sticky",
					"prefer-sticky",
				}, false),
				Description: "The behavior for selecting proxy numbers. Valid values are `avoid-sticky` or `prefer-sticky`. Defaults to `prefer-sticky`",
			},
			"intercept_callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL to receive intercept callback events for the Proxy service",
			},
			"out_of_session_callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL to receive out-of-session callback events for the Proxy service",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy service was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy service was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the Proxy service resource",
			},
		},
	}
}

func resourceProxyServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	createInput := &services.CreateServiceInput{
		UniqueName:              d.Get("unique_name").(string),
		DefaultTtl:              utils.OptionalIntWith0OnChange(d, "default_ttl"),
		CallbackURL:             utils.OptionalStringWithEmptyStringOnChange(d, "callback_url"),
		GeoMatchLevel:           utils.OptionalString(d, "geo_match_level"),
		NumberSelectionBehavior: utils.OptionalString(d, "number_selection_behavior"),
		InterceptCallbackURL:    utils.OptionalStringWithEmptyStringOnChange(d, "intercept_callback_url"),
		OutOfSessionCallbackURL: utils.OptionalStringWithEmptyStringOnChange(d, "out_of_session_callback_url"),
		ChatInstanceSid:         utils.OptionalStringWithEmptyStringOnChange(d, "chat_instance_sid"),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create proxy service: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceProxyServiceRead(ctx, d, meta)
}

func resourceProxyServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read proxy service: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("chat_instance_sid", getResponse.ChatInstanceSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("default_ttl", getResponse.DefaultTtl)
	d.Set("callback_url", getResponse.CallbackURL)
	d.Set("geo_match_level", getResponse.GeoMatchLevel)
	d.Set("number_selection_behavior", getResponse.NumberSelectionBehavior)
	d.Set("intercept_callback_url", getResponse.InterceptCallbackURL)
	d.Set("out_of_session_callback_url", getResponse.OutOfSessionCallbackURL)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceProxyServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	updateInput := &service.UpdateServiceInput{
		UniqueName:              utils.OptionalString(d, "unique_name"),
		DefaultTtl:              utils.OptionalIntWith0OnChange(d, "default_ttl"),
		CallbackURL:             utils.OptionalStringWithEmptyStringOnChange(d, "callback_url"),
		GeoMatchLevel:           utils.OptionalString(d, "geo_match_level"),
		NumberSelectionBehavior: utils.OptionalString(d, "number_selection_behavior"),
		InterceptCallbackURL:    utils.OptionalStringWithEmptyStringOnChange(d, "intercept_callback_url"),
		OutOfSessionCallbackURL: utils.OptionalStringWithEmptyStringOnChange(d, "out_of_session_callback_url"),
		ChatInstanceSid:         utils.OptionalStringWithEmptyStringOnChange(d, "chat_instance_sid"),
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update proxy service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyServiceRead(ctx, d, meta)
}

func resourceProxyServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete proxy service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
