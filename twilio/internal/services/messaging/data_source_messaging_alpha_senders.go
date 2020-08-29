package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMessagingAlphaSenders() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMessagingAlphaSendersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alpha_senders": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alpha_sender": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"capabilities": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
					},
				},
			},
		},
	}
}

func dataSourceMessagingAlphaSendersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Messaging
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).AlphaSenders.NewAlphaSendersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No alpha senders were found for messaging service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to list messaging alpha senders: %s", err)
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	alphaSenders := make([]interface{}, 0)

	for _, alphaSender := range paginator.AlphaSenders {
		d.Set("account_sid", alphaSender.AccountSid)

		alphaSenderMap := make(map[string]interface{})

		alphaSenderMap["sid"] = alphaSender.Sid
		alphaSenderMap["capabilities"] = alphaSender.Capabilities
		alphaSenderMap["alpha_sender"] = alphaSender.AlphaSender
		alphaSenderMap["date_created"] = alphaSender.DateCreated.Format(time.RFC3339)

		if alphaSender.DateUpdated != nil {
			alphaSenderMap["date_updated"] = alphaSender.DateUpdated.Format(time.RFC3339)
		}

		alphaSenderMap["url"] = alphaSender.URL

		alphaSenders = append(alphaSenders, alphaSenderMap)
	}

	d.Set("alpha_senders", &alphaSenders)

	return nil
}
