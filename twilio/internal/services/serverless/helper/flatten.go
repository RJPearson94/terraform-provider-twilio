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
		results = append(results, map[string]interface{}{
			"sid":          prop.Sid,
			"account_sid":  prop.AccountSid,
			"service_sid":  prop.ServiceSid,
			"asset_sid":    prop.AssetSid,
			"date_created": prop.DateCreated.Format(time.RFC3339),
			"path":         prop.Path,
			"visibility":   prop.Visibility,
		})
	}

	return &results
}

func FlattenFunctionVersions(input *[]build.FetchFunctionVersion) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	for _, prop := range *input {
		results = append(results, map[string]interface{}{
			"sid":          prop.Sid,
			"account_sid":  prop.AccountSid,
			"service_sid":  prop.ServiceSid,
			"function_sid": prop.FunctionSid,
			"date_created": prop.DateCreated.Format(time.RFC3339),
			"path":         prop.Path,
			"visibility":   prop.Visibility,
		})
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
