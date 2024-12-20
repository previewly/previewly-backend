package process

import (
	"errors"

	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"
)

type (
	Process interface {
		Run(path.PathData) (*path.PathData, error)
	}
	ProcessFactory interface {
		Create(types.ImageProcessType, []types.ImageProcessOption) (Process, error)
	}
	processFactoryImpl struct{}
)

func NewProcessFactory() ProcessFactory { return processFactoryImpl{} }

func (p processFactoryImpl) Create(processType types.ImageProcessType, options []types.ImageProcessOption) (Process, error) {
	if processType == types.Resize {
		return NewResizeProcess(options), nil
	}
	return nil, errors.New("unsupported process type")
}
