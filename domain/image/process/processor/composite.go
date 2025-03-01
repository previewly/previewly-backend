package processor

import (
	"strings"

	"wsw/backend/domain/image/path"
	"wsw/backend/lib/utils"
)

type (
	compositeProcessor struct {
		processors []Processor
	}
)

func NewCompositeProcessor(processors []Processor) Processor {
	return compositeProcessor{processors: processors}
}

// Run implements Processor.
func (c compositeProcessor) Run(from path.PathData, filename string) (*path.PathData, error) {
	for _, process := range c.processors {
		to, err := process.Run(from, filename)
		if err != nil {
			return nil, err
		}
		from = *to
	}
	return &from, nil
}

func (c compositeProcessor) GetHash() string {
	var sb strings.Builder
	for _, process := range c.processors {
		sb.WriteString(process.GetHash())
	}
	return utils.GetMD5Hash(sb.String())
}

// GetPathPrefix implements Processor.
func (c compositeProcessor) GetPathPrefix() string {
	panic("Count not get path prefix for composite processor")
}
