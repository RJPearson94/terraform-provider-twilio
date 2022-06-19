package verify

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configurations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVerifyMessagingConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVerifyMessagingConfigurationCreate,
		ReadContext:   resourceVerifyMessagingConfigurationRead,
		UpdateContext: resourceVerifyMessagingConfigurationUpdate,
		DeleteContext: resourceVerifyMessagingConfigurationDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/MessagingConfigurations/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("country_code", match[2])
				d.SetId(match[1] + "/" + match[2])
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
			},
			"messaging_service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
			},
			"country_code": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
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

func resourceVerifyMessagingConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	createInput := &messaging_configurations.CreateMessagingConfigurationInput{
		MessagingServiceSid: d.Get("messaging_service_sid").(string),
		Country:             d.Get("country_code").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).MessagingConfigurations.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create messaging configuration: %s", err.Error())
	}

	d.SetId(createResult.ServiceSid + "/" + createResult.Country)
	return resourceVerifyMessagingConfigurationRead(ctx, d, meta)
}

func resourceVerifyMessagingConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	getResponse, err := client.Service(d.Get("service_sid").(string)).MessagingConfiguration(d.Get("country_code").(string)).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read messaging configuration: %s", err.Error())
	}

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

func resourceVerifyMessagingConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	updateInput := &messaging_configuration.UpdateMessagingConfigurationInput{
		MessagingServiceSid: d.Get("messaging_service_sid").(string),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).MessagingConfiguration(d.Get("country_code").(string)).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update messaging configuration: %s", err.Error())
	}

	d.SetId(updateResp.ServiceSid + "/" + updateResp.Country)
	return resourceVerifyMessagingConfigurationRead(ctx, d, meta)
}

func resourceVerifyMessagingConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	if err := client.Service(d.Get("service_sid").(string)).MessagingConfiguration(d.Get("country_code").(string)).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete messaging configuration: %s", err.Error())
	}

	d.SetId("")
	return nil
}
