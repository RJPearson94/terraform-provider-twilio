package studio

import (
	"context"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetState() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetStateRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"transitions": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event": {
							Type:     schema.TypeString,
							Required: true,
						},
						"next": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"conditions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"arguments": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"friendly_name": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
									"value": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"properties": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func dataSourceStudioFlowWidgetStateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	transitions := []flow.Transition{}
	for _, match := range d.Get("transitions").([]interface{}) {
		transitionMap := match.(map[string]interface{})

		transition := flow.Transition{
			Event: transitionMap["event"].(string),
		}

		if transitionMap["next"] != nil && transitionMap["next"].(string) != "" {
			transition.Next = sdkUtils.String(transitionMap["next"].(string))
		}

		if transitionMap["conditions"] != nil && len(transitionMap["conditions"].([]interface{})) > 0 {
			conditions := []flow.Condition{}

			for _, condition := range transitionMap["conditions"].([]interface{}) {
				conditionsMap := condition.(map[string]interface{})

				conditions = append(conditions, flow.Condition{
					Arguments:    utils.ConvertToStringSlice(conditionsMap["arguments"].([]interface{})),
					FriendlyName: conditionsMap["friendly_name"].(string),
					Type:         conditionsMap["type"].(string),
					Value:        conditionsMap["value"].(string),
				})
			}

			transition.Conditions = &conditions
		}

		transitions = append(transitions, transition)
	}

	state := flow.State{
		Name:        name,
		Properties:  d.Get("properties").(map[string]interface{}),
		Transitions: transitions,
		Type:        d.Get("type").(string),
	}

	if err := state.Validate(); err != nil {
		return diag.Errorf("State failed validation: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal state to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
