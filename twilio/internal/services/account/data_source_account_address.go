package account

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccountAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountAddressRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AddressSidValidation(),
				Description:  "The SID of the address to look up",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this address",
			},
			"customer_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the customer associated with the address",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the address",
			},
			"street": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The street address",
			},
			"street_secondary": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The secondary street address information (e.g., suite or apartment number)",
			},
			"city": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The city of the address",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state or region of the address",
			},
			"postal_code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The postal code of the address",
			},
			"iso_country": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ISO 3166-1 alpha-2 country code of the address",
			},
			"emergency_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether emergency calling is enabled for the address",
			},
			"validated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the address has been validated by Twilio",
			},
			"verified": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the address has been verified by the customer",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the address was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the address was last updated, in RFC 3339 format",
			},
		},
	}
}

func dataSourceAccountAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Account(accountSid).Address(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Address with sid (%s) was not found in account (%s)", sid, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to read address: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("customer_name", getResponse.CustomerName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("street", getResponse.Street)
	d.Set("street_secondary", getResponse.StreetSecondary)
	d.Set("city", getResponse.City)
	d.Set("region", getResponse.Region)
	d.Set("postal_code", getResponse.PostalCode)
	d.Set("iso_country", getResponse.IsoCountry)
	d.Set("emergency_enabled", getResponse.EmergencyEnabled)
	d.Set("validated", getResponse.Validated)
	d.Set("verified", getResponse.Verified)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
