package process

import (
	"context"

	"wsw/backend/ent/types"
	"wsw/backend/graph/model"
	"wsw/backend/lib/utils"
	"wsw/backend/model/image"
)

type (
	Resolver interface {
		Resolve(context.Context, int, []*model.ImageProcessesInput) (*model.ImageProcesses, error)
	}

	resolverImpl struct {
		imagesModel image.UploadedImage
	}
)

func NewProcessResolver(imagesModel image.UploadedImage) Resolver {
	return resolverImpl{imagesModel: imagesModel}
}

// Resolve implements Resolver.
func (r resolverImpl) Resolve(ctx context.Context, imageID int, processes []*model.ImageProcessesInput) (*model.ImageProcesses, error) {
	imageEntity, err := r.imagesModel.GetByID(imageID)
	if err != nil {
		return nil, err
	}

	imageProcesses := r.validateProcesses(processes)

	utils.D(imageEntity, imageProcesses)
	panic("unimplemented")
}

func (r resolverImpl) validateProcesses(processes []*model.ImageProcessesInput) []*types.ImageProcess {
	validedProcesses := make([]*types.ImageProcess, 0, len(processes))
	for _, process := range processes {
		validProcess := r.createValidProcess(process)
		if validProcess != nil {
			validedProcesses = append(validedProcesses, validProcess)
		}

	}
	return utils.FilterNil(validedProcesses)
}

func (r resolverImpl) createValidProcess(input *model.ImageProcessesInput) *types.ImageProcess {
	processType := types.NewImageProcessType(input.Type.String())

	options := make([]types.ImageProcessOption, 0, len(input.Options))
	for _, option := range input.Options {
		newOption := types.NewImageProcessOption(option.Key, option.Value)
		if newOption != nil {
			options = append(options, *newOption)
		}
	}

	return types.NewImageProcess(processType, options)
}
