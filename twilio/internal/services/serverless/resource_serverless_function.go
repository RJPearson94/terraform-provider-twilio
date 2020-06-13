package serverless

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/functions"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessFunctionCreate,
		Read:   resourceServerlessFunctionRead,
		Update: resourceServerlessFunctionUpdate,
		Delete: resourceServerlessFunctionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceServerlessFunctionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &functions.CreateFunctionInput{
		FriendlyName: d.Get("friendly_name").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Functions.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless function: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceServerlessFunctionRead(d, meta)
}

func resourceServerlessFunctionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Get("service_sid").(string)).Function(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless function: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	updateInput := &function.UpdateFunctionInput{
		FriendlyName: d.Get("friendly_name").(string),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Function(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update serverless function: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceServerlessFunctionRead(d, meta)
}

func resourceServerlessFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	if err := client.Service(d.Get("service_sid").(string)).Function(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete serverless function: %s", err.Error())
	}
	d.SetId("")
	return nil
}
