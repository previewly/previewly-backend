package utils

import (
	"slices"
)

func FilterNil[T any](values []*T) []*T {
	filtered := slices.DeleteFunc(
		values,
		func(thing *T) bool {
			return thing == nil
		},
	)
	return slices.Clip(filtered)
}
