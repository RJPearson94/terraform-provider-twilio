package account

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAccountDetails() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountDetailsRead,
		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner_account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
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
		},
	}
}

func dataSourceAccountDetailsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API

	sid := d.Get("sid").(string)
	getResponse, err := client.Account(sid).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Account with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read account details: %s", err)
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("owner_account_sid", getResponse.OwnerAccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("status", getResponse.Status)
	d.Set("type", getResponse.Type)
	d.Set("auth_token", getResponse.AuthToken)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
