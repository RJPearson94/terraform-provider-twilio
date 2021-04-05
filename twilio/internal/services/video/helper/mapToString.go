package helper

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
)

func FlattenJsonToStringOrEmptyObjectString(jsonMap map[string]interface{}) (string, error) {
	if len(jsonMap) == 0 {
		return "{}", nil
	}
	return structure.FlattenJsonToString(jsonMap)

}
