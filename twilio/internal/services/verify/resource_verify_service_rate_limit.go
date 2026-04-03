package verify

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVerifyServiceRateLimit() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVerifyServiceRateLimitCreate,
		ReadContext:   resourceVerifyServiceRateLimitRead,
		UpdateContext: resourceVerifyServiceRateLimitUpdate,
		DeleteContext: resourceVerifyServiceRateLimitDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/RateLimits/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this service rate limit by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this service rate limit",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
				Description:  "The SID of the Verify service. Changing this forces a new resource",
			},
			"unique_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A unique name for the rate limit. Changing this forces a new resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the rate limit",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the service rate limit was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the service rate limit was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the service rate limit resource",
			},
		},
	}
}

func resourceVerifyServiceRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	createInput := &rate_limits.CreateRateLimitInput{
		UniqueName:  d.Get("unique_name").(string),
		Description: utils.OptionalStringWithEmptyStringOnChange(d, "description"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).RateLimits.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create service rate limit: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceVerifyServiceRateLimitRead(ctx, d, meta)
}

func resourceVerifyServiceRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	getResponse, err := client.Service(d.Get("service_sid").(string)).RateLimit(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read service rate limit: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("description", getResponse.Description)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}

func resourceVerifyServiceRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	updateInput := &rate_limit.UpdateRateLimitInput{
		Description: utils.OptionalStringWithEmptyStringOnChange(d, "description"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).RateLimit(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update service rate limit: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceVerifyServiceRateLimitRead(ctx, d, meta)
}

func resourceVerifyServiceRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	if err := client.Service(d.Get("service_sid").(string)).RateLimit(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete service rate limit: %s", err.Error())
	}

	d.SetId("")
	return nil
}
