package studio

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceStudioFlow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStudioFlowRead,
		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"definition": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"commit_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"valid": {
				Type:     schema.TypeBool,
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
			"webhook_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceStudioFlowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Studio

	sid := d.Get("sid").(string)
	getResponse, err := client.Flow(sid).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Studio flow with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read studio flow: %s", err)
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("definition", getResponse.Definition)
	d.Set("status", getResponse.Status)
	d.Set("revision", getResponse.Revision)
	d.Set("commit_message", getResponse.CommitMessage)
	d.Set("valid", getResponse.Valid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)
	d.Set("webhook_url", getResponse.WebhookURL)

	return nil
}
