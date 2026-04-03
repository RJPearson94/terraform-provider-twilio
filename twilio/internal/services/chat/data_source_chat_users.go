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

func dataSourceChatUsers() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatUsersRead,

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
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the users",
			},
			"users": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of users",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to the user by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the user",
						},
						"attributes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A JSON string of custom attributes for the user",
						},
						"identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identity string of the user",
						},
						"is_notifiable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the user has a potentially valid push registration for the chat service",
						},
						"is_online": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the user is actively connected to the chat service",
						},
						"joined_channels_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of channels the user has joined",
						},
						"role_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the role assigned to the user",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the user was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the user was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the user resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceChatUsersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Users.NewUsersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return diag.Errorf("No users were found for chat service with sid (%s)", serviceSid)
			}
		}
		return diag.Errorf("Failed to list chat users: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	users := make([]interface{}, 0)

	for _, user := range paginator.Users {
		d.Set("account_sid", user.AccountSid)

		userMap := make(map[string]interface{})

		userMap["sid"] = user.Sid
		userMap["friendly_name"] = user.FriendlyName
		userMap["attributes"] = user.Attributes
		userMap["identity"] = user.Identity
		userMap["is_notifiable"] = user.IsNotifiable
		userMap["is_online"] = user.IsOnline
		userMap["joined_channels_count"] = user.JoinedChannelsCount
		userMap["role_sid"] = user.RoleSid
		userMap["date_created"] = user.DateCreated.Format(time.RFC3339)

		if user.DateUpdated != nil {
			userMap["date_updated"] = user.DateUpdated.Format(time.RFC3339)
		}

		userMap["url"] = user.URL

		users = append(users, userMap)
	}

	d.Set("users", &users)

	return nil
}
