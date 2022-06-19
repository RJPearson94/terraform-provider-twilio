package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVerifyServiceRateLimitBucket() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyServiceRateLimitBucketRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyRateLimitBucketSidValidation(),
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
			},
			"rate_limit_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyRateLimitSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interval": {
				Type:     schema.TypeInt,
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

func dataSourceVerifyServiceRateLimitBucketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	sid := d.Get("sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)
	serviceSid := d.Get("service_sid").(string)
	getResponse, err := client.Service(serviceSid).RateLimit(rateLimitSid).Bucket(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Service rate limit bucket with sid (%s) was not found for service with sid (%s) and service rate limit with sid (%s)", sid, serviceSid, rateLimitSid)
		}
		return diag.Errorf("Failed to read verify service rate limit bucket : %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("rate_limit_sid", getResponse.RateLimitSid)
	d.Set("max", getResponse.Max)
	d.Set("interval", getResponse.Interval)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
