package messaging

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMessagingPhoneNumbers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMessagingPhoneNumbersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
				Description:  "The SID of the messaging service to retrieve phone numbers for",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these phone numbers",
			},
			"phone_numbers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of phone numbers associated with the messaging service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this phone number by Twilio",
						},
						"capabilities": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of capabilities for the phone number",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"country_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The two-character ISO country code of the phone number",
						},
						"phone_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The phone number in E.164 format",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the phone number was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the phone number was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the phone number resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceMessagingPhoneNumbersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).PhoneNumbers.NewPhoneNumbersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No phone numbers were found for messaging service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list messaging phone numbers: %s", err.Error())
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
