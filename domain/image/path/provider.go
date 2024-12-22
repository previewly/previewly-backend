package path

import "strings"

type (
	PathData struct {
		RelativeDirectory string
		RelativeFullPath  string
		Directory         string
		FullPath          string
	}
	PathProvider interface {
		Provide(directory string, filename string) *PathData
	}
	pathProviderImpl struct {
		baseDirectory string
	}
)

func (p pathProviderImpl) Provide(directory string, filename string) *PathData {
	absoluteDirectory := strings.Join([]string{
		strings.TrimSuffix(p.baseDirectory, "/"),
		strings.TrimSuffix(directory, "/"),
	}, "/")
	return &PathData{
		RelativeDirectory: directory,
		RelativeFullPath:  strings.Join([]string{strings.TrimSuffix(directory, "/"), filename}, "/"),
		Directory:         absoluteDirectory,
		FullPath:          strings.Join([]string{absoluteDirectory, filename}, "/"),
	}
}

func NewPathProvider(baseDirectory string) PathProvider {
	return pathProviderImpl{baseDirectory: baseDirectory}
}
