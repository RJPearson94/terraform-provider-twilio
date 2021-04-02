package voice

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/voice/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVoiceQueue() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVoiceQueueRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: helper.QueueSidValidation(),
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"average_wait_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"current_size": {
				Type:     schema.TypeInt,
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

func dataSourceVoiceQueueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Account(accountSid).Queue(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Queue with sid (%s) was not found in account (%s)", sid, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to read queue: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("max_size", getResponse.MaxSize)
	d.Set("average_wait_time", getResponse.AverageWaitTime)
	d.Set("current_size", getResponse.CurrentSize)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}
