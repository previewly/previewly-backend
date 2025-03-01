package processor

import "wsw/backend/domain/image/path"

type (
	Processor interface {
		Run(from path.PathData, filename string) (*path.PathData, error)
		GetHash() string
	}
)
