package account

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/address"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/addresses"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAccountAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountAddressCreate,
		Read:   resourceAccountAddressRead,
		Update: resourceAccountAddressUpdate,
		Delete: resourceAccountAddressDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/Addresses/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("account_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"customer_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"street": {
				Type:     schema.TypeString,
				Required: true,
			},
			"street_secondary": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"postal_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iso_country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"emergency_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
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
	}
}

func resourceAccountAddressCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &addresses.CreateAddressInput{
		City:             d.Get("city").(string),
		CustomerName:     d.Get("customer_name").(string),
		EmergencyEnabled: utils.OptionalBool(d, "emergency_enabled"),
		FriendlyName:     utils.OptionalString(d, "friendly_name"),
		IsoCountry:       d.Get("iso_country").(string),
		PostalCode:       d.Get("postal_code").(string),
		Region:           d.Get("region").(string),
		Street:           d.Get("street").(string),
		StreetSecondary:  utils.OptionalString(d, "street_secondary"),
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).Addresses.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create address: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAccountAddressRead(d, meta)
}

func resourceAccountAddressRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Account(d.Get("account_sid").(string)).Address(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return fmt.Errorf("[ERROR] Failed to read address: %s", err.Error())
	}

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

func resourceAccountAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &address.UpdateAddressInput{
		City:             utils.OptionalString(d, "city"),
		CustomerName:     utils.OptionalString(d, "customer_name"),
		EmergencyEnabled: utils.OptionalBool(d, "emergency_enabled"),
		FriendlyName:     utils.OptionalString(d, "friendly_name"),
		PostalCode:       utils.OptionalString(d, "postal_code"),
		Region:           utils.OptionalString(d, "region"),
		Street:           utils.OptionalString(d, "street"),
		StreetSecondary:  utils.OptionalString(d, "street_secondary"),
	}

	updateResp, err := client.Account(d.Get("account_sid").(string)).Address(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update address: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAccountAddressRead(d, meta)
}

func resourceAccountAddressDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Account(d.Get("account_sid").(string)).Address(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete address: %s", err.Error())
	}

	d.SetId("")
	return nil
}
