package studio

import (
	"context"
	"encoding/json"

	sdkStudio "github.com/RJPearson94/twilio-sdk-go/studio"
	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowDefinition() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowDefinitionRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"flags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_concurrent_calls": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"initial_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"states": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"json": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     validation.StringIsJSON,
							DiffSuppressFunc: structure.SuppressJsonDiff,
						},
					},
				},
			},
		},
	}
}

func dataSourceStudioFlowDefinitionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var flags *sdkStudio.FlowFlags
	if _, ok := d.GetOk("flags"); ok {
		flags = &sdkStudio.FlowFlags{
			AllowConcurrentCalls: d.Get("flags.0.allow_concurrent_calls").(bool),
		}
	}

	states := []flow.State{}
	for _, stateDefinition := range d.Get("states").([]interface{}) {
		stateDefinitionMap := stateDefinition.(map[string]interface{})

		state := flow.State{}
		if err := json.Unmarshal([]byte(stateDefinitionMap["json"].(string)), &state); err != nil {
			return diag.Errorf("Failed to unmarshal json to state struct %s", err.Error())
		}

		states = append(states, state)
	}

	flow := sdkStudio.Flow{
		Description:  d.Get("description").(string),
		Flags:        flags,
		InitialState: d.Get("initial_state").(string),
		States:       states,
	}

	if err := flow.Validate(); err != nil {
		return diag.Errorf("Flow defintion failed validation: %s", err.Error())
	}

	json, jsonErr := flow.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal flow defintion to JSON: %s", jsonErr.Error())
	}

	d.SetId(resource.UniqueId())
	d.Set("json", json)

	return nil
}
