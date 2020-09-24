package utils

import "strings"

func ConvertToStringSlice(input []interface{}) []string {
	stringArray := make([]string, len(input))
	for index, value := range input {
		stringArray[index] = value.(string)
	}
	return stringArray
}

func ConvertSliceToSeperatedString(input []interface{}, separator string) string {
	stringSlice := ConvertToStringSlice(input)
	return strings.Join(stringSlice[:], separator)
}
