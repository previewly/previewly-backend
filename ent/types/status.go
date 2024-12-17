package types

type StatusEnum string

const (
	Success StatusEnum = "success"
	Error   StatusEnum = "error"
	Pending StatusEnum = "pending"
)

// Values provides list valid values for Enum.
func (StatusEnum) Values() (kinds []string) {
	for _, s := range []StatusEnum{Success, Error, Pending} {
		kinds = append(kinds, string(s))
	}
	return
}
