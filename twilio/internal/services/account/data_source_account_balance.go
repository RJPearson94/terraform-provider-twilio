package account

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccountBalance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountBalanceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account to retrieve the balance for",
			},
			"balance": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current balance of the account",
			},
			"currency": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The currency unit of the balance (e.g., USD)",
			},
		},
	}
}

func dataSourceAccountBalanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	sid := d.Get("account_sid").(string)
	getResponse, err := client.Account(sid).Balance().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Account balance with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read account balance: %s", err.Error())
	}

	d.SetId(getResponse.AccountSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("balance", getResponse.Balance)
	d.Set("currency", getResponse.Currency)
	return nil
}
