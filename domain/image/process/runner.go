package process

import (
	"wsw/backend/ent"
	"wsw/backend/ent/types"
)

type (
	ProcessRunner interface {
		Start(*ent.UploadImage, []types.ImageProcess) (types.StatusEnum, error)
	}
	processRunnerimpl struct{}
)

// Start implements ProcessRunner.
func (p processRunnerimpl) Start(*ent.UploadImage, []types.ImageProcess) (types.StatusEnum, error) {
	return types.Success, nil
}

func NewProcessRunner() ProcessRunner { return processRunnerimpl{} }
