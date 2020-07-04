package utils

func ConvertToStringSlice(input []interface{}) []string {
	stringArray := make([]string, len(input))
	for index, value := range input {
		stringArray[index] = value.(string)
	}
	return stringArray
}
