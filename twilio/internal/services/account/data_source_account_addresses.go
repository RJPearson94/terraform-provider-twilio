package account

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccountAddresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountAddressesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account to list addresses for",
			},
			"addresses": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of addresses associated with the account",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this address by Twilio",
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
				},
			},
		},
	}
}

func dataSourceAccountAddressesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	paginator := client.Account(accountSid).Addresses.NewAddressesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to list addresses: %s", err.Error())
	}

	d.SetId(accountSid)
	d.Set("account_sid", accountSid)

	addresses := make([]interface{}, 0)

	for _, address := range paginator.Addresses {
		addressMap := make(map[string]interface{})

		addressMap["sid"] = address.Sid
		addressMap["customer_name"] = address.CustomerName
		addressMap["friendly_name"] = address.FriendlyName
		addressMap["street"] = address.Street
		addressMap["street_secondary"] = address.StreetSecondary
		addressMap["city"] = address.City
		addressMap["region"] = address.Region
		addressMap["postal_code"] = address.PostalCode
		addressMap["iso_country"] = address.IsoCountry
		addressMap["emergency_enabled"] = address.EmergencyEnabled
		addressMap["validated"] = address.Validated
		addressMap["verified"] = address.Verified
		addressMap["date_created"] = address.DateCreated.Format(time.RFC3339)

		if address.DateUpdated != nil {
			addressMap["date_updated"] = address.DateUpdated.Format(time.RFC3339)
		}

		addresses = append(addresses, addressMap)
	}

	d.Set("addresses", &addresses)

	return nil
}
