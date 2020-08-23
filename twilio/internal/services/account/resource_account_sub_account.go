package account

import (
	"fmt"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/accounts"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAccountSubAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountSubAccountCreate,
		Read:   resourceAccountSubAccountRead,
		Update: resourceAccountSubAccountUpdate,
		Delete: resourceAccountSubAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceAccountSubAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API

	createInput := &accounts.CreateAccountInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	createResult, err := client.Accounts.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create sub account: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceAccountSubAccountRead(d, meta)
}

func resourceAccountSubAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read account: %s", err)
	}

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

func resourceAccountSubAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API

	updateInput := &account.UpdateAccountInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		Status:       utils.OptionalString(d, "status"),
	}

	updateResp, err := client.Account(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update account: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAccountSubAccountRead(d, meta)
}

func resourceAccountSubAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API

	log.Println("[INFO] Accounts can only be closed and will be deleted after 30 days. So updating the account to close it")

	updateInput := &account.UpdateAccountInput{
		Status: sdkUtils.String("closed"),
	}

	if _, err := client.Account(d.Id()).Update(updateInput); err != nil {
		return fmt.Errorf("Failed to close account: %s", err.Error())
	}

	d.SetId("")
	return nil
}
