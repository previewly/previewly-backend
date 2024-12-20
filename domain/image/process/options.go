package process

import (
	"strconv"

	"wsw/backend/ent/types"
)

func GetIntOption(options []types.ImageProcessOption, key string) *int {
	value := GetStringOption(options, key)
	if value != nil {
		intValue, err := strconv.Atoi(*value)
		if err != nil {
			return nil
		}
		return &intValue
	}
	return nil
}

func GetStringOption(options []types.ImageProcessOption, key string) *string {
	for _, option := range options {
		if option.Key == key {
			return option.Value
		}
	}
	return nil
}
