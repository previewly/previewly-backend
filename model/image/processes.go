package image

type (
	ImageProcesses     interface{}
	imageProcessesImpl struct{}
)

func NewImageProcesses() ImageProcesses {
	return imageProcessesImpl{}
}
