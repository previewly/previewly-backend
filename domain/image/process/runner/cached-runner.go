package runner

import (
	"wsw/backend/domain/image/process/input"
	"wsw/backend/domain/image/process/processor"
	"wsw/backend/domain/image/process/runner/result"
	"wsw/backend/domain/image/url"
	"wsw/backend/model/image"
)

type (
	cachedRunnerImpl struct {
		runner         ProcessRunner
		urlProvider    url.Provider
		processesModel image.ImageProcesses
		resultFactory  result.Factory
	}
)

func NewCachedRunner(runner ProcessRunner, urlProvider url.Provider, processesModel image.ImageProcesses, resultFactory result.Factory) ProcessRunner {
	return cachedRunnerImpl{
		runner:         runner,
		urlProvider:    urlProvider,
		processesModel: processesModel,
		resultFactory:  resultFactory,
	}
}

// Start implements ProcessRunner.
func (c cachedRunnerImpl) Start(input input.Input, processor processor.Processor) (*result.Result, error) {
	processEntity, err := c.processesModel.TryGetByHash(input.Image.ID, processor.GetHash())
	if err != nil {
		return nil, err
	}

	if processEntity != nil {
		return c.resultFactory.New(input, processEntity)
	} else {
		return c.runner.Start(input, processor)
	}
}
