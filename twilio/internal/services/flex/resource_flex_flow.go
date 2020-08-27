package flex

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFlexFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceFlexFlowCreate,
		Read:   resourceFlexFlowRead,
		Update: resourceFlexFlowUpdate,
		Delete: resourceFlexFlowDelete,

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

func resourceFlexFlowCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

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
		createInput.IntegrationChannel = utils.OptionalString(d, "integration.0.channel")
		createInput.IntegrationCreationOnMessage = utils.OptionalBool(d, "integration.0.creation_on_message")
		createInput.IntegrationFlowSid = utils.OptionalString(d, "integration.0.flow_sid")
		createInput.IntegrationPriority = utils.OptionalInt(d, "integration.0.priority")
		createInput.IntegrationRetryCount = utils.OptionalInt(d, "integration.0.retry_count")
		createInput.IntegrationTimeout = utils.OptionalInt(d, "integration.0.timeout")
		createInput.IntegrationURL = utils.OptionalString(d, "integration.0.url")
		createInput.IntegrationWorkflowSid = utils.OptionalString(d, "integration.0.workflow_sid")
		createInput.IntegrationWorkspaceSid = utils.OptionalString(d, "integration.0.workspace_sid")
	}

	createResult, err := client.FlexFlows.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create flex flow: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceFlexFlowRead(d, meta)
}

func resourceFlexFlowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.FlexFlow(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read flex channel: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("channel_type", getResponse.ChannelType)
	d.Set("chat_service_sid", getResponse.ChatServiceSid)
	d.Set("contact_identity", getResponse.ContactIdentity)
	d.Set("enabled", getResponse.Enabled)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("integration", flatternIntegration(getResponse.Integration))
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

func resourceFlexFlowUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

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
		updateInput.IntegrationChannel = utils.OptionalString(d, "integration.0.channel")
		updateInput.IntegrationCreationOnMessage = utils.OptionalBool(d, "integration.0.creation_on_message")
		updateInput.IntegrationFlowSid = utils.OptionalString(d, "integration.0.flow_sid")
		updateInput.IntegrationPriority = utils.OptionalInt(d, "integration.0.priority")
		updateInput.IntegrationRetryCount = utils.OptionalInt(d, "integration.0.retry_count")
		updateInput.IntegrationTimeout = utils.OptionalInt(d, "integration.0.timeout")
		updateInput.IntegrationURL = utils.OptionalString(d, "integration.0.url")
		updateInput.IntegrationWorkflowSid = utils.OptionalString(d, "integration.0.workflowSid")
		updateInput.IntegrationWorkspaceSid = utils.OptionalString(d, "integration.0.workspace_sid")
	}

	createResult, err := client.FlexFlow(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to update flex flow: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceFlexFlowRead(d, meta)
}

func resourceFlexFlowDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.FlexFlow(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete flex flow: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func flatternIntegration(integration *flex_flow.FetchFlexFlowResponseIntegration) *[]interface{} {
	if integration == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["channel"] = integration.Channel
	result["creation_on_message"] = integration.CreationOnMessage
	result["flow_sid"] = integration.FlowSid
	result["priority"] = integration.Priority
	result["retry_count"] = integration.RetryCount
	result["timeout"] = integration.Timeout
	result["url"] = integration.URL
	result["workspace_sid"] = integration.WorkspaceSid

	results = append(results, result)
	return &results
}
