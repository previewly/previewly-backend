package process

import "wsw/backend/ent"

type (
	ProcessRunner interface {
		Start(*ent.ImageProcess) (*ent.ImageProcess, error)
	}
	processRunnerimpl struct{}
)

func (p processRunnerimpl) Start(processEntity *ent.ImageProcess) (*ent.ImageProcess, error) {
	panic("unimplemented")
}

func NewProcessRunner() ProcessRunner { return processRunnerimpl{} }
