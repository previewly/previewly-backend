package process

import (
	"wsw/backend/ent"
	"wsw/backend/ent/types"
)

type (
	RunnerResult struct {
		PrefixPath *string
		Status     types.StatusEnum
		Error      error
	}
	ProcessRunner interface {
		Start(*ent.UploadImage, []types.ImageProcess) RunnerResult
	}
	processRunnerimpl struct{}
)

// Start implements ProcessRunner.
func (p processRunnerimpl) Start(*ent.UploadImage, []types.ImageProcess) RunnerResult {
	return p.createSuccessResult("")
}

func (p processRunnerimpl) createSuccessResult(prefix string) RunnerResult {
	return RunnerResult{PrefixPath: &prefix, Status: types.Success, Error: nil}
}

func NewProcessRunner() ProcessRunner { return processRunnerimpl{} }
