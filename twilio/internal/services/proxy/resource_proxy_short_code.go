package proxy

import (
	"fmt"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_code"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_codes"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceProxyShortCode() *schema.Resource {
	return &schema.Resource{
		Create: resourceProxyShortCodeCreate,
		Read:   resourceProxyShortCodeRead,
		Update: resourceProxyShortCodeUpdate,
		Delete: resourceProxyShortCodeDelete,
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
				Computed: true,
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
			"short_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_country": {
				Type:     schema.TypeString,
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

func resourceProxyShortCodeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	createInput := &short_codes.CreateShortCodeInput{
		Sid: d.Get("sid").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).ShortCodes.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create proxy short code resource: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if _, ok := d.GetOkExists("is_reserved"); ok {
		log.Println("[INFO] Is reserved can only be set on update, so updating the proxy short code resource to set the `is_reserved` flag")
		return resourceProxyShortCodeUpdate(d, meta)
	}
	return resourceProxyShortCodeRead(d, meta)
}

func resourceProxyShortCodeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	getResponse, err := client.Service(d.Get("service_sid").(string)).ShortCode(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read proxy short code resource: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("short_code", getResponse.ShortCode)
	d.Set("iso_country", getResponse.IsoCountry)
	d.Set("is_reserved", getResponse.IsReserved)
	d.Set("capabilities", flatternShortCodeCapabilities(getResponse.Capabilities))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceProxyShortCodeUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	updateInput := &short_code.UpdateShortCodeInput{
		IsReserved: utils.OptionalBool(d, "is_reserved"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).ShortCode(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update proxy short code resource: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyShortCodeRead(d, meta)
}

func resourceProxyShortCodeDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy

	if err := client.Service(d.Get("service_sid").(string)).ShortCode(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete proxy short code resource: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func flatternShortCodeCapabilities(capabilities *short_code.FetchShortCodeResponseCapabilities) *[]interface{} {
	if capabilities == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["fax_inbound"] = capabilities.FaxInbound
	result["fax_outbound"] = capabilities.FaxOutbound
	result["mms_inbound"] = capabilities.MmsInbound
	result["mms_outbound"] = capabilities.MmsOutbound
	result["restriction_fax_domestic"] = capabilities.RestrictionFaxDomestic
	result["restriction_mms_domestic"] = capabilities.RestrictionMmsDomestic
	result["restriction_sms_domestic"] = capabilities.RestrictionSmsDomestic
	result["restriction_voice_domestic"] = capabilities.RestrictionVoiceDomestic
	result["sip_trunking"] = capabilities.SipTrunking
	result["sms_inbound"] = capabilities.SmsInbound
	result["sms_outbound"] = capabilities.SmsOutbound
	result["voice_inbound"] = capabilities.VoiceInbound
	result["voice_outbound"] = capabilities.VoiceOutbound

	results = append(results, result)
	return &results
}
