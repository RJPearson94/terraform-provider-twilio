package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMessagingPhoneNumbers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMessagingPhoneNumbersRead,

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
			"phone_numbers": &schema.Schema{
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
						"phone_number": &schema.Schema{
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

func dataSourceMessagingPhoneNumbersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).PhoneNumbers.NewPhoneNumbersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No phone numbers were found for messaging service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to list messaging phone numbers: %s", err)
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	phoneNumbers := make([]interface{}, 0)

	for _, phoneNumber := range paginator.PhoneNumbers {
		d.Set("account_sid", phoneNumber.AccountSid)

		phoneNumberMap := make(map[string]interface{})

		phoneNumberMap["sid"] = phoneNumber.Sid
		phoneNumberMap["capabilities"] = phoneNumber.Capabilities
		phoneNumberMap["phone_number"] = phoneNumber.PhoneNumber
		phoneNumberMap["country_code"] = phoneNumber.CountryCode
		phoneNumberMap["date_created"] = phoneNumber.DateCreated.Format(time.RFC3339)

		if phoneNumber.DateUpdated != nil {
			phoneNumberMap["date_updated"] = phoneNumber.DateUpdated.Format(time.RFC3339)
		}

		phoneNumberMap["url"] = phoneNumber.URL

		phoneNumbers = append(phoneNumbers, phoneNumberMap)
	}

	d.Set("phone_numbers", &phoneNumbers)

	return nil
}
