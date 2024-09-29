package gowitness

import (
	"wsw/backend/domain/path/screenshot/relative"
	"wsw/backend/domain/url"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"

	"github.com/sensepost/gowitness/pkg/models"
	"github.com/sensepost/gowitness/pkg/writers"
)

type (
	writerImpl struct {
		url                  *ent.Url
		repository           repository.Url
		relativePathProvider relative.Provider
	}
)

// Write implements writers.Writer.
func (w writerImpl) Write(result *models.Result) error {
	_, err := w.repository.Update(w.relativePathProvider.Provide(result.Filename), url.Success, w.url.ID, nil)
	return err
}

func NewRunnerWriter(url *ent.Url, repository repository.Url, relativePathProvider relative.Provider) writers.Writer {
	return writerImpl{
		url:                  url,
		repository:           repository,
		relativePathProvider: relativePathProvider,
	}
}
