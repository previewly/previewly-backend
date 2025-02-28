package runner

import (
	"wsw/backend/domain/image/process/input"
	"wsw/backend/domain/image/process/processor"
	"wsw/backend/domain/image/process/runner/result"
)

type (
	ProcessRunner interface {
		Start(input input.Input, processor processor.Processor) (*result.Result, error)
	}
)
