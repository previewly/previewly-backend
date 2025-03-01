package processor

import (
	"errors"

	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"
)

type (
	Factory interface {
		NewProcessor(processes []types.ImageProcess) (Processor, error)
	}

	processorFactoryImpl struct {
		pathProvider  path.PathProvider
		pathGenerator path.FilenameGenerator
	}
)

// NewProcessor implements ProcessorFactory.
func (p processorFactoryImpl) NewProcessor(processes []types.ImageProcess) (Processor, error) {
	processors := make([]Processor, 0, len(processes))
	for _, processInput := range processes {
		process, err := p.createProcessor(processInput.Type, processInput.Options)
		if err != nil {
			return nil, err
		}
		processors = append(processors, process)
	}
	return NewCompositeProcessor(processors), nil
}

func (p processorFactoryImpl) createProcessor(processType types.ImageProcessType, options []types.ImageProcessOption) (Processor, error) {
	if processType == types.Resize {
		return NewResizeProcessor(p.pathProvider, p.pathGenerator, options)
	}
	return nil, errors.New("process type not found")
}

func NewProcessorFactory(pathProvider path.PathProvider, pathGenerator path.FilenameGenerator) Factory {
	return processorFactoryImpl{pathProvider: pathProvider, pathGenerator: pathGenerator}
}
