package verify

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/bucket"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/buckets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVerifyServiceRateLimitBucket() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVerifyServiceRateLimitBucketCreate,
		ReadContext:   resourceVerifyServiceRateLimitBucketRead,
		UpdateContext: resourceVerifyServiceRateLimitBucketUpdate,
		DeleteContext: resourceVerifyServiceRateLimitBucketDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/RateLimits/(.*)/Buckets/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("rate_limit_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
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
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
			},
			"rate_limit_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.VerifyRateLimitSidValidation(),
			},
			"max": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"interval": {
				Type:     schema.TypeInt,
				Required: true,
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

func resourceVerifyServiceRateLimitBucketCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	createInput := &buckets.CreateBucketInput{
		Max:      d.Get("max").(int),
		Interval: d.Get("interval").(int),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).RateLimit(d.Get("rate_limit_sid").(string)).Buckets.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create service rate limit: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceVerifyServiceRateLimitBucketRead(ctx, d, meta)
}

func resourceVerifyServiceRateLimitBucketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	getResponse, err := client.Service(d.Get("service_sid").(string)).RateLimit(d.Get("rate_limit_sid").(string)).Bucket(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read service rate limit: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("rate_limit_sid", getResponse.RateLimitSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("max", getResponse.Max)
	d.Set("interval", getResponse.Interval)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}

func resourceVerifyServiceRateLimitBucketUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	updateInput := &bucket.UpdateBucketInput{
		Max:      utils.OptionalInt(d, "max"),
		Interval: utils.OptionalInt(d, "interval"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).RateLimit(d.Get("rate_limit_sid").(string)).Bucket(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update service rate limit bucket: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceVerifyServiceRateLimitBucketRead(ctx, d, meta)
}

func resourceVerifyServiceRateLimitBucketDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	if err := client.Service(d.Get("service_sid").(string)).RateLimit(d.Get("rate_limit_sid").(string)).Bucket(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete service rate limit bucket: %s", err.Error())
	}

	d.SetId("")
	return nil
}
