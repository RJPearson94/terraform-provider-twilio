package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limit/buckets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVerifyServiceRateLimitBuckets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyServiceRateLimitBucketsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
				Description:  "The SID of the Verify service",
			},
			"rate_limit_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyRateLimitSidValidation(),
				Description:  "The SID of the rate limit that the buckets belong to",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these rate limit buckets",
			},
			"rate_limit_buckets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of rate limit buckets for the service rate limit",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this rate limit bucket by Twilio",
						},
						"max": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of requests permitted in the given time interval",
						},
						"interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The time interval in seconds for the rate limit bucket",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the rate limit bucket was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the rate limit bucket was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the rate limit bucket resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceVerifyServiceRateLimitBucketsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.Verify

	options := &buckets.BucketsPageOptions{}

	rateLimitSid := d.Get("rate_limit_sid").(string)
	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).RateLimit(rateLimitSid).Buckets.NewBucketsPaginatorWithOptions(options)
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No service rate limit buckets were found for Verify service with sid (%s) and rate limit with sid (%s)", serviceSid, rateLimitSid)
		}
		return diag.Errorf("Failed to list Verify service rate limit buckets: %s", err.Error())
	}

	d.SetId(serviceSid + "/" + rateLimitSid)
	d.Set("account_sid", twilioClient.AccountSid)
	d.Set("service_sid", serviceSid)
	d.Set("rate_limit_sid", rateLimitSid)

	rateLimitBuckets := make([]interface{}, 0)

	for _, rateLimitBucket := range paginator.Buckets {
		rateLimitBucketMap := make(map[string]interface{})

		rateLimitBucketMap["sid"] = rateLimitBucket.Sid
		rateLimitBucketMap["max"] = rateLimitBucket.Max
		rateLimitBucketMap["interval"] = rateLimitBucket.Interval
		rateLimitBucketMap["date_created"] = rateLimitBucket.DateCreated.Format(time.RFC3339)

		if rateLimitBucket.DateUpdated != nil {
			rateLimitBucketMap["date_updated"] = rateLimitBucket.DateUpdated.Format(time.RFC3339)
		}
		rateLimitBucketMap["url"] = rateLimitBucket.URL

		rateLimitBuckets = append(rateLimitBuckets, rateLimitBucketMap)
	}

	d.Set("rate_limit_buckets", &rateLimitBuckets)

	return nil
}
