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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProxyPhoneNumber() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProxyPhoneNumberCreate,
		ReadContext:   resourceProxyPhoneNumberRead,
		UpdateContext: resourceProxyPhoneNumberUpdate,
		DeleteContext: resourceProxyPhoneNumberDelete,

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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Proxy phone number",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ProxyServiceSidValidation(),
				Description:  "The SID of the Proxy service. Changing this forces a new resource",
			},
			"sid": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"phone_number"},
				ValidateFunc:  utils.PhoneNumberSidValidation(),
				Description:   "The SID of the Twilio phone number to add to the Proxy service. Conflicts with `phone_number`. Changing this forces a new resource",
			},
			"phone_number": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"sid"},
				ValidateFunc:  utils.PhoneNumberValidation(),
				Description:   "The phone number in E.164 format to add to the Proxy service. Conflicts with `sid`. Changing this forces a new resource",
			},
			"is_reserved": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the phone number is reserved and not assigned to a Proxy session",
			},
			"capabilities": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The capabilities of the Proxy phone number",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive faxes",
						},
						"fax_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can send faxes",
						},
						"mms_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive MMS messages",
						},
						"mms_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can send MMS messages",
						},
						"restriction_fax_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether fax is restricted to domestic only",
						},
						"restriction_mms_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether MMS is restricted to domestic only",
						},
						"restriction_sms_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether SMS is restricted to domestic only",
						},
						"restriction_voice_domestic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether voice is restricted to domestic only",
						},
						"sip_trunking": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number supports SIP trunking",
						},
						"sms_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive SMS messages",
						},
						"sms_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can send SMS messages",
						},
						"voice_inbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can receive voice calls",
						},
						"voice_outbound": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number can make voice calls",
						},
					},
				},
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the Proxy phone number",
			},
			"iso_country": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ISO country code of the phone number",
			},
			"in_use": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of active Proxy sessions assigned to this phone number",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy phone number was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy phone number was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the Proxy phone number resource",
			},
		},
	}
}

func resourceProxyPhoneNumberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	createInput := &phone_numbers.CreatePhoneNumberInput{
		Sid:         utils.OptionalString(d, "sid"),
		PhoneNumber: utils.OptionalString(d, "phone_number"),
		IsReserved:  utils.OptionalBool(d, "is_reserved"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).PhoneNumbers.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create proxy phone number resource: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceProxyPhoneNumberRead(ctx, d, meta)
}

func resourceProxyPhoneNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	getResponse, err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read proxy phone number resource: %s", err.Error())
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

func resourceProxyPhoneNumberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	updateInput := &phone_number.UpdatePhoneNumberInput{
		IsReserved: utils.OptionalBool(d, "is_reserved"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update proxy phone number resource: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyPhoneNumberRead(ctx, d, meta)
}

func resourceProxyPhoneNumberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	if err := client.Service(d.Get("service_sid").(string)).PhoneNumber(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete proxy phone number resource: %s", err.Error())
	}
	d.SetId("")
	return nil
}
