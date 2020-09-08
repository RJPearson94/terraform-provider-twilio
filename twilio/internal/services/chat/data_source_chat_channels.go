package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatChannels() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChatChannelsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
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
			"channels": {
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
						"unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"members_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"messages_count": {
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

func dataSourceChatChannelsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Channels.NewChannelsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return diag.Errorf("No channels were found for chat service with sid (%s)", serviceSid)
			}
		}
		return diag.Errorf("Failed to read chat channel: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	channels := make([]interface{}, 0)

	for _, channel := range paginator.Channels {
		d.Set("account_sid", channel.AccountSid)

		channelMap := make(map[string]interface{})

		channelMap["sid"] = channel.Sid
		channelMap["friendly_name"] = channel.FriendlyName
		channelMap["unique_name"] = channel.UniqueName
		channelMap["attributes"] = channel.Attributes
		channelMap["type"] = channel.Type
		channelMap["created_by"] = channel.CreatedBy
		channelMap["members_count"] = channel.MembersCount
		channelMap["messages_count"] = channel.MessagesCount
		channelMap["date_created"] = channel.DateCreated.Format(time.RFC3339)

		if channel.DateUpdated != nil {
			channelMap["date_updated"] = channel.DateUpdated.Format(time.RFC3339)
		}

		channelMap["url"] = channel.URL

		channels = append(channels, channelMap)
	}

	d.Set("channels", &channels)

	return nil
}
