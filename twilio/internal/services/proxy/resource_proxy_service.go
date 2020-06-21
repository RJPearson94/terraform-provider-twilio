package proxy

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/services"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceProxyService() *schema.Resource {
	return &schema.Resource{
		Create: resourceProxyServiceCreate,
		Read:   resourceProxyServiceRead,
		Update: resourceProxyServiceUpdate,
		Delete: resourceProxyServiceDelete,
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
			"chat_instance_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"chat_service_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_match_level": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"number_selection_behavior": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intercept_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"out_of_session_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceProxyServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	createInput := &services.CreateServiceInput{
		UniqueName:              d.Get("unique_name").(string),
		DefaultTtl:              utils.OptionalInt(d, "default_ttl"),
		CallbackUrl:             utils.OptionalString(d, "callback_url"),
		GeoMatchLevel:           utils.OptionalString(d, "geo_match_level"),
		NumberSelectionBehavior: utils.OptionalString(d, "number_selection_behavior"),
		InterceptCallbackUrl:    utils.OptionalString(d, "intercept_callback_url"),
		OutOfSessionCallbackUrl: utils.OptionalString(d, "out_of_session_callback_url"),
		ChatInstanceSid:         utils.OptionalString(d, "chat_instance_sid"),
	}

	createResult, err := client.Services.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create proxy service: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceProxyServiceRead(d, meta)
}

func resourceProxyServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	getResponse, err := client.Service(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read proxy service: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("chat_instance_sid", getResponse.ChatInstanceSid)
	d.Set("chat_service_sid", getResponse.ChatServiceSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("default_ttl", getResponse.DefaultTtl)
	d.Set("callback_url", getResponse.CallbackUrl)
	d.Set("geo_match_level", getResponse.GeoMatchLevel)
	d.Set("number_selection_behavior", getResponse.NumberSelectionBehavior)
	d.Set("intercept_callback_url", getResponse.InterceptCallbackUrl)
	d.Set("out_of_session_callback_url", getResponse.OutOfSessionCallbackUrl)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceProxyServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	updateInput := &service.UpdateServiceInput{
		UniqueName:              utils.OptionalString(d, "unique_name"),
		DefaultTtl:              utils.OptionalInt(d, "default_ttl"),
		CallbackUrl:             utils.OptionalString(d, "callback_url"),
		GeoMatchLevel:           utils.OptionalString(d, "geo_match_level"),
		NumberSelectionBehavior: utils.OptionalString(d, "number_selection_behavior"),
		InterceptCallbackUrl:    utils.OptionalString(d, "intercept_callback_url"),
		OutOfSessionCallbackUrl: utils.OptionalString(d, "out_of_session_callback_url"),
		ChatInstanceSid:         utils.OptionalString(d, "chat_instance_sid"),
	}

	updateResp, err := client.Service(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update proxy service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyServiceRead(d, meta)
}

func resourceProxyServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	if err := client.Service(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete proxy service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
