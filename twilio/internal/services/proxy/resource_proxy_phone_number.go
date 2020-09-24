package proxy

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/proxy/helper"
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
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/PhoneNumbers/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
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
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"phone_number"},
			},
			"phone_number": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"sid"},
			},
			"is_reserved": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"capabilities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"fax_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_fax_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_mms_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_sms_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"restriction_voice_domestic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sip_trunking": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sms_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sms_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice_inbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice_outbound": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
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
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &phone_numbers.CreatePhoneNumberInput{
		Sid:         utils.OptionalString(d, "sid"),
		PhoneNumber: utils.OptionalString(d, "phone_number"),
		IsReserved:  utils.OptionalBool(d, "is_reserved"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).PhoneNumbers.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create proxy phone number resource: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceProxyPhoneNumberRead(d, meta)
}

func resourceProxyPhoneNumberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read proxy phone number resource: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("phone_number", getResponse.PhoneNumber)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("iso_country", getResponse.IsoCountry)
	d.Set("is_reserved", getResponse.IsReserved)
	d.Set("capabilities", helper.FlattenPhoneNumberCapabilities(getResponse.Capabilities))
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
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &phone_number.UpdatePhoneNumberInput{
		IsReserved: utils.OptionalBool(d, "is_reserved"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update proxy phone number resource: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyPhoneNumberRead(d, meta)
}

func resourceProxyPhoneNumberDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete proxy phone number resource: %s", err.Error())
	}
	d.SetId("")
	return nil
}
