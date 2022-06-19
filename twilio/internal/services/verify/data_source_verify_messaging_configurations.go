package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service/messaging_configurations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVerifyMessagingConfigurations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyMessagingConfigurationsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"messaging_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"messaging_service_sid": {
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

func dataSourceVerifyMessagingConfigurationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.Verify

	options := &messaging_configurations.MessagingConfigurationsPageOptions{}

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).MessagingConfigurations.NewMessagingConfigurationsPaginatorWithOptions(options)
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No messaging configurations were found for Verify service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list Verify messaging configurations: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("account_sid", twilioClient.AccountSid)
	d.Set("service_sid", serviceSid)

	messagingConfigurations := make([]interface{}, 0)

	for _, messagingConfiguration := range paginator.MessagingConfigurations {
		messagingConfigurationMap := make(map[string]interface{})
		messagingConfigurationMap["messaging_service_sid"] = messagingConfiguration.MessagingServiceSid
		messagingConfigurationMap["country_code"] = messagingConfiguration.Country
		messagingConfigurationMap["date_created"] = messagingConfiguration.DateCreated.Format(time.RFC3339)

		if messagingConfiguration.DateUpdated != nil {
			messagingConfigurationMap["date_updated"] = messagingConfiguration.DateUpdated.Format(time.RFC3339)
		}
		messagingConfigurationMap["url"] = messagingConfiguration.URL

		messagingConfigurations = append(messagingConfigurations, messagingConfigurationMap)
	}

	d.Set("messaging_configurations", &messagingConfigurations)

	return nil
}
