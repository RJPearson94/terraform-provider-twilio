package serverless

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/services"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceServerlessService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerlessServiceCreate,
		ReadContext:   resourceServerlessServiceRead,
		UpdateContext: resourceServerlessServiceUpdate,
		DeleteContext: resourceServerlessServiceDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
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
				Computed: true,
			},
			"unique_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 50),
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"include_credentials": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ui_editable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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

func resourceServerlessServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &services.CreateServiceInput{
		UniqueName:         d.Get("unique_name").(string),
		FriendlyName:       d.Get("friendly_name").(string),
		IncludeCredentials: utils.OptionalBool(d, "include_credentials"),
		UiEditable:         utils.OptionalBool(d, "ui_editable"),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create serverless service: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceServerlessServiceRead(ctx, d, meta)
}

func resourceServerlessServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read serverless service: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("include_credentials", getResponse.IncludeCredentials)
	d.Set("ui_editable", getResponse.UiEditable)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	updateInput := &service.UpdateServiceInput{
		FriendlyName:       utils.OptionalString(d, "friendly_name"),
		IncludeCredentials: utils.OptionalBool(d, "include_credentials"),
		UiEditable:         utils.OptionalBool(d, "ui_editable"),
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update serverless service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceServerlessServiceRead(ctx, d, meta)
}

func resourceServerlessServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete serverless service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
