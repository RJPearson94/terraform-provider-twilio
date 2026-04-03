package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatChannels() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatChannelsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatServiceSidValidation(),
				Description:  "The SID of the Programmable Chat service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the channels",
			},
			"channels": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of channels",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to the channel by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the channel",
						},
						"unique_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique name for the channel",
						},
						"attributes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A JSON string of custom attributes for the channel",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The visibility of the channel. Values are `public` or `private`",
						},
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identity of the user that created the channel",
						},
						"members_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of members in the channel",
						},
						"messages_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of messages in the channel",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the channel was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the channel was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the channel resource",
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
