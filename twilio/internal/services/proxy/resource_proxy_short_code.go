package proxy

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/proxy/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_code"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_codes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProxyShortCode() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProxyShortCodeCreate,
		ReadContext:   resourceProxyShortCodeRead,
		UpdateContext: resourceProxyShortCodeUpdate,
		DeleteContext: resourceProxyShortCodeDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/ShortCodes/(.*)"
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

func resourceProxyShortCodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	createInput := &short_codes.CreateShortCodeInput{
		Sid: d.Get("sid").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).ShortCodes.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create proxy short code resource: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if _, ok := d.GetOkExists("is_reserved"); ok {
		log.Println("[INFO] Is reserved can only be set on update, so updating the proxy short code resource to set the `is_reserved` flag")
		return resourceProxyShortCodeUpdate(ctx, d, meta)
	}
	return resourceProxyShortCodeRead(ctx, d, meta)
}

func resourceProxyShortCodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	getResponse, err := client.Service(d.Get("service_sid").(string)).ShortCode(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read proxy short code resource: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("short_code", getResponse.ShortCode)
	d.Set("iso_country", getResponse.IsoCountry)
	d.Set("is_reserved", getResponse.IsReserved)
	d.Set("capabilities", helper.FlattenShortCodeCapabilities(getResponse.Capabilities))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceProxyShortCodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	updateInput := &short_code.UpdateShortCodeInput{
		IsReserved: utils.OptionalBool(d, "is_reserved"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).ShortCode(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update proxy short code resource: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceProxyShortCodeRead(ctx, d, meta)
}

func resourceProxyShortCodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	if err := client.Service(d.Get("service_sid").(string)).ShortCode(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete proxy short code resource: %s", err.Error())
	}
	d.SetId("")
	return nil
}
