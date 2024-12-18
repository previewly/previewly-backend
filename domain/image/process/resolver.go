package process

import (
	"context"

	"wsw/backend/graph/model"
)

type (
	Resolver interface {
		Resolve(context.Context, int, []*model.ImageProcessesInput) (*model.ImageProcesses, error)
	}

	resolverImpl struct{}
)

// Resolve implements Resolver.
func (r resolverImpl) Resolve(context.Context, int, []*model.ImageProcessesInput) (*model.ImageProcesses, error) {
	panic("unimplemented")
}

func NewProcessResolver() Resolver {
	return resolverImpl{}
}
