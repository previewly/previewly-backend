package url

type Status string

const (
	Success Status = "success"
	Error   Status = "error"
	Pending Status = "pending"
)

// Values provides list valid values for Enum.
func (Status) Values() (kinds []string) {
	for _, s := range []Status{Success, Error, Pending} {
		kinds = append(kinds, string(s))
	}
	return
}
