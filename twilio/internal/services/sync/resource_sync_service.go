package sync

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/services"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSyncService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSyncServiceCreate,
		ReadContext:   resourceSyncServiceRead,
		UpdateContext: resourceSyncServiceUpdate,
		DeleteContext: resourceSyncServiceDelete,

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
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"acl_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reachability_debouncing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"reachability_debouncing_window": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1000, 30000),
				Default:      5000,
			},
			"reachability_webhooks_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"webhook_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"webhooks_from_rest_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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

func resourceSyncServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Sync

	createInput := &services.CreateServiceInput{
		AclEnabled:                    utils.OptionalBool(d, "acl_enabled"),
		FriendlyName:                  utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		ReachabilityDebouncingEnabled: utils.OptionalBool(d, "webhooks_from_rest_enabled"),
		ReachabilityDebouncingWindow:  utils.OptionalInt(d, "reachability_debouncing_window"),
		ReachabilityWebhooksEnabled:   utils.OptionalBool(d, "webhooks_from_rest_enabled"),
		WebhookURL:                    utils.OptionalStringWithEmptyStringOnChange(d, "webhook_url"),
		WebhooksFromRestEnabled:       utils.OptionalBool(d, "webhooks_from_rest_enabled"),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create service: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSyncServiceRead(ctx, d, meta)
}

func resourceSyncServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Sync

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read service: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("acl_enabled", getResponse.AclEnabled)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("reachability_debouncing_enabled", getResponse.ReachabilityDebouncingEnabled)
	d.Set("reachability_debouncing_window", getResponse.ReachabilityDebouncingWindow)
	d.Set("reachability_webhooks_enabled", getResponse.ReachabilityWebhooksEnabled)
	d.Set("webhook_url", getResponse.WebhookURL)
	d.Set("webhooks_from_rest_enabled", getResponse.WebhooksFromRestEnabled)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}

func resourceSyncServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Sync

	updateInput := &service.UpdateServiceInput{
		AclEnabled:                    utils.OptionalBool(d, "acl_enabled"),
		FriendlyName:                  utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		ReachabilityDebouncingEnabled: utils.OptionalBool(d, "webhooks_from_rest_enabled"),
		ReachabilityDebouncingWindow:  utils.OptionalInt(d, "reachability_debouncing_window"),
		ReachabilityWebhooksEnabled:   utils.OptionalBool(d, "webhooks_from_rest_enabled"),
		WebhookURL:                    utils.OptionalStringWithEmptyStringOnChange(d, "webhook_url"),
		WebhooksFromRestEnabled:       utils.OptionalBool(d, "webhooks_from_rest_enabled"),
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceSyncServiceRead(ctx, d, meta)
}

func resourceSyncServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Sync

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete service: %s", err.Error())
	}

	d.SetId("")
	return nil
}
