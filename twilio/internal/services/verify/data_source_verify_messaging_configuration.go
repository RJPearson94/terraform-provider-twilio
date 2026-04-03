package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceVerifyMessagingConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyMessagingConfigurationRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"country_code": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The ISO-3166-1 country code of the messaging configuration to fetch",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
				Description:  "The SID of the Verify service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this messaging configuration",
			},
			"messaging_service_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the messaging service associated with this configuration",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the messaging configuration was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the messaging configuration was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the messaging configuration resource",
			},
		},
	}
}

func dataSourceVerifyMessagingConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	countryCode := d.Get("country_code").(string)
	serviceSid := d.Get("service_sid").(string)
	getResponse, err := client.Service(serviceSid).MessagingConfiguration(countryCode).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Messaging configuration with country code (%s) was not found for Verify service with sid (%s)", countryCode, serviceSid)
		}
		return diag.Errorf("Failed to read Verify messaging configuration: %s", err.Error())
	}

	d.SetId(getResponse.ServiceSid + "/" + getResponse.Country)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("messaging_service_sid", getResponse.MessagingServiceSid)
	d.Set("country_code", getResponse.Country)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
