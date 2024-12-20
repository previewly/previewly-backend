package process

import (
	"wsw/backend/domain/image/path"
	"wsw/backend/ent"
	"wsw/backend/ent/types"

	"github.com/xorcare/pointer"
)

type (
	RunnerResult struct {
		PrefixPath string
		ImageName  *string
		ImageURL   *string
		Status     types.StatusEnum
		Error      error
	}
	ProcessRunner interface {
		Start(*ent.UploadImage, []types.ImageProcess) RunnerResult
	}
	processRunnerimpl struct {
		pathProvider path.PathProvider
	}
)

func NewProcessRunner(pathProvider path.PathProvider) ProcessRunner {
	return processRunnerimpl{pathProvider: pathProvider}
}

// Start implements ProcessRunner.
func (p processRunnerimpl) Start(image *ent.UploadImage, processes []types.ImageProcess) RunnerResult {
	var imagePath *path.PathData
	imagePath = p.pathProvider.Provide(image.DestinationPath, image.Filename)
	for _, processInput := range processes {
		processFactory, err := GetProcessFactory(processInput.Type)
		if err != nil {
			return p.createError(err)
		}
		process, err := processFactory.Create(processInput.Options)
		if err != nil {
			return p.createError(err)
		}
		updatedImagePath, err := process.Run(*imagePath)
		if err != nil {
			return p.createError(err)
		}
		imagePath = updatedImagePath
	}
	return p.createSuccessResult("", pointer.String(image.Filename), nil)
}

func (p processRunnerimpl) createError(err error) RunnerResult {
	return RunnerResult{PrefixPath: "", Status: types.Error, Error: err, ImageName: nil, ImageURL: nil}
}

func (p processRunnerimpl) createSuccessResult(prefix string, imageName *string, imageURL *string) RunnerResult {
	return RunnerResult{PrefixPath: prefix, Status: types.Success, Error: nil, ImageName: imageName, ImageURL: imageURL}
}
