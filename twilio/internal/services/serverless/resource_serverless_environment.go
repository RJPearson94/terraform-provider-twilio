package serverless

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environments"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessEnvironmentCreate,
		Read:   resourceServerlessEnvironmentRead,
		Delete: resourceServerlessEnvironmentDelete,
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
			"build_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_suffix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"domain_name": {
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

func resourceServerlessEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &environments.CreateEnvironmentInput{
		UniqueName:   d.Get("unique_name").(string),
		DomainSuffix: sdkUtils.String(d.Get("domain_suffix").(string)),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Environments.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless environment: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceServerlessEnvironmentRead(d, meta)
}

func resourceServerlessEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Get("service_sid").(string)).Environment(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless environment: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("build_sid", getResponse.BuildSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("domain_suffix", getResponse.DomainSuffix)
	d.Set("domain_name", getResponse.DomainName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	if err := client.Service(d.Get("service_sid").(string)).Environment(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete serverless service: %s", err.Error())
	}
	d.SetId("")
	return nil
}
