package process

import (
	"context"

	"wsw/backend/domain/image/process/input"
	"wsw/backend/domain/image/process/processor"
	"wsw/backend/domain/image/process/runner"
	"wsw/backend/ent"
	"wsw/backend/ent/types"
	"wsw/backend/graph/model"
	"wsw/backend/lib/utils"
	"wsw/backend/model/image"
)

type (
	Resolver interface {
		Resolve(context.Context, int, []*model.ImageProcessesInput) (*model.ImageProcess, error)
	}

	resolverImpl struct {
		imagesModel      image.Model
		gqlConvertor     Convertor
		runner           runner.ProcessRunner
		processorFactory processor.Factory
	}
)

func NewProcessResolver(
	imagesModel image.Model,
	gqlConvertor Convertor,
	runner runner.ProcessRunner,
	processorFactory processor.Factory,
) Resolver {
	return resolverImpl{
		imagesModel:      imagesModel,
		gqlConvertor:     gqlConvertor,
		runner:           runner,
		processorFactory: processorFactory,
	}
}

// Resolve implements Resolver.
func (r resolverImpl) Resolve(ctx context.Context, imageID int, processes []*model.ImageProcessesInput) (*model.ImageProcess, error) {
	imageEntity, err := r.imagesModel.GetByID(imageID)
	if err != nil {
		return nil, err
	}

	imageProcesses := r.validateProcesses(processes)
	return r.createImageProcess(imageEntity, imageProcesses)
}

func (r resolverImpl) createImageProcess(imageEntity *ent.Image, imageProcesses []types.ImageProcess) (*model.ImageProcess, error) {
	inputArg := input.Input{Image: imageEntity, Processes: imageProcesses}

	processor, err := r.processorFactory.NewProcessor(imageProcesses)
	if err != nil {
		return nil, err
	}

	runnerResult, err := r.runner.Start(inputArg, processor)
	if err != nil {
		return nil, err
	}
	return r.gqlConvertor.Convert(*runnerResult), nil
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
