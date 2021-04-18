package conversations

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation/webhooks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConversationsConversationStudioWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConversationsConversationStudioWebhookCreate,
		ReadContext:   resourceConversationsConversationStudioWebhookRead,
		UpdateContext: resourceConversationsConversationStudioWebhookUpdate,
		DeleteContext: resourceConversationsConversationStudioWebhookDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Conversations/(.*)/Webhooks/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("conversation_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
			},
			"conversation_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ConversationSidValidation(),
			},
			"target": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.StudioFlowSidValidation(),
			},
			"replay_after": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
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
	}
}

func resourceConversationsConversationStudioWebhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	createInput := &webhooks.CreateWebhookInput{
		Configuration: &webhooks.CreateWebhookConfigurationInput{
			FlowSid:     utils.OptionalString(d, "flow_sid"),
			ReplayAfter: utils.OptionalInt(d, "replay_after"),
		},
		Target: "studio",
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Conversation(d.Get("conversation_sid").(string)).Webhooks.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create conversation webhook: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceConversationsConversationStudioWebhookRead(ctx, d, meta)
}

func resourceConversationsConversationStudioWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Service(d.Get("service_sid").(string)).Conversation(d.Get("conversation_sid").(string)).Webhook(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read conversation webhook: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("conversation_sid", getResponse.ConversationSid)
	d.Set("target", getResponse.Target)
	d.Set("flow_sid", getResponse.Configuration.FlowSid)
	d.Set("replay_after", getResponse.Configuration.ReplayAfter)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceConversationsConversationStudioWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	updateInput := &webhook.UpdateWebhookInput{
		Configuration: &webhook.UpdateWebhookConfigurationInput{
			FlowSid: utils.OptionalString(d, "flow_sid"),
		},
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Conversation(d.Get("conversation_sid").(string)).Webhook(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update conversation webhook: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceConversationsConversationStudioWebhookRead(ctx, d, meta)
}

func resourceConversationsConversationStudioWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	if err := client.Service(d.Get("service_sid").(string)).Conversation(d.Get("conversation_sid").(string)).Webhook(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete conversation webhook: %s", err.Error())
	}
	d.SetId("")
	return nil
}
