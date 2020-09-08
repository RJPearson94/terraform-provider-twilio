package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceChatChannels() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceChatChannelsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"channels": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unique_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"members_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"messages_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"date_created": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"date_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceChatChannelsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Channels.NewChannelsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return fmt.Errorf("[ERROR] No channels were found for chat service with sid (%s)", serviceSid)
			}
		}
		return fmt.Errorf("[ERROR] Failed to read chat channel: %s", err)
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
