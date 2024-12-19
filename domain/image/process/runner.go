package process

import (
	"wsw/backend/ent"
	"wsw/backend/ent/types"

	"github.com/xorcare/pointer"
)

type (
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
	processRunnerimpl struct{}
)

// Start implements ProcessRunner.
func (p processRunnerimpl) Start(image *ent.UploadImage, processes []types.ImageProcess) RunnerResult {
	return p.createSuccessResult("", pointer.String(image.Filename), nil)
}

func (p processRunnerimpl) createSuccessResult(prefix string, imageName *string, imageURL *string) RunnerResult {
	return RunnerResult{PrefixPath: &prefix, Status: types.Success, Error: nil, ImageName: imageName, ImageURL: imageURL}
}

func NewProcessRunner() ProcessRunner { return processRunnerimpl{} }
