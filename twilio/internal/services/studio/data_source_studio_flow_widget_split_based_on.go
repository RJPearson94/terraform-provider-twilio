package studio

import (
	"context"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/transition"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetSplitBasedOn() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetSplitBasedOnRead,

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
						"matches": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"next": {
										Type:     schema.TypeString,
										Required: true,
									},
									"conditions": {
										Type:     schema.TypeList,
										Required: true,
										MinItems: 1,
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
													Type:     schema.TypeString,
													Required: true,
													ValidateFunc: validation.StringInSlice([]string{
														"equal_to",
														"not_equal_to",
														"matches_any_of",
														"does_not_match_any_of",
														"is_blank",
														"is_not_blank",
														"regex",
														"contains",
														"does_not_contain",
														"starts_with",
														"does_not_start_with",
														"less_than",
														"greater_than",
														"is_before_time",
														"is_after_time",
														"is_before_date",
														"is_after_date",
													}, false),
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
						"no_match": {
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
			"input": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func dataSourceStudioFlowWidgetSplitBasedOnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.SplitBasedOnNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Matches = getMatches(d)
		nextTransitions.NoMatch = utils.OptionalString(d, "transitions.0.no_match")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := widgets.SplitBasedOn{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.SplitBasedOnProperties{
			Input:  d.Get("input").(string),
			Offset: offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Split based on widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create split based on widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal split based on widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}

func getMatches(d *schema.ResourceData) *[]transition.Conditional {
	if v, ok := d.GetOk("transitions.0.matches"); ok {
		matches := []transition.Conditional{}

		for _, match := range v.([]interface{}) {
			matchesMap := match.(map[string]interface{})

			conditions := []flow.Condition{}
			for _, condition := range matchesMap["conditions"].([]interface{}) {
				conditionsMap := condition.(map[string]interface{})

				conditions = append(conditions, flow.Condition{
					Arguments:    utils.ConvertToStringSlice(conditionsMap["arguments"].([]interface{})),
					FriendlyName: conditionsMap["friendly_name"].(string),
					Type:         conditionsMap["type"].(string),
					Value:        conditionsMap["value"].(string),
				})
			}

			matches = append(matches, transition.Conditional{
				Next:       matchesMap["next"].(string),
				Conditions: &conditions,
			})
		}
		return &matches
	}
	return nil
}
