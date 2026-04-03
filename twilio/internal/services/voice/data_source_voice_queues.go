package voice

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVoiceQueues() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVoiceQueuesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account to list queues for",
			},
			"queues": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of voice queues in the account",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID of the queue",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The human-readable label for the queue",
						},
						"max_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of calls that can be in the queue at one time",
						},
						"average_wait_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The average wait time of calls currently in the queue, in seconds",
						},
						"current_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The current number of calls in the queue",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the queue was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the queue was last updated, in RFC 3339 format",
						},
					},
				},
			},
		},
	}
}

func dataSourceVoiceQueuesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	paginator := client.Account(accountSid).Queues.NewQueuesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to list queues: %s", err.Error())
	}

	d.SetId(accountSid)
	d.Set("account_sid", accountSid)

	queues := make([]interface{}, 0)

	for _, queue := range paginator.Queues {
		queueMap := make(map[string]interface{})

		queueMap["sid"] = queue.Sid
		queueMap["friendly_name"] = queue.FriendlyName
		queueMap["max_size"] = queue.MaxSize
		queueMap["average_wait_time"] = queue.AverageWaitTime
		queueMap["current_size"] = queue.CurrentSize
		queueMap["date_created"] = queue.DateCreated.Format(time.RFC3339)

		if queue.DateUpdated != nil {
			queueMap["date_updated"] = queue.DateUpdated.Format(time.RFC3339)
		}

		queues = append(queues, queueMap)
	}

	d.Set("queues", &queues)

	return nil
}
