package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsRoles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsRolesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"roles": {
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
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"permissions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

func dataSourceConversationsRolesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Roles.NewRolesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No roles were found for conversations service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list conversations roles: %s", err.Error())
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
