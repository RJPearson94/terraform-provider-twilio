package studio

import (
	"context"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetTrigger() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetTriggerRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"transitions": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"incoming_call": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"incoming_message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"incoming_request": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"incoming_parent": {
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
		},
	}
}

func dataSourceStudioFlowWidgetTriggerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.TriggerNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.IncomingCall = utils.OptionalString(d, "transitions.0.incoming_call")
		nextTransitions.IncomingMessage = utils.OptionalString(d, "transitions.0.incoming_message")
		nextTransitions.IncomingRequest = utils.OptionalString(d, "transitions.0.incoming_request")
		nextTransitions.IncomingParent = utils.OptionalString(d, "transitions.0.incoming_parent")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := widgets.Trigger{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.TriggerProperties{
			Offset: offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Trigger widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create trigger widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal trigger widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
