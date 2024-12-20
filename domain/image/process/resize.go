package process

import (
	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"
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
func (r resizeProcessImpl) Run(imagePath path.PathData) (path.PathData, error) {
	return path.PathData{
		FullPath:  "Sss",
		Directory: "sxsdsdsd",
	}, nil
}
