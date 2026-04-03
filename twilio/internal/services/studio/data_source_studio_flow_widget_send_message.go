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

func dataSourceStudioFlowWidgetSendMessage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetSendMessageRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source",
			},
			"transitions": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The next widget(s) to transition to after this widget",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"failed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the message fails to send",
						},
						"sent": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the message is sent successfully",
						},
					},
				},
			},
			"offset": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The position of this widget in the Studio visual editor",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"x": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The x-axis position. Defaults to 0",
						},
						"y": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The y-axis position. Defaults to 0",
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The unique name of this widget within the flow, used to reference it in transitions",
			},
			"to": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "{{contact.channel.address}}",
				Description: "The recipient address for the message. Defaults to `{{contact.channel.address}}`",
			},
			"from": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "{{flow.channel.address}}",
				Description: "The sender address for the message. Defaults to `{{flow.channel.address}}`",
			},
			"body": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The text body of the message to send. Supports Liquid template expressions",
			},
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
				Description:      "A JSON string of custom attributes to attach to the message",
			},
			"service_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.ChatServiceSidValidation(),
				),
				Description: "The SID of the Programmable Chat service to use for sending the message",
			},
			"channel_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.ChatChannelSidValidation(),
				),
				Description: "The SID of the Programmable Chat channel to send the message to",
			},
			"media_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				Description: "The HTTP/HTTPS URL of a media file to include with the message (e.g. an image or PDF)",
			},
		},
	}
}

func dataSourceStudioFlowWidgetSendMessageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.SendMessageNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Failed = utils.OptionalString(d, "transitions.0.failed")
		nextTransitions.Sent = utils.OptionalString(d, "transitions.0.sent")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := widgets.SendMessage{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.SendMessageProperties{
			To:         d.Get("to").(string),
			From:       d.Get("from").(string),
			Body:       d.Get("body").(string),
			Channel:    utils.OptionalString(d, "channel_sid"),
			Service:    utils.OptionalString(d, "service_sid"),
			Attributes: utils.OptionalJSONString(d, "attributes"),
			MediaURL:   utils.OptionalString(d, "media_url"),
			Offset:     offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Send message widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create send message widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal send message to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
