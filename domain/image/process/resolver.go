package process

import (
	"context"

	"wsw/backend/ent"
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
		imagesModel    image.UploadedImage
		processesModel image.ImageProcesses
		gqlConvertor   Convertor
		runner         ProcessRunner
	}
)

func NewProcessResolver(imagesModel image.UploadedImage, processesModel image.ImageProcesses, gqlConvertor Convertor, runner ProcessRunner) Resolver {
	return resolverImpl{imagesModel: imagesModel, processesModel: processesModel, gqlConvertor: gqlConvertor, runner: runner}
}

// Resolve implements Resolver.
func (r resolverImpl) Resolve(ctx context.Context, imageID int, processes []*model.ImageProcessesInput) (*model.ImageProcesses, error) {
	imageEntity, err := r.imagesModel.GetByID(imageID)
	if err != nil {
		return nil, err
	}

	imageProcesses := r.validateProcesses(processes)
	return r.createImageProcess(imageEntity, imageProcesses)
}

func (r resolverImpl) createImageProcess(imageEntity *ent.UploadImage, imageProcesses []types.ImageProcess) (*model.ImageProcesses, error) {
	processEntity, err := r.processesModel.Create(imageEntity, imageProcesses)
	if err != nil {
		return nil, err
	}

	runnerResult := r.runner.Start(imageEntity, imageProcesses)

	processEntity, errSaving := r.processesModel.Update(processEntity, runnerResult.PrefixPath, runnerResult.Status, runnerResult.Error)
	if errSaving != nil {
		return nil, errSaving
	}
	return r.gqlConvertor.Convert(processEntity), nil
}

func (r resolverImpl) validateProcesses(processes []*model.ImageProcessesInput) []types.ImageProcess {
	validedProcesses := make([]*types.ImageProcess, 0, len(processes))
	for _, process := range processes {
		validProcess := r.createValidProcess(process)
		if validProcess != nil {
			validedProcesses = append(validedProcesses, validProcess)
		}

	}
	validedProcesses = utils.FilterNil(validedProcesses)

	result := make([]types.ImageProcess, 0, len(validedProcesses))
	for _, process := range validedProcesses {
		result = append(result, *process)
	}
	return result
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
