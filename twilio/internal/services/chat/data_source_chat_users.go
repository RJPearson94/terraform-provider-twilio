package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceChatUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceChatUsersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
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
			"users": &schema.Schema{
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
						"attributes": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_notifiable": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_online": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"joined_channels_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"role_sid": &schema.Schema{
							Type:     schema.TypeString,
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

func dataSourceChatUsersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Users.NewUsersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return fmt.Errorf("[ERROR] No users were found for chat service with sid (%s)", serviceSid)
			}
		}
		return fmt.Errorf("[ERROR] Failed to list chat users: %s", err)
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
