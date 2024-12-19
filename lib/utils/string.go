package utils

func ToString(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}
