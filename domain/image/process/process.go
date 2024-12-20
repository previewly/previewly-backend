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
		Create([]types.ImageProcessOption) (Process, error)
	}
)

func GetProcessFactory(processType types.ImageProcessType) (ProcessFactory, error) {
	if processType == types.Resize {
		return resizeProcessFactoryImpl{}, nil
	}
	return nil, errors.New("process type not found")
}
