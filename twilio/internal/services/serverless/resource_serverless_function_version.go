package serverless

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function/versions"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessFunctionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessFunctionVersionCreate,
		Read:   resourceServerlessFunctionVersionRead,
		Delete: resourceServerlessFunctionVersionDelete,
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
			"function_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content_body": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"date_created": {
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

func resourceServerlessFunctionVersionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &versions.CreateVersionInput{
		Content: versions.ContentDetails{
			Body:        d.Get("content_body").(string),
			ContentType: d.Get("content_type").(string),
			FileName:    d.Get("file_name").(string),
		},
		Path:       d.Get("path").(string),
		Visibility: d.Get("visibility").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Function(d.Get("function_sid").(string)).Versions.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless function: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceServerlessFunctionVersionRead(d, meta)
}

func resourceServerlessFunctionVersionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Get("service_sid").(string)).Function(d.Get("function_sid").(string)).Version(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless function version: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("function_sid", getResponse.FunctionSid)
	d.Set("content_body", d.Get("content_body").(string))
	d.Set("content_type", d.Get("content_type").(string))
	d.Set("file_name", d.Get("file_name").(string))
	d.Set("path", getResponse.Path)
	d.Set("visibility", getResponse.Visibility)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessFunctionVersionDelete(d *schema.ResourceData, meta interface{}) error {
	fmt.Printf("[INFO] Serverless function versions cannot be deleted. So the resource will remain until the function resource has been removed")

	d.SetId("")
	return nil
}
