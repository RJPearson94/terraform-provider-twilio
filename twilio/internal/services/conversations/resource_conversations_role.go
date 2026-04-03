package conversations

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/role"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/roles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceConversationsRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConversationsRoleCreate,
		ReadContext:   resourceConversationsRoleRead,
		UpdateContext: resourceConversationsRoleUpdate,
		DeleteContext: resourceConversationsRoleDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Roles/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[2])
				d.Set("service_sid", match[1])
				d.SetId(match[2])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this role by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this role",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service. Changing this forces a new resource",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 256),
				Description:  "A human-readable label for the role. Changing this forces a new resource",
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"conversation",
					"service",
				}, false),
				Description: "The type of role. Valid values are `conversation` or `service`. Changing this forces a new resource",
			},
			"permissions": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The list of permissions granted to the role",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"editOwnMessage",
						"deleteAnyMessage",
						"addParticipant",
						"editConversationAttributes",
						"editAnyParticipantAttributes",
						"editAnyMessage",
						"editConversationName",
						"editAnyMessageAttributes",
						"deleteOwnMessage",
						"editOwnMessageAttributes",
						"removeParticipant",
						"addNonChatParticipant",
						"editOwnParticipantAttributes",
						"deleteConversation",

						// Conversation
						"editNotificationLevel",
						"sendMessage",
						"leaveConversation",
						"sendMediaMessage",

						// Service
						"editAnyUserInfo",
						"removeParticipant",
						"createConversation",
						"editOwnUserInfo",
						"joinConversation",
					}, false),
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
	}
}

func resourceConversationsRoleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	createInput := &roles.CreateRoleInput{
		FriendlyName: d.Get("friendly_name").(string),
		Type:         d.Get("type").(string),
		Permissions:  utils.ConvertToStringSlice(d.Get("permissions").([]interface{})),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Roles.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create conversations role: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceConversationsRoleRead(ctx, d, meta)
}

func resourceConversationsRoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Service(d.Get("service_sid").(string)).Role(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read conversations role: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("type", getResponse.Type)
	d.Set("permissions", getResponse.Permissions)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceConversationsRoleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	updateInput := &role.UpdateRoleInput{
		Permissions: utils.ConvertToStringSlice(d.Get("permissions").([]interface{})),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Role(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update conversations role: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceConversationsRoleRead(ctx, d, meta)
}

func resourceConversationsRoleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	if err := client.Service(d.Get("service_sid").(string)).Role(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete conversations role: %s", err.Error())
	}
	d.SetId("")
	return nil
}
