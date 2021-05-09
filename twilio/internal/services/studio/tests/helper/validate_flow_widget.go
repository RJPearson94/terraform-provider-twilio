package helper

import (
	"encoding/json"
	"fmt"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow_validation"
	sdkStudio "github.com/RJPearson94/twilio-sdk-go/studio"
	"github.com/RJPearson94/twilio-sdk-go/studio/flow"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func ValidateFlowWidget(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		state := flow.State{}
		if err := json.Unmarshal([]byte(rs.Primary.Attributes["json"]), &state); err != nil {
			return fmt.Errorf("Failed to unmarshal json to state struct %s", err.Error())
		}

		flow := sdkStudio.Flow{
			Description:  "widget testing",
			InitialState: "Trigger",
			States: []flow.State{
				{
					Name: "Trigger",
					Type: "trigger",
					Transitions: []flow.Transition{
						{
							Event: "incomingCall",
						},
						{
							Event: "incomingMessage",
							Next:  &state.Name,
						},
						{
							Event: "incomingRequest",
						},
					},
					Properties: map[string]interface{}{},
				},
				state,
			},
		}

		if err := flow.Validate(); err != nil {
			return fmt.Errorf("Flow defintion failed validation: %s", err.Error())
		}

		json, jsonErr := flow.ToString()
		if jsonErr != nil {
			return fmt.Errorf("Failed to marshal flow defintion to JSON: %s", jsonErr.Error())
		}

		return validateFlow(*json)
	}
}

func ValidateFlowWidgetTrigger(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		state := flow.State{}
		if err := json.Unmarshal([]byte(rs.Primary.Attributes["json"]), &state); err != nil {
			return fmt.Errorf("Failed to unmarshal json to state struct %s", err.Error())
		}

		flow := sdkStudio.Flow{
			Description:  "widget testing",
			InitialState: "Trigger",
			States: []flow.State{
				state,
				{
					Name: "Next",
					Type: "run-function",
					Transitions: []flow.Transition{
						{
							Event: "fail",
						},
						{
							Event: "success",
						},
					},
					Properties: map[string]interface{}{
						"url": "https://test-function.twil.io/test-function",
					},
				},
			},
		}

		if err := flow.Validate(); err != nil {
			return fmt.Errorf("Flow defintion failed validation: %s", err.Error())
		}

		json, jsonErr := flow.ToString()
		if jsonErr != nil {
			return fmt.Errorf("Failed to marshal flow defintion to JSON: %s", jsonErr.Error())
		}

		return validateFlow(*json)
	}
}

func ValidateFlowDefinition(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		return validateFlow(rs.Primary.Attributes["json"])
	}
}

func validateFlow(flowJSON string) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Studio

	resp, err := client.FlowValidation.Validate(&flow_validation.ValidateFlowInput{
		FriendlyName: "Widget Validation",
		Status:       "draft",
		Definition:   flowJSON,
	})

	if err != nil {
		return fmt.Errorf("Error occurred when validating the flow definition %s", err.Error())
	}

	if !resp.Valid {
		return fmt.Errorf("Flow definition is invalid")
	}

	return nil
}
