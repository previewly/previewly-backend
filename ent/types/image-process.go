package types

type ImageProcess string

const (
	Resize ImageProcess = "resize"
)

// Values provides list valid values for Enum.
func (ImageProcess) Values() (kinds []string) {
	for _, s := range []ImageProcess{Resize} {
		kinds = append(kinds, string(s))
	}
	return
}
