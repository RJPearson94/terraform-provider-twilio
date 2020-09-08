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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceMessagingService() *schema.Resource {
	return &schema.Resource{
		Create: resourceMessagingServiceCreate,
		Read:   resourceMessagingServiceRead,
		Update: resourceMessagingServiceUpdate,
		Delete: resourceMessagingServiceDelete,

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
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"area_code_geomatch": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"fallback_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"POST",
					"GET",
				}, false),
			},
			"fallback_to_long_code": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"fallback_url": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"inbound_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"POST",
					"GET",
				}, false),
			},
			"inbound_request_url": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"mms_converter": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"smart_encoding": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"status_callback_url": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"sticky_sender": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"validity_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

func resourceMessagingServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &services.CreateServiceInput{
		FriendlyName:       d.Get("friendly_name").(string),
		AreaCodeGeomatch:   utils.OptionalBool(d, "area_code_geomatch"),
		FallbackMethod:     utils.OptionalString(d, "fallback_method"),
		FallbackToLongCode: utils.OptionalBool(d, "fallback_to_long_code"),
		FallbackURL:        utils.OptionalString(d, "fallback_url"),
		InboundMethod:      utils.OptionalString(d, "inbound_method"),
		InboundRequestURL:  utils.OptionalString(d, "inbound_request_url"),
		MmsConverter:       utils.OptionalBool(d, "mms_converter"),
		SmartEncoding:      utils.OptionalBool(d, "smart_encoding"),
		StatusCallback:     utils.OptionalString(d, "status_callback_url"),
		StickySender:       utils.OptionalBool(d, "sticky_sender"),
		ValidityPeriod:     utils.OptionalInt(d, "validity_period"),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create messaging service: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceMessagingServiceRead(d, meta)
}

func resourceMessagingServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read messaging service: %s", err)
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
	d.Set("validity_period", getResponse.ValidityPeriod)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceMessagingServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &service.UpdateServiceInput{
		FriendlyName:       utils.OptionalString(d, "friendly_name"),
		AreaCodeGeomatch:   utils.OptionalBool(d, "area_code_geomatch"),
		FallbackMethod:     utils.OptionalString(d, "fallback_method"),
		FallbackToLongCode: utils.OptionalBool(d, "fallback_to_long_code"),
		FallbackURL:        utils.OptionalString(d, "fallback_url"),
		InboundMethod:      utils.OptionalString(d, "inbound_method"),
		InboundRequestURL:  utils.OptionalString(d, "inbound_request_url"),
		MmsConverter:       utils.OptionalBool(d, "mms_converter"),
		SmartEncoding:      utils.OptionalBool(d, "smart_encoding"),
		StatusCallback:     utils.OptionalString(d, "status_callback_url"),
		StickySender:       utils.OptionalBool(d, "sticky_sender"),
		ValidityPeriod:     utils.OptionalInt(d, "validity_period"),
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update messaging service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceMessagingServiceRead(d, meta)
}

func resourceMessagingServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete messaging service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
