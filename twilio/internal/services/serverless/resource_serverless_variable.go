package serverless

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variable"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variables"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceServerlessVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerlessVariableCreate,
		ReadContext:   resourceServerlessVariableRead,
		UpdateContext: resourceServerlessVariableUpdate,
		DeleteContext: resourceServerlessVariableDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Environments/(.*)/Variables/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("environment_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
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
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
			},
			"environment_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ServerlessEnvironmentSidValidation(),
			},
			"key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"value": {
				Type:         schema.TypeString,
				Required:     true,
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

func resourceServerlessVariableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &variables.CreateVariableInput{
		Key:   d.Get("key").(string),
		Value: d.Get("value").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variables.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create serverless variable: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceServerlessVariableRead(ctx, d, meta)
}

func resourceServerlessVariableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variable(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read serverless variable: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("environment_sid", getResponse.EnvironmentSid)
	d.Set("key", getResponse.Key)
	d.Set("value", getResponse.Value)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessVariableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	updateInput := &variable.UpdateVariableInput{
		Key:   utils.OptionalString(d, "key"),
		Value: utils.OptionalString(d, "value"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variable(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update serverless variable: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceServerlessVariableRead(ctx, d, meta)
}

func resourceServerlessVariableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	if err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variable(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete serverless variable: %s", err.Error())
	}
	d.SetId("")
	return nil
}
