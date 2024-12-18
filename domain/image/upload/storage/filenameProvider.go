package storage

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type (
	FilenameProvider interface {
		GenerateFilename(string) string
		GenerateFilepath(*string) string
	}
	filenameProviderImpl struct{}
)

// GenerateFilepath implements FilenameProvider.
func (f filenameProviderImpl) GenerateFilepath(prefix *string) string {
	path := fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	if prefix != nil {
		return strings.Join([]string{path, *prefix}, "/")
	} else {
		return path
	}
}

// GenerateFilename implements FilenameProvider.
func (f filenameProviderImpl) GenerateFilename(filename string) string {
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(filename))
}

func NewFilenameProvider() FilenameProvider {
	return filenameProviderImpl{}
}
