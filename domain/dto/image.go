package dto

type (
	Image interface {
		Filename() string
		RelativeDirectory() string
		RelativeFullPath() string
	}
	imageImpl struct {
		filename          string
		relativeDirectory string
	}
)

func (i *imageImpl) Filename() string { return i.filename }

func (i *imageImpl) RelativeDirectory() string { return i.relativeDirectory }

func (i *imageImpl) RelativeFullPath() string { return i.relativeDirectory + i.filename }

func NewImage(filename string, relativeDirectory string) Image {
	return &imageImpl{filename: filename, relativeDirectory: relativeDirectory}
}
