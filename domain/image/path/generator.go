package path

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type (
	FilenameGenerator interface {
		GenerateFilename(string) string
		GenerateFilepath(*string) string
	}
	filenameGeneratorImpl struct{}
)

// GenerateFilepath implements FilenameProvider.
func (f filenameGeneratorImpl) GenerateFilepath(prefix *string) string {
	path := fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	if prefix != nil {
		return strings.Join([]string{path, *prefix}, "/")
	} else {
		return path
	}
}

// GenerateFilename implements FilenameProvider.
func (f filenameGeneratorImpl) GenerateFilename(filename string) string {
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(filename))
}

func NewFilenameProvider() FilenameGenerator {
	return filenameGeneratorImpl{}
}
