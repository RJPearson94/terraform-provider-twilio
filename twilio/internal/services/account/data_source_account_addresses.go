package account

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAccountAddresses() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountAddressesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"customer_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"street": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"street_secondary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"city": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"postal_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iso_country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"emergency_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"validated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"verified": {
							Type:     schema.TypeBool,
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
					},
				},
			},
		},
	}
}

func dataSourceAccountAddressesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	accountSid := d.Get("account_sid").(string)
	paginator := client.Account(accountSid).Addresses.NewAddressesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return fmt.Errorf("[ERROR] Failed to list addresses: %s", err.Error())
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
