package messaging

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMessagingAlphaSenders() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMessagingAlphaSendersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
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

func dataSourceMessagingAlphaSendersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).AlphaSenders.NewAlphaSendersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No alpha senders were found for messaging service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list messaging alpha senders: %s", err.Error())
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
