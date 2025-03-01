package path

import (
	"wsw/backend/domain/image/path"

	"github.com/xorcare/pointer"
)

type (
	Provider interface {
		Get(prefix string, filename string) *path.PathData
	}

	providerImpl struct {
		pathProvider  path.PathProvider
		pathGenerator path.FilenameGenerator
	}
)

func NewProvider(pathProvider path.PathProvider, pathGenerator path.FilenameGenerator) Provider {
	return providerImpl{pathProvider: pathProvider, pathGenerator: pathGenerator}
}

// Get implements Provider.
func (p providerImpl) Get(prefix string, filename string) *path.PathData {
	return p.pathProvider.Provide(
		p.pathGenerator.GenerateFilepath(pointer.String(prefix)),
		filename,
	)
}
