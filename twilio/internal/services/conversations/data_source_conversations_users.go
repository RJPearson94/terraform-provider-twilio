package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsUsersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these users",
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
							Description: "The unique SID assigned to this user by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the user",
						},
						"attributes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A JSON string of attributes associated with the user",
						},
						"identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identity string for the user",
						},
						"is_notifiable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the user has a potentially valid push channel registration for notifications",
						},
						"is_online": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the user is actively connected to the service",
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

func dataSourceConversationsUsersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Users.NewUsersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No users were found for conversations service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list conversations users: %s", err.Error())
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
