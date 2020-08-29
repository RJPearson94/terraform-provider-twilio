package helper

func FlattenNotifications(input map[string]interface{}) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["added_to_channel"] = FlattenChannelUserModification(input["added_to_channel"].(map[string]interface{}))
	result["invited_to_channel"] = FlattenChannelUserModification(input["invited_to_channel"].(map[string]interface{}))
	result["removed_from_channel"] = FlattenChannelUserModification(input["removed_from_channel"].(map[string]interface{}))
	result["new_message"] = FlattenNewMesage(input["new_message"].(map[string]interface{}))
	result["log_enabled"] = input["log_enabled"]
	results = append(results, result)

	return &results
}

func FlattenChannelUserModification(input map[string]interface{}) *[]interface{} {
	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["enabled"] = input["enabled"]
	result["template"] = input["template"]
	result["sound"] = input["sound"]

	results = append(results, result)
	return &results
}

func FlattenNewMesage(input map[string]interface{}) *[]interface{} {
	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["enabled"] = input["enabled"]
	result["template"] = input["template"]
	result["sound"] = input["sound"]
	result["badge_count_enabled"] = input["badge_count_enabled"]

	results = append(results, result)
	return &results
}

func FlattenLimits(input map[string]interface{}) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["user_channels"] = input["user_channels"]
	result["channel_members"] = input["channel_members"]

	results = append(results, result)
	return &results
}

func FlattenMedia(input map[string]interface{}) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["compatibility_message"] = input["compatibility_message"]
	result["size_limit_mb"] = input["size_limit_mb"]

	results = append(results, result)
	return &results
}
