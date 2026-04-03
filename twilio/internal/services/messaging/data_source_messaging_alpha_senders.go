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
				Description:  "The SID of the messaging service to retrieve alpha senders for",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these alpha senders",
			},
			"alpha_senders": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of alpha senders associated with the messaging service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this alpha sender by Twilio",
						},
						"alpha_sender": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The alphanumeric sender ID string",
						},
						"capabilities": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of capabilities for the alpha sender",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the alpha sender was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the alpha sender was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the alpha sender resource",
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
