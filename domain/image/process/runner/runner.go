package runner

import (
	"wsw/backend/domain/image/path"
	"wsw/backend/ent"
	"wsw/backend/ent/types"
	"wsw/backend/lib/utils"

	"github.com/xorcare/pointer"
)

type (
	Process        interface{}
	ProcessFactory interface {
		Create(types.ImageProcess) Process
	}
	RunnerResult struct {
		PrefixPath *string
		ImageName  *string
		ImageURL   *string
		Status     types.StatusEnum
		Error      error
	}
	ProcessRunner interface {
		Start(*ent.UploadImage, []types.ImageProcess) RunnerResult
	}
	processRunnerimpl struct {
		pathProvider   path.PathProvider
		processFactory ProcessFactory
	}
)

func NewProcessRunner(pathProvider path.PathProvider, processFactory ProcessFactory) ProcessRunner {
	return processRunnerimpl{pathProvider: pathProvider, processFactory: processFactory}
}

// Start implements ProcessRunner.
func (p processRunnerimpl) Start(image *ent.UploadImage, processes []types.ImageProcess) RunnerResult {
	// var imagePath path.PathData
	utils.D(p.pathProvider.Provide(image.DestinationPath, image.Filename))
	for _, processInput := range processes {
		process := p.processFactory.Create(processInput)
	}
	return p.createSuccessResult("", pointer.String(image.Filename), nil)
}

func (p processRunnerimpl) createSuccessResult(prefix string, imageName *string, imageURL *string) RunnerResult {
	return RunnerResult{PrefixPath: &prefix, Status: types.Success, Error: nil, ImageName: imageName, ImageURL: imageURL}
}
