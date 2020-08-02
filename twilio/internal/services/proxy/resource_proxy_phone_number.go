package proxy

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceProxyPhoneNumber() *schema.Resource {
	return &schema.Resource{
		Create: resourceProxyPhoneNumberCreate,
		Read:   resourceProxyPhoneNumberRead,
		Update: resourceProxyPhoneNumberUpdate,
		Delete: resourceProxyPhoneNumberDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sid": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"is_reserved": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_country": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"in_use": {
				Type:     schema.TypeInt,
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
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceProxyPhoneNumberCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	createInput := &phone_numbers.CreatePhoneNumberInput{
		Sid:        utils.OptionalString(d, "sid"),
		IsReserved: utils.OptionalBool(d, "is_reserved"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).PhoneNumbers.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create proxy phone number: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceProxyPhoneNumberRead(d, meta)
}

func resourceProxyPhoneNumberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	getResponse, err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read proxy phone number: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("phone_number", getResponse.PhoneNumber)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("iso_country", getResponse.IsoCountry)
	d.Set("is_reserved", getResponse.IsReserved)
	// TODO set capabilities
	d.Set("in_use", getResponse.InUse)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceProxyPhoneNumberUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	updateInput := &phone_number.UpdatePhoneNumberInput{
		IsReserved: utils.OptionalBool(d, "is_reserved"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update proxy phone number: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyPhoneNumberRead(d, meta)
}

func resourceProxyPhoneNumberDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	if err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete proxy phone number: %s", err.Error())
	}
	d.SetId("")
	return nil
}
