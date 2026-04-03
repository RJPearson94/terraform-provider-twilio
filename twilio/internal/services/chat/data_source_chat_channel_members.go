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

func dataSourceChatChannelMembers() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatChannelMembersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatServiceSidValidation(),
				Description:  "The SID of the Programmable Chat service",
			},
			"channel_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatChannelSidValidation(),
				Description:  "The SID of the chat channel",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the channel members",
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of channel members",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to the channel member by Twilio",
						},
						"attributes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A JSON string of custom attributes for the channel member",
						},
						"identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identity string of the member",
						},
						"role_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the role assigned to the channel member",
						},
						"last_consumed_message_index": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The index of the last message the member has read in the channel",
						},
						"last_consumption_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the member last consumed a message, in RFC 3339 format",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the channel member was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the channel member was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the channel member resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceChatChannelMembersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	paginator := client.Service(serviceSid).Channel(channelSid).Members.NewChannelMembersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return diag.Errorf("No channel members were found for chat service with sid (%s) and channel with sid (%s)", serviceSid, channelSid)
			}
		}
		return diag.Errorf("Failed to list chat channel members: %s", err.Error())
	}

	d.SetId(serviceSid + "/" + channelSid)
	d.Set("service_sid", serviceSid)
	d.Set("channel_sid", channelSid)

	members := make([]interface{}, 0)

	for _, member := range paginator.Members {
		d.Set("account_sid", member.AccountSid)

		memberMap := make(map[string]interface{})

		memberMap["sid"] = member.Sid
		memberMap["attributes"] = member.Attributes
		memberMap["identity"] = member.Identity
		memberMap["role_sid"] = member.RoleSid
		memberMap["last_consumed_message_index"] = member.LastConsumedMessageIndex
		if member.LastConsumedTimestamp != nil {
			memberMap["last_consumption_timestamp"] = member.LastConsumedTimestamp.Format(time.RFC3339)
		}
		memberMap["date_created"] = member.DateCreated.Format(time.RFC3339)

		if member.DateUpdated != nil {
			memberMap["date_updated"] = member.DateUpdated.Format(time.RFC3339)
		}

		memberMap["url"] = member.URL

		members = append(members, memberMap)
	}

	d.Set("members", &members)

	return nil
}
