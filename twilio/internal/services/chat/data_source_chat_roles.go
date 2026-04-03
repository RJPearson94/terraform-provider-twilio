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

func dataSourceChatRoles() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatRolesRead,

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
				Description: "The SID of the account that owns the roles",
			},
			"roles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of roles",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to the role by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the role",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of role. Values are `channel` or `deployment`",
						},
						"permissions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of permissions granted to the role",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the role was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the role was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the role resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceChatRolesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Roles.NewRolesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return diag.Errorf("No roles were found for chat service with sid (%s)", serviceSid)
			}
		}
		return diag.Errorf("Failed to read chat role: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	roles := make([]interface{}, 0)

	for _, role := range paginator.Roles {
		d.Set("account_sid", role.AccountSid)

		roleMap := make(map[string]interface{})

		roleMap["sid"] = role.Sid
		roleMap["friendly_name"] = role.FriendlyName
		roleMap["type"] = role.Type
		roleMap["permissions"] = role.Permissions
		roleMap["date_created"] = role.DateCreated.Format(time.RFC3339)

		if role.DateUpdated != nil {
			roleMap["date_updated"] = role.DateUpdated.Format(time.RFC3339)
		}

		roleMap["url"] = role.URL

		roles = append(roles, roleMap)
	}

	d.Set("roles", &roles)

	return nil
}
