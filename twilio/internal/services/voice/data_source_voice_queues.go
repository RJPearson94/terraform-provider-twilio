package voice

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
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
				Type:     schema.TypeString,
				Required: true,
			},
			"queues": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
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
		queueMap["max_size"] = queue.FriendlyName
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
