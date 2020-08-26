package serverless

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/services"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessServiceCreate,
		Read:   resourceServerlessServiceRead,
		Update: resourceServerlessServiceUpdate,
		Delete: resourceServerlessServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include_credentials": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ui_editable": {
				Type:     schema.TypeBool,
				Optional: true,
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

func resourceServerlessServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &services.CreateServiceInput{
		UniqueName:         d.Get("unique_name").(string),
		FriendlyName:       d.Get("friendly_name").(string),
		IncludeCredentials: utils.OptionalBool(d, "include_credentials"),
		UiEditable:         utils.OptionalBool(d, "ui_editable"),
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless service: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceServerlessServiceRead(d, meta)
}

func resourceServerlessServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless service: %s", err.Error())
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

func resourceServerlessServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &service.UpdateServiceInput{
		FriendlyName:       utils.OptionalString(d, "friendly_name"),
		IncludeCredentials: utils.OptionalBool(d, "include_credentials"),
		UiEditable:         utils.OptionalBool(d, "ui_editable"),
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update serverless service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceServerlessServiceRead(d, meta)
}

func resourceServerlessServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete serverless service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
