package conversations

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/conversations/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceConversationsConversation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConversationsConversationCreate,
		ReadContext:   resourceConversationsConversationRead,
		UpdateContext: resourceConversationsConversationUpdate,
		DeleteContext: resourceConversationsConversationDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Conversations/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("sid", match[2])
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
				Description: "The unique SID assigned to this conversation by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this conversation",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service. Changing this forces a new resource",
			},
			"unique_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A unique, developer-assigned name for the conversation",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 256),
				Description:  "A human-readable label for the conversation",
			},
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "{}",
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
				Description:      "A JSON string of attributes associated with the conversation. Defaults to `{}`",
			},
			"messaging_service_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
				Description:  "The SID of the messaging service associated with the conversation",
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "active",
				ValidateFunc: validation.StringInSlice([]string{
					"active",
					"inactive",
					"closed",
				}, false),
				Description: "The state of the conversation. Valid values are `active`, `inactive`, or `closed`. Defaults to `active`",
			},
			"timers": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Timer settings for the conversation",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"closed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ISO 8601 duration after which the conversation will be automatically closed",
						},
						"date_closed": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the conversation will be closed, in RFC 3339 format",
						},
						"inactive": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ISO 8601 duration after which the conversation will be marked as inactive",
						},
						"date_inactive": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the conversation will be marked as inactive, in RFC 3339 format",
						},
					},
				},
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the conversation was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the conversation was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the conversation resource",
			},
		},
	}
}

func resourceConversationsConversationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	createInput := &conversations.CreateConversationInput{
		Attributes:          utils.OptionalJSONString(d, "attributes"),
		UniqueName:          utils.OptionalStringWithEmptyStringOnChange(d, "unique_name"),
		FriendlyName:        utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		MessagingServiceSid: utils.OptionalString(d, "messaging_service_sid"),
		State:               utils.OptionalString(d, "state"),
	}

	if _, ok := d.GetOk("timers"); ok {
		createInput.Timers = &conversations.CreateConversationTimersInput{
			Closed:   utils.OptionalString(d, "timers.0.closed"),
			Inactive: utils.OptionalString(d, "timers.0.inactive"),
		}
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Conversations.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create conversations conversation: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceConversationsConversationRead(ctx, d, meta)
}

func resourceConversationsConversationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Service(d.Get("service_sid").(string)).Conversation(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read conversations conversation: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("attributes", getResponse.Attributes)
	d.Set("messaging_service_sid", getResponse.MessagingServiceSid)
	d.Set("state", getResponse.State)
	d.Set("timers", helper.FlattenTimers(d, getResponse.Timers))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceConversationsConversationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	updateInput := &conversation.UpdateConversationInput{
		Attributes:          utils.OptionalJSONString(d, "attributes"),
		UniqueName:          utils.OptionalStringWithEmptyStringOnChange(d, "unique_name"),
		FriendlyName:        utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		MessagingServiceSid: utils.OptionalString(d, "messaging_service_sid"),
		State:               utils.OptionalString(d, "state"),
	}

	if _, ok := d.GetOk("timers"); ok {
		updateInput.Timers = &conversation.UpdateConversationTimersInput{
			Closed:   utils.OptionalString(d, "timers.0.closed"),
			Inactive: utils.OptionalString(d, "timers.0.inactive"),
		}
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Conversation(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update conversations conversation: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceConversationsConversationRead(ctx, d, meta)
}

func resourceConversationsConversationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	if err := client.Service(d.Get("service_sid").(string)).Conversation(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete conversations conversation: %s", err.Error())
	}
	d.SetId("")
	return nil
}
