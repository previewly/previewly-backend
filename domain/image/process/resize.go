package process

import (
	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"

	"github.com/h2non/bimg"
)

type (
	resizeProcessFactoryImpl struct{}
	resizeProcessImpl        struct {
		options []types.ImageProcessOption
	}
)

// Create implements ProcessFactory.
func (r resizeProcessFactoryImpl) Create(options []types.ImageProcessOption) (Process, error) {
	panic("unimplemented")
}

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
