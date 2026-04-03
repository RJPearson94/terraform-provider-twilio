package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/rate_limits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVerifyServiceRateLimits() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyServiceRateLimitsRead,

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
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these service rate limits",
			},
			"rate_limits": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of rate limits for the Verify service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this service rate limit by Twilio",
						},
						"unique_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique name of the rate limit",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
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
				},
			},
		},
	}
}

func dataSourceVerifyServiceRateLimitsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.Verify

	options := &rate_limits.RateLimitsPageOptions{}

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).RateLimits.NewRateLimitsPaginatorWithOptions(options)
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No service rate limits were found for Verify service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list Verify service rate limits: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("account_sid", twilioClient.AccountSid)
	d.Set("service_sid", serviceSid)

	rateLimits := make([]interface{}, 0)

	for _, rateLimit := range paginator.RateLimits {
		rateLimitMap := make(map[string]interface{})

		rateLimitMap["sid"] = rateLimit.Sid
		rateLimitMap["unique_name"] = rateLimit.UniqueName
		rateLimitMap["description"] = rateLimit.Description
		rateLimitMap["date_created"] = rateLimit.DateCreated.Format(time.RFC3339)

		if rateLimit.DateUpdated != nil {
			rateLimitMap["date_updated"] = rateLimit.DateUpdated.Format(time.RFC3339)
		}
		rateLimitMap["url"] = rateLimit.URL

		rateLimits = append(rateLimits, rateLimitMap)
	}

	d.Set("rate_limits", &rateLimits)

	return nil
}
