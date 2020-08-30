package helper

import (
	"time"

	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/build"
)

func FlattenAssetVersions(input *[]build.FetchAssetVersion) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	for _, prop := range *input {
		result := make(map[string]interface{})
		result["sid"] = prop.Sid
		result["account_sid"] = prop.AccountSid
		result["service_sid"] = prop.ServiceSid
		result["asset_sid"] = prop.AssetSid
		result["date_created"] = prop.DateCreated.Format(time.RFC3339)
		result["path"] = prop.Path
		result["visibility"] = prop.Visibility

		results = append(results, result)
	}

	return &results
}

func FlattenFunctionVersions(input *[]build.FetchFunctionVersion) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	for _, prop := range *input {
		result := make(map[string]interface{})
		result["sid"] = prop.Sid
		result["account_sid"] = prop.AccountSid
		result["service_sid"] = prop.ServiceSid
		result["function_sid"] = prop.FunctionSid
		result["date_created"] = prop.DateCreated.Format(time.RFC3339)
		result["path"] = prop.Path
		result["visibility"] = prop.Visibility

		results = append(results, result)
	}

	return &results
}

func FlattenDependencies(input *[]build.FetchDependency) map[string]string {
	if input == nil {
		return nil
	}

	results := make(map[string]string, len(*input))

	for _, prop := range *input {
		results[prop.Name] = prop.Version
	}

	return results
}
