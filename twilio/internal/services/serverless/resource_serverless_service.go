package serverless

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/services"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
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
				Default:  false,
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

func resourceServerlessServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &services.CreateServiceInput{
		UniqueName:         d.Get("unique_name").(string),
		FriendlyName:       d.Get("friendly_name").(string),
		IncludeCredentials: sdkUtils.Bool(d.Get("include_credentials").(bool)),
		UiEditable:         sdkUtils.Bool(d.Get("ui_editable").(bool)),
	}

	createResult, err := client.Services.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless service: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceServerlessServiceRead(d, meta)
}

func resourceServerlessServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless service: %s", err)
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

	updateInput := &service.UpdateServiceInput{
		FriendlyName:       d.Get("friendly_name").(string),
		IncludeCredentials: sdkUtils.Bool(d.Get("include_credentials").(bool)),
		UiEditable:         sdkUtils.Bool(d.Get("ui_editable").(bool)),
	}

	updateResp, err := client.Service(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update serverless service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceServerlessServiceRead(d, meta)
}

func resourceServerlessServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	if err := client.Service(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete serverless service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
