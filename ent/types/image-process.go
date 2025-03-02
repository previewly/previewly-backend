package types

type (
	ImageProcessType   string
	ImageProcessOption struct {
		Key   string
		Value *string
	}
	ImageProcess struct {
		Type    ImageProcessType
		Options []ImageProcessOption
	}
)

const (
	Resize ImageProcessType = "resize"
	Crop   ImageProcessType = "crop"
)

// Values provides list valid values for Enum.
func (ImageProcessType) Values() (kinds []string) {
	for _, s := range []ImageProcessType{Resize} {
		kinds = append(kinds, string(s))
	}
	return
}

func NewImageProcessOption(key string, value *string) *ImageProcessOption {
	return &ImageProcessOption{Key: key, Value: value}
}

func NewImageProcessType(value string) ImageProcessType {
	return ImageProcessType(value)
}

func NewImageProcess(processType ImageProcessType, options []ImageProcessOption) *ImageProcess {
	return &ImageProcess{
		Type:    processType,
		Options: options,
	}
}
