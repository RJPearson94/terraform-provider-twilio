package messaging

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/services"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceMessagingService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMessagingServiceCreate,
		ReadContext:   resourceMessagingServiceRead,
		UpdateContext: resourceMessagingServiceUpdate,
		DeleteContext: resourceMessagingServiceDelete,

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
				Description: "The unique SID assigned to this messaging service by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this messaging service",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A human-readable label for the messaging service",
			},
			"area_code_geomatch": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable area code geomatch on the messaging service. Defaults to `true`",
			},
			"fallback_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "POST",
				ValidateFunc: validation.StringInSlice([]string{
					"POST",
					"GET",
				}, false),
				Description: "The HTTP method used to call the fallback URL. Valid values are `POST` or `GET`. Defaults to `POST`",
			},
			"fallback_to_long_code": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable fallback to long code for the messaging service. Defaults to `true`",
			},
			"fallback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL to call when an inbound message error is received. Must be a valid HTTP or HTTPS URL",
			},
			"inbound_method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "POST",
				ValidateFunc: validation.StringInSlice([]string{
					"POST",
					"GET",
				}, false),
				Description: "The HTTP method used to call the inbound request URL. Valid values are `POST` or `GET`. Defaults to `POST`",
			},
			"inbound_request_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL to call when a message is received by any phone number or short code in the messaging service. Must be a valid HTTP or HTTPS URL",
			},
			"mms_converter": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable the MMS converter for the messaging service. Defaults to `true`",
			},
			"smart_encoding": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable smart encoding for the messaging service. Defaults to `true`",
			},
			"status_callback_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The URL to call when a message status change event is received. Must be a valid HTTP or HTTPS URL",
			},
			"sticky_sender": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable sticky sender on the messaging service. Defaults to `true`",
			},
			"use_inbound_webhook_on_number": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to use the inbound webhook on the phone number for incoming messages. Defaults to `false`",
			},
			"validity_period": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      36000,
				ValidateFunc: validation.IntBetween(1, 36000),
				Description:  "The number of seconds that the messaging service will keep a message in the sending queue before it is considered failed. Value must be between `1` and `36000`. Defaults to `36000`",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the messaging service was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the messaging service was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the messaging service resource",
			},
		},
	}
}

func resourceMessagingServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	createInput := &services.CreateServiceInput{
		FriendlyName:              d.Get("friendly_name").(string),
		AreaCodeGeomatch:          utils.OptionalBool(d, "area_code_geomatch"),
		FallbackMethod:            utils.OptionalString(d, "fallback_method"),
		FallbackToLongCode:        utils.OptionalBool(d, "fallback_to_long_code"),
		FallbackURL:               utils.OptionalStringWithEmptyStringOnChange(d, "fallback_url"),
		InboundMethod:             utils.OptionalString(d, "inbound_method"),
		InboundRequestURL:         utils.OptionalStringWithEmptyStringOnChange(d, "inbound_request_url"),
		MmsConverter:              utils.OptionalBool(d, "mms_converter"),
		SmartEncoding:             utils.OptionalBool(d, "smart_encoding"),
		StatusCallback:            utils.OptionalStringWithEmptyStringOnChange(d, "status_callback_url"),
		StickySender:              utils.OptionalBool(d, "sticky_sender"),
		UseInboundWebhookOnNumber: utils.OptionalBool(d, "use_inbound_webhook_on_number"),
		ValidityPeriod:            utils.OptionalInt(d, "validity_period"),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create messaging service: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceMessagingServiceRead(ctx, d, meta)
}

func resourceMessagingServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read messaging service: %s", err.Error())
	}

	d.Set("account_sid", getResponse.AccountSid)
	d.Set("area_code_geomatch", getResponse.AreaCodeGeomatch)
	d.Set("fallback_method", getResponse.FallbackMethod)
	d.Set("fallback_to_long_code", getResponse.FallbackToLongCode)
	d.Set("fallback_url", getResponse.FallbackURL)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("inbound_method", getResponse.InboundMethod)
	d.Set("inbound_request_url", getResponse.InboundRequestURL)
	d.Set("mms_converter", getResponse.MmsConverter)
	d.Set("sid", getResponse.Sid)
	d.Set("smart_encoding", getResponse.SmartEncoding)
	d.Set("status_callback_url", getResponse.StatusCallback)
	d.Set("sticky_sender", getResponse.StickySender)
	d.Set("use_inbound_webhook_on_number", getResponse.UseInboundWebhookOnNumber)
	d.Set("validity_period", getResponse.ValidityPeriod)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceMessagingServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	updateInput := &service.UpdateServiceInput{
		FriendlyName:              utils.OptionalString(d, "friendly_name"),
		AreaCodeGeomatch:          utils.OptionalBool(d, "area_code_geomatch"),
		FallbackMethod:            utils.OptionalString(d, "fallback_method"),
		FallbackToLongCode:        utils.OptionalBool(d, "fallback_to_long_code"),
		FallbackURL:               utils.OptionalStringWithEmptyStringOnChange(d, "fallback_url"),
		InboundMethod:             utils.OptionalString(d, "inbound_method"),
		InboundRequestURL:         utils.OptionalStringWithEmptyStringOnChange(d, "inbound_request_url"),
		MmsConverter:              utils.OptionalBool(d, "mms_converter"),
		SmartEncoding:             utils.OptionalBool(d, "smart_encoding"),
		StatusCallback:            utils.OptionalStringWithEmptyStringOnChange(d, "status_callback_url"),
		StickySender:              utils.OptionalBool(d, "sticky_sender"),
		UseInboundWebhookOnNumber: utils.OptionalBool(d, "use_inbound_webhook_on_number"),
		ValidityPeriod:            utils.OptionalInt(d, "validity_period"),
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update messaging service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceMessagingServiceRead(ctx, d, meta)
}

func resourceMessagingServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete messaging service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
