package types

type (
	ImageProcessType    string
	ImageProcessOptions struct {
		Key   string
		Value *string
	}
	ImageProcess struct {
		Type    ImageProcessType
		Options []ImageProcessOptions
	}
)

const (
	Resize ImageProcessType = "resize"
)

// Values provides list valid values for Enum.
func (ImageProcessType) Values() (kinds []string) {
	for _, s := range []ImageProcessType{Resize} {
		kinds = append(kinds, string(s))
	}
	return
}
