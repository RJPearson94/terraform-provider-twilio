package messaging

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMessagingShortCodes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMessagingShortCodesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
				Description:  "The SID of the messaging service to retrieve short codes for",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these short codes",
			},
			"short_codes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of short codes associated with the messaging service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this short code by Twilio",
						},
						"capabilities": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of capabilities for the short code",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"country_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The two-character ISO country code of the short code",
						},
						"short_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The short code value",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the short code was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the short code was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the short code resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceMessagingShortCodesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).ShortCodes.NewShortCodesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No short codes were found for messaging service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to read messaging short code: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	shortCodes := make([]interface{}, 0)

	for _, shortCode := range paginator.ShortCodes {
		d.Set("account_sid", shortCode.AccountSid)

		shortCodeMap := make(map[string]interface{})

		shortCodeMap["sid"] = shortCode.Sid
		shortCodeMap["capabilities"] = shortCode.Capabilities
		shortCodeMap["country_code"] = shortCode.CountryCode
		shortCodeMap["short_code"] = shortCode.ShortCode
		shortCodeMap["date_created"] = shortCode.DateCreated.Format(time.RFC3339)

		if shortCode.DateUpdated != nil {
			shortCodeMap["date_updated"] = shortCode.DateUpdated.Format(time.RFC3339)
		}

		shortCodeMap["url"] = shortCode.URL

		shortCodes = append(shortCodes, shortCodeMap)
	}

	d.Set("short_codes", &shortCodes)

	return nil
}
