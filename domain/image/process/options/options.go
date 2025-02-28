package options

import (
	"strconv"

	"wsw/backend/ent/types"
)

func ExtractIntOption(options []types.ImageProcessOption, key string) *int {
	value := ExtractStringOption(options, key)
	if value != nil {
		intValue, err := strconv.Atoi(*value)
		if err != nil {
			return nil
		}
		return &intValue
	}
	return nil
}

func ExtractStringOption(options []types.ImageProcessOption, key string) *string {
	for _, option := range options {
		if option.Key == key {
			return option.Value
		}
	}
	return nil
}
