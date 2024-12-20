package process

import (
	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"

	"github.com/h2non/bimg"
)

type (
	resizeProcessImpl struct {
		options []types.ImageProcessOption
	}
)

func NewResizeProcess(options []types.ImageProcessOption) Process {
	return resizeProcessImpl{options: options}
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
