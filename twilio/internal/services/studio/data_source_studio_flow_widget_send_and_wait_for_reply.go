package studio

import (
	"context"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetSendAndWaitForReply() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetSendAndWaitForReplyRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"transitions": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delivery_failure": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"incoming_message": {
							Type:     schema.TypeString,
							Required: true,
						},
						"timeout": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"offset": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"x": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"y": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"from": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "{{flow.channel.address}}",
			},
			"body": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"service_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.ChatServiceSidValidation(),
				),
			},
			"channel_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.ChatChannelSidValidation(),
				),
			},
			"media_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
			},
			"timeout": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "3600",
			},
		},
	}
}

func dataSourceStudioFlowWidgetSendAndWaitForReplyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.SendAndWaitForReplyNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.DeliveryFailure = utils.OptionalString(d, "transitions.0.delivery_failure")
		nextTransitions.IncomingMessage = d.Get("transitions.0.incoming_message").(string)
		nextTransitions.Timeout = utils.OptionalString(d, "transitions.0.timeout")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := widgets.SendAndWaitForReply{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.SendAndWaitForReplyProperties{
			From:       d.Get("from").(string),
			Body:       d.Get("body").(string),
			Channel:    utils.OptionalString(d, "channel_sid"),
			Service:    utils.OptionalString(d, "service_sid"),
			Attributes: utils.OptionalJSONString(d, "attributes"),
			MediaURL:   utils.OptionalString(d, "media_url"),
			Offset:     offset,
			Timeout:    d.Get("timeout").(string),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Send and wait for reply widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create send and wait for reply widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal send and wait for reply widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
