package account

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAccountBalance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountBalanceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"balance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"currency": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAccountBalanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).API
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	sid := d.Get("account_sid").(string)
	getResponse, err := client.Account(sid).Balance().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Account balance with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read account balance: %s", err)
	}

	d.SetId(getResponse.AccountSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("balance", getResponse.Balance)
	d.Set("currency", getResponse.Currency)
	return nil
}
