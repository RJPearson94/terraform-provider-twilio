package serverless

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variable"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variables"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessVariable() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessVariableCreate,
		Read:   resourceServerlessVariableRead,
		Update: resourceServerlessVariableUpdate,
		Delete: resourceServerlessVariableDelete,
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
			"environment_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
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

func resourceServerlessVariableCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &variables.CreateVariableInput{
		Key:   d.Get("key").(string),
		Value: d.Get("value").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variables.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless variable: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceServerlessVariableRead(d, meta)
}

func resourceServerlessVariableRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variable(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless variable: %s", err)
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

func resourceServerlessVariableUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	updateInput := &variable.UpdateVariableInput{
		Key:   utils.OptionalString(d, "key"),
		Value: utils.OptionalString(d, "value"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variable(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update serverless variable: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceServerlessVariableRead(d, meta)
}

func resourceServerlessVariableDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	if err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Variable(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete serverless variable: %s", err.Error())
	}
	d.SetId("")
	return nil
}
