package iam

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApiKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiKeyCreate,
		Read:   resourceApiKeyRead,
		Update: resourceApiKeyUpdate,
		Delete: resourceApiKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceApiKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Twilio
	context := context.Background()

	createKeyInput := url.Values{}
	createKeyInput.Add("FriendlyName", d.Get("friendly_name").(string))

	createResult, err := client.Keys.Create(context, createKeyInput)

	if err != nil {
		return fmt.Errorf("Failed to create key: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	d.Set("secret", createResult.Secret)

	return resourceApiKeyRead(d, meta)
}

func resourceApiKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Twilio
	context := context.Background()

	sid := d.Id()

	keyResponse, err := client.Keys.Get(context, sid)

	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("%s.json was not found", sid)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Failed to read key: %s", err.Error())
	}

	d.Set("sid", keyResponse.Sid)
	d.Set("friendly_name", keyResponse.FriendlyName)
	d.Set("date_created", keyResponse.DateCreated.Time.Format(time.RFC3339))
	d.Set("date_updated", keyResponse.DateUpdated.Time.Format(time.RFC3339))

	return nil
}

func resourceApiKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Twilio
	context := context.Background()

	updateKeyInput := url.Values{}

	if d.HasChange("friendly_name") {
		updateKeyInput.Set("FriendlyName", d.Get("friendly_name").(string))
	}

	updateKeyResult, err := client.Keys.Update(context, d.Id(), updateKeyInput)

	if err != nil {
		return fmt.Errorf("Failed to Update key: %s", err.Error())
	}

	d.SetId(updateKeyResult.Sid)

	return resourceApiKeyRead(d, meta)
}

func resourceApiKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Twilio
	context := context.Background()

	err := client.Keys.Delete(context, d.Id())

	if err != nil {
		return fmt.Errorf("Failed to delete key: %s", err.Error())
	}

	return nil
}
