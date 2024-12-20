package process

import (
	"errors"

	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"

	"github.com/h2non/bimg"
)

type (
	resizeProcessFactoryImpl struct{}
	resizeProcessImpl        struct {
		width  *int
		height *int
	}
)

// Run implements Process.
func (r resizeProcessImpl) Run(imagePath path.PathData) (*path.PathData, error) {
	_, err := bimg.Read(imagePath.FullPath)
	if err != nil {
		return nil, err
	}
	return &path.PathData{
		FullPath:  "Sss",
		Directory: "sxsdsdsd",
	}, nil
}

// Create implements ProcessFactory.
func (r resizeProcessFactoryImpl) Create(options []types.ImageProcessOption) (Process, error) {
	width := GetIntOption(options, "width")
	height := GetIntOption(options, "height")

	if width != nil || height != nil {
		return resizeProcessImpl{width: width, height: height}, nil
	}
	return nil, errors.New("width or height should be provided")
}
