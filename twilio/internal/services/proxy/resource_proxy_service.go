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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceProxyService() *schema.Resource {
	return &schema.Resource{
		Create: resourceProxyServiceCreate,
		Read:   resourceProxyServiceRead,
		Update: resourceProxyServiceUpdate,
		Delete: resourceProxyServiceDelete,

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
			"chat_instance_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"chat_service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"callback_url": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"geo_match_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"number_selection_behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intercept_callback_url": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"out_of_session_callback_url": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
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

func resourceProxyServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &services.CreateServiceInput{
		UniqueName:              d.Get("unique_name").(string),
		DefaultTtl:              utils.OptionalInt(d, "default_ttl"),
		CallbackURL:             utils.OptionalString(d, "callback_url"),
		GeoMatchLevel:           utils.OptionalString(d, "geo_match_level"),
		NumberSelectionBehavior: utils.OptionalString(d, "number_selection_behavior"),
		InterceptCallbackURL:    utils.OptionalString(d, "intercept_callback_url"),
		OutOfSessionCallbackURL: utils.OptionalString(d, "out_of_session_callback_url"),
		ChatInstanceSid:         utils.OptionalString(d, "chat_instance_sid"),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create proxy service: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceProxyServiceRead(d, meta)
}

func resourceProxyServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read proxy service: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("chat_instance_sid", getResponse.ChatInstanceSid)
	d.Set("chat_service_sid", getResponse.ChatServiceSid)
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

func resourceProxyServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &service.UpdateServiceInput{
		UniqueName:              utils.OptionalString(d, "unique_name"),
		DefaultTtl:              utils.OptionalInt(d, "default_ttl"),
		CallbackURL:             utils.OptionalString(d, "callback_url"),
		GeoMatchLevel:           utils.OptionalString(d, "geo_match_level"),
		NumberSelectionBehavior: utils.OptionalString(d, "number_selection_behavior"),
		InterceptCallbackURL:    utils.OptionalString(d, "intercept_callback_url"),
		OutOfSessionCallbackURL: utils.OptionalString(d, "out_of_session_callback_url"),
		ChatInstanceSid:         utils.OptionalString(d, "chat_instance_sid"),
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update proxy service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyServiceRead(d, meta)
}

func resourceProxyServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete proxy service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
