package runner

import (
	"wsw/backend/domain/image/path"
	"wsw/backend/domain/image/process/input"
	"wsw/backend/domain/image/process/processor"
	"wsw/backend/domain/image/process/runner/result"
	"wsw/backend/domain/image/url"
	"wsw/backend/ent"
	"wsw/backend/ent/types"
	"wsw/backend/model/image"
)

type (
	processRunnerImpl struct {
		pathProvider   path.PathProvider
		urlProvider    url.Provider
		processesModel image.ImageProcesses
		resultFactory  result.Factory
	}
)

func NewProcessRunner(pathProvider path.PathProvider, urlProvider url.Provider, processesModel image.ImageProcesses,
	resultFactory result.Factory,
) ProcessRunner {
	return processRunnerImpl{
		pathProvider:   pathProvider,
		urlProvider:    urlProvider,
		processesModel: processesModel,
		resultFactory:  resultFactory,
	}
}

// Start implements ProcessRunner.
func (p processRunnerImpl) Start(input input.Input, processor processor.Processor) (*result.Result, error) {
	processEntity, err := p.doProcess(input.Image, input.Processes, processor)
	if err != nil {
		return nil, err
	}
	updatedProcessEntity, err := p.processesModel.Update(processEntity, processEntity.PathPrefix, processEntity.Status, processEntity.Error)
	if err != nil {
		return nil, err
	}
	return p.resultFactory.New(input, updatedProcessEntity)
}

func (p processRunnerImpl) doProcess(image *ent.Image, processes []types.ImageProcess, processor processor.Processor) (*ent.ImageProcess, error) {
	processEntity, err := p.createProcessEntity(image, processes, processor.GetHash())
	if err != nil {
		return nil, err
	}
	imagePath, err := p.runProcesses(image, processor)
	if err != nil {
		processEntity.Status = types.Error
		processEntity.Error = err.Error()

	} else {
		processEntity.PathPrefix = imagePath.RelativeDirectory
		processEntity.Status = types.Success
	}
	return processEntity, nil
}

func (p processRunnerImpl) createProcessEntity(image *ent.Image, processes []types.ImageProcess, hash string) (*ent.ImageProcess, error) {
	return p.processesModel.Create(image, processes, hash)
}

func (p processRunnerImpl) runProcesses(image *ent.Image, processor processor.Processor) (*path.PathData, error) {
	from := p.pathProvider.Provide(image.DestinationPath, image.Filename)

	toPath, processError := processor.Run(*from, image.Filename)
	if processError != nil {
		return nil, processError
	}

	return toPath, nil
}
