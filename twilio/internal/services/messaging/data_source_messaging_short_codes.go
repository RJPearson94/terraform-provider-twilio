package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMessagingShortCodes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMessagingShortCodesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"short_codes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capabilities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"country_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"short_code": &schema.Schema{
							Type:     schema.TypeString,
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
				},
			},
		},
	}
}

func dataSourceMessagingShortCodesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).ShortCodes.NewShortCodesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No short codes were found for messaging service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read messaging short code: %s", err)
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
