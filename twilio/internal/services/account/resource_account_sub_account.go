package account

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/accounts"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccountSubAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccountSubAccountCreate,
		ReadContext:   resourceAccountSubAccountRead,
		UpdateContext: resourceAccountSubAccountUpdate,
		DeleteContext: resourceAccountSubAccountDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
				return []*schema.ResourceData{d}, nil
			},
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

func resourceAccountSubAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &accounts.CreateAccountInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	createResult, err := client.Accounts.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create sub account: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAccountSubAccountRead(ctx, d, meta)
}

func resourceAccountSubAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read account: %s", err.Error())
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

func resourceAccountSubAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	updateInput := &account.UpdateAccountInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		Status:       utils.OptionalString(d, "status"),
	}

	updateResp, err := client.Account(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update account: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAccountSubAccountRead(ctx, d, meta)
}

func resourceAccountSubAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	log.Println("[INFO] Accounts can only be closed and will be deleted after 30 days. So updating the account to close it")

	updateInput := &account.UpdateAccountInput{
		Status: sdkUtils.String("closed"),
	}

	if _, err := client.Account(d.Id()).UpdateWithContext(ctx, updateInput); err != nil {
		return diag.Errorf("Failed to close account: %s", err.Error())
	}

	d.SetId("")
	return nil
}
