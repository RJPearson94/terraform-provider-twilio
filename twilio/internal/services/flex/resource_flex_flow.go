package flex

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/flex/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFlexFlow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlexFlowCreate,
		ReadContext:   resourceFlexFlowRead,
		UpdateContext: resourceFlexFlowUpdate,
		DeleteContext: resourceFlexFlowDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/FlexFlows/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
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
			"channel_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"web",
					"sms",
					"facebook",
					"whatsapp",
					"line",
					"custom",
				}, false),
			},
			"chat_service_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"contact_identity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"integration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"creation_on_message": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"flow_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"retry_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"workflow_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"workspace_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"integration_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"studio",
					"external",
					"task",
				}, false),
			},
			"janitor_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"long_lived": {
				Type:     schema.TypeBool,
				Optional: true,
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

func resourceFlexFlowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	createInput := &flex_flows.CreateFlexFlowInput{
		ChannelType:     d.Get("channel_type").(string),
		ChatServiceSid:  d.Get("chat_service_sid").(string),
		ContactIdentity: utils.OptionalString(d, "contact_identity"),
		Enabled:         utils.OptionalBool(d, "enabled"),
		FriendlyName:    d.Get("friendly_name").(string),
		IntegrationType: utils.OptionalString(d, "integration_type"),
		JanitorEnabled:  utils.OptionalBool(d, "janitor_enabled"),
		LongLived:       utils.OptionalBool(d, "long_lived"),
	}

	if _, ok := d.GetOk("integration"); ok {
		createInput.Integration = &flex_flows.CreateFlexFlowIntegrationInput{
			Channel:           utils.OptionalString(d, "integration.0.channel"),
			CreationOnMessage: utils.OptionalBool(d, "integration.0.creation_on_message"),
			FlowSid:           utils.OptionalString(d, "integration.0.flow_sid"),
			Priority:          utils.OptionalInt(d, "integration.0.priority"),
			RetryCount:        utils.OptionalInt(d, "integration.0.retry_count"),
			Timeout:           utils.OptionalInt(d, "integration.0.timeout"),
			URL:               utils.OptionalString(d, "integration.0.url"),
			WorkflowSid:       utils.OptionalString(d, "integration.0.workflow_sid"),
			WorkspaceSid:      utils.OptionalString(d, "integration.0.workspace_sid"),
		}
	}

	createResult, err := client.FlexFlows.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create flex flow: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceFlexFlowRead(ctx, d, meta)
}

func resourceFlexFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	getResponse, err := client.FlexFlow(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read flex channel: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("channel_type", getResponse.ChannelType)
	d.Set("chat_service_sid", getResponse.ChatServiceSid)
	d.Set("contact_identity", getResponse.ContactIdentity)
	d.Set("enabled", getResponse.Enabled)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("integration", helper.FlattenIntegration(getResponse.Integration))
	d.Set("integration_type", getResponse.IntegrationType)
	d.Set("janitor_enabled", getResponse.JanitorEnabled)
	d.Set("long_lived", getResponse.LongLived)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceFlexFlowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	updateInput := &flex_flow.UpdateFlexFlowInput{
		ChannelType:     utils.OptionalString(d, "channel_type"),
		ChatServiceSid:  utils.OptionalString(d, "chat_service_sid"),
		ContactIdentity: utils.OptionalString(d, "contact_identity"),
		Enabled:         utils.OptionalBool(d, "enabled"),
		FriendlyName:    utils.OptionalString(d, "friendly_name"),
		IntegrationType: utils.OptionalString(d, "integration_type"),
		JanitorEnabled:  utils.OptionalBool(d, "janitor_enabled"),
		LongLived:       utils.OptionalBool(d, "long_lived"),
	}

	if _, ok := d.GetOk("integration"); ok {
		updateInput.Integration = &flex_flow.UpdateFlexFlowIntegrationInput{
			Channel:           utils.OptionalString(d, "integration.0.channel"),
			CreationOnMessage: utils.OptionalBool(d, "integration.0.creation_on_message"),
			FlowSid:           utils.OptionalString(d, "integration.0.flow_sid"),
			Priority:          utils.OptionalInt(d, "integration.0.priority"),
			RetryCount:        utils.OptionalInt(d, "integration.0.retry_count"),
			Timeout:           utils.OptionalInt(d, "integration.0.timeout"),
			URL:               utils.OptionalString(d, "integration.0.url"),
			WorkflowSid:       utils.OptionalString(d, "integration.0.workflow_sid"),
			WorkspaceSid:      utils.OptionalString(d, "integration.0.workspace_sid"),
		}
	}

	createResult, err := client.FlexFlow(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update flex flow: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceFlexFlowRead(ctx, d, meta)
}

func resourceFlexFlowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	if err := client.FlexFlow(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete flex flow: %s", err.Error())
	}
	d.SetId("")
	return nil
}
